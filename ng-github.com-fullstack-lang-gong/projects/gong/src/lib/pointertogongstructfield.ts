// generated code - do not edit

import { PointerToGongStructFieldAPI } from './pointertogongstructfield-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongStruct } from './gongstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PointerToGongStructField {

	static GONGSTRUCT_NAME = "PointerToGongStructField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""
	IsType: boolean = false

	// insertion point for pointers and slices of pointers declarations
	GongStruct?: GongStruct

}

export function CopyPointerToGongStructFieldToPointerToGongStructFieldAPI(pointertogongstructfield: PointerToGongStructField, pointertogongstructfieldAPI: PointerToGongStructFieldAPI) {

	pointertogongstructfieldAPI.CreatedAt = pointertogongstructfield.CreatedAt
	pointertogongstructfieldAPI.DeletedAt = pointertogongstructfield.DeletedAt
	pointertogongstructfieldAPI.ID = pointertogongstructfield.ID

	// insertion point for basic fields copy operations
	pointertogongstructfieldAPI.Name = pointertogongstructfield.Name
	pointertogongstructfieldAPI.Index = pointertogongstructfield.Index
	pointertogongstructfieldAPI.CompositeStructName = pointertogongstructfield.CompositeStructName
	pointertogongstructfieldAPI.IsType = pointertogongstructfield.IsType

	// insertion point for pointer fields encoding
	pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding.GongStructID.Valid = true
	if (pointertogongstructfield.GongStruct != undefined) {
		pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding.GongStructID.Int64 = pointertogongstructfield.GongStruct.ID  
	} else {
		pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding.GongStructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyPointerToGongStructFieldAPIToPointerToGongStructField update basic, pointers and slice of pointers fields of pointertogongstructfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pointertogongstructfieldAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPointerToGongStructFieldAPIToPointerToGongStructField(pointertogongstructfieldAPI: PointerToGongStructFieldAPI, pointertogongstructfield: PointerToGongStructField, frontRepo: FrontRepo) {

	pointertogongstructfield.CreatedAt = pointertogongstructfieldAPI.CreatedAt
	pointertogongstructfield.DeletedAt = pointertogongstructfieldAPI.DeletedAt
	pointertogongstructfield.ID = pointertogongstructfieldAPI.ID

	// insertion point for basic fields copy operations
	pointertogongstructfield.Name = pointertogongstructfieldAPI.Name
	pointertogongstructfield.Index = pointertogongstructfieldAPI.Index
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldAPI.CompositeStructName
	pointertogongstructfield.IsType = pointertogongstructfieldAPI.IsType

	// insertion point for pointer fields encoding
	pointertogongstructfield.GongStruct = frontRepo.map_ID_GongStruct.get(pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding.GongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
