// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AstructAPIs []*AstructAPI

	AstructBstruct2UseAPIs []*AstructBstruct2UseAPI

	AstructBstructUseAPIs []*AstructBstructUseAPI

	BstructAPIs []*BstructAPI

	DstructAPIs []*DstructAPI

	FstructAPIs []*FstructAPI

	GstructAPIs []*GstructAPI
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

	for _, astructbstruct2useDB := range backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UseDB {

		var astructbstruct2useAPI AstructBstruct2UseAPI
		astructbstruct2useAPI.ID = astructbstruct2useDB.ID
		astructbstruct2useAPI.AstructBstruct2UsePointersEncoding = astructbstruct2useDB.AstructBstruct2UsePointersEncoding
		astructbstruct2useDB.CopyBasicFieldsToAstructBstruct2Use_WOP(&astructbstruct2useAPI.AstructBstruct2Use_WOP)

		backRepoData.AstructBstruct2UseAPIs = append(backRepoData.AstructBstruct2UseAPIs, &astructbstruct2useAPI)
	}

	for _, astructbstructuseDB := range backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUseDB {

		var astructbstructuseAPI AstructBstructUseAPI
		astructbstructuseAPI.ID = astructbstructuseDB.ID
		astructbstructuseAPI.AstructBstructUsePointersEncoding = astructbstructuseDB.AstructBstructUsePointersEncoding
		astructbstructuseDB.CopyBasicFieldsToAstructBstructUse_WOP(&astructbstructuseAPI.AstructBstructUse_WOP)

		backRepoData.AstructBstructUseAPIs = append(backRepoData.AstructBstructUseAPIs, &astructbstructuseAPI)
	}

	for _, bstructDB := range backRepo.BackRepoBstruct.Map_BstructDBID_BstructDB {

		var bstructAPI BstructAPI
		bstructAPI.ID = bstructDB.ID
		bstructAPI.BstructPointersEncoding = bstructDB.BstructPointersEncoding
		bstructDB.CopyBasicFieldsToBstruct_WOP(&bstructAPI.Bstruct_WOP)

		backRepoData.BstructAPIs = append(backRepoData.BstructAPIs, &bstructAPI)
	}

	for _, dstructDB := range backRepo.BackRepoDstruct.Map_DstructDBID_DstructDB {

		var dstructAPI DstructAPI
		dstructAPI.ID = dstructDB.ID
		dstructAPI.DstructPointersEncoding = dstructDB.DstructPointersEncoding
		dstructDB.CopyBasicFieldsToDstruct_WOP(&dstructAPI.Dstruct_WOP)

		backRepoData.DstructAPIs = append(backRepoData.DstructAPIs, &dstructAPI)
	}

	for _, fstructDB := range backRepo.BackRepoFstruct.Map_FstructDBID_FstructDB {

		var fstructAPI FstructAPI
		fstructAPI.ID = fstructDB.ID
		fstructAPI.FstructPointersEncoding = fstructDB.FstructPointersEncoding
		fstructDB.CopyBasicFieldsToFstruct_WOP(&fstructAPI.Fstruct_WOP)

		backRepoData.FstructAPIs = append(backRepoData.FstructAPIs, &fstructAPI)
	}

	for _, gstructDB := range backRepo.BackRepoGstruct.Map_GstructDBID_GstructDB {

		var gstructAPI GstructAPI
		gstructAPI.ID = gstructDB.ID
		gstructAPI.GstructPointersEncoding = gstructDB.GstructPointersEncoding
		gstructDB.CopyBasicFieldsToGstruct_WOP(&gstructAPI.Gstruct_WOP)

		backRepoData.GstructAPIs = append(backRepoData.GstructAPIs, &gstructAPI)
	}

}
