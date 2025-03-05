// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	SplitAreaAPIs []*SplitAreaAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, splitareaDB := range backRepo.BackRepoSplitArea.Map_SplitAreaDBID_SplitAreaDB {

		var splitareaAPI SplitAreaAPI
		splitareaAPI.ID = splitareaDB.ID
		splitareaAPI.SplitAreaPointersEncoding = splitareaDB.SplitAreaPointersEncoding
		splitareaDB.CopyBasicFieldsToSplitArea_WOP(&splitareaAPI.SplitArea_WOP)

		backRepoData.SplitAreaAPIs = append(backRepoData.SplitAreaAPIs, &splitareaAPI)
	}

}
