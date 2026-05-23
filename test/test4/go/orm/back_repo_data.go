// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AstructAPIs []*AstructAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, astructDB := range backRepo.BackRepoAstruct.Map_AstructDBID_AstructDB {

		var astructAPI AstructAPI
		astructAPI.ID = astructDB.ID
		astructAPI.AstructPointersEncoding = astructDB.AstructPointersEncoding
		astructDB.CopyBasicFieldsToAstruct_WOP(&astructAPI.Astruct_WOP)

		backRepoData.AstructAPIs = append(backRepoData.AstructAPIs, &astructAPI)
	}

}
