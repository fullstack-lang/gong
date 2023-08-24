// insertion point for imports
import { CellStringDB } from './cellstring-db'
import { CellFloat64DB } from './cellfloat64-db'
import { CellIntDB } from './cellint-db'
import { CellBooleanDB } from './cellboolean-db'
import { CellIconDB } from './cellicon-db'
import { RowDB } from './row-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellDB {

	static GONGSTRUCT_NAME = "Cell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	CellString?: CellStringDB
	CellStringID: NullInt64 = new NullInt64 // if pointer is null, CellString.ID = 0

	CellFloat64?: CellFloat64DB
	CellFloat64ID: NullInt64 = new NullInt64 // if pointer is null, CellFloat64.ID = 0

	CellInt?: CellIntDB
	CellIntID: NullInt64 = new NullInt64 // if pointer is null, CellInt.ID = 0

	CellBool?: CellBooleanDB
	CellBoolID: NullInt64 = new NullInt64 // if pointer is null, CellBool.ID = 0

	CellIcon?: CellIconDB
	CellIconID: NullInt64 = new NullInt64 // if pointer is null, CellIcon.ID = 0

	Row_CellsDBID: NullInt64 = new NullInt64
	Row_CellsDBID_Index: NullInt64  = new NullInt64 // store the index of the cell instance in Row.Cells
	Row_Cells_reverse?: RowDB 

}
