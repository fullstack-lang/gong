package models

import (
	"math"
	"sort"
)

type layoutNode[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType] struct {
	shape    CT
	children []*layoutNode[AT, CT]
	parent   *layoutNode[AT, CT]
}

func layoutGenericShapes[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType, ACT LayoutAssociationType](
	shapes []CT,
	rootElements []AT,
	getChildren func(AT) []AT,
	getParent func(AT) AT,
	compositionShapes []ACT,
	getLinkElement func(ACT) AT,
	getName func(AT) string,
	nextX *float64,
	startY float64,
	margin float64,
) (hasChanged bool) {
	if len(shapes) == 0 {
		return false
	}

	nodesByElement := make(map[AT]*layoutNode[AT, CT])
	var rootNodes []*layoutNode[AT, CT]

	for _, shape := range shapes {
		if any(shape) == nil || shape.GetAbstractElement() == nil {
			continue
		}
		// Type assert to get the AT
		abstractElement := shape.GetAbstractElement().(AT)
		nodesByElement[abstractElement] = &layoutNode[AT, CT]{shape: shape}
	}

	for _, shape := range shapes {
		if any(shape) == nil || shape.GetAbstractElement() == nil {
			continue
		}
		abstractElement := shape.GetAbstractElement().(AT)
		node := nodesByElement[abstractElement]
		isRoot := true

		parent := getParent(abstractElement)
		// Go interfaces comparison against nil is tricky. We can check if it's valid.
		// Let's assume parent != nil implies it has a parent if it exists in the map
		if any(parent) != nil {
			if _, exists := nodesByElement[parent]; exists {
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children respecting order
	for _, node := range nodesByElement {
		abstractElement := node.shape.GetAbstractElement().(AT)
		for _, subElement := range getChildren(abstractElement) {
			if childNode, exists := nodesByElement[subElement]; exists {
				node.children = append(node.children, childNode)
				childNode.parent = node
			}
		}
	}

	rootElementIndex := make(map[AT]int)
	for i, sub := range rootElements {
		rootElementIndex[sub] = i
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		idxI, okI := rootElementIndex[rootNodes[i].shape.GetAbstractElement().(AT)]
		idxJ, okJ := rootElementIndex[rootNodes[j].shape.GetAbstractElement().(AT)]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return getName(rootNodes[i].shape.GetAbstractElement().(AT)) < getName(rootNodes[j].shape.GetAbstractElement().(AT))
	})

	for _, root := range rootNodes {
		maxX, _, rootChanged := layoutGenericDFS(root, *nextX, startY, margin)
		if rootChanged {
			hasChanged = true
		}
		*nextX = maxX
	}

	// Find the parent of each element
	parentByElement := make(map[AT]*layoutNode[AT, CT])
	for _, node := range nodesByElement {
		abstractElement := node.shape.GetAbstractElement().(AT)
		for _, subElement := range getChildren(abstractElement) {
			parentByElement[subElement] = node
		}
	}

	for _, link := range compositionShapes {
		if any(link) == nil {
			continue
		}
		if math.Abs(link.GetCornerOffsetRatio()-1.5) > 1e-4 {
			link.SetCornerOffsetRatio(1.5)
			hasChanged = true
		}
		element := getLinkElement(link)
		if any(element) != nil {
			if parentNode, ok := parentByElement[element]; ok {
				layoutDirection := Vertical
				if treeNode, ok := parentNode.shape.GetAbstractElement().(TreeAbstractType); ok {
					layoutDirection = treeNode.GetLayoutDirection()
				}
				if parentNode.shape.GetOverideLayoutDirection() {
					layoutDirection = parentNode.shape.GetConcreteLayoutDirection()
				}
				var expectedStartOrientation, expectedEndOrientation OrientationType
				if layoutDirection == Horizontal {
					expectedStartOrientation = ORIENTATION_VERTICAL
					expectedEndOrientation = ORIENTATION_HORIZONTAL
				} else {
					expectedStartOrientation = ORIENTATION_VERTICAL
					expectedEndOrientation = ORIENTATION_VERTICAL
				}
				if link.GetStartOrientation() != expectedStartOrientation {
					link.SetStartOrientation(expectedStartOrientation)
					hasChanged = true
				}
				if link.GetEndOrientation() != expectedEndOrientation {
					link.SetEndOrientation(expectedEndOrientation)
					hasChanged = true
				}
			}
		}
	}

	return hasChanged
}

func layoutGenericDFS[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType](node *layoutNode[AT, CT], currentX float64, currentY float64, margin float64) (float64, float64, bool) {
	hasChanged := false
	if math.Abs(node.shape.GetX()-currentX) > 1e-4 {
		node.shape.SetX(currentX)
		hasChanged = true
	}
	if math.Abs(node.shape.GetY()-currentY) > 1e-4 {
		node.shape.SetY(currentY)
		hasChanged = true
	}

	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		return currentX + w + margin, currentY + h + margin, hasChanged
	}

	var maxX float64 = currentX + w + margin
	var maxY float64 = currentY + h + margin

	layoutDirection := Vertical
	if treeNode, ok := node.shape.GetAbstractElement().(TreeAbstractType); ok {
		layoutDirection = treeNode.GetLayoutDirection()
	}
	if node.shape.GetOverideLayoutDirection() {
		layoutDirection = node.shape.GetConcreteLayoutDirection()
	}

	if layoutDirection == Vertical {
		// Children are arranged horizontally.
		childX := currentX

		isParentHorizontal := false
		if node.parent != nil {
			parentLayout := Vertical
			if treeParent, ok := node.parent.shape.GetAbstractElement().(TreeAbstractType); ok {
				parentLayout = treeParent.GetLayoutDirection()
			}
			if node.parent.shape.GetOverideLayoutDirection() {
				parentLayout = node.parent.shape.GetConcreteLayoutDirection()
			}
			if parentLayout == Horizontal {
				isParentHorizontal = true
			}
		}

		if isParentHorizontal {
			childX = currentX + w/2.0 + margin
		}

		for _, child := range node.children {
			childMaxX, childMaxY, childChanged := layoutGenericDFS(child, childX, currentY+h*2.0, margin)
			if childChanged {
				hasChanged = true
			}
			childX = childMaxX

			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		if !isParentHorizontal {
			newParentX := node.children[0].shape.GetX()
			if math.Abs(node.shape.GetX()-newParentX) > 1e-4 {
				node.shape.SetX(newParentX)
				hasChanged = true
			}
		}
		maxX = childX
	} else {
		// Children are arranged vertically (indented tree)
		verticalMargin := 15.0
		childY := currentY + h + verticalMargin
		for _, child := range node.children {
			childMaxX, childMaxY, childChanged := layoutGenericDFS(child, currentX+w/2.0+margin, childY, margin)
			if childChanged {
				hasChanged = true
			}
			childY = (childMaxY - margin) + verticalMargin

			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		// Do not align parent vertically with its children. Parent stays at (currentX, currentY).
		maxY = (childY - verticalMargin) + margin
	}

	return maxX, maxY, hasChanged
}

func (diagram *Diagram) Layout(stager *Stager) {
	layoutDiagram(diagram, stager)
}

func layoutDiagram(diagram *Diagram, stager *Stager) (hasChanged bool) {
	stager.enforceParentAssociation()

	var owningLibrary *Library
	for _, lib := range stager.stage.GetInstancesSorted[*Library]() {
		for _, diag := range lib.Diagrams {
			if diag == diagram {
				owningLibrary = lib
				break
			}
		}
		if owningLibrary != nil {
			break
		}
	}
	if owningLibrary == nil {
		owningLibrary = stager.getRootLibrary()
	}

	nextX := 50.0
	startY := 50.0
	margin := 50.0

	var rootProducts []*Product
	var rootTasks []*Task
	var rootResources []*Resource

	if owningLibrary != nil {
		rootProducts = owningLibrary.RootProducts
		rootTasks = owningLibrary.RootTasks
		rootResources = owningLibrary.RootResources
	}

	// Layout Products
	var productShapes []*ProductShape
	for _, shape := range diagram.Product_Shapes {
		productShapes = append(productShapes, shape)
	}
	var productLinks []*ProductCompositionShape
	for _, link := range diagram.ProductComposition_Shapes {
		productLinks = append(productLinks, link)
	}

	if layoutGenericShapes(
		productShapes,
		rootProducts,
		func(p *Product) []*Product { return p.SubProducts },
		func(p *Product) *Product { return p.parentProduct },
		productLinks,
		func(l *ProductCompositionShape) *Product { return l.Product },
		func(p *Product) string { return p.Name },
		&nextX, startY, margin,
	) {
		hasChanged = true
	}

	// Layout Tasks
	var taskShapes []*TaskShape
	for _, shape := range diagram.Task_Shapes {
		taskShapes = append(taskShapes, shape)
	}
	var taskLinks []*TaskCompositionShape
	for _, link := range diagram.TaskComposition_Shapes {
		taskLinks = append(taskLinks, link)
	}

	if layoutGenericShapes(
		taskShapes,
		rootTasks,
		func(t *Task) []*Task { return t.SubTasks },
		func(t *Task) *Task { return t.parentTask },
		taskLinks,
		func(l *TaskCompositionShape) *Task { return l.Task },
		func(t *Task) string { return t.Name },
		&nextX, startY, margin,
	) {
		hasChanged = true
	}

	// Layout Resources
	var resourceShapes []*ResourceShape
	for _, shape := range diagram.Resource_Shapes {
		resourceShapes = append(resourceShapes, shape)
	}
	var resourceLinks []*ResourceCompositionShape
	for _, link := range diagram.ResourceComposition_Shapes {
		resourceLinks = append(resourceLinks, link)
	}

	if layoutGenericShapes(
		resourceShapes,
		rootResources,
		func(r *Resource) []*Resource { return r.SubResources },
		func(r *Resource) *Resource { return r.parentResource },
		resourceLinks,
		func(l *ResourceCompositionShape) *Resource { return l.Resource },
		func(r *Resource) string { return r.Name },
		&nextX, startY, margin,
	) {
		hasChanged = true
	}

	return hasChanged
}

