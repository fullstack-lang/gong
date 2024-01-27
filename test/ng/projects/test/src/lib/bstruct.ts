// generated code - do not edit

import { BstructDB } from './bstruct-db'
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

export function CopyBstructToBstructDB(bstruct: Bstruct, bstructDB: BstructDB) {

	bstructDB.CreatedAt = bstruct.CreatedAt
	bstructDB.DeletedAt = bstruct.DeletedAt
	bstructDB.ID = bstruct.ID

	// insertion point for basic fields copy operations
	bstructDB.Name = bstruct.Name
	bstructDB.Floatfield = bstruct.Floatfield
	bstructDB.Floatfield2 = bstruct.Floatfield2
	bstructDB.Intfield = bstruct.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBstructDBToBstruct update basic, pointers and slice of pointers fields of bstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of bstructDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBstructDBToBstruct(bstructDB: BstructDB, bstruct: Bstruct, frontRepo: FrontRepo) {

	bstruct.CreatedAt = bstructDB.CreatedAt
	bstruct.DeletedAt = bstructDB.DeletedAt
	bstruct.ID = bstructDB.ID

	// insertion point for basic fields copy operations
	bstruct.Name = bstructDB.Name
	bstruct.Floatfield = bstructDB.Floatfield
	bstruct.Floatfield2 = bstructDB.Floatfield2
	bstruct.Intfield = bstructDB.Intfield

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
