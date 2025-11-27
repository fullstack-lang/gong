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
// ButtonToggleOrchestrator
type ButtonToggleOrchestrator struct {
}

func (orchestrator *ButtonToggleOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedButtonToggle, backRepoButtonToggle *ButtonToggle) {

	stagedButtonToggle.OnAfterUpdate(gongsvgStage, stagedButtonToggle, backRepoButtonToggle)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Button:
		stage.OnAfterButtonUpdateCallback = new(ButtonOrchestrator)
	case ButtonToggle:
		stage.OnAfterButtonToggleUpdateCallback = new(ButtonToggleOrchestrator)

	}

}
