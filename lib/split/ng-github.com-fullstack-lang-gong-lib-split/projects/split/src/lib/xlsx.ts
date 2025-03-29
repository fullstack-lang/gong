// generated code - do not edit

import { XlsxAPI } from './xlsx-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Xlsx {

	static GONGSTRUCT_NAME = "Xlsx"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyXlsxToXlsxAPI(xlsx: Xlsx, xlsxAPI: XlsxAPI) {

	xlsxAPI.CreatedAt = xlsx.CreatedAt
	xlsxAPI.DeletedAt = xlsx.DeletedAt
	xlsxAPI.ID = xlsx.ID

	// insertion point for basic fields copy operations
	xlsxAPI.Name = xlsx.Name
	xlsxAPI.StackName = xlsx.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyXlsxAPIToXlsx update basic, pointers and slice of pointers fields of xlsx
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlsxAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXlsxAPIToXlsx(xlsxAPI: XlsxAPI, xlsx: Xlsx, frontRepo: FrontRepo) {

	xlsx.CreatedAt = xlsxAPI.CreatedAt
	xlsx.DeletedAt = xlsxAPI.DeletedAt
	xlsx.ID = xlsxAPI.ID

	// insertion point for basic fields copy operations
	xlsx.Name = xlsxAPI.Name
	xlsx.StackName = xlsxAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
