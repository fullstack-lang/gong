package models

// FileToDownload is the content that need to be downloaded
type FileToDownload struct {

	//gong:width 600
	Name string

	//gong:text
	//gong:width 600 gong:height 2000
	Content string
}

type FileToUploadProxy interface {
	OnFileUpload(*FileToUpload)
}

// FileToUpload is the content that need to be downloaded
// The instance must be created by the back, otherwise the front will not work
type FileToUpload struct {

	//gong:width 600
	Name string

	//gong:text
	//gong:width 600 gong:height 2000
	Content string

	FileToUploadProxy FileToUploadProxy
}

func (fileToUpload *FileToUpload) OnAfterUpdate(stage *Stage, staged, front *FileToUpload) {

	if fileToUpload.FileToUploadProxy != nil {
		fileToUpload.FileToUploadProxy.OnFileUpload(front)
	}
}
