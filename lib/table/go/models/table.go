package models

type Table struct {
	Name             string
	DisplayedColumns []*DisplayedColumn
	Rows             []*Row

	HasFiltering     bool
	HasColumnSorting bool
	HasPaginator     bool

	HasCheckableRows bool

	HasSaveButton bool

	CanDragDropRows bool
	HasCloseButton  bool // Is used in case of drag drop since drag drop operation save the result

	// SavingInProgress is true when rows are being saved
	// it is set to true by the front at the begining and set back
	// to false by the front
	// This information is used by the back to compute when all rows to be updated
	// have been updated
	SavingInProgress bool

	// The `sticky` property in CSS positions an element based on user's scroll,
	// allowing it to be fixed within its parent after a scroll point is reached.
	// NbOfStickyColumns defines the number of sticky column at the left of the table
	NbOfStickyColumns int

	// swagger:ignore
	Impl TableImplInterface
}

type TableImplInterface interface {

	// TableUpdated function is called each time a Table is modified
	TableUpdated(stage *StageStruct, table, updatedTable *Table)
}

// OnAfterUpdate is called when there is an update to the table
// note the play on words "font table"
func (table *Table) OnAfterUpdate(stage *StageStruct, _, frontTable *Table) {

	if table.Impl != nil {
		table.Impl.TableUpdated(stage, table, frontTable)
	}
}
