// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *FileToDownload:
		if stage.OnAfterFileToDownloadCreateCallback != nil {
			stage.OnAfterFileToDownloadCreateCallback.OnAfterCreate(stage, target)
		}
	case *FileToUpload:
		if stage.OnAfterFileToUploadCreateCallback != nil {
			stage.OnAfterFileToUploadCreateCallback.OnAfterCreate(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageCreateCallback != nil {
			stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *FileToDownload:
		newTarget := any(new).(*FileToDownload)
		if stage.OnAfterFileToDownloadUpdateCallback != nil {
			stage.OnAfterFileToDownloadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FileToUpload:
		newTarget := any(new).(*FileToUpload)
		if stage.OnAfterFileToUploadUpdateCallback != nil {
			stage.OnAfterFileToUploadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Message:
		newTarget := any(new).(*Message)
		if stage.OnAfterMessageUpdateCallback != nil {
			stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *FileToDownload:
		if stage.OnAfterFileToDownloadDeleteCallback != nil {
			staged := any(staged).(*FileToDownload)
			stage.OnAfterFileToDownloadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FileToUpload:
		if stage.OnAfterFileToUploadDeleteCallback != nil {
			staged := any(staged).(*FileToUpload)
			stage.OnAfterFileToUploadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Message:
		if stage.OnAfterMessageDeleteCallback != nil {
			staged := any(staged).(*Message)
			stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
