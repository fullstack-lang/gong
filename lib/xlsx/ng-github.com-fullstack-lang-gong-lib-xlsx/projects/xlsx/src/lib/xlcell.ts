// generated code - do not edit

import { XLCellAPI } from './xlcell-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLCell {

	static GONGSTRUCT_NAME = "XLCell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyXLCellToXLCellAPI(xlcell: XLCell, xlcellAPI: XLCellAPI) {

	xlcellAPI.CreatedAt = xlcell.CreatedAt
	xlcellAPI.DeletedAt = xlcell.DeletedAt
	xlcellAPI.ID = xlcell.ID

	// insertion point for basic fields copy operations
	xlcellAPI.Name = xlcell.Name
	xlcellAPI.X = xlcell.X
	xlcellAPI.Y = xlcell.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyXLCellAPIToXLCell update basic, pointers and slice of pointers fields of xlcell
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlcellAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLCellAPIToXLCell(xlcellAPI: XLCellAPI, xlcell: XLCell, frontRepo: FrontRepo) {

	xlcell.CreatedAt = xlcellAPI.CreatedAt
	xlcell.DeletedAt = xlcellAPI.DeletedAt
	xlcell.ID = xlcellAPI.ID

	// insertion point for basic fields copy operations
	xlcell.Name = xlcellAPI.Name
	xlcell.X = xlcellAPI.X
	xlcell.Y = xlcellAPI.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
