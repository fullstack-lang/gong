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

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Camera:
		stage.OnAfterCameraUpdateCallback = new(CameraOrchestrator)
	case Canvas:
		stage.OnAfterCanvasUpdateCallback = new(CanvasOrchestrator)

	}

}
