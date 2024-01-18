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

export function CopyPointerToGongStructFieldDBToPointerToGongStructField(pointertogongstructfieldDB: PointerToGongStructFieldDB, pointertogongstructfield: PointerToGongStructField, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations
	pointertogongstructfield.Name = pointertogongstructfieldDB.Name
	pointertogongstructfield.Index = pointertogongstructfieldDB.Index
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldDB.CompositeStructName

	// insertion point for pointer fields encoding
	pointertogongstructfield.GongStruct = frontRepo.GongStructs.get(pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding.GongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
