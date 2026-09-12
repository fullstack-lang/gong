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

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case FileToUpload:
		stage.OnAfterFileToUploadUpdateCallback = new(FileToUploadOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
