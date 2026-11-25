package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func setDataStoreNodeAcronymPrefix(dataStore *DataStore, node *tree.Node, diagramProcess *DiagramProcess) {
	if dataStore == nil || node == nil {
		return
	}
	prefix := ""
	if diagramProcess != nil && diagramProcess.GetIsShowPrefix() && dataStore.GetComputedPrefix() != "" {
		prefix = dataStore.GetComputedPrefix()
	}
	if dataStore.Acronym != "" {
		if prefix != "" {
			prefix = prefix + " " + dataStore.Acronym + " /"
		} else {
			prefix = dataStore.Acronym + " /"
		}
	}
	if prefix != "" {
		node.IsWithPrefix = true
		node.Prefix = prefix
	}
}

func (stager *Stager) treeDataStoreWithinLibrary(
	library *Library,
	dataStore *DataStore,
	parentNode *tree.Node,
) {
	dataStoreNode := &tree.Node{
		Name:            dataStore.GetName(),
		IsExpanded:      slices.Contains(library.DataStoresWhoseNodeIsExpanded, dataStore),
		IsNodeClickable: true,
		IsInEditMode:    dataStore.GetIsInRenameMode(),
	}
	setDataStoreNodeAcronymPrefix(dataStore, dataStoreNode, nil)
	parentNode.Children = append(parentNode.Children, dataStoreNode)

	addRenameButton(dataStore, dataStoreNode, stager)

	dataStoreNode.OnNameChange = stager.onNameChange(dataStore)
	dataStoreNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataStore, &library.DataStoresWhoseNodeIsExpanded)
	dataStoreNode.OnClick = onNodeClicked(stager, dataStore)

	// Data stored in the DataStore
	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      dataStore.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	dataStoreNode.Children = append(dataStoreNode.Children, datasNode)
	datasNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&dataStore.IsDatasNodeExpanded)
	datasNode.OnClick = onNodeClicked(stager, dataStore)

	for _, data := range topLevelDatasInSlice(dataStore.Datas) {
		stager.treeData(library, data, datasNode)
	}
}
