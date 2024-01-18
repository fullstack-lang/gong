// generated code - do not edit

import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'
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

export function CopySliceOfPointerToGongStructFieldToSliceOfPointerToGongStructFieldDB(sliceofpointertogongstructfield: SliceOfPointerToGongStructField, sliceofpointertogongstructfieldDB: SliceOfPointerToGongStructFieldDB) {

	// insertion point for basic fields copy operations
	sliceofpointertogongstructfieldDB.Name = sliceofpointertogongstructfield.Name
	sliceofpointertogongstructfieldDB.Index = sliceofpointertogongstructfield.Index
	sliceofpointertogongstructfieldDB.CompositeStructName = sliceofpointertogongstructfield.CompositeStructName

	// insertion point for pointer fields encoding
    sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Valid = true
	if (sliceofpointertogongstructfield.GongStruct != undefined) {
		sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64 = sliceofpointertogongstructfield.GongStruct.ID  
    } else {
		sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopySliceOfPointerToGongStructFieldDBToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldDB: SliceOfPointerToGongStructFieldDB, sliceofpointertogongstructfield: SliceOfPointerToGongStructField, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations
	sliceofpointertogongstructfield.Name = sliceofpointertogongstructfieldDB.Name
	sliceofpointertogongstructfield.Index = sliceofpointertogongstructfieldDB.Index
	sliceofpointertogongstructfield.CompositeStructName = sliceofpointertogongstructfieldDB.CompositeStructName

	// insertion point for pointer fields encoding
	sliceofpointertogongstructfield.GongStruct = frontRepo.GongStructs.get(sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
