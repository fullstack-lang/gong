// insertion point for imports
import { DisplayedColumnAPI } from './displayedcolumn-api'
import { RowAPI } from './row-api'
import { ButtonAPI } from './button-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableAPI {

	static GONGSTRUCT_NAME = "Table"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasFiltering: boolean = false
	HasColumnSorting: boolean = false
	HasPaginator: boolean = false
	HasCheckableRows: boolean = false
	HasSaveButton: boolean = false
	SaveButtonLabel: string = ""
	HasBulkDeleteButton: boolean = false
	BulkDeleteButtonTooltip: string = ""
	BulkDeleteSelectedRowsIDsJson: string = ""
	CanDragDropRows: boolean = false
	HasCloseButton: boolean = false
	SavingInProgress: boolean = false
	NbOfStickyColumns: number = 0

	// insertion point for other decls

	TablePointersEncoding: TablePointersEncoding = new TablePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class TablePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	DisplayedColumns: number[] = []
	Rows: number[] = []
	Buttons: number[] = []
}
