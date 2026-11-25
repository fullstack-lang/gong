package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

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
