// generated code - do not edit

import { AAPI } from './a-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class A {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyAToAAPI(a: A, aAPI: AAPI) {

	aAPI.CreatedAt = a.CreatedAt
	aAPI.DeletedAt = a.DeletedAt
	aAPI.ID = a.ID

	// insertion point for basic fields copy operations
	aAPI.Name = a.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAAPIToA update basic, pointers and slice of pointers fields of a
// from respectively the basic fields and encoded fields of pointers and slices of pointers of aAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAAPIToA(aAPI: AAPI, a: A, frontRepo: FrontRepo) {

	a.CreatedAt = aAPI.CreatedAt
	a.DeletedAt = aAPI.DeletedAt
	a.ID = aAPI.ID

	// insertion point for basic fields copy operations
	a.Name = aAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
