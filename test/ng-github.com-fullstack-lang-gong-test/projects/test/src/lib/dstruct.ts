// generated code - do not edit

import { DstructAPI } from './dstruct-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Bstruct } from './bstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Dstruct {

	static GONGSTRUCT_NAME = "Dstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Anarrayofb: Array<Bstruct> = []
}

export function CopyDstructToDstructAPI(dstruct: Dstruct, dstructAPI: DstructAPI) {

	dstructAPI.CreatedAt = dstruct.CreatedAt
	dstructAPI.DeletedAt = dstruct.DeletedAt
	dstructAPI.ID = dstruct.ID

	// insertion point for basic fields copy operations
	dstructAPI.Name = dstruct.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	dstructAPI.DstructPointersEncoding.Anarrayofb = []
	for (let _bstruct of dstruct.Anarrayofb) {
		dstructAPI.DstructPointersEncoding.Anarrayofb.push(_bstruct.ID)
	}

}

// CopyDstructAPIToDstruct update basic, pointers and slice of pointers fields of dstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of dstructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDstructAPIToDstruct(dstructAPI: DstructAPI, dstruct: Dstruct, frontRepo: FrontRepo) {

	dstruct.CreatedAt = dstructAPI.CreatedAt
	dstruct.DeletedAt = dstructAPI.DeletedAt
	dstruct.ID = dstructAPI.ID

	// insertion point for basic fields copy operations
	dstruct.Name = dstructAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	dstruct.Anarrayofb = new Array<Bstruct>()
	for (let _id of dstructAPI.DstructPointersEncoding.Anarrayofb) {
		let _bstruct = frontRepo.map_ID_Bstruct.get(_id)
		if (_bstruct != undefined) {
			dstruct.Anarrayofb.push(_bstruct!)
		}
	}
}
