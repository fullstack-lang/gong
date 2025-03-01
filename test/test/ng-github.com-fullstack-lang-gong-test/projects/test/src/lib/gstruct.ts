// generated code - do not edit

import { GstructAPI } from './gstruct-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Gstruct {

	static GONGSTRUCT_NAME = "Gstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGstructToGstructAPI(gstruct: Gstruct, gstructAPI: GstructAPI) {

	gstructAPI.CreatedAt = gstruct.CreatedAt
	gstructAPI.DeletedAt = gstruct.DeletedAt
	gstructAPI.ID = gstruct.ID

	// insertion point for basic fields copy operations
	gstructAPI.Name = gstruct.Name
	gstructAPI.Floatfield = gstruct.Floatfield
	gstructAPI.Floatfield2 = gstruct.Floatfield2
	gstructAPI.Intfield = gstruct.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGstructAPIToGstruct update basic, pointers and slice of pointers fields of gstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gstructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGstructAPIToGstruct(gstructAPI: GstructAPI, gstruct: Gstruct, frontRepo: FrontRepo) {

	gstruct.CreatedAt = gstructAPI.CreatedAt
	gstruct.DeletedAt = gstructAPI.DeletedAt
	gstruct.ID = gstructAPI.ID

	// insertion point for basic fields copy operations
	gstruct.Name = gstructAPI.Name
	gstruct.Floatfield = gstructAPI.Floatfield
	gstruct.Floatfield2 = gstructAPI.Floatfield2
	gstruct.Intfield = gstructAPI.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
