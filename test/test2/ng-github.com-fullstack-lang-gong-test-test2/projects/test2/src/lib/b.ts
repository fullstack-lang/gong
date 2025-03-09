// generated code - do not edit

import { BAPI } from './b-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class B {

	static GONGSTRUCT_NAME = "B"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyBToBAPI(b: B, bAPI: BAPI) {

	bAPI.CreatedAt = b.CreatedAt
	bAPI.DeletedAt = b.DeletedAt
	bAPI.ID = b.ID

	// insertion point for basic fields copy operations
	bAPI.Name = b.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBAPIToB update basic, pointers and slice of pointers fields of b
// from respectively the basic fields and encoded fields of pointers and slices of pointers of bAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBAPIToB(bAPI: BAPI, b: B, frontRepo: FrontRepo) {

	b.CreatedAt = bAPI.CreatedAt
	b.DeletedAt = bAPI.DeletedAt
	b.ID = bAPI.ID

	// insertion point for basic fields copy operations
	b.Name = bAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
