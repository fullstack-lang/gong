// generated code - do not edit
package models

// insertion point
// FormEditAssocButtonOrchestrator
type FormEditAssocButtonOrchestrator struct {
}

func (orchestrator *FormEditAssocButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedFormEditAssocButton, backRepoFormEditAssocButton *FormEditAssocButton) {

	stagedFormEditAssocButton.OnAfterUpdate(gongsvgStage, stagedFormEditAssocButton, backRepoFormEditAssocButton)
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

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonUpdateCallback = new(FormEditAssocButtonOrchestrator)
	case FormGroup:
		stage.OnAfterFormGroupUpdateCallback = new(FormGroupOrchestrator)
	case FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonUpdateCallback = new(FormSortAssocButtonOrchestrator)

	}

}
