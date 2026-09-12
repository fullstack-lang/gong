// generated code - do not edit
package models

// insertion point
// CheckboxOrchestrator
type CheckboxOrchestrator struct {
}

func (orchestrator *CheckboxOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCheckbox, backRepoCheckbox *Checkbox) {

	stagedCheckbox.OnAfterUpdate(gongsvgStage, stagedCheckbox, backRepoCheckbox)
}

// SliderOrchestrator
type SliderOrchestrator struct {
}

func (orchestrator *SliderOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedSlider, backRepoSlider *Slider) {

	stagedSlider.OnAfterUpdate(gongsvgStage, stagedSlider, backRepoSlider)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Checkbox:
		stage.OnAfterCheckboxUpdateCallback = new(CheckboxOrchestrator)
	case Slider:
		stage.OnAfterSliderUpdateCallback = new(SliderOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
