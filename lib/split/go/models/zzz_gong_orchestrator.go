// generated code - do not edit
package models

// insertion point
// ViewOrchestrator
type ViewOrchestrator struct {
}

func (orchestrator *ViewOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedView, backRepoView *View) {

	stagedView.OnAfterUpdate(gongsvgStage, stagedView, backRepoView)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case View:
		stage.OnAfterViewUpdateCallback = new(ViewOrchestrator)

	}

}
