// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	Probe2APIs []*Probe2API

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, probe2DB := range backRepo.BackRepoProbe2.Map_Probe2DBID_Probe2DB {

		var probe2API Probe2API
		probe2API.ID = probe2DB.ID
		probe2API.Probe2PointersEncoding = probe2DB.Probe2PointersEncoding
		probe2DB.CopyBasicFieldsToProbe2_WOP(&probe2API.Probe2_WOP)

		backRepoData.Probe2APIs = append(backRepoData.Probe2APIs, &probe2API)
	}

}
