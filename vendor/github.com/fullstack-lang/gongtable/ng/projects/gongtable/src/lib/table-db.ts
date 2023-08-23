// insertion point for imports
import { DisplayedColumnDB } from './displayedcolumn-db'
import { RowDB } from './row-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableDB {

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

	// insertion point for other declarations
	DisplayedColumns?: Array<DisplayedColumnDB>
	Rows?: Array<RowDB>
}
