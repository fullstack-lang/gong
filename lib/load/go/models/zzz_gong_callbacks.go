// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (filetodownload *FileToDownload) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFileToDownloadCreateCallback != nil {
		stage.OnAfterFileToDownloadCreateCallback.OnAfterCreate(stage, filetodownload)
	}
}

func (filetodownload *FileToDownload) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToDownloadUpdateCallback != nil {
		var frontFileToDownload *FileToDownload
		if front != nil {
			frontFileToDownload, _ = front.(*FileToDownload)
		}
		stage.OnAfterFileToDownloadUpdateCallback.OnAfterUpdate(stage, filetodownload, frontFileToDownload)
	}
}

func (filetodownload *FileToDownload) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToDownloadDeleteCallback != nil {
		var frontFileToDownload *FileToDownload
		if front != nil {
			frontFileToDownload, _ = front.(*FileToDownload)
		}
		stage.OnAfterFileToDownloadDeleteCallback.OnAfterDelete(stage, filetodownload, frontFileToDownload)
	}
}

func (filetoupload *FileToUpload) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFileToUploadCreateCallback != nil {
		stage.OnAfterFileToUploadCreateCallback.OnAfterCreate(stage, filetoupload)
	}
}

func (filetoupload *FileToUpload) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToUploadUpdateCallback != nil {
		var frontFileToUpload *FileToUpload
		if front != nil {
			frontFileToUpload, _ = front.(*FileToUpload)
		}
		stage.OnAfterFileToUploadUpdateCallback.OnAfterUpdate(stage, filetoupload, frontFileToUpload)
	}
}

func (filetoupload *FileToUpload) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToUploadDeleteCallback != nil {
		var frontFileToUpload *FileToUpload
		if front != nil {
			frontFileToUpload, _ = front.(*FileToUpload)
		}
		stage.OnAfterFileToUploadDeleteCallback.OnAfterDelete(stage, filetoupload, frontFileToUpload)
	}
}

func (message *Message) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMessageCreateCallback != nil {
		stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, message)
	}
}

func (message *Message) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageUpdateCallback != nil {
		var frontMessage *Message
		if front != nil {
			frontMessage, _ = front.(*Message)
		}
		stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, message, frontMessage)
	}
}

func (message *Message) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageDeleteCallback != nil {
		var frontMessage *Message
		if front != nil {
			frontMessage, _ = front.(*Message)
		}
		stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, message, frontMessage)
	}
}

