// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	Form2APIs []*Form2API

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, form2DB := range backRepo.BackRepoForm2.Map_Form2DBID_Form2DB {

		var form2API Form2API
		form2API.ID = form2DB.ID
		form2API.Form2PointersEncoding = form2DB.Form2PointersEncoding
		form2DB.CopyBasicFieldsToForm2_WOP(&form2API.Form2_WOP)

		backRepoData.Form2APIs = append(backRepoData.Form2APIs, &form2API)
	}

}
