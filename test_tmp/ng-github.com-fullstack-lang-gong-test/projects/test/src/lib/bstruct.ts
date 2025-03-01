// generated code - do not edit

import { BstructAPI } from './bstruct-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Bstruct {

	static GONGSTRUCT_NAME = "Bstruct"

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

export function CopyBstructToBstructAPI(bstruct: Bstruct, bstructAPI: BstructAPI) {

	bstructAPI.CreatedAt = bstruct.CreatedAt
	bstructAPI.DeletedAt = bstruct.DeletedAt
	bstructAPI.ID = bstruct.ID

	// insertion point for basic fields copy operations
	bstructAPI.Name = bstruct.Name
	bstructAPI.Floatfield = bstruct.Floatfield
	bstructAPI.Floatfield2 = bstruct.Floatfield2
	bstructAPI.Intfield = bstruct.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBstructAPIToBstruct update basic, pointers and slice of pointers fields of bstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of bstructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBstructAPIToBstruct(bstructAPI: BstructAPI, bstruct: Bstruct, frontRepo: FrontRepo) {

	bstruct.CreatedAt = bstructAPI.CreatedAt
	bstruct.DeletedAt = bstructAPI.DeletedAt
	bstruct.ID = bstructAPI.ID

	// insertion point for basic fields copy operations
	bstruct.Name = bstructAPI.Name
	bstruct.Floatfield = bstructAPI.Floatfield
	bstruct.Floatfield2 = bstructAPI.Floatfield2
	bstruct.Intfield = bstructAPI.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
