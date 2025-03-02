// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ButtonAPIs []*ButtonAPI

	NodeAPIs []*NodeAPI

	SVGIconAPIs []*SVGIconAPI

	TreeAPIs []*TreeAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, buttonDB := range backRepo.BackRepoButton.Map_ButtonDBID_ButtonDB {

		var buttonAPI ButtonAPI
		buttonAPI.ID = buttonDB.ID
		buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
		buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)

		backRepoData.ButtonAPIs = append(backRepoData.ButtonAPIs, &buttonAPI)
	}

	for _, nodeDB := range backRepo.BackRepoNode.Map_NodeDBID_NodeDB {

		var nodeAPI NodeAPI
		nodeAPI.ID = nodeDB.ID
		nodeAPI.NodePointersEncoding = nodeDB.NodePointersEncoding
		nodeDB.CopyBasicFieldsToNode_WOP(&nodeAPI.Node_WOP)

		backRepoData.NodeAPIs = append(backRepoData.NodeAPIs, &nodeAPI)
	}

	for _, svgiconDB := range backRepo.BackRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {

		var svgiconAPI SVGIconAPI
		svgiconAPI.ID = svgiconDB.ID
		svgiconAPI.SVGIconPointersEncoding = svgiconDB.SVGIconPointersEncoding
		svgiconDB.CopyBasicFieldsToSVGIcon_WOP(&svgiconAPI.SVGIcon_WOP)

		backRepoData.SVGIconAPIs = append(backRepoData.SVGIconAPIs, &svgiconAPI)
	}

	for _, treeDB := range backRepo.BackRepoTree.Map_TreeDBID_TreeDB {

		var treeAPI TreeAPI
		treeAPI.ID = treeDB.ID
		treeAPI.TreePointersEncoding = treeDB.TreePointersEncoding
		treeDB.CopyBasicFieldsToTree_WOP(&treeAPI.Tree_WOP)

		backRepoData.TreeAPIs = append(backRepoData.TreeAPIs, &treeAPI)
	}

}
