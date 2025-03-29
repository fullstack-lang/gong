// generated code - do not edit
package models

// insertion point
// ButtonOrchestrator
type ButtonOrchestrator struct {
}

func (orchestrator *ButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedButton, backRepoButton *Button) {

	stagedButton.OnAfterUpdate(gongsvgStage, stagedButton, backRepoButton)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Button:
		stage.OnAfterButtonUpdateCallback = new(ButtonOrchestrator)

	}

}
