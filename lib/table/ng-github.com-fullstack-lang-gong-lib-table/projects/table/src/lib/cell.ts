// generated code - do not edit

import { CellAPI } from './cell-api'
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

export function CopyCellToCellAPI(cell: Cell, cellAPI: CellAPI) {

	cellAPI.CreatedAt = cell.CreatedAt
	cellAPI.DeletedAt = cell.DeletedAt
	cellAPI.ID = cell.ID

	// insertion point for basic fields copy operations
	cellAPI.Name = cell.Name

	// insertion point for pointer fields encoding
	cellAPI.CellPointersEncoding.CellStringID.Valid = true
	if (cell.CellString != undefined) {
		cellAPI.CellPointersEncoding.CellStringID.Int64 = cell.CellString.ID  
	} else {
		cellAPI.CellPointersEncoding.CellStringID.Int64 = 0 		
	}

	cellAPI.CellPointersEncoding.CellFloat64ID.Valid = true
	if (cell.CellFloat64 != undefined) {
		cellAPI.CellPointersEncoding.CellFloat64ID.Int64 = cell.CellFloat64.ID  
	} else {
		cellAPI.CellPointersEncoding.CellFloat64ID.Int64 = 0 		
	}

	cellAPI.CellPointersEncoding.CellIntID.Valid = true
	if (cell.CellInt != undefined) {
		cellAPI.CellPointersEncoding.CellIntID.Int64 = cell.CellInt.ID  
	} else {
		cellAPI.CellPointersEncoding.CellIntID.Int64 = 0 		
	}

	cellAPI.CellPointersEncoding.CellBoolID.Valid = true
	if (cell.CellBool != undefined) {
		cellAPI.CellPointersEncoding.CellBoolID.Int64 = cell.CellBool.ID  
	} else {
		cellAPI.CellPointersEncoding.CellBoolID.Int64 = 0 		
	}

	cellAPI.CellPointersEncoding.CellIconID.Valid = true
	if (cell.CellIcon != undefined) {
		cellAPI.CellPointersEncoding.CellIconID.Int64 = cell.CellIcon.ID  
	} else {
		cellAPI.CellPointersEncoding.CellIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyCellAPIToCell update basic, pointers and slice of pointers fields of cell
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellAPIToCell(cellAPI: CellAPI, cell: Cell, frontRepo: FrontRepo) {

	cell.CreatedAt = cellAPI.CreatedAt
	cell.DeletedAt = cellAPI.DeletedAt
	cell.ID = cellAPI.ID

	// insertion point for basic fields copy operations
	cell.Name = cellAPI.Name

	// insertion point for pointer fields encoding
	cell.CellString = frontRepo.map_ID_CellString.get(cellAPI.CellPointersEncoding.CellStringID.Int64)
	cell.CellFloat64 = frontRepo.map_ID_CellFloat64.get(cellAPI.CellPointersEncoding.CellFloat64ID.Int64)
	cell.CellInt = frontRepo.map_ID_CellInt.get(cellAPI.CellPointersEncoding.CellIntID.Int64)
	cell.CellBool = frontRepo.map_ID_CellBoolean.get(cellAPI.CellPointersEncoding.CellBoolID.Int64)
	cell.CellIcon = frontRepo.map_ID_CellIcon.get(cellAPI.CellPointersEncoding.CellIconID.Int64)

	// insertion point for slice of pointers fields encoding
}
