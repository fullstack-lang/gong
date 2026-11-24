package models

import (
	"slices"

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
	stager.treeDataWithinDiagramStructureWithinDataFlowVisited(
		dataFlowNode,
		diagramStructure,
		dataFlow,
		dataFlowShape,
		isDataFlowShapePresent,
		data,
		make(map[*Data]bool),
	)
}

func (stager *Stager) treeDataWithinDiagramStructureWithinDataFlowVisited(
	parentNode *tree.Node,
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	dataFlowShape *DataFlowShape,
	isDataFlowShapePresent bool,
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
		dataFlow:      dataFlow,
		dataFlowShape: dataFlowShape,
		data:          data,
	}
	var dataShape *DataShape
	var isDataShapePresent bool
	if isDataFlowShapePresent && dataFlowShape != nil && diagramStructure.map_DataShapeKey_DataShape != nil {
		dataShape, isDataShapePresent = diagramStructure.map_DataShapeKey_DataShape[key]
	}

	dataNode := &tree.Node{
		Name:                    data.GetName(),
		IsExpanded:              slices.Contains(diagramStructure.DatasWhoseNodeIsExpanded, data),
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
	parentNode.Children = append(parentNode.Children, dataNode)
	addRenameButton(data, dataNode, stager)

	if isDataFlowShapePresent && dataFlowShape != nil && len(data.SubDatas) > 0 {
		if dataNode.Menu == nil {
			dataNode.Menu = &tree.Menu{Name: "Menu"}
		}
		displayAllBtn := &tree.Button{
			Name:            "Display data & sub-datas",
			Icon:            string(buttons.BUTTON_add_circle),
			ToolTipText:     "Display this data and all its sub-datas on this data flow",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				dfName := dataFlowShape.GetName()
				walkDatasPreOrder([]*Data{data}, func(d *Data) {
					k := dataShapeKey{dataFlow: dataFlow, dataFlowShape: dataFlowShape, data: d}
					if _, exists := diagramStructure.map_DataShapeKey_DataShape[k]; !exists {
						ds := (&DataShape{
							Name:          d.GetName() + "-" + dfName + "-" + diagramStructure.GetName(),
							Data:          d,
							DataFlow:      dataFlow,
							DataFlowShape: dataFlowShape,
						}).Stage(stage)
						diagramStructure.Data_Shapes = append(diagramStructure.Data_Shapes, ds)
						diagramStructure.map_DataShapeKey_DataShape[k] = ds
					}
					if len(d.SubDatas) > 0 && !slices.Contains(diagramStructure.DatasWhoseNodeIsExpanded, d) {
						diagramStructure.DatasWhoseNodeIsExpanded = append(diagramStructure.DatasWhoseNodeIsExpanded, d)
					}
				})
				stage.Commit()
			},
		}
		removeAllBtn := &tree.Button{
			Name:            "Remove data & sub-datas",
			Icon:            string(buttons.BUTTON_remove_circle),
			ToolTipText:     "Remove this data and all its sub-datas from this data flow",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				walkDatasPreOrder([]*Data{data}, func(d *Data) {
					k := dataShapeKey{dataFlow: dataFlow, dataFlowShape: dataFlowShape, data: d}
					if ds, exists := diagramStructure.map_DataShapeKey_DataShape[k]; exists && ds != nil {
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
	dataNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, data, &diagramStructure.DatasWhoseNodeIsExpanded)

	dataNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isDataShapePresent {
			dfName := dataFlow.GetName()
			if dataFlowShape != nil {
				dfName = dataFlowShape.GetName()
			}
			newDataShape := (&DataShape{
				Name:          data.GetName() + "-" + dfName + "-" + diagramStructure.GetName(),
				Data:          data,
				DataFlow:      dataFlow,
				DataFlowShape: dataFlowShape,
			}).Stage(stager.stage)
			diagramStructure.Data_Shapes = append(diagramStructure.Data_Shapes, newDataShape)
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
		stager.treeDataWithinDiagramStructureWithinDataFlowVisited(
			dataNode,
			diagramStructure,
			dataFlow,
			dataFlowShape,
			isDataFlowShapePresent,
			subData,
			visited,
		)
	}
}
