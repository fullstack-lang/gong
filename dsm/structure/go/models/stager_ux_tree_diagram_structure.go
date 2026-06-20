package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramStructure(
	system *System,
	diagram *DiagramStructure,
	parentNode *tree.Node,
	expandedSlice *[]*DiagramStructure,
) {
	diagNode := &tree.Node{
		Name:              diagram.Name,
		IsExpanded:        diagram.IsExpanded,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsInEditMode:      diagram.isInRenameMode,
		IsChecked:         diagram.IsChecked,
	}
	if diagNode.Name == "" {
		diagNode.Name = "Diagram"
	}
	parentNode.Children = append(parentNode.Children, diagNode)

	addRenameButton(diagram, diagNode, stager)

	diagCopy := diagram
	diagNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			// uncheck all diagrams
			for d := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stager.stage) {
				d.IsChecked = false
			}
			// check this diagram
			diagCopy.IsChecked = true
			stager.diagram = diagCopy
			stager.stage.Commit()
		} else {
			diagCopy.IsChecked = false
			stager.diagram = nil
			stager.stage.Commit()
		}
	}

	diagNode.OnClick = onNodeClicked(stager, diagram)
	diagNode.OnNameChange = stager.onNameChange(diagram)
	diagNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)

	// prefix button
	{
		showPrefixButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagram.IsShowPrefix = !diagram.IsShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagram.IsShowPrefix {
			showPrefixButton.Icon = string(buttons.BUTTON_label)
			showPrefixButton.ToolTipText = "Show Prefix"
		} else {
			showPrefixButton.Icon = string(buttons.BUTTON_label_off)
			showPrefixButton.ToolTipText = "Hide Prefix"
		}
		diagNode.Buttons = append(diagNode.Buttons, showPrefixButton)
	}

	// Parts
	partsNode := &tree.Node{
		Name:            "Parts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsPartsNodeExpanded,
		IsNodeClickable: true,
	}
	diagNode.Children = append(diagNode.Children, partsNode)
	partsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsPartsNodeExpanded)

	for _, part := range system.Parts {
		stager.treePartInDiagram(diagram, part, partsNode)
	}

	// Links
	linksNode := &tree.Node{
		Name:            "Links",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsLinksNodeExpanded,
		IsNodeClickable: true,
	}
	diagNode.Children = append(diagNode.Children, linksNode)
	linksNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsLinksNodeExpanded)

	for _, link := range system.Links {
		stager.treeLinkInDiagram(diagram, link, linksNode)
	}
}

func (stager *Stager) treePartInDiagram(
	diagram *DiagramStructure,
	part *Part,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagram.map_Part_PartShape[part]
	node := &tree.Node{
		Name:                    part.Name,
		IsExpanded:              slices.Contains(diagram.PartsWhoseNodeIsExpanded, part),
		IsNodeClickable:         true,
		IsInEditMode:            part.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the part shape"
			}
			return "Click to create a part shape for this part within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(part, node, stager)

	if shape, ok := diagram.map_Part_PartShape[part]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagram.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnNameChange = stager.onNameChange(part)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, &diagram.PartsWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already present shape")
			}
			shape = newShapeToDiagram(part, diagram, &diagram.Part_Shapes, stage)
			stage.Commit()
			return
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape from diagram")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagram.Part_Shapes, shape)
			diagram.Part_Shapes = slices.Delete(diagram.Part_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, part)
}

func (stager *Stager) treeLinkInDiagram(
	diagram *DiagramStructure,
	link *Link,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the linkShape (if any)
	linkShape, isShapePresent := diagram.map_Link_LinkAssociationShape[link]

	isStartShapePresent := false
	isEndShapePresent := false

	var startName, endName string
	var startElement, endElement AbstractType

	if link.Source != nil {
		_, isStartShapePresent = diagram.map_Part_PartShape[link.Source]
		startName = link.Source.GetName()
		startElement = link.Source
	}
	if link.Target != nil {
		_, isEndShapePresent = diagram.map_Part_PartShape[link.Target]
		endName = link.Target.GetName()
		endElement = link.Target
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	linkNode := &tree.Node{
		Name:                    link.GetName(),
		IsExpanded:              slices.Contains(diagram.LinksWhoseNodeIsExpanded, link),
		IsNodeClickable:         true,
		IsInEditMode:            link.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "A link shape cannot be created if the start or end shape is not present from the diagram"
			}
			if isShapePresent {
				return "Click to remove the link shape"
			}
			return "Click to create a link shape for this link within this diagram"
		}(),
	}

	if isCheckboxDisabled {
		linkNode.CheckboxHasToolTip = true
		linkNode.CheckboxToolTipText = "Start or end shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, linkNode)

	addRenameButton(link, linkNode, stager)

	if linkShape, ok := diagram.map_Link_LinkAssociationShape[link]; ok {
		linkNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagram.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				linkShape.SetIsHidden(!linkShape.GetIsHidden())
				stage.Commit()
			},
		}
		if linkShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		linkNode.Buttons = append(linkNode.Buttons, visibilityButton)
	}

	linkNode.OnNameChange = stager.onNameChange(link)
	linkNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, link, &diagram.LinksWhoseNodeIsExpanded)
	linkNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if linkShape != nil {
				log.Panic("adding a shape to an already present shape")
			}
			linkShape := (&LinkAssociationShape{
				Link: link,
			}).Stage(stage)
			linkShape.SetName(startName + " to " + endName)
			linkShape.SetAbstractStartElement(startElement)
			linkShape.SetAbstractEndElement(endElement)
			linkShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
			linkShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

			linkShape.SetCornerOffsetRatio(1.1)
			linkShape.SetStartRatio(0.5)
			linkShape.SetEndRatio(0.5)
			diagram.Link_Shapes = append(diagram.Link_Shapes, linkShape)

			stage.Commit()
			return
		}
		if !isChecked && isShapePresent {
			isShapePresent = isChecked
			if linkShape == nil {
				log.Panic("remove a non existing shape from diagram")
			}
			linkShape.UnstageVoid(stage)

			idx := slices.Index(diagram.Link_Shapes, linkShape)
			diagram.Link_Shapes = slices.Delete(diagram.Link_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	linkNode.OnClick = onNodeClicked(stager, link)
}
