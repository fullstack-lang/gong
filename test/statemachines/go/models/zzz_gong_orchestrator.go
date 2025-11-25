// generated code - do not edit
package models

// insertion point
// TransitionOrchestrator
type TransitionOrchestrator struct {
}

func (orchestrator *TransitionOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedTransition, backRepoTransition *Transition) {

	stagedTransition.OnAfterUpdate(gongsvgStage, stagedTransition, backRepoTransition)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Transition:
		stage.OnAfterTransitionUpdateCallback = new(TransitionOrchestrator)

	}

}
