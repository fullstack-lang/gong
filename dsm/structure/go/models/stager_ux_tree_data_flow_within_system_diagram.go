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
	stage := stager.stage

	// find the dataFlowShape (if any)
	dataFlowShape, isShapePresent := diagramStructure.map_DataFlow_DataFlowShape[dataFlow]

	isStartShapePresent := false
	isEndShapePresent := false

	var startName, endName string
	var startElement, endElement AbstractType

	switch dataFlow.Type {
	case DataFlow_Port2Port:
		if dataFlow.StartPort != nil {
			_, isStartShapePresent = diagramStructure.map_Port_PortShape[dataFlow.StartPort]
			startName = dataFlow.StartPort.GetName()
			startElement = dataFlow.StartPort
		}
		if dataFlow.EndPort != nil {
			_, isEndShapePresent = diagramStructure.map_Port_PortShape[dataFlow.EndPort]
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
			_, isEndShapePresent = diagramStructure.map_Port_PortShape[dataFlow.EndPort]
			endName = dataFlow.EndPort.GetName()
			endElement = dataFlow.EndPort
		}
	case DataFlow_Port2ExternalPart:
		if dataFlow.StartPort != nil {
			_, isStartShapePresent = diagramStructure.map_Port_PortShape[dataFlow.StartPort]
			startName = dataFlow.StartPort.GetName()
			startElement = dataFlow.StartPort
		}
		if dataFlow.EndExternalPart != nil {
			_, isEndShapePresent = diagramStructure.map_Part_ExternalPartShape[dataFlow.EndExternalPart]
			endName = dataFlow.EndExternalPart.GetName()
			endElement = dataFlow.EndExternalPart
		}
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	dataFlowNode := &tree.Node{
		Name:                    dataFlow.GetName(),
		IsExpanded:              slices.Contains(diagramStructure.DataFlowsWhoseNodeIsExpanded, dataFlow),
		IsNodeClickable:         true,
		IsInEditMode:            dataFlow.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "A data flow cannot be created if the start or end shape is not present from the diagram"
			}
			if isShapePresent {
				return "Click to remove the data flow shape"
			}
			return "Click to create a data flow shape for this data flow within this diagram"
		}(),
	}

	if isCheckboxDisabled {
		dataFlowNode.CheckboxHasToolTip = true
		dataFlowNode.CheckboxToolTipText = "Start or end shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	if shape, ok := diagramStructure.map_DataFlow_DataFlowShape[dataFlow]; ok {
		dataFlowNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
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
		dataFlowNode.Buttons = append(dataFlowNode.Buttons, visibilityButton)
	}

	dataFlowNode.OnNameChange = stager.onNameChange(dataFlow)
	dataFlowNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataFlow, &diagramStructure.DataFlowsWhoseNodeIsExpanded)
	dataFlowNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			// shape = newShapeToDiagram(controlflow, diagramStructure, &diagramStructure.ControlFlowShapes, stage)
			// addAssociationShapeToDiagram(stager, controlflow.Start, controlflow.End, &diagramStructure.ControlFlowShapes)
			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stage)
			dataFlowShape.SetName(startName + " to " + endName)
			dataFlowShape.SetAbstractStartElement(startElement)
			dataFlowShape.SetAbstractEndElement(endElement)
			dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

			dataFlowShape.SetCornerOffsetRatio(1.1)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramStructure.DataFlow_Shapes = append(diagramStructure.DataFlow_Shapes, dataFlowShape)

			stage.Commit()
			return
		}
		if !isChecked && isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			dataFlowShape.UnstageVoid(stage)

			// not necessary since there is a semantic rule (gong clean) that remove the shape from the slice when it is unstaged
			idx := slices.Index(diagramStructure.DataFlow_Shapes, dataFlowShape)
			diagramStructure.DataFlow_Shapes = slices.Delete(diagramStructure.DataFlow_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	dataFlowNode.OnClick = onNodeClicked(stager, dataFlow)

	for _, data := range dataFlow.Datas {
		stager.treeDataWithinDiagramStructureWithinDataFlow(dataFlowNode, diagramStructure, dataFlow, dataFlowShape, isShapePresent, data)
	}
}
