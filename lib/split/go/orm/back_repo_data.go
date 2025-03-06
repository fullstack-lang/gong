// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AsSplitAPIs []*AsSplitAPI

	AsSplitAreaAPIs []*AsSplitAreaAPI

	ViewAPIs []*ViewAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, assplitDB := range backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitDB {

		var assplitAPI AsSplitAPI
		assplitAPI.ID = assplitDB.ID
		assplitAPI.AsSplitPointersEncoding = assplitDB.AsSplitPointersEncoding
		assplitDB.CopyBasicFieldsToAsSplit_WOP(&assplitAPI.AsSplit_WOP)

		backRepoData.AsSplitAPIs = append(backRepoData.AsSplitAPIs, &assplitAPI)
	}

	for _, assplitareaDB := range backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaDB {

		var assplitareaAPI AsSplitAreaAPI
		assplitareaAPI.ID = assplitareaDB.ID
		assplitareaAPI.AsSplitAreaPointersEncoding = assplitareaDB.AsSplitAreaPointersEncoding
		assplitareaDB.CopyBasicFieldsToAsSplitArea_WOP(&assplitareaAPI.AsSplitArea_WOP)

		backRepoData.AsSplitAreaAPIs = append(backRepoData.AsSplitAreaAPIs, &assplitareaAPI)
	}

	for _, viewDB := range backRepo.BackRepoView.Map_ViewDBID_ViewDB {

		var viewAPI ViewAPI
		viewAPI.ID = viewDB.ID
		viewAPI.ViewPointersEncoding = viewDB.ViewPointersEncoding
		viewDB.CopyBasicFieldsToView_WOP(&viewAPI.View_WOP)

		backRepoData.ViewAPIs = append(backRepoData.ViewAPIs, &viewAPI)
	}

}
