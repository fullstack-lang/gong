// generated code - do not edit

import { GongStructDB } from './gongstruct-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongBasicField } from './gongbasicfield'
import { GongTimeField } from './gongtimefield'
import { PointerToGongStructField } from './pointertogongstructfield'
import { SliceOfPointerToGongStructField } from './sliceofpointertogongstructfield'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongStruct {

	static GONGSTRUCT_NAME = "GongStruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	HasOnAfterUpdateSignature: boolean = false
	IsIgnoredForFront: boolean = false

	// insertion point for pointers and slices of pointers declarations
	GongBasicFields: Array<GongBasicField> = []
	GongTimeFields: Array<GongTimeField> = []
	PointerToGongStructFields: Array<PointerToGongStructField> = []
	SliceOfPointerToGongStructFields: Array<SliceOfPointerToGongStructField> = []
}

export function CopyGongStructToGongStructDB(gongstruct: GongStruct, gongstructDB: GongStructDB) {

	gongstructDB.CreatedAt = gongstruct.CreatedAt
	gongstructDB.DeletedAt = gongstruct.DeletedAt
	gongstructDB.ID = gongstruct.ID

	// insertion point for basic fields copy operations
	gongstructDB.Name = gongstruct.Name
	gongstructDB.HasOnAfterUpdateSignature = gongstruct.HasOnAfterUpdateSignature
	gongstructDB.IsIgnoredForFront = gongstruct.IsIgnoredForFront

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongstructDB.GongStructPointersEncoding.GongBasicFields = []
	for (let _gongbasicfield of gongstruct.GongBasicFields) {
		gongstructDB.GongStructPointersEncoding.GongBasicFields.push(_gongbasicfield.ID)
	}

	gongstructDB.GongStructPointersEncoding.GongTimeFields = []
	for (let _gongtimefield of gongstruct.GongTimeFields) {
		gongstructDB.GongStructPointersEncoding.GongTimeFields.push(_gongtimefield.ID)
	}

	gongstructDB.GongStructPointersEncoding.PointerToGongStructFields = []
	for (let _pointertogongstructfield of gongstruct.PointerToGongStructFields) {
		gongstructDB.GongStructPointersEncoding.PointerToGongStructFields.push(_pointertogongstructfield.ID)
	}

	gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields = []
	for (let _sliceofpointertogongstructfield of gongstruct.SliceOfPointerToGongStructFields) {
		gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield.ID)
	}

}

// CopyGongStructDBToGongStruct update basic, pointers and slice of pointers fields of gongstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongstructDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongStructDBToGongStruct(gongstructDB: GongStructDB, gongstruct: GongStruct, frontRepo: FrontRepo) {

	gongstruct.CreatedAt = gongstructDB.CreatedAt
	gongstruct.DeletedAt = gongstructDB.DeletedAt
	gongstruct.ID = gongstructDB.ID

	// insertion point for basic fields copy operations
	gongstruct.Name = gongstructDB.Name
	gongstruct.HasOnAfterUpdateSignature = gongstructDB.HasOnAfterUpdateSignature
	gongstruct.IsIgnoredForFront = gongstructDB.IsIgnoredForFront

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongstruct.GongBasicFields = new Array<GongBasicField>()
	for (let _id of gongstructDB.GongStructPointersEncoding.GongBasicFields) {
		let _gongbasicfield = frontRepo.map_ID_GongBasicField.get(_id)
		if (_gongbasicfield != undefined) {
			gongstruct.GongBasicFields.push(_gongbasicfield!)
		}
	}
	gongstruct.GongTimeFields = new Array<GongTimeField>()
	for (let _id of gongstructDB.GongStructPointersEncoding.GongTimeFields) {
		let _gongtimefield = frontRepo.map_ID_GongTimeField.get(_id)
		if (_gongtimefield != undefined) {
			gongstruct.GongTimeFields.push(_gongtimefield!)
		}
	}
	gongstruct.PointerToGongStructFields = new Array<PointerToGongStructField>()
	for (let _id of gongstructDB.GongStructPointersEncoding.PointerToGongStructFields) {
		let _pointertogongstructfield = frontRepo.map_ID_PointerToGongStructField.get(_id)
		if (_pointertogongstructfield != undefined) {
			gongstruct.PointerToGongStructFields.push(_pointertogongstructfield!)
		}
	}
	gongstruct.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructField>()
	for (let _id of gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields) {
		let _sliceofpointertogongstructfield = frontRepo.map_ID_SliceOfPointerToGongStructField.get(_id)
		if (_sliceofpointertogongstructfield != undefined) {
			gongstruct.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield!)
		}
	}
}
