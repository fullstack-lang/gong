package models

import (
	"math/rand/v2"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func newDataStoreShapeOnDiagram(
	stager *Stager,
	diagramProcess *DiagramProcess,
	dataStore *DataStore,
	clientOnY float64,
) *DataStoreShape {
	shape := (&DataStoreShape{
		Name:      dataStore.GetName() + "-" + diagramProcess.GetName(),
		DataStore: dataStore,
	}).Stage(stager.stage)
	shape.SetWidth(DataStoreDefaultWidth)
	shape.SetHeight(DataStoreDefaultHeight)

	zoom := 1.0
	panX := 0.0
	panY := 0.0
	if svgObj := stager.GetSvgObject(); svgObj != nil {
		if svgObj.Zoom != 0 {
			zoom = svgObj.Zoom
		}
		panX = svgObj.PanX
		panY = svgObj.PanY
	}

	shape.SetX((100+panX)/zoom + rand.Float64()*100.0)
	if clientOnY != 0 {
		shape.SetY((clientOnY + panY) / zoom)
	} else {
		shape.SetY((100+panY)/zoom + rand.Float64()*100.0)
	}

	diagramProcess.DataStore_Shapes = append(diagramProcess.DataStore_Shapes, shape)
	if diagramProcess.map_DataStore_DataStoreShape != nil {
		diagramProcess.map_DataStore_DataStoreShape[dataStore] = shape
	}
	return shape
}

func (stager *Stager) treeDataStoreWithinDiagramProcess(
	diagramProcess *DiagramProcess,
	dataStore *DataStore,
	parentNode *tree.Node,
) {
	stage := stager.stage
	shape, ok := diagramProcess.map_DataStore_DataStoreShape[dataStore]

	node := &tree.Node{
		Name:                    dataStore.GetName(),
		IsExpanded:              slices.Contains(diagramProcess.DataStoresWhoseNodeIsExpanded, dataStore),
		IsNodeClickable:         true,
		IsInEditMode:            dataStore.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the data store shape"
			}
			return "Click to create a data store shape for this data store within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(dataStore, node, stager)

	if ok && shape != nil {
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
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

	node.OnNameChange = stager.onNameChange(dataStore)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataStore, &diagramProcess.DataStoresWhoseNodeIsExpanded)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(dataStore, "DataStore")
		if ok && shape != nil {
			fillUpShapeAppearanceForm(stager, shape)
		} else {
			clearShapeAppearanceForm(stager)
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			newDataStoreShapeOnDiagram(stager, diagramProcess, dataStore, node.ClientOnY)
			stage.Commit()
			return
		}
		if !isChecked && ok && shape != nil {
			idx := slices.Index(diagramProcess.DataStore_Shapes, shape)
			if idx != -1 {
				diagramProcess.DataStore_Shapes = slices.Delete(diagramProcess.DataStore_Shapes, idx, idx+1)
			}
			shape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}

	// Data stored in the DataStore
	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      dataStore.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, datasNode)
	datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&dataStore.IsDatasNodeExpanded)
	datasNode.OnClick = onNodeClicked(stager, dataStore)

	owningLib := dataStore.GetOwningLibrary()
	if owningLib == nil {
		owningLib = stager.GetRootLibrary()
	}
	for _, data := range topLevelDatasInSlice(dataStore.Datas) {
		stager.treeData(owningLib, data, datasNode)
	}

	// out data flows and in data flows
	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramProcess.DataStoresWhoseOutDataFlowsNodeIsExpanded, dataStore),
	}
	node.Children = append(node.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataStore, &diagramProcess.DataStoresWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range dataStore.outDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramProcess.DataStoresWhoseInDataFlowsNodeIsExpanded, dataStore),
	}
	node.Children = append(node.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataStore, &diagramProcess.DataStoresWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range dataStore.inDataFlows {
		stager.treeDataFlowsWithinDiagramProcessWithinTask(diagramProcess, dataFlow, nodeInDataFlows)
	}
}
