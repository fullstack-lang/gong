// insertion point for imports
import { CellStringDB } from './cellstring-db'
import { CellFloat64DB } from './cellfloat64-db'
import { CellIntDB } from './cellint-db'
import { CellBooleanDB } from './cellboolean-db'
import { CellIconDB } from './cellicon-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellDB {

	static GONGSTRUCT_NAME = "Cell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	CellString?: CellStringDB

	CellFloat64?: CellFloat64DB

	CellInt?: CellIntDB

	CellBool?: CellBooleanDB

	CellIcon?: CellIconDB


	CellPointersEncoding: CellPointersEncoding = new CellPointersEncoding
}

export class CellPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	CellStringID: NullInt64 = new NullInt64 // if pointer is null, CellString.ID = 0

	CellFloat64ID: NullInt64 = new NullInt64 // if pointer is null, CellFloat64.ID = 0

	CellIntID: NullInt64 = new NullInt64 // if pointer is null, CellInt.ID = 0

	CellBoolID: NullInt64 = new NullInt64 // if pointer is null, CellBool.ID = 0

	CellIconID: NullInt64 = new NullInt64 // if pointer is null, CellIcon.ID = 0

}
