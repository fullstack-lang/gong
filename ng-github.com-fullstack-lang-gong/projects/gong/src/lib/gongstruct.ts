// generated code - do not edit

import { GongStructAPI } from './gongstruct-api'
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

export function CopyGongStructToGongStructAPI(gongstruct: GongStruct, gongstructAPI: GongStructAPI) {

	gongstructAPI.CreatedAt = gongstruct.CreatedAt
	gongstructAPI.DeletedAt = gongstruct.DeletedAt
	gongstructAPI.ID = gongstruct.ID

	// insertion point for basic fields copy operations
	gongstructAPI.Name = gongstruct.Name
	gongstructAPI.HasOnAfterUpdateSignature = gongstruct.HasOnAfterUpdateSignature
	gongstructAPI.IsIgnoredForFront = gongstruct.IsIgnoredForFront

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongstructAPI.GongStructPointersEncoding.GongBasicFields = []
	for (let _gongbasicfield of gongstruct.GongBasicFields) {
		gongstructAPI.GongStructPointersEncoding.GongBasicFields.push(_gongbasicfield.ID)
	}

	gongstructAPI.GongStructPointersEncoding.GongTimeFields = []
	for (let _gongtimefield of gongstruct.GongTimeFields) {
		gongstructAPI.GongStructPointersEncoding.GongTimeFields.push(_gongtimefield.ID)
	}

	gongstructAPI.GongStructPointersEncoding.PointerToGongStructFields = []
	for (let _pointertogongstructfield of gongstruct.PointerToGongStructFields) {
		gongstructAPI.GongStructPointersEncoding.PointerToGongStructFields.push(_pointertogongstructfield.ID)
	}

	gongstructAPI.GongStructPointersEncoding.SliceOfPointerToGongStructFields = []
	for (let _sliceofpointertogongstructfield of gongstruct.SliceOfPointerToGongStructFields) {
		gongstructAPI.GongStructPointersEncoding.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield.ID)
	}

}

// CopyGongStructAPIToGongStruct update basic, pointers and slice of pointers fields of gongstruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongstructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongStructAPIToGongStruct(gongstructAPI: GongStructAPI, gongstruct: GongStruct, frontRepo: FrontRepo) {

	gongstruct.CreatedAt = gongstructAPI.CreatedAt
	gongstruct.DeletedAt = gongstructAPI.DeletedAt
	gongstruct.ID = gongstructAPI.ID

	// insertion point for basic fields copy operations
	gongstruct.Name = gongstructAPI.Name
	gongstruct.HasOnAfterUpdateSignature = gongstructAPI.HasOnAfterUpdateSignature
	gongstruct.IsIgnoredForFront = gongstructAPI.IsIgnoredForFront

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(gongstructAPI.GongStructPointersEncoding.GongBasicFields)) {
		console.error('Rects is not an array:', gongstructAPI.GongStructPointersEncoding.GongBasicFields);
		return;
	}

	gongstruct.GongBasicFields = new Array<GongBasicField>()
	for (let _id of gongstructAPI.GongStructPointersEncoding.GongBasicFields) {
		let _gongbasicfield = frontRepo.map_ID_GongBasicField.get(_id)
		if (_gongbasicfield != undefined) {
			gongstruct.GongBasicFields.push(_gongbasicfield!)
		}
	}
	if (!Array.isArray(gongstructAPI.GongStructPointersEncoding.GongTimeFields)) {
		console.error('Rects is not an array:', gongstructAPI.GongStructPointersEncoding.GongTimeFields);
		return;
	}

	gongstruct.GongTimeFields = new Array<GongTimeField>()
	for (let _id of gongstructAPI.GongStructPointersEncoding.GongTimeFields) {
		let _gongtimefield = frontRepo.map_ID_GongTimeField.get(_id)
		if (_gongtimefield != undefined) {
			gongstruct.GongTimeFields.push(_gongtimefield!)
		}
	}
	if (!Array.isArray(gongstructAPI.GongStructPointersEncoding.PointerToGongStructFields)) {
		console.error('Rects is not an array:', gongstructAPI.GongStructPointersEncoding.PointerToGongStructFields);
		return;
	}

	gongstruct.PointerToGongStructFields = new Array<PointerToGongStructField>()
	for (let _id of gongstructAPI.GongStructPointersEncoding.PointerToGongStructFields) {
		let _pointertogongstructfield = frontRepo.map_ID_PointerToGongStructField.get(_id)
		if (_pointertogongstructfield != undefined) {
			gongstruct.PointerToGongStructFields.push(_pointertogongstructfield!)
		}
	}
	if (!Array.isArray(gongstructAPI.GongStructPointersEncoding.SliceOfPointerToGongStructFields)) {
		console.error('Rects is not an array:', gongstructAPI.GongStructPointersEncoding.SliceOfPointerToGongStructFields);
		return;
	}

	gongstruct.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructField>()
	for (let _id of gongstructAPI.GongStructPointersEncoding.SliceOfPointerToGongStructFields) {
		let _sliceofpointertogongstructfield = frontRepo.map_ID_SliceOfPointerToGongStructField.get(_id)
		if (_sliceofpointertogongstructfield != undefined) {
			gongstruct.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield!)
		}
	}
}
