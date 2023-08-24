// insertion point for imports
import { CellDB } from './cell-db'
import { TableDB } from './table-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RowDB {

	static GONGSTRUCT_NAME = "Row"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsChecked: boolean = false

	// insertion point for other declarations
	Cells?: Array<CellDB>
	Table_RowsDBID: NullInt64 = new NullInt64
	Table_RowsDBID_Index: NullInt64  = new NullInt64 // store the index of the row instance in Table.Rows
	Table_Rows_reverse?: TableDB 

}
