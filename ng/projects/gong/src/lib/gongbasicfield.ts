// generated code - do not edit

import { GongBasicFieldDB } from './gongbasicfield-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongEnum } from './gongenum'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongBasicField {

	static GONGSTRUCT_NAME = "GongBasicField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	BasicKindName: string = ""
	DeclaredType: string = ""
	CompositeStructName: string = ""
	Index: number = 0
	IsDocLink: boolean = false
	IsTextArea: boolean = false
	IsBespokeWidth: boolean = false
	BespokeWidth: number = 0
	IsBespokeHeight: boolean = false
	BespokeHeight: number = 0

	// insertion point for pointers and slices of pointers declarations
	GongEnum?: GongEnum

}

export function CopyGongBasicFieldToGongBasicFieldDB(gongbasicfield: GongBasicField, gongbasicfieldDB: GongBasicFieldDB) {

	// insertion point for basic fields copy operations
	gongbasicfieldDB.Name = gongbasicfield.Name
	gongbasicfieldDB.BasicKindName = gongbasicfield.BasicKindName
	gongbasicfieldDB.DeclaredType = gongbasicfield.DeclaredType
	gongbasicfieldDB.CompositeStructName = gongbasicfield.CompositeStructName
	gongbasicfieldDB.Index = gongbasicfield.Index
	gongbasicfieldDB.IsDocLink = gongbasicfield.IsDocLink
	gongbasicfieldDB.IsTextArea = gongbasicfield.IsTextArea
	gongbasicfieldDB.IsBespokeWidth = gongbasicfield.IsBespokeWidth
	gongbasicfieldDB.BespokeWidth = gongbasicfield.BespokeWidth
	gongbasicfieldDB.IsBespokeHeight = gongbasicfield.IsBespokeHeight
	gongbasicfieldDB.BespokeHeight = gongbasicfield.BespokeHeight

	// insertion point for pointer fields encoding
    gongbasicfieldDB.GongBasicFieldPointersEncoding.GongEnumID.Valid = true
	if (gongbasicfield.GongEnum != undefined) {
		gongbasicfieldDB.GongBasicFieldPointersEncoding.GongEnumID.Int64 = gongbasicfield.GongEnum.ID  
    } else {
		gongbasicfieldDB.GongBasicFieldPointersEncoding.GongEnumID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyGongBasicFieldDBToGongBasicField(gongbasicfieldDB: GongBasicFieldDB, gongbasicfield: GongBasicField, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations
	gongbasicfield.Name = gongbasicfieldDB.Name
	gongbasicfield.BasicKindName = gongbasicfieldDB.BasicKindName
	gongbasicfield.DeclaredType = gongbasicfieldDB.DeclaredType
	gongbasicfield.CompositeStructName = gongbasicfieldDB.CompositeStructName
	gongbasicfield.Index = gongbasicfieldDB.Index
	gongbasicfield.IsDocLink = gongbasicfieldDB.IsDocLink
	gongbasicfield.IsTextArea = gongbasicfieldDB.IsTextArea
	gongbasicfield.IsBespokeWidth = gongbasicfieldDB.IsBespokeWidth
	gongbasicfield.BespokeWidth = gongbasicfieldDB.BespokeWidth
	gongbasicfield.IsBespokeHeight = gongbasicfieldDB.IsBespokeHeight
	gongbasicfield.BespokeHeight = gongbasicfieldDB.BespokeHeight

	// insertion point for pointer fields encoding
	gongbasicfield.GongEnum = frontRepo.GongEnums.get(gongbasicfieldDB.GongBasicFieldPointersEncoding.GongEnumID.Int64)

	// insertion point for slice of pointers fields encoding
}
