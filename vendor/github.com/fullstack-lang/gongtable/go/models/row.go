package models

type Row struct {
	Name  string
	Cells []*Cell

	IsChecked bool

	// swagger:ignore
	Impl RowImplInterface
}

type RowImplInterface interface {

	// RowUpdated function is called each time a Row is modified
	RowUpdated(stage *StageStruct, row, updatedRow *Row)
}

// OnAfterUpdate is called when there is an update to the row
// note the play on words "font row"
func (row *Row) OnAfterUpdate(stage *StageStruct, _, frontRow *Row) {

	if row.Impl != nil {
		row.Impl.RowUpdated(stage, row, frontRow)
	}
}
