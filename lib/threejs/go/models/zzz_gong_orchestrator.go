// generated code - do not edit
package models

// insertion point
// CameraOrchestrator
type CameraOrchestrator struct {
}

func (orchestrator *CameraOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCamera, backRepoCamera *Camera) {

	stagedCamera.OnAfterUpdate(gongsvgStage, stagedCamera, backRepoCamera)
}

// CanvasOrchestrator
type CanvasOrchestrator struct {
}

func (orchestrator *CanvasOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCanvas, backRepoCanvas *Canvas) {

	stagedCanvas.OnAfterUpdate(gongsvgStage, stagedCanvas, backRepoCanvas)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Camera:
		stage.OnAfterCameraUpdateCallback = new(CameraOrchestrator)
	case Canvas:
		stage.OnAfterCanvasUpdateCallback = new(CanvasOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
