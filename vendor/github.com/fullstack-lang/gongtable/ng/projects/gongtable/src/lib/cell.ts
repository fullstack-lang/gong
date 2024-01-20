// generated code - do not edit

import { CellDB } from './cell-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { CellString } from './cellstring'
import { CellFloat64 } from './cellfloat64'
import { CellInt } from './cellint'
import { CellBoolean } from './cellboolean'
import { CellIcon } from './cellicon'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Cell {

	static GONGSTRUCT_NAME = "Cell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	CellString?: CellString

	CellFloat64?: CellFloat64

	CellInt?: CellInt

	CellBool?: CellBoolean

	CellIcon?: CellIcon

}

export function CopyCellToCellDB(cell: Cell, cellDB: CellDB) {

	cellDB.CreatedAt = cell.CreatedAt
	cellDB.DeletedAt = cell.DeletedAt
	cellDB.ID = cell.ID
	
	// insertion point for basic fields copy operations
	cellDB.Name = cell.Name

	// insertion point for pointer fields encoding
    cellDB.CellPointersEncoding.CellStringID.Valid = true
	if (cell.CellString != undefined) {
		cellDB.CellPointersEncoding.CellStringID.Int64 = cell.CellString.ID  
    } else {
		cellDB.CellPointersEncoding.CellStringID.Int64 = 0 		
	}

    cellDB.CellPointersEncoding.CellFloat64ID.Valid = true
	if (cell.CellFloat64 != undefined) {
		cellDB.CellPointersEncoding.CellFloat64ID.Int64 = cell.CellFloat64.ID  
    } else {
		cellDB.CellPointersEncoding.CellFloat64ID.Int64 = 0 		
	}

    cellDB.CellPointersEncoding.CellIntID.Valid = true
	if (cell.CellInt != undefined) {
		cellDB.CellPointersEncoding.CellIntID.Int64 = cell.CellInt.ID  
    } else {
		cellDB.CellPointersEncoding.CellIntID.Int64 = 0 		
	}

    cellDB.CellPointersEncoding.CellBoolID.Valid = true
	if (cell.CellBool != undefined) {
		cellDB.CellPointersEncoding.CellBoolID.Int64 = cell.CellBool.ID  
    } else {
		cellDB.CellPointersEncoding.CellBoolID.Int64 = 0 		
	}

    cellDB.CellPointersEncoding.CellIconID.Valid = true
	if (cell.CellIcon != undefined) {
		cellDB.CellPointersEncoding.CellIconID.Int64 = cell.CellIcon.ID  
    } else {
		cellDB.CellPointersEncoding.CellIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyCellDBToCell(cellDB: CellDB, cell: Cell, frontRepo: FrontRepo) {

	cell.CreatedAt = cellDB.CreatedAt
	cell.DeletedAt = cellDB.DeletedAt
	cell.ID = cellDB.ID
	
	// insertion point for basic fields copy operations
	cell.Name = cellDB.Name

	// insertion point for pointer fields encoding
	cell.CellString = frontRepo.CellStrings.get(cellDB.CellPointersEncoding.CellStringID.Int64)
	cell.CellFloat64 = frontRepo.CellFloat64s.get(cellDB.CellPointersEncoding.CellFloat64ID.Int64)
	cell.CellInt = frontRepo.CellInts.get(cellDB.CellPointersEncoding.CellIntID.Int64)
	cell.CellBool = frontRepo.CellBooleans.get(cellDB.CellPointersEncoding.CellBoolID.Int64)
	cell.CellIcon = frontRepo.CellIcons.get(cellDB.CellPointersEncoding.CellIconID.Int64)

	// insertion point for slice of pointers fields encoding
}
