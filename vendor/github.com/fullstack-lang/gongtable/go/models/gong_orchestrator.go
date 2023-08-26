// generated code - do not edit
package models

// insertion point
// CellIconOrchestrator
type CellIconOrchestrator struct {
}

func (orchestrator *CellIconOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedCellIcon, backRepoCellIcon *CellIcon) {

	stagedCellIcon.OnAfterUpdate(gongsvgStage, stagedCellIcon, backRepoCellIcon)
}
// FormEditAssocButtonOrchestrator
type FormEditAssocButtonOrchestrator struct {
}

func (orchestrator *FormEditAssocButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedFormEditAssocButton, backRepoFormEditAssocButton *FormEditAssocButton) {

	stagedFormEditAssocButton.OnAfterUpdate(gongsvgStage, stagedFormEditAssocButton, backRepoFormEditAssocButton)
}
// FormSortAssocButtonOrchestrator
type FormSortAssocButtonOrchestrator struct {
}

func (orchestrator *FormSortAssocButtonOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedFormSortAssocButton, backRepoFormSortAssocButton *FormSortAssocButton) {

	stagedFormSortAssocButton.OnAfterUpdate(gongsvgStage, stagedFormSortAssocButton, backRepoFormSortAssocButton)
}
// RowOrchestrator
type RowOrchestrator struct {
}

func (orchestrator *RowOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRow, backRepoRow *Row) {

	stagedRow.OnAfterUpdate(gongsvgStage, stagedRow, backRepoRow)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case CellIcon:
		stage.OnAfterCellIconUpdateCallback = new(CellIconOrchestrator)
	case FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonUpdateCallback = new(FormEditAssocButtonOrchestrator)
	case FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonUpdateCallback = new(FormSortAssocButtonOrchestrator)
	case Row:
		stage.OnAfterRowUpdateCallback = new(RowOrchestrator)

	}

}
