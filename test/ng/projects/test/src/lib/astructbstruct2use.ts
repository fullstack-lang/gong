// generated code - do not edit

import { AstructBstruct2UseAPI } from './astructbstruct2use-api'
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

export function CopyAstructBstruct2UseToAstructBstruct2UseAPI(astructbstruct2use: AstructBstruct2Use, astructbstruct2useAPI: AstructBstruct2UseAPI) {

	astructbstruct2useAPI.CreatedAt = astructbstruct2use.CreatedAt
	astructbstruct2useAPI.DeletedAt = astructbstruct2use.DeletedAt
	astructbstruct2useAPI.ID = astructbstruct2use.ID

	// insertion point for basic fields copy operations
	astructbstruct2useAPI.Name = astructbstruct2use.Name

	// insertion point for pointer fields encoding
	astructbstruct2useAPI.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Valid = true
	if (astructbstruct2use.Bstrcut2 != undefined) {
		astructbstruct2useAPI.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64 = astructbstruct2use.Bstrcut2.ID  
	} else {
		astructbstruct2useAPI.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyAstructBstruct2UseAPIToAstructBstruct2Use update basic, pointers and slice of pointers fields of astructbstruct2use
// from respectively the basic fields and encoded fields of pointers and slices of pointers of astructbstruct2useAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAstructBstruct2UseAPIToAstructBstruct2Use(astructbstruct2useAPI: AstructBstruct2UseAPI, astructbstruct2use: AstructBstruct2Use, frontRepo: FrontRepo) {

	astructbstruct2use.CreatedAt = astructbstruct2useAPI.CreatedAt
	astructbstruct2use.DeletedAt = astructbstruct2useAPI.DeletedAt
	astructbstruct2use.ID = astructbstruct2useAPI.ID

	// insertion point for basic fields copy operations
	astructbstruct2use.Name = astructbstruct2useAPI.Name

	// insertion point for pointer fields encoding
	astructbstruct2use.Bstrcut2 = frontRepo.map_ID_Bstruct.get(astructbstruct2useAPI.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64)

	// insertion point for slice of pointers fields encoding
}
