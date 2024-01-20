// generated code - do not edit

import { TableDB } from './table-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { DisplayedColumn } from './displayedcolumn'
import { Row } from './row'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Table {

	static GONGSTRUCT_NAME = "Table"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasFiltering: boolean = false
	HasColumnSorting: boolean = false
	HasPaginator: boolean = false
	HasCheckableRows: boolean = false
	HasSaveButton: boolean = false
	CanDragDropRows: boolean = false
	HasCloseButton: boolean = false
	SavingInProgress: boolean = false
	NbOfStickyColumns: number = 0

	// insertion point for pointers and slices of pointers declarations
	DisplayedColumns: Array<DisplayedColumn> = []
	Rows: Array<Row> = []
}

export function CopyTableToTableDB(table: Table, tableDB: TableDB) {

	tableDB.CreatedAt = table.CreatedAt
	tableDB.DeletedAt = table.DeletedAt
	tableDB.ID = table.ID
	
	// insertion point for basic fields copy operations
	tableDB.Name = table.Name
	tableDB.HasFiltering = table.HasFiltering
	tableDB.HasColumnSorting = table.HasColumnSorting
	tableDB.HasPaginator = table.HasPaginator
	tableDB.HasCheckableRows = table.HasCheckableRows
	tableDB.HasSaveButton = table.HasSaveButton
	tableDB.CanDragDropRows = table.CanDragDropRows
	tableDB.HasCloseButton = table.HasCloseButton
	tableDB.SavingInProgress = table.SavingInProgress
	tableDB.NbOfStickyColumns = table.NbOfStickyColumns

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	tableDB.TablePointersEncoding.DisplayedColumns = []
    for (let _displayedcolumn of table.DisplayedColumns) {
		tableDB.TablePointersEncoding.DisplayedColumns.push(_displayedcolumn.ID)
    }
	
	tableDB.TablePointersEncoding.Rows = []
    for (let _row of table.Rows) {
		tableDB.TablePointersEncoding.Rows.push(_row.ID)
    }
	
}

export function CopyTableDBToTable(tableDB: TableDB, table: Table, frontRepo: FrontRepo) {

	table.CreatedAt = tableDB.CreatedAt
	table.DeletedAt = tableDB.DeletedAt
	table.ID = tableDB.ID
	
	// insertion point for basic fields copy operations
	table.Name = tableDB.Name
	table.HasFiltering = tableDB.HasFiltering
	table.HasColumnSorting = tableDB.HasColumnSorting
	table.HasPaginator = tableDB.HasPaginator
	table.HasCheckableRows = tableDB.HasCheckableRows
	table.HasSaveButton = tableDB.HasSaveButton
	table.CanDragDropRows = tableDB.CanDragDropRows
	table.HasCloseButton = tableDB.HasCloseButton
	table.SavingInProgress = tableDB.SavingInProgress
	table.NbOfStickyColumns = tableDB.NbOfStickyColumns

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	table.DisplayedColumns = new Array<DisplayedColumn>()
	for (let _id of tableDB.TablePointersEncoding.DisplayedColumns) {
	  let _displayedcolumn = frontRepo.DisplayedColumns.get(_id)
	  if (_displayedcolumn != undefined) {
		table.DisplayedColumns.push(_displayedcolumn!)
	  }
	}
	table.Rows = new Array<Row>()
	for (let _id of tableDB.TablePointersEncoding.Rows) {
	  let _row = frontRepo.Rows.get(_id)
	  if (_row != undefined) {
		table.Rows.push(_row!)
	  }
	}
}
