package models

import "slices"

// reachesSystemViaPartTypes returns true if start == target or if following
// System -> Parts -> Type edges from start reaches target.
// Uses a visited map to remain safe even if a cycle currently exists.
func reachesSystemViaPartTypes(start *System, target *System) bool {
	if start == nil || target == nil {
		return false
	}
	visited := make(map[*System]bool)
	var dfs func(cur *System) bool
	dfs = func(cur *System) bool {
		if cur == nil {
			return false
		}
		if cur == target {
			return true
		}
		if visited[cur] {
			return false
		}
		visited[cur] = true
		for _, p := range cur.Parts {
			if p != nil && p.Type != nil {
				if dfs(p.Type) {
					return true
				}
			}
		}
		return false
	}
	return dfs(start)
}

// partShapeDepth returns 0 for top-level PartShapes (ParentPartShape == nil),
// positive depth for nested PartShapes, or -1 if ParentPartShape has a cycle.
func partShapeDepth(ps *PartShape) int {
	if ps == nil {
		return -1
	}
	depth := 0
	visited := make(map[*PartShape]bool)
	cur := ps
	for cur.ParentPartShape != nil {
		if visited[cur] {
			return -1
		}
		visited[cur] = true
		depth++
		cur = cur.ParentPartShape
	}
	return depth
}

// sortedPartShapesTopDown returns a new slice ordered by ascending nesting depth
// so parent PartShapes are always processed/clamped/drawn before child PartShapes.
func sortedPartShapesTopDown(shapes []*PartShape) []*PartShape {
	type indexedShape struct {
		shape *PartShape
		depth int
		index int
	}
	items := make([]indexedShape, len(shapes))
	for i, s := range shapes {
		items[i] = indexedShape{
			shape: s,
			depth: partShapeDepth(s),
			index: i,
		}
	}
	slices.SortFunc(items, func(a, b indexedShape) int {
		if a.depth != b.depth {
			return a.depth - b.depth
		}
		return a.index - b.index
	})
	res := make([]*PartShape, len(items))
	for i, item := range items {
		res[i] = item.shape
	}
	return res
}

// partShapeAncestors returns strict ancestors of ps ordered from immediate parent to outermost root.
func partShapeAncestors(ps *PartShape) []*PartShape {
	var ancestors []*PartShape
	if ps == nil {
		return ancestors
	}
	visited := make(map[*PartShape]bool)
	visited[ps] = true
	cur := ps.ParentPartShape
	for cur != nil {
		if visited[cur] {
			break
		}
		visited[cur] = true
		ancestors = append(ancestors, cur)
		cur = cur.ParentPartShape
	}
	return ancestors
}

// isPartShapeOrAncestorHidden returns true if ps is nil, hidden, missing from diagram.Part_Shapes,
// or if any of its ParentPartShape ancestors is hidden, cyclic, or missing from diagram.Part_Shapes.
func isPartShapeOrAncestorHidden(diagram *DiagramStructure, ps *PartShape) bool {
	if ps == nil || diagram == nil {
		return true
	}
	inDiagram := make(map[*PartShape]bool, len(diagram.Part_Shapes))
	for _, candidate := range diagram.Part_Shapes {
		inDiagram[candidate] = true
	}
	visited := make(map[*PartShape]bool)
	cur := ps
	for cur != nil {
		if visited[cur] {
			return true
		}
		visited[cur] = true
		if !inDiagram[cur] || cur.GetIsHidden() {
			return true
		}
		cur = cur.ParentPartShape
	}
	return false
}

// partShapeNameInDiagram computes a descriptive unique shape name for a PartShape in a diagram.
func partShapeNameInDiagram(diagram *DiagramStructure, part *Part, parentPartShape *PartShape) string {
	var chain []string
	for _, anc := range partShapeAncestors(&PartShape{ParentPartShape: parentPartShape}) {
		if anc.Part != nil {
			chain = append([]string{anc.Part.GetName()}, chain...)
		}
	}
	chain = append(chain, part.GetName())
	prefix := ""
	for i, name := range chain {
		if i > 0 {
			prefix += "-"
		}
		prefix += name
	}
	return prefix + "-" + diagram.GetName()
}

// effectivePortParentPartShape returns portShape.ParentPartShape when explicitly set,
// or falls back to the top-level PartShape (ParentPartShape == nil) of portShape.Port.owningPart
// for backward compatibility with non-nested diagrams.
func effectivePortParentPartShape(diagram *DiagramStructure, portShape *PortShape) *PartShape {
	if portShape == nil || diagram == nil {
		return nil
	}
	if portShape.ParentPartShape != nil {
		return portShape.ParentPartShape
	}
	if portShape.Port != nil && portShape.Port.owningPart != nil && diagram.map_PartShapeKey_PartShape != nil {
		return diagram.map_PartShapeKey_PartShape[partShapeKey{
			part:            portShape.Port.owningPart,
			parentPartShape: nil,
		}]
	}
	return nil
}

// effectiveControlFlowParentPartShape returns controlFlowShape.ParentPartShape when explicitly set,
// or falls back to the top-level PartShape (ParentPartShape == nil) of Start.owningPart.
func effectiveControlFlowParentPartShape(diagram *DiagramStructure, controlFlowShape *ControlFlowShape) *PartShape {
	if controlFlowShape == nil || diagram == nil {
		return nil
	}
	if controlFlowShape.ParentPartShape != nil {
		return controlFlowShape.ParentPartShape
	}
	if controlFlowShape.ControlFlow != nil &&
		controlFlowShape.ControlFlow.Start != nil &&
		controlFlowShape.ControlFlow.Start.owningPart != nil &&
		diagram.map_PartShapeKey_PartShape != nil {
		return diagram.map_PartShapeKey_PartShape[partShapeKey{
			part:            controlFlowShape.ControlFlow.Start.owningPart,
			parentPartShape: nil,
		}]
	}
	return nil
}

// effectiveDataFlowShape returns explicit if non-nil, or falls back to the top-level DataFlowShape
// (ParentPartShape == nil) for df in diagram.
func effectiveDataFlowShape(diagram *DiagramStructure, df *DataFlow, explicit *DataFlowShape) *DataFlowShape {
	if explicit != nil {
		return explicit
	}
	if diagram == nil || df == nil || diagram.map_DataFlowShapeKey_DataFlowShape == nil {
		return nil
	}
	return diagram.map_DataFlowShapeKey_DataFlowShape[dataFlowShapeKey{
		dataFlow:        df,
		parentPartShape: nil,
	}]
}

// resolvePortShapeInScope finds the PortShape of port belonging to a child PartShape
// whose ParentPartShape == enclosingPartShape (nil for top-level diagram Parts).
func resolvePortShapeInScope(diagram *DiagramStructure, port *Port, enclosingPartShape *PartShape) (*PortShape, bool) {
	if diagram == nil || port == nil || port.owningPart == nil {
		return nil, false
	}
	if diagram.map_PartShapeKey_PartShape == nil || diagram.map_PortShapeKey_PortShape == nil {
		return nil, false
	}
	targetPartShape, ok := diagram.map_PartShapeKey_PartShape[partShapeKey{
		part:            port.owningPart,
		parentPartShape: enclosingPartShape,
	}]
	if !ok || targetPartShape == nil {
		return nil, false
	}
	ps, ok := diagram.map_PortShapeKey_PortShape[portShapeKey{
		port:            port,
		parentPartShape: targetPartShape,
	}]
	return ps, ok && ps != nil
}

// resolveDataFlowEndpointPortShapes resolves the concrete start and end PortShapes for a DataFlowShape.
func resolveDataFlowEndpointPortShapes(diagram *DiagramStructure, dataFlowShape *DataFlowShape) (
	startPortShape *PortShape, startOK bool,
	endPortShape *PortShape, endOK bool,
) {
	if diagram == nil || dataFlowShape == nil || dataFlowShape.DataFlow == nil {
		return nil, false, nil, false
	}
	df := dataFlowShape.DataFlow
	resolveOne := func(explicit *PortShape, port *Port) (*PortShape, bool) {
		if port == nil {
			return nil, false
		}
		if explicit != nil {
			if explicit.Port == port && slices.Contains(diagram.Port_Shapes, explicit) {
				return explicit, true
			}
			return nil, false
		}
		return resolvePortShapeInScope(diagram, port, dataFlowShape.ParentPartShape)
	}
	startPortShape, startOK = resolveOne(dataFlowShape.StartPortShape, df.StartPort)
	endPortShape, endOK = resolveOne(dataFlowShape.EndPortShape, df.EndPort)
	return
}

// hasVisibleChildPartShape returns true if parentPartShape has at least one non-hidden
// child PartShape in diagram.Part_Shapes.
func hasVisibleChildPartShape(diagram *DiagramStructure, parentPartShape *PartShape) bool {
	if diagram == nil || parentPartShape == nil {
		return false
	}
	for _, ps := range diagram.Part_Shapes {
		if ps != nil && ps.ParentPartShape == parentPartShape && !ps.GetIsHidden() {
			return true
		}
	}
	return false
}
