// generated code - do not edit

import { DstructDB } from './dstruct-db'
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

export function CopyDstructToDstructDB(dstruct: Dstruct, dstructDB: DstructDB) {

	// insertion point for basic fields copy operations
	dstructDB.Name = dstruct.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	dstructDB.DstructPointersEncoding.Anarrayofb = []
    for (let _bstruct of dstruct.Anarrayofb) {
		dstructDB.DstructPointersEncoding.Anarrayofb.push(_bstruct.ID)
    }
	
}

export function CopyDstructDBToDstruct(dstructDB: DstructDB, dstruct: Dstruct, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations
	dstruct.Name = dstructDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	dstruct.Anarrayofb = new Array<Bstruct>()
	for (let _id of dstructDB.DstructPointersEncoding.Anarrayofb) {
	  let _bstruct = frontRepo.Bstructs.get(_id)
	  if (_bstruct != undefined) {
		dstruct.Anarrayofb.push(_bstruct!)
	  }
	}
}
