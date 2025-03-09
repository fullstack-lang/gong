// generated code - do not edit

import { AAPI } from './a-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { B } from './b'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class A {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NumberField: number = 0

	// insertion point for pointers and slices of pointers declarations
	B?: B

	Bs: Array<B> = []
}

export function CopyAToAAPI(a: A, aAPI: AAPI) {

	aAPI.CreatedAt = a.CreatedAt
	aAPI.DeletedAt = a.DeletedAt
	aAPI.ID = a.ID

	// insertion point for basic fields copy operations
	aAPI.Name = a.Name
	aAPI.NumberField = a.NumberField

	// insertion point for pointer fields encoding
	aAPI.APointersEncoding.BID.Valid = true
	if (a.B != undefined) {
		aAPI.APointersEncoding.BID.Int64 = a.B.ID  
	} else {
		aAPI.APointersEncoding.BID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	aAPI.APointersEncoding.Bs = []
	for (let _b of a.Bs) {
		aAPI.APointersEncoding.Bs.push(_b.ID)
	}

}

// CopyAAPIToA update basic, pointers and slice of pointers fields of a
// from respectively the basic fields and encoded fields of pointers and slices of pointers of aAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAAPIToA(aAPI: AAPI, a: A, frontRepo: FrontRepo) {

	a.CreatedAt = aAPI.CreatedAt
	a.DeletedAt = aAPI.DeletedAt
	a.ID = aAPI.ID

	// insertion point for basic fields copy operations
	a.Name = aAPI.Name
	a.NumberField = aAPI.NumberField

	// insertion point for pointer fields encoding
	a.B = frontRepo.map_ID_B.get(aAPI.APointersEncoding.BID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(aAPI.APointersEncoding.Bs)) {
		console.error('Rects is not an array:', aAPI.APointersEncoding.Bs);
		return;
	}

	a.Bs = new Array<B>()
	for (let _id of aAPI.APointersEncoding.Bs) {
		let _b = frontRepo.map_ID_B.get(_id)
		if (_b != undefined) {
			a.Bs.push(_b!)
		}
	}
}
