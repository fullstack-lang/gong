// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FileToDownloadAPIs []*FileToDownloadAPI

	FileToUploadAPIs []*FileToUploadAPI

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

	for _, filetouploadDB := range backRepo.BackRepoFileToUpload.Map_FileToUploadDBID_FileToUploadDB {

		var filetouploadAPI FileToUploadAPI
		filetouploadAPI.ID = filetouploadDB.ID
		filetouploadAPI.FileToUploadPointersEncoding = filetouploadDB.FileToUploadPointersEncoding
		filetouploadDB.CopyBasicFieldsToFileToUpload_WOP(&filetouploadAPI.FileToUpload_WOP)

		backRepoData.FileToUploadAPIs = append(backRepoData.FileToUploadAPIs, &filetouploadAPI)
	}

}
