// generated code - do not edit

import { XLFileAPI } from './xlfile-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLSheet } from './xlsheet'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLFile {

	static GONGSTRUCT_NAME = "XLFile"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NbSheets: number = 0

	// insertion point for pointers and slices of pointers declarations
	Sheets: Array<XLSheet> = []
}

export function CopyXLFileToXLFileAPI(xlfile: XLFile, xlfileAPI: XLFileAPI) {

	xlfileAPI.CreatedAt = xlfile.CreatedAt
	xlfileAPI.DeletedAt = xlfile.DeletedAt
	xlfileAPI.ID = xlfile.ID

	// insertion point for basic fields copy operations
	xlfileAPI.Name = xlfile.Name
	xlfileAPI.NbSheets = xlfile.NbSheets

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlfileAPI.XLFilePointersEncoding.Sheets = []
	for (let _xlsheet of xlfile.Sheets) {
		xlfileAPI.XLFilePointersEncoding.Sheets.push(_xlsheet.ID)
	}

}

// CopyXLFileAPIToXLFile update basic, pointers and slice of pointers fields of xlfile
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlfileAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLFileAPIToXLFile(xlfileAPI: XLFileAPI, xlfile: XLFile, frontRepo: FrontRepo) {

	xlfile.CreatedAt = xlfileAPI.CreatedAt
	xlfile.DeletedAt = xlfileAPI.DeletedAt
	xlfile.ID = xlfileAPI.ID

	// insertion point for basic fields copy operations
	xlfile.Name = xlfileAPI.Name
	xlfile.NbSheets = xlfileAPI.NbSheets

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(xlfileAPI.XLFilePointersEncoding.Sheets)) {
		console.error('Rects is not an array:', xlfileAPI.XLFilePointersEncoding.Sheets);
		return;
	}

	xlfile.Sheets = new Array<XLSheet>()
	for (let _id of xlfileAPI.XLFilePointersEncoding.Sheets) {
		let _xlsheet = frontRepo.map_ID_XLSheet.get(_id)
		if (_xlsheet != undefined) {
			xlfile.Sheets.push(_xlsheet!)
		}
	}
}
