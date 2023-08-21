// insertion point for imports
import { TableDB } from './table-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplayedColumnDB {

	static GONGSTRUCT_NAME = "DisplayedColumn"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Table_DisplayedColumnsDBID: NullInt64 = new NullInt64
	Table_DisplayedColumnsDBID_Index: NullInt64  = new NullInt64 // store the index of the displayedcolumn instance in Table.DisplayedColumns
	Table_DisplayedColumns_reverse?: TableDB 

}
