// generated code - do not edit

import { DstructAPI } from './dstruct-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Bstruct } from './bstruct'
import { Gstruct } from './gstruct'

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
	Gstruct?: Gstruct

	Gstructs: Array<Gstruct> = []
}

export function CopyDstructToDstructAPI(dstruct: Dstruct, dstructAPI: DstructAPI) {

	dstructAPI.CreatedAt = dstruct.CreatedAt
	dstructAPI.DeletedAt = dstruct.DeletedAt
	dstructAPI.ID = dstruct.ID

	// insertion point for basic fields copy operations
	dstructAPI.Name = dstruct.Name

	// insertion point for pointer fields encoding
	dstructAPI.DstructPointersEncoding.GstructID.Valid = true
	if (dstruct.Gstruct != undefined) {
		dstructAPI.DstructPointersEncoding.GstructID.Int64 = dstruct.Gstruct.ID  
	} else {
		dstructAPI.DstructPointersEncoding.GstructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	dstructAPI.DstructPointersEncoding.Anarrayofb = []
	for (let _bstruct of dstruct.Anarrayofb) {
		dstructAPI.DstructPointersEncoding.Anarrayofb.push(_bstruct.ID)
	}

	dstructAPI.DstructPointersEncoding.Gstructs = []
	for (let _gstruct of dstruct.Gstructs) {
		dstructAPI.DstructPointersEncoding.Gstructs.push(_gstruct.ID)
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
	dstruct.Gstruct = frontRepo.map_ID_Gstruct.get(dstructAPI.DstructPointersEncoding.GstructID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(dstructAPI.DstructPointersEncoding.Anarrayofb)) {
		console.error('Rects is not an array:', dstructAPI.DstructPointersEncoding.Anarrayofb);
		return;
	}

	dstruct.Anarrayofb = new Array<Bstruct>()
	for (let _id of dstructAPI.DstructPointersEncoding.Anarrayofb) {
		let _bstruct = frontRepo.map_ID_Bstruct.get(_id)
		if (_bstruct != undefined) {
			dstruct.Anarrayofb.push(_bstruct!)
		}
	}
	if (!Array.isArray(dstructAPI.DstructPointersEncoding.Gstructs)) {
		console.error('Rects is not an array:', dstructAPI.DstructPointersEncoding.Gstructs);
		return;
	}

	dstruct.Gstructs = new Array<Gstruct>()
	for (let _id of dstructAPI.DstructPointersEncoding.Gstructs) {
		let _gstruct = frontRepo.map_ID_Gstruct.get(_id)
		if (_gstruct != undefined) {
			dstruct.Gstructs.push(_gstruct!)
		}
	}
}
