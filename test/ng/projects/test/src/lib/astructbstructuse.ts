// generated code - do not edit

import { AstructBstructUseDB } from './astructbstructuse-db'

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

	// insertion point for basic fields copy operations
	astructbstructuseDB.Name = astructbstructuse.Name
	
	// insertion point for pointer fields encoding
    astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Valid = true
	if (astructbstructuse.Bstruct2 != undefined) {
      astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = astructbstructuse.Bstruct2.ID  
    } else {
      astructbstructuseDB.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = 0 		
	}

}
