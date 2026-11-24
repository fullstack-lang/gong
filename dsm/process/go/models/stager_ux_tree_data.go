package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataWithinLibrary(
	library *Library,
	data *Data,
	parentNode *tree.Node,
) {
	stager.treeDataWithinLibraryVisited(library, data, parentNode, make(map[*Data]bool))
}

func (stager *Stager) treeDataWithinLibraryVisited(
	library *Library,
	data *Data,
	parentNode *tree.Node,
	visited map[*Data]bool,
) {
	if data == nil || visited[data] {
		return
	}
	visited[data] = true
	defer delete(visited, data)

	dataNode := &tree.Node{
		Name:            data.GetName(),
		IsExpanded:      slices.Contains(library.DatasWhoseNodeIsExpanded, data),
		IsNodeClickable: true,
		IsInEditMode:    data.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, dataNode)

	addRenameButton(data, dataNode, stager)

	dataNode.OnNameChange = stager.onNameChange(data)
	dataNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, data, &library.DatasWhoseNodeIsExpanded)
	dataNode.OnClick = onNodeClicked(stager, data)

	confSubDatas := ItemButtonConfiguration[
		Data, *Data,
		Data, *Data,
	]{
		parentNode:                         dataNode,
		sliceForNewAddedItem:               &data.SubDatas,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
		parentNodeExpansionSliceEncoding:   &library.DatasWhoseNodeIsExpanded,
		parentElement:                      data,
		IsButtonInMenu:                     true,
	}
	callbacksSubDatas := addCreateItemButton(stager, confSubDatas)
	callbacksSubDatas.button.Name = "Add SubData"
	callbacksSubDatas.button.ToolTipText = "Add a SubData to \"" + data.GetName() + "\""
	callbacksSubDatas.OnBeforeCommit = func() {
		data.IsSubDatasNodeExpanded = true
		if callbacksSubDatas.createdItem != nil {
			callbacksSubDatas.createdItem.SetOwningLibrary(library)
		}
	}
	EnsureRenameIsFirst(dataNode)

	// Sub Datas category node (only if items are present)
	if len(data.SubDatas) > 0 {
		subDatasNode := &tree.Node{
			Name:            "Sub Datas",
			FontStyle:       tree.ITALIC,
			IsExpanded:      data.IsSubDatasNodeExpanded,
			IsNodeClickable: true,
		}
		dataNode.Children = append(dataNode.Children, subDatasNode)
		subDatasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&data.IsSubDatasNodeExpanded)
		subDatasNode.OnClick = onNodeClicked(stager, data)

		confSubDatasNode := ItemButtonConfiguration[
			Data, *Data,
			Data, *Data,
		]{
			parentNode:                         subDatasNode,
			sliceForNewAddedItem:               &data.SubDatas,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &data.IsSubDatasNodeExpanded,
			parentElement:                      data,
			IsButtonInMenu:                     true,
		}
		callbacksSubDatasNode := addCreateItemButton(stager, confSubDatasNode)
		callbacksSubDatasNode.button.Name = "Add SubData"
		callbacksSubDatasNode.button.ToolTipText = "Add a SubData to \"" + data.GetName() + "\""
		callbacksSubDatasNode.OnBeforeCommit = func() {
			data.IsSubDatasNodeExpanded = true
			if callbacksSubDatasNode.createdItem != nil {
				callbacksSubDatasNode.createdItem.SetOwningLibrary(library)
			}
		}

		for _, subData := range data.SubDatas {
			stager.treeDataWithinLibraryVisited(library, subData, subDatasNode, visited)
		}
	}

	for _, dataFlow := range stager.rm_Data_DataFlows[data] {
		stager.treeDataFlowWithinLibrary(library, dataFlow, dataNode)
	}
}

func (stager *Stager) treeData(
	library *Library,
	data *Data,
	parentNode *tree.Node,
) {
	stager.treeDataVisited(library, data, parentNode, make(map[*Data]bool))
}

func (stager *Stager) treeDataVisited(
	library *Library,
	data *Data,
	parentNode *tree.Node,
	visited map[*Data]bool,
) {
	if data == nil || visited[data] {
		return
	}
	visited[data] = true
	defer delete(visited, data)

	dataNode := &tree.Node{
		Name:            data.GetName(),
		IsExpanded:      slices.Contains(library.DatasWhoseNodeIsExpanded, data),
		IsNodeClickable: true,
		IsInEditMode:    data.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, dataNode)

	addRenameButton(data, dataNode, stager)

	dataNode.OnNameChange = stager.onNameChange(data)
	dataNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, data, &library.DatasWhoseNodeIsExpanded)
	dataNode.OnClick = onNodeClicked(stager, data)

	for _, subData := range data.SubDatas {
		stager.treeDataVisited(library, subData, dataNode, visited)
	}
}
