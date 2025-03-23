// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FileToDownloadAPIs []*FileToDownloadAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, filetodownloadDB := range backRepo.BackRepoFileToDownload.Map_FileToDownloadDBID_FileToDownloadDB {

		var filetodownloadAPI FileToDownloadAPI
		filetodownloadAPI.ID = filetodownloadDB.ID
		filetodownloadAPI.FileToDownloadPointersEncoding = filetodownloadDB.FileToDownloadPointersEncoding
		filetodownloadDB.CopyBasicFieldsToFileToDownload_WOP(&filetodownloadAPI.FileToDownload_WOP)

		backRepoData.FileToDownloadAPIs = append(backRepoData.FileToDownloadAPIs, &filetodownloadAPI)
	}

}
