// generated code - do not edit

import { AstructBstruct2UseDB } from './astructbstruct2use-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Bstruct } from './bstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstruct2Use {

	static GONGSTRUCT_NAME = "AstructBstruct2Use"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Bstrcut2?: Bstruct

}

export function CopyAstructBstruct2UseToAstructBstruct2UseDB(astructbstruct2use: AstructBstruct2Use, astructbstruct2useDB: AstructBstruct2UseDB) {

	astructbstruct2useDB.CreatedAt = astructbstruct2use.CreatedAt
	astructbstruct2useDB.DeletedAt = astructbstruct2use.DeletedAt
	astructbstruct2useDB.ID = astructbstruct2use.ID
	
	// insertion point for basic fields copy operations
	astructbstruct2useDB.Name = astructbstruct2use.Name

	// insertion point for pointer fields encoding
    astructbstruct2useDB.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Valid = true
	if (astructbstruct2use.Bstrcut2 != undefined) {
		astructbstruct2useDB.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64 = astructbstruct2use.Bstrcut2.ID  
    } else {
		astructbstruct2useDB.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyAstructBstruct2UseDBToAstructBstruct2Use(astructbstruct2useDB: AstructBstruct2UseDB, astructbstruct2use: AstructBstruct2Use, frontRepo: FrontRepo) {

	astructbstruct2use.CreatedAt = astructbstruct2useDB.CreatedAt
	astructbstruct2use.DeletedAt = astructbstruct2useDB.DeletedAt
	astructbstruct2use.ID = astructbstruct2useDB.ID
	
	// insertion point for basic fields copy operations
	astructbstruct2use.Name = astructbstruct2useDB.Name

	// insertion point for pointer fields encoding
	astructbstruct2use.Bstrcut2 = frontRepo.Bstructs.get(astructbstruct2useDB.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64)

	// insertion point for slice of pointers fields encoding
}
