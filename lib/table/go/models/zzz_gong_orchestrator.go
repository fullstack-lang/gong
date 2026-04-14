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
// CellIconOrchestrator
type CellIconOrchestrator struct {
}

func (orchestrator *CellIconOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCellIcon, backRepoCellIcon *CellIcon) {

	stagedCellIcon.OnAfterUpdate(gongsvgStage, stagedCellIcon, backRepoCellIcon)
}
// RowOrchestrator
type RowOrchestrator struct {
}

func (orchestrator *RowOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedRow, backRepoRow *Row) {

	stagedRow.OnAfterUpdate(gongsvgStage, stagedRow, backRepoRow)
}
// TableOrchestrator
type TableOrchestrator struct {
}

func (orchestrator *TableOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedTable, backRepoTable *Table) {

	stagedTable.OnAfterUpdate(gongsvgStage, stagedTable, backRepoTable)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Button:
		stage.OnAfterButtonUpdateCallback = new(ButtonOrchestrator)
	case CellIcon:
		stage.OnAfterCellIconUpdateCallback = new(CellIconOrchestrator)
	case Row:
		stage.OnAfterRowUpdateCallback = new(RowOrchestrator)
	case Table:
		stage.OnAfterTableUpdateCallback = new(TableOrchestrator)

	}

}
