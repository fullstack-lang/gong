package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataWithinDiagramProcessWithinControlFlow(
	controlFlowNode *tree.Node,
	diagramProcess *DiagramProcess,
	controlFlow *ControlFlow,
	controlFlowShape *ControlFlowShape,
	isControlFlowShapePresent bool,
	data *Data,
) {
	stager.treeDataWithinDiagramProcessWithinControlFlowVisited(
		controlFlowNode,
		diagramProcess,
		controlFlow,
		controlFlowShape,
		isControlFlowShapePresent,
		data,
		make(map[*Data]bool),
	)
}

func (stager *Stager) treeDataWithinDiagramProcessWithinControlFlowVisited(
	parentNode *tree.Node,
	diagramProcess *DiagramProcess,
	controlFlow *ControlFlow,
	controlFlowShape *ControlFlowShape,
	isControlFlowShapePresent bool,
	data *Data,
	visited map[*Data]bool,
) {
	if data == nil || visited[data] {
		return
	}
	visited[data] = true
	defer delete(visited, data)

	stage := stager.stage

	key := dataShapeKey{
		controlFlow: controlFlow,
		data:        data,
	}
	var dataShape *DataShape
	var isDataShapePresent bool
	if isControlFlowShapePresent && diagramProcess.map_DataShapeKey_DataShape != nil {
		dataShape, isDataShapePresent = diagramProcess.map_DataShapeKey_DataShape[key]
	}

	dataNode := &tree.Node{
		Name:                    data.GetName(),
		IsExpanded:              slices.Contains(diagramProcess.DatasWhoseNodeIsExpanded, data),
		IsNodeClickable:         true,
		IsInEditMode:            data.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isDataShapePresent,
		IsCheckboxDisabled:      !isControlFlowShapePresent,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if !isControlFlowShapePresent {
				return "A data shape cannot be created if the control flow shape is not present"
			}
			if isDataShapePresent {
				return "Click to remove the data shape"
			}
			return "Click to create a data shape for this data within this control flow"
		}(),
	}
	setDataNodeAcronymPrefix(data, dataNode, diagramProcess)
	setDataNodeAcronymPrefix(data, dataNode, diagramProcess)
	parentNode.Children = append(parentNode.Children, dataNode)
	addRenameButton(data, dataNode, stager)

	if isControlFlowShapePresent && len(data.SubDatas) > 0 {
		if dataNode.Menu == nil {
			dataNode.Menu = &tree.Menu{Name: "Menu"}
		}
		displayAllBtn := &tree.Button{
			Name:            "Display data & sub-datas",
			Icon:            string(buttons.BUTTON_add_circle),
			ToolTipText:     "Display this data and all its sub-datas on this control flow",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				cfName := controlFlow.GetName()
				if controlFlowShape != nil && controlFlowShape.GetName() != "" {
					cfName = controlFlowShape.GetName()
				}
				walkDatasPreOrder([]*Data{data}, func(d *Data) {
					k := dataShapeKey{controlFlow: controlFlow, data: d}
					if _, exists := diagramProcess.map_DataShapeKey_DataShape[k]; !exists {
						ds := (&DataShape{
							Name:        d.GetName() + "-" + cfName + "-" + diagramProcess.GetName(),
							Data:        d,
							ControlFlow: controlFlow,
						}).Stage(stage)
						diagramProcess.Data_Shapes = append(diagramProcess.Data_Shapes, ds)
						diagramProcess.map_DataShapeKey_DataShape[k] = ds
					}
					if len(d.SubDatas) > 0 && !slices.Contains(diagramProcess.DatasWhoseNodeIsExpanded, d) {
						diagramProcess.DatasWhoseNodeIsExpanded = append(diagramProcess.DatasWhoseNodeIsExpanded, d)
					}
				})
				stage.Commit()
			},
		}
		removeAllBtn := &tree.Button{
			Name:            "Remove data & sub-datas",
			Icon:            string(buttons.BUTTON_remove_circle),
			ToolTipText:     "Remove this data and all its sub-datas from this control flow",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				walkDatasPreOrder([]*Data{data}, func(d *Data) {
					k := dataShapeKey{controlFlow: controlFlow, data: d}
					if ds, exists := diagramProcess.map_DataShapeKey_DataShape[k]; exists && ds != nil {
						ds.UnstageVoid(stage)
					}
				})
				stage.Commit()
			},
		}
		dataNode.Menu.Buttons = append(dataNode.Menu.Buttons, displayAllBtn, removeAllBtn)
		EnsureRenameIsFirst(dataNode)
	}

	if isDataShapePresent && dataShape != nil {
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
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
	dataNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, data, &diagramProcess.DatasWhoseNodeIsExpanded)

	dataNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isDataShapePresent {
			cfName := controlFlow.GetName()
			if controlFlowShape != nil && controlFlowShape.GetName() != "" {
				cfName = controlFlowShape.GetName()
			}
			newDataShape := (&DataShape{
				Name:        data.GetName() + "-" + cfName + "-" + diagramProcess.GetName(),
				Data:        data,
				ControlFlow: controlFlow,
			}).Stage(stager.stage)
			diagramProcess.Data_Shapes = append(diagramProcess.Data_Shapes, newDataShape)
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

	for _, subData := range data.SubDatas {
		stager.treeDataWithinDiagramProcessWithinControlFlowVisited(
			dataNode,
			diagramProcess,
			controlFlow,
			controlFlowShape,
			isControlFlowShapePresent,
			subData,
			visited,
		)
	}
}
