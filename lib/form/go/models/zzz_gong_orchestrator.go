// generated code - do not edit
package models

// insertion point
// CheckBoxOrchestrator
type CheckBoxOrchestrator struct {
}

func (orchestrator *CheckBoxOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCheckBox, backRepoCheckBox *CheckBox) {

	stagedCheckBox.OnAfterUpdate(gongsvgStage, stagedCheckBox, backRepoCheckBox)
}

// FormEditAssocButtonOrchestrator
type FormEditAssocButtonOrchestrator struct {
}

func (orchestrator *FormEditAssocButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormEditAssocButton, backRepoFormEditAssocButton *FormEditAssocButton) {

	stagedFormEditAssocButton.OnAfterUpdate(gongsvgStage, stagedFormEditAssocButton, backRepoFormEditAssocButton)
}

// FormFieldFloat64Orchestrator
type FormFieldFloat64Orchestrator struct {
}

func (orchestrator *FormFieldFloat64Orchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormFieldFloat64, backRepoFormFieldFloat64 *FormFieldFloat64) {

	stagedFormFieldFloat64.OnAfterUpdate(gongsvgStage, stagedFormFieldFloat64, backRepoFormFieldFloat64)
}

// FormFieldIntOrchestrator
type FormFieldIntOrchestrator struct {
}

func (orchestrator *FormFieldIntOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormFieldInt, backRepoFormFieldInt *FormFieldInt) {

	stagedFormFieldInt.OnAfterUpdate(gongsvgStage, stagedFormFieldInt, backRepoFormFieldInt)
}

// FormFieldSelectOrchestrator
type FormFieldSelectOrchestrator struct {
}

func (orchestrator *FormFieldSelectOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormFieldSelect, backRepoFormFieldSelect *FormFieldSelect) {

	stagedFormFieldSelect.OnAfterUpdate(gongsvgStage, stagedFormFieldSelect, backRepoFormFieldSelect)
}

// FormGroupOrchestrator
type FormGroupOrchestrator struct {
}

func (orchestrator *FormGroupOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormGroup, backRepoFormGroup *FormGroup) {

	stagedFormGroup.OnAfterUpdate(gongsvgStage, stagedFormGroup, backRepoFormGroup)
}

// FormSortAssocButtonOrchestrator
type FormSortAssocButtonOrchestrator struct {
}

func (orchestrator *FormSortAssocButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormSortAssocButton, backRepoFormSortAssocButton *FormSortAssocButton) {

	stagedFormSortAssocButton.OnAfterUpdate(gongsvgStage, stagedFormSortAssocButton, backRepoFormSortAssocButton)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case CheckBox:
		stage.OnAfterCheckBoxUpdateCallback = new(CheckBoxOrchestrator)
	case FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonUpdateCallback = new(FormEditAssocButtonOrchestrator)
	case FormFieldFloat64:
		stage.OnAfterFormFieldFloat64UpdateCallback = new(FormFieldFloat64Orchestrator)
	case FormFieldInt:
		stage.OnAfterFormFieldIntUpdateCallback = new(FormFieldIntOrchestrator)
	case FormFieldSelect:
		stage.OnAfterFormFieldSelectUpdateCallback = new(FormFieldSelectOrchestrator)
	case FormGroup:
		stage.OnAfterFormGroupUpdateCallback = new(FormGroupOrchestrator)
	case FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonUpdateCallback = new(FormSortAssocButtonOrchestrator)
	}

}
