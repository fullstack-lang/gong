// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AstructDBs []*AstructDB

	AstructBstruct2UseDBs []*AstructBstruct2UseDB

	AstructBstructUseDBs []*AstructBstructUseDB

	BstructDBs []*BstructDB

	DstructDBs []*DstructDB

	FstructDBs []*FstructDB
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, astructDB := range backRepo.BackRepoAstruct.Map_AstructDBID_AstructDB {
		backRepoData.AstructDBs = append(backRepoData.AstructDBs, astructDB)
	}

	for _, astructbstruct2useDB := range backRepo.BackRepoAstructBstruct2Use.Map_AstructBstruct2UseDBID_AstructBstruct2UseDB {
		backRepoData.AstructBstruct2UseDBs = append(backRepoData.AstructBstruct2UseDBs, astructbstruct2useDB)
	}

	for _, astructbstructuseDB := range backRepo.BackRepoAstructBstructUse.Map_AstructBstructUseDBID_AstructBstructUseDB {
		backRepoData.AstructBstructUseDBs = append(backRepoData.AstructBstructUseDBs, astructbstructuseDB)
	}

	for _, bstructDB := range backRepo.BackRepoBstruct.Map_BstructDBID_BstructDB {
		backRepoData.BstructDBs = append(backRepoData.BstructDBs, bstructDB)
	}

	for _, dstructDB := range backRepo.BackRepoDstruct.Map_DstructDBID_DstructDB {
		backRepoData.DstructDBs = append(backRepoData.DstructDBs, dstructDB)
	}

	for _, fstructDB := range backRepo.BackRepoFstruct.Map_FstructDBID_FstructDB {
		backRepoData.FstructDBs = append(backRepoData.FstructDBs, fstructDB)
	}

}
