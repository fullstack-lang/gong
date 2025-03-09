// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AAPIs []*AAPI

	BAPIs []*BAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, aDB := range backRepo.BackRepoA.Map_ADBID_ADB {

		var aAPI AAPI
		aAPI.ID = aDB.ID
		aAPI.APointersEncoding = aDB.APointersEncoding
		aDB.CopyBasicFieldsToA_WOP(&aAPI.A_WOP)

		backRepoData.AAPIs = append(backRepoData.AAPIs, &aAPI)
	}

	for _, bDB := range backRepo.BackRepoB.Map_BDBID_BDB {

		var bAPI BAPI
		bAPI.ID = bDB.ID
		bAPI.BPointersEncoding = bDB.BPointersEncoding
		bDB.CopyBasicFieldsToB_WOP(&bAPI.B_WOP)

		backRepoData.BAPIs = append(backRepoData.BAPIs, &bAPI)
	}

}
