// generated code - do not edit

import { TableAPI } from './table-api'
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

export function CopyTableToTableAPI(table: Table, tableAPI: TableAPI) {

	tableAPI.CreatedAt = table.CreatedAt
	tableAPI.DeletedAt = table.DeletedAt
	tableAPI.ID = table.ID

	// insertion point for basic fields copy operations
	tableAPI.Name = table.Name
	tableAPI.HasFiltering = table.HasFiltering
	tableAPI.HasColumnSorting = table.HasColumnSorting
	tableAPI.HasPaginator = table.HasPaginator
	tableAPI.HasCheckableRows = table.HasCheckableRows
	tableAPI.HasSaveButton = table.HasSaveButton
	tableAPI.CanDragDropRows = table.CanDragDropRows
	tableAPI.HasCloseButton = table.HasCloseButton
	tableAPI.SavingInProgress = table.SavingInProgress
	tableAPI.NbOfStickyColumns = table.NbOfStickyColumns

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	tableAPI.TablePointersEncoding.DisplayedColumns = []
	for (let _displayedcolumn of table.DisplayedColumns) {
		tableAPI.TablePointersEncoding.DisplayedColumns.push(_displayedcolumn.ID)
	}

	tableAPI.TablePointersEncoding.Rows = []
	for (let _row of table.Rows) {
		tableAPI.TablePointersEncoding.Rows.push(_row.ID)
	}

}

// CopyTableAPIToTable update basic, pointers and slice of pointers fields of table
// from respectively the basic fields and encoded fields of pointers and slices of pointers of tableAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTableAPIToTable(tableAPI: TableAPI, table: Table, frontRepo: FrontRepo) {

	table.CreatedAt = tableAPI.CreatedAt
	table.DeletedAt = tableAPI.DeletedAt
	table.ID = tableAPI.ID

	// insertion point for basic fields copy operations
	table.Name = tableAPI.Name
	table.HasFiltering = tableAPI.HasFiltering
	table.HasColumnSorting = tableAPI.HasColumnSorting
	table.HasPaginator = tableAPI.HasPaginator
	table.HasCheckableRows = tableAPI.HasCheckableRows
	table.HasSaveButton = tableAPI.HasSaveButton
	table.CanDragDropRows = tableAPI.CanDragDropRows
	table.HasCloseButton = tableAPI.HasCloseButton
	table.SavingInProgress = tableAPI.SavingInProgress
	table.NbOfStickyColumns = tableAPI.NbOfStickyColumns

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(tableAPI.TablePointersEncoding.DisplayedColumns)) {
		console.error('Rects is not an array:', tableAPI.TablePointersEncoding.DisplayedColumns);
		return;
	}

	table.DisplayedColumns = new Array<DisplayedColumn>()
	for (let _id of tableAPI.TablePointersEncoding.DisplayedColumns) {
		let _displayedcolumn = frontRepo.map_ID_DisplayedColumn.get(_id)
		if (_displayedcolumn != undefined) {
			table.DisplayedColumns.push(_displayedcolumn!)
		}
	}
	if (!Array.isArray(tableAPI.TablePointersEncoding.Rows)) {
		console.error('Rects is not an array:', tableAPI.TablePointersEncoding.Rows);
		return;
	}

	table.Rows = new Array<Row>()
	for (let _id of tableAPI.TablePointersEncoding.Rows) {
		let _row = frontRepo.map_ID_Row.get(_id)
		if (_row != undefined) {
			table.Rows.push(_row!)
		}
	}
}
