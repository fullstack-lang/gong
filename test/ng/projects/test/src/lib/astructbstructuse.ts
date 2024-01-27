// generated code - do not edit

import { AstructBstructUseDB } from './astructbstructuse-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Bstruct } from './bstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstructUse {

	static GONGSTRUCT_NAME = "AstructBstructUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Bstruct2?: Bstruct

}

export function CopyAstructBstructUseToAstructBstructUseDB(astructbstructuse: AstructBstructUse, astructbstructuseDB: AstructBstructUseDB) {

	astructbstructuseDB.CreatedAt = astructbstructuse.CreatedAt
	astructbstructuseDB.DeletedAt = astructbstructuse.DeletedAt
	astructbstructuseDB.ID = astructbstructuse.ID

	// insertion point for basic fields copy operations
	astructbstructuseDB.Name = astructbstructuse.Name

	// insertion point for pointer fields encoding
	astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Valid = true
	if (astructbstructuse.Bstruct2 != undefined) {
		astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = astructbstructuse.Bstruct2.ID  
	} else {
		astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyAstructBstructUseDBToAstructBstructUse update basic, pointers and slice of pointers fields of astructbstructuse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of astructbstructuseDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB: AstructBstructUseDB, astructbstructuse: AstructBstructUse, frontRepo: FrontRepo) {

	astructbstructuse.CreatedAt = astructbstructuseDB.CreatedAt
	astructbstructuse.DeletedAt = astructbstructuseDB.DeletedAt
	astructbstructuse.ID = astructbstructuseDB.ID

	// insertion point for basic fields copy operations
	astructbstructuse.Name = astructbstructuseDB.Name

	// insertion point for pointer fields encoding
	astructbstructuse.Bstruct2 = frontRepo.map_ID_Bstruct.get(astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64)

	// insertion point for slice of pointers fields encoding
}
