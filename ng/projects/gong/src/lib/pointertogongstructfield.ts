// generated code - do not edit

import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
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

	// insertion point for pointers and slices of pointers declarations
	GongStruct?: GongStruct

}

export function CopyPointerToGongStructFieldToPointerToGongStructFieldDB(pointertogongstructfield: PointerToGongStructField, pointertogongstructfieldDB: PointerToGongStructFieldDB) {

	pointertogongstructfieldDB.CreatedAt = pointertogongstructfield.CreatedAt
	pointertogongstructfieldDB.DeletedAt = pointertogongstructfield.DeletedAt
	pointertogongstructfieldDB.ID = pointertogongstructfield.ID

	// insertion point for basic fields copy operations
	pointertogongstructfieldDB.Name = pointertogongstructfield.Name
	pointertogongstructfieldDB.Index = pointertogongstructfield.Index
	pointertogongstructfieldDB.CompositeStructName = pointertogongstructfield.CompositeStructName

	// insertion point for pointer fields encoding
	pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding.GongStructID.Valid = true
	if (pointertogongstructfield.GongStruct != undefined) {
		pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding.GongStructID.Int64 = pointertogongstructfield.GongStruct.ID  
	} else {
		pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding.GongStructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyPointerToGongStructFieldDBToPointerToGongStructField update basic, pointers and slice of pointers fields of pointertogongstructfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pointertogongstructfieldDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPointerToGongStructFieldDBToPointerToGongStructField(pointertogongstructfieldDB: PointerToGongStructFieldDB, pointertogongstructfield: PointerToGongStructField, frontRepo: FrontRepo) {

	pointertogongstructfield.CreatedAt = pointertogongstructfieldDB.CreatedAt
	pointertogongstructfield.DeletedAt = pointertogongstructfieldDB.DeletedAt
	pointertogongstructfield.ID = pointertogongstructfieldDB.ID

	// insertion point for basic fields copy operations
	pointertogongstructfield.Name = pointertogongstructfieldDB.Name
	pointertogongstructfield.Index = pointertogongstructfieldDB.Index
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldDB.CompositeStructName

	// insertion point for pointer fields encoding
	pointertogongstructfield.GongStruct = frontRepo.map_ID_GongStruct.get(pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding.GongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
