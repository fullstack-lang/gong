// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *FileToDownload:
		if stage.OnAfterFileToDownloadCreateCallback != nil {
			stage.OnAfterFileToDownloadCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *FileToDownload:
		newTarget := any(new).(*FileToDownload)
		if stage.OnAfterFileToDownloadUpdateCallback != nil {
			stage.OnAfterFileToDownloadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *FileToDownload:
		if stage.OnAfterFileToDownloadDeleteCallback != nil {
			staged := any(staged).(*FileToDownload)
			stage.OnAfterFileToDownloadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *FileToDownload:
		if stage.OnAfterFileToDownloadReadCallback != nil {
			stage.OnAfterFileToDownloadReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *FileToDownload:
		stage.OnAfterFileToDownloadUpdateCallback = any(callback).(OnAfterUpdateInterface[FileToDownload])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *FileToDownload:
		stage.OnAfterFileToDownloadCreateCallback = any(callback).(OnAfterCreateInterface[FileToDownload])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *FileToDownload:
		stage.OnAfterFileToDownloadDeleteCallback = any(callback).(OnAfterDeleteInterface[FileToDownload])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *FileToDownload:
		stage.OnAfterFileToDownloadReadCallback = any(callback).(OnAfterReadInterface[FileToDownload])
	
	}
}
