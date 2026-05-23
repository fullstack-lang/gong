// generated code - do not edit

import { AstructAPI } from './astruct-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Astruct {

	static GONGSTRUCT_NAME = "Astruct"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyAstructToAstructAPI(astruct: Astruct, astructAPI: AstructAPI) {

	astructAPI.CreatedAt = astruct.CreatedAt
	astructAPI.DeletedAt = astruct.DeletedAt
	astructAPI.ID = astruct.ID

	// insertion point for basic fields copy operations
	astructAPI.Name = astruct.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAstructAPIToAstruct update basic, pointers and slice of pointers fields of astruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of astructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAstructAPIToAstruct(astructAPI: AstructAPI, astruct: Astruct, frontRepo: FrontRepo) {

	astruct.CreatedAt = astructAPI.CreatedAt
	astruct.DeletedAt = astructAPI.DeletedAt
	astruct.ID = astructAPI.ID

	// insertion point for basic fields copy operations
	astruct.Name = astructAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
