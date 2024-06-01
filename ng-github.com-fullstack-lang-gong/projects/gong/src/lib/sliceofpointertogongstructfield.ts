// generated code - do not edit

import { SliceOfPointerToGongStructFieldAPI } from './sliceofpointertogongstructfield-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongStruct } from './gongstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliceOfPointerToGongStructField {

	static GONGSTRUCT_NAME = "SliceOfPointerToGongStructField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for pointers and slices of pointers declarations
	GongStruct?: GongStruct

}

export function CopySliceOfPointerToGongStructFieldToSliceOfPointerToGongStructFieldAPI(sliceofpointertogongstructfield: SliceOfPointerToGongStructField, sliceofpointertogongstructfieldAPI: SliceOfPointerToGongStructFieldAPI) {

	sliceofpointertogongstructfieldAPI.CreatedAt = sliceofpointertogongstructfield.CreatedAt
	sliceofpointertogongstructfieldAPI.DeletedAt = sliceofpointertogongstructfield.DeletedAt
	sliceofpointertogongstructfieldAPI.ID = sliceofpointertogongstructfield.ID

	// insertion point for basic fields copy operations
	sliceofpointertogongstructfieldAPI.Name = sliceofpointertogongstructfield.Name
	sliceofpointertogongstructfieldAPI.Index = sliceofpointertogongstructfield.Index
	sliceofpointertogongstructfieldAPI.CompositeStructName = sliceofpointertogongstructfield.CompositeStructName

	// insertion point for pointer fields encoding
	sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Valid = true
	if (sliceofpointertogongstructfield.GongStruct != undefined) {
		sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64 = sliceofpointertogongstructfield.GongStruct.ID  
	} else {
		sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySliceOfPointerToGongStructFieldAPIToSliceOfPointerToGongStructField update basic, pointers and slice of pointers fields of sliceofpointertogongstructfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of sliceofpointertogongstructfieldAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySliceOfPointerToGongStructFieldAPIToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldAPI: SliceOfPointerToGongStructFieldAPI, sliceofpointertogongstructfield: SliceOfPointerToGongStructField, frontRepo: FrontRepo) {

	sliceofpointertogongstructfield.CreatedAt = sliceofpointertogongstructfieldAPI.CreatedAt
	sliceofpointertogongstructfield.DeletedAt = sliceofpointertogongstructfieldAPI.DeletedAt
	sliceofpointertogongstructfield.ID = sliceofpointertogongstructfieldAPI.ID

	// insertion point for basic fields copy operations
	sliceofpointertogongstructfield.Name = sliceofpointertogongstructfieldAPI.Name
	sliceofpointertogongstructfield.Index = sliceofpointertogongstructfieldAPI.Index
	sliceofpointertogongstructfield.CompositeStructName = sliceofpointertogongstructfieldAPI.CompositeStructName

	// insertion point for pointer fields encoding
	sliceofpointertogongstructfield.GongStruct = frontRepo.map_ID_GongStruct.get(sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
