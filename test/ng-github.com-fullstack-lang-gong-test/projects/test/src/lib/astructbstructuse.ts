// generated code - do not edit

import { AstructBstructUseAPI } from './astructbstructuse-api'
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

export function CopyAstructBstructUseToAstructBstructUseAPI(astructbstructuse: AstructBstructUse, astructbstructuseAPI: AstructBstructUseAPI) {

	astructbstructuseAPI.CreatedAt = astructbstructuse.CreatedAt
	astructbstructuseAPI.DeletedAt = astructbstructuse.DeletedAt
	astructbstructuseAPI.ID = astructbstructuse.ID

	// insertion point for basic fields copy operations
	astructbstructuseAPI.Name = astructbstructuse.Name

	// insertion point for pointer fields encoding
	astructbstructuseAPI.AstructBstructUsePointersEncoding.Bstruct2ID.Valid = true
	if (astructbstructuse.Bstruct2 != undefined) {
		astructbstructuseAPI.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = astructbstructuse.Bstruct2.ID  
	} else {
		astructbstructuseAPI.AstructBstructUsePointersEncoding.Bstruct2ID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyAstructBstructUseAPIToAstructBstructUse update basic, pointers and slice of pointers fields of astructbstructuse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of astructbstructuseAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAstructBstructUseAPIToAstructBstructUse(astructbstructuseAPI: AstructBstructUseAPI, astructbstructuse: AstructBstructUse, frontRepo: FrontRepo) {

	astructbstructuse.CreatedAt = astructbstructuseAPI.CreatedAt
	astructbstructuse.DeletedAt = astructbstructuseAPI.DeletedAt
	astructbstructuse.ID = astructbstructuseAPI.ID

	// insertion point for basic fields copy operations
	astructbstructuse.Name = astructbstructuseAPI.Name

	// insertion point for pointer fields encoding
	astructbstructuse.Bstruct2 = frontRepo.map_ID_Bstruct.get(astructbstructuseAPI.AstructBstructUsePointersEncoding.Bstruct2ID.Int64)

	// insertion point for slice of pointers fields encoding
}
