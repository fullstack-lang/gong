package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataWithinDiagramStructureWithinDataFlow(
	dataFlowNode *tree.Node,
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	dataFlowShape *DataFlowShape,
	isDataFlowShapePresent bool,
	data *Data,
) {
	stage := stager.stage

	// find the dataShape (if any)
	dataShapeKey := dataShapeKey{
		dataFlow: dataFlow,
		data:     data,
	}
	dataShape, isDataShapePresent := diagramStructure.map_DataShapeKey_DataShape[dataShapeKey]

	dataNode := &tree.Node{
		Name:                    data.GetName(),
		IsNodeClickable:         true,
		IsInEditMode:            data.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isDataShapePresent,
		IsCheckboxDisabled:      !isDataFlowShapePresent,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if !isDataFlowShapePresent {
				return "A data shape cannot be created if the data flow shape is not present"
			}
			if isDataShapePresent {
				return "Click to remove the data shape"
			}
			return "Click to create a data shape for this data within this data flow"
		}(),
	}
	dataFlowNode.Children = append(dataFlowNode.Children, dataNode)
	addRenameButton(data, dataNode, stager)

	if isDataShapePresent && dataShape != nil {
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				dataShape.SetIsHidden(!dataShape.GetIsHidden())
				stage.Commit()
			},
		}
		if dataShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		dataNode.Buttons = append(dataNode.Buttons, visibilityButton)
	}

	dataNode.OnNameChange = stager.onNameChange(data)

	dataNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isDataShapePresent {
			dataShape := (&DataShape{
				Name:     dataFlow.Name + "-" + dataFlowShape.GetName() + "-" + diagramStructure.GetName(),
				Data:     data,
				DataFlow: dataFlow,
			}).Stage(stager.stage)
			dataShape.SetWidth(180.0)
			diagramStructure.Data_Shapes = append(diagramStructure.Data_Shapes, dataShape)
			stage.Commit()
			return
		}
		if !isChecked && isDataShapePresent {
			dataShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}
	dataNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(data, "Data")
		if isDataShapePresent && dataShape != nil {
			fillUpShapeAppearanceForm(stager, dataShape)
		} else {
			clearShapeAppearanceForm(stager)
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}
}

