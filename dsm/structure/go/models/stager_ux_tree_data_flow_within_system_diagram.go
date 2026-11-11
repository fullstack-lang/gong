package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowsWithinSystemDiagram(
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stager.treeDataFlowsScoped(
		diagramStructure,
		dataFlow,
		nil,
		false,
		&diagramStructure.DataFlowsWhoseNodeIsExpanded,
		parentNode,
	)
}

func (stager *Stager) treeDataFlowsScoped(
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	enclosingPartShape *PartShape,
	isNested bool,
	dataFlowsWhoseNodeIsExpanded *[]*DataFlow,
	parentNode *tree.Node,
) {
	stage := stager.stage

	isEnclosingMissing := isNested && enclosingPartShape == nil

	var dataFlowShape *DataFlowShape
	var isShapePresent bool
	if !isEnclosingMissing && diagramStructure.map_DataFlowShapeKey_DataFlowShape != nil {
		dataFlowShape, isShapePresent = diagramStructure.map_DataFlowShapeKey_DataFlowShape[dataFlowShapeKey{
			dataFlow:        dataFlow,
			parentPartShape: enclosingPartShape,
		}]
	}

	isStartShapePresent := false
	isEndShapePresent := false
	var resolvedStartPortShape, resolvedEndPortShape *PortShape

	var startName, endName string
	var startElement, endElement AbstractType

	if !isEnclosingMissing {
		switch dataFlow.Type {
		case DataFlow_Port2Port:
			if dataFlow.StartPort != nil {
				resolvedStartPortShape, isStartShapePresent = resolvePortShapeInScope(diagramStructure, dataFlow.StartPort, enclosingPartShape)
				startName = dataFlow.StartPort.GetName()
				startElement = dataFlow.StartPort
			}
			if dataFlow.EndPort != nil {
				resolvedEndPortShape, isEndShapePresent = resolvePortShapeInScope(diagramStructure, dataFlow.EndPort, enclosingPartShape)
				endName = dataFlow.EndPort.GetName()
				endElement = dataFlow.EndPort
			}
		case DataFlow_ExternalPart2Port:
			if dataFlow.StartExternalPart != nil {
				_, isStartShapePresent = diagramStructure.map_Part_ExternalPartShape[dataFlow.StartExternalPart]
				startName = dataFlow.StartExternalPart.GetName()
				startElement = dataFlow.StartExternalPart
			}
			if dataFlow.EndPort != nil {
				resolvedEndPortShape, isEndShapePresent = resolvePortShapeInScope(diagramStructure, dataFlow.EndPort, enclosingPartShape)
				endName = dataFlow.EndPort.GetName()
				endElement = dataFlow.EndPort
			}
		case DataFlow_Port2ExternalPart:
			if dataFlow.StartPort != nil {
				resolvedStartPortShape, isStartShapePresent = resolvePortShapeInScope(diagramStructure, dataFlow.StartPort, enclosingPartShape)
				startName = dataFlow.StartPort.GetName()
				startElement = dataFlow.StartPort
			}
			if dataFlow.EndExternalPart != nil {
				_, isEndShapePresent = diagramStructure.map_Part_ExternalPartShape[dataFlow.EndExternalPart]
				endName = dataFlow.EndExternalPart.GetName()
				endElement = dataFlow.EndExternalPart
			}
		}
	}
	isCheckboxDisabled := isEnclosingMissing || !(isStartShapePresent && isEndShapePresent)

	dataFlowNode := &tree.Node{
		Name:                    dataFlow.GetName(),
		IsExpanded:              slices.Contains(*dataFlowsWhoseNodeIsExpanded, dataFlow),
		IsNodeClickable:         true,
		IsInEditMode:            dataFlow.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "Start or end shape is missing from the diagram"
			}
			if isShapePresent {
				return "Click to remove the data flow shape"
			}
			return "Click to create a data flow shape for this data flow within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	if isShapePresent && dataFlowShape != nil {
		dataFlowNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				dataFlowShape.SetIsHidden(!dataFlowShape.GetIsHidden())
				stage.Commit()
			},
		}
		if dataFlowShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		dataFlowNode.Buttons = append(dataFlowNode.Buttons, visibilityButton)
	}

	dataFlowNode.OnNameChange = stager.onNameChange(dataFlow)
	dataFlowNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataFlow, dataFlowsWhoseNodeIsExpanded)
	dataFlowNode.OnIsCheckedChanged = func(isChecked bool) {
		if isEnclosingMissing {
			return
		}
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			newDFShape := (&DataFlowShape{
				DataFlow:        dataFlow,
				ParentPartShape: enclosingPartShape,
				StartPortShape:  resolvedStartPortShape,
				EndPortShape:    resolvedEndPortShape,
			}).Stage(stage)
			newDFShape.SetName(startName + " to " + endName)
			newDFShape.SetAbstractStartElement(startElement)
			newDFShape.SetAbstractEndElement(endElement)
			newDFShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
			newDFShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

			newDFShape.SetCornerOffsetRatio(1.1)
			newDFShape.SetStartRatio(0.5)
			newDFShape.SetEndRatio(0.5)
			diagramStructure.DataFlow_Shapes = append(diagramStructure.DataFlow_Shapes, newDFShape)

			stage.Commit()
			return
		}
		if !isChecked && isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			dataFlowShape.UnstageVoid(stage)
			idx := slices.Index(diagramStructure.DataFlow_Shapes, dataFlowShape)
			if idx >= 0 {
				diagramStructure.DataFlow_Shapes = slices.Delete(diagramStructure.DataFlow_Shapes, idx, idx+1)
			}
			stage.Commit()
			return
		}
	}
	dataFlowNode.OnClick = onNodeClicked(stager, dataFlow)

	for _, data := range dataFlow.Datas {
		stager.treeDataWithinDiagramStructureWithinDataFlow(dataFlowNode, diagramStructure, dataFlow, dataFlowShape, isShapePresent, data)
	}
	if dataFlow.Protocol != nil {
		stager.treeProtocolWithinDiagramStructureWithinDataFlow(dataFlowNode, diagramStructure, dataFlow, dataFlowShape, isShapePresent, dataFlow.Protocol)
	}
}
