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
// FormGroupOrchestrator
type FormGroupOrchestrator struct {
}

func (orchestrator *FormGroupOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedFormGroup, backRepoFormGroup *FormGroup) {

	stagedFormGroup.OnAfterUpdate(gongsvgStage, stagedFormGroup, backRepoFormGroup)
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
// TableOrchestrator
type TableOrchestrator struct {
}

func (orchestrator *TableOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedTable, backRepoTable *Table) {

	stagedTable.OnAfterUpdate(gongsvgStage, stagedTable, backRepoTable)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case CellIcon:
		stage.OnAfterCellIconUpdateCallback = new(CellIconOrchestrator)
	case FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonUpdateCallback = new(FormEditAssocButtonOrchestrator)
	case FormGroup:
		stage.OnAfterFormGroupUpdateCallback = new(FormGroupOrchestrator)
	case FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonUpdateCallback = new(FormSortAssocButtonOrchestrator)
	case Row:
		stage.OnAfterRowUpdateCallback = new(RowOrchestrator)
	case Table:
		stage.OnAfterTableUpdateCallback = new(TableOrchestrator)

	}

}
