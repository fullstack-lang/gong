// generated code - do not edit
package models

// insertion point
// FileToUploadOrchestrator
type FileToUploadOrchestrator struct {
}

func (orchestrator *FileToUploadOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFileToUpload, backRepoFileToUpload *FileToUpload) {

	stagedFileToUpload.OnAfterUpdate(gongsvgStage, stagedFileToUpload, backRepoFileToUpload)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case FileToUpload:
		stage.OnAfterFileToUploadUpdateCallback = new(FileToUploadOrchestrator)

	}

}
