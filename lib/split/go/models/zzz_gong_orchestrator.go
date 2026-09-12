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

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case View:
		stage.OnAfterViewUpdateCallback = new(ViewOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
