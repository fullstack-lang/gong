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

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Checkbox:
		stage.OnAfterCheckboxUpdateCallback = new(CheckboxOrchestrator)
	case Slider:
		stage.OnAfterSliderUpdateCallback = new(SliderOrchestrator)

	}

}
