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

export function CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB: AstructBstructUseDB, astructbstructuse: AstructBstructUse, frontRepo: FrontRepo) {

	astructbstructuse.ID = astructbstructuseDB.ID
	
	// insertion point for basic fields copy operations
	astructbstructuse.Name = astructbstructuseDB.Name

	// insertion point for pointer fields encoding
	astructbstructuse.Bstruct2 = frontRepo.Bstructs.get(astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64)

	// insertion point for slice of pointers fields encoding
}
