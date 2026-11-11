package models

import (
	"fmt"
	"log"
	"math"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePartWithinLibrary(
	library *Library,
	part *Part,
	parentNode *tree.Node,
) {
	nodeName := partDisplayLabel(part)
	if part.GetIsInRenameMode() {
		nodeName = part.GetName()
	}
	partNode := &tree.Node{
		Name:            nodeName,
		IsExpanded:      slices.Contains(library.PartsWhoseNodeIsExpanded, part),
		IsNodeClickable: true,
		IsInEditMode:    part.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, partNode)

	addRenameButton(part, partNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete the part",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system := part.GetOwningSystem()
			if system != nil {
				system.Parts = slices.DeleteFunc(system.Parts, func(p *Part) bool { return p == part })
				system.ExternalParts = slices.DeleteFunc(system.ExternalParts, func(p *Part) bool { return p == part })
			}
			part.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if partNode.Menu == nil {
		partNode.Menu = &tree.Menu{Name: "Menu"}
	}
	partNode.Menu.Buttons = append(partNode.Menu.Buttons, deleteButton)

	partNode.OnNameChange = stager.onNameChange(part)
	partNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, &library.PartsWhoseNodeIsExpanded)
	partNode.OnClick = onNodeClicked(stager, part)
}

func (stager *Stager) treeParts(
	diagramStructure *DiagramStructure,
	part *Part,
	parentNode *tree.Node,
) {
	visited := make(map[*System]bool)
	if diagramStructure.owningSystem != nil {
		visited[diagramStructure.owningSystem] = true
	}
	stager.treePartsWithEnclosing(
		diagramStructure,
		part,
		nil,
		false,
		&diagramStructure.PartWhoseNodeIsExpanded,
		visited,
		parentNode,
	)
}

func (stager *Stager) treePartsWithEnclosing(
	diagramStructure *DiagramStructure,
	part *Part,
	parentPartShape *PartShape,
	isNested bool,
	partsWhoseNodeIsExpanded *[]*Part,
	visitedSystems map[*System]bool,
	parentNode *tree.Node,
) {
	stage := stager.stage

	isEnclosingMissing := isNested && parentPartShape == nil

	var partShape *PartShape
	var ok bool
	if !isEnclosingMissing && diagramStructure.map_PartShapeKey_PartShape != nil {
		partShape, ok = diagramStructure.map_PartShapeKey_PartShape[partShapeKey{
			part:            part,
			parentPartShape: parentPartShape,
		}]
	}

	nodeName := partDisplayLabel(part)
	if part.isInRenameMode {
		nodeName = part.Name
	}

	node := &tree.Node{
		Name:                    nodeName,
		IsExpanded:              slices.Contains(*partsWhoseNodeIsExpanded, part),
		IsNodeClickable:         true,
		IsInEditMode:            part.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		IsCheckboxDisabled:      isEnclosingMissing,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isEnclosingMissing {
				return "Enclosing part shape must be present in the diagram"
			}
			if ok {
				return "Click to remove the part shape"
			}
			return "Click to create a part shape for this part within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(part, node, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete the part",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system := part.GetOwningSystem()
			if system != nil {
				system.Parts = slices.DeleteFunc(system.Parts, func(p *Part) bool { return p == part })
				system.ExternalParts = slices.DeleteFunc(system.ExternalParts, func(p *Part) bool { return p == part })
			}
			part.Unstage(stager.stage)
			if partShape != nil {
				partShape.UnstageVoid(stager.stage)
				diagramStructure.Part_Shapes = slices.DeleteFunc(diagramStructure.Part_Shapes, func(ps *PartShape) bool { return ps == partShape })
			}
			stager.stage.Commit()
		},
	}
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}
	node.Menu.Buttons = append(node.Menu.Buttons, deleteButton)

	if ok && partShape != nil {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				partShape.SetIsHidden(!partShape.GetIsHidden())
				stage.Commit()
			},
		}
		if partShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnNameChange = stager.onNameChange(part)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, partsWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isEnclosingMissing {
			return
		}
		if isChecked {
			if partShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			partShape = newShapeToDiagram(part, diagramStructure, &diagramStructure.Part_Shapes, stager, node.ClientOnY)
			partShape.ParentPartShape = parentPartShape
			partShape.SetName(partShapeNameInDiagram(diagramStructure, part, parentPartShape))

			if parentPartShape != nil {
				depth := partShapeDepth(partShape)
				if depth < 1 {
					depth = 1
				}
				defaultW := diagramStructure.GetDefaultBoxWidth()
				if defaultW <= 0 {
					defaultW = 250.0
				}
				defaultH := diagramStructure.GetDefaultBoxHeigth()
				if defaultH <= 0 {
					defaultH = 70.0
				}

				const pad = 20.0
				const headerHeight = 40.0
				const gap = 15.0

				siblingCount := 0
				for _, ps := range diagramStructure.Part_Shapes {
					if ps != nil && ps != partShape && ps.ParentPartShape == parentPartShape {
						siblingCount++
					}
				}

				childW := math.Max(120.0, defaultW-float64(depth)*40.0)
				childH := defaultH
				partShape.X = parentPartShape.X + pad
				partShape.Y = parentPartShape.Y + headerHeight + float64(siblingCount)*(childH+gap)
				partShape.Width = childW
				partShape.Height = childH

				// Expand ancestor PartShapes as needed so nested shape fits cleanly below their top header
				childCur := partShape
				for _, anc := range partShapeAncestors(partShape) {
					reqW := (childCur.X - anc.X) + childCur.Width + pad
					reqH := (childCur.Y - anc.Y) + childCur.Height + pad
					if anc.Width < reqW {
						anc.Width = reqW
					}
					if anc.Height < reqH {
						anc.Height = reqH
					}
					childCur = anc
				}
			}

			stage.Commit()
			return
		} else {
			if partShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			partShape.UnstageVoid(stage)
			idx := slices.Index(diagramStructure.Part_Shapes, partShape)
			if idx >= 0 {
				diagramStructure.Part_Shapes = slices.Delete(diagramStructure.Part_Shapes, idx, idx+1)
			}
			stage.Commit()
			return
		}
	}
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(part, "Part")
		if partShape != nil {
			fillUpShapeAppearanceForm(stager, partShape)
		} else {
			clearShapeAppearanceForm(stager)
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}

	// Determine scoped expansion state pointers (per PartShape instance when present, fallback to Part/Diagram)
	portsExpandedBool := &part.IsPortsNodeExpanded
	portsWhoseExpanded := &diagramStructure.PortsWhoseNodeIsExpanded
	cfExpandedBool := &part.IsControlFlowsNodeExpanded
	cfWhoseExpanded := &diagramStructure.ControlFlowsWhoseNodeIsExpanded
	typePartsExpandedBool := &part.IsTypePartsNodeExpanded
	subPartsWhoseExpanded := &diagramStructure.PartWhoseNodeIsExpanded
	typeDFExpandedBool := &part.IsTypeDataFlowsNodeExpanded
	dfWhoseExpanded := &diagramStructure.DataFlowsWhoseNodeIsExpanded

	if partShape != nil {
		portsExpandedBool = &partShape.IsPortsNodeExpanded
		portsWhoseExpanded = &partShape.PortsWhoseNodeIsExpanded
		cfExpandedBool = &partShape.IsControlFlowsNodeExpanded
		cfWhoseExpanded = &partShape.ControlFlowsWhoseNodeIsExpanded
		typePartsExpandedBool = &partShape.IsTypePartsNodeExpanded
		subPartsWhoseExpanded = &partShape.SubPartsWhoseNodeIsExpanded
		typeDFExpandedBool = &partShape.IsTypeDataFlowsNodeExpanded
		dfWhoseExpanded = &partShape.DataFlowsWhoseNodeIsExpanded
	}

	// Ports
	portsNode := &tree.Node{
		Name:            "Ports",
		FontStyle:       tree.ITALIC,
		IsExpanded:      *portsExpandedBool || (!isNested && part.IsPortsNodeExpanded),
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, portsNode)
	if !isNested {
		portsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsPortsNodeExpanded)
	} else {
		portsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(portsExpandedBool)
	}

	for _, port := range part.Ports {
		stager.treePort(diagramStructure, part, partShape, port, portsNode, portsWhoseExpanded)
	}

	if partShape != nil {
		expansionTarget := portsExpandedBool
		if !isNested {
			expansionTarget = &part.IsPortsNodeExpanded
		}
		conf := ItemAndShapeButtonConfiguration[
			Port, *Port,
			Part, *Part,
			PortShape, *PortShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Port, *Port,
				Part, *Part,
			]{
				parentNode:                         portsNode,
				sliceForNewAddedItem:               &part.Ports,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    expansionTarget,
				parentElement:                      part,
			},
			receivingDiagram:      diagramStructure,
			sliceForNewAddedShape: &diagramStructure.Port_Shapes,
		}
		callback := addCreateItemAndShapeButton(stager, conf)
		capturedPartShape := partShape
		callback.OnBeforeCommit = func() {
			newShape := diagramStructure.Port_Shapes[len(diagramStructure.Port_Shapes)-1]
			newShape.ParentPartShape = capturedPartShape
			newShape.SetX(capturedPartShape.X - defaultPortWidth/2)
			newShape.SetY(capturedPartShape.Y + capturedPartShape.Height/2)
			newShape.SetWidth(defaultPortWidth)
			newShape.SetHeight(defaultPortHeight)

			port := callback.createdItem
			port.SetName(fmt.Sprintf("P%d", len(part.Ports)))
			port.SetIsInRenameMode(false)
		}
	} else {
		addCreateItemButton(stager, ItemButtonConfiguration[
			Port, *Port,
			Part, *Part,
		]{
			parentNode:                         portsNode,
			sliceForNewAddedItem:               &part.Ports,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    portsExpandedBool,
			parentElement:                      part,
		})
	}

	// ControlFlows
	controlflowsNode := &tree.Node{
		Name:            "ControlFlows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      *cfExpandedBool || (!isNested && part.IsControlFlowsNodeExpanded),
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, controlflowsNode)
	if !isNested {
		controlflowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsControlFlowsNodeExpanded)
	} else {
		controlflowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(cfExpandedBool)
	}

	for _, controlflow := range part.ControlFlows {
		stager.treeControlFlowsScoped(diagramStructure, partShape, controlflow, controlflowsNode, cfWhoseExpanded)
	}

	// Recursive Part.Type (*System) elements: Parts and Data Flows
	if part.Type != nil && !visitedSystems[part.Type] {
		childVisited := make(map[*System]bool, len(visitedSystems)+1)
		for k, v := range visitedSystems {
			childVisited[k] = v
		}
		childVisited[part.Type] = true

		typeSystem := part.Type

		// Nested Parts of part.Type
		typePartsNode := &tree.Node{
			Name:            "Parts",
			FontStyle:       tree.ITALIC,
			IsExpanded:      *typePartsExpandedBool,
			IsNodeClickable: true,
		}
		node.Children = append(node.Children, typePartsNode)
		typePartsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(typePartsExpandedBool)
		typePartsNode.OnClick = onNodeClicked(stager, typeSystem)

		for _, subPart := range typeSystem.Parts {
			stager.treePartsWithEnclosing(
				diagramStructure,
				subPart,
				partShape,
				true,
				subPartsWhoseExpanded,
				childVisited,
				typePartsNode,
			)
		}

		addCreateItemButton(stager, ItemButtonConfiguration[
			Part, *Part,
			System, *System,
		]{
			parentNode:                         typePartsNode,
			sliceForNewAddedItem:               &typeSystem.Parts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    typePartsExpandedBool,
			parentElement:                      typeSystem,
		})

		// Internal Data Flows of part.Type
		typeDataFlowsNode := &tree.Node{
			Name:            "Data Flows",
			FontStyle:       tree.ITALIC,
			IsExpanded:      *typeDFExpandedBool,
			IsNodeClickable: true,
		}
		node.Children = append(node.Children, typeDataFlowsNode)
		typeDataFlowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(typeDFExpandedBool)
		typeDataFlowsNode.OnClick = onNodeClicked(stager, typeSystem)

		for _, dataFlow := range typeSystem.DataFlows {
			stager.treeDataFlowsScoped(
				diagramStructure,
				dataFlow,
				partShape,
				true,
				dfWhoseExpanded,
				typeDataFlowsNode,
			)
		}
	}
}
