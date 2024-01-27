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

	gongbasicfieldDB.CreatedAt = gongbasicfield.CreatedAt
	gongbasicfieldDB.DeletedAt = gongbasicfield.DeletedAt
	gongbasicfieldDB.ID = gongbasicfield.ID

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

// CopyGongBasicFieldDBToGongBasicField update basic, pointers and slice of pointers fields of gongbasicfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongbasicfieldDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongBasicFieldDBToGongBasicField(gongbasicfieldDB: GongBasicFieldDB, gongbasicfield: GongBasicField, frontRepo: FrontRepo) {

	gongbasicfield.CreatedAt = gongbasicfieldDB.CreatedAt
	gongbasicfield.DeletedAt = gongbasicfieldDB.DeletedAt
	gongbasicfield.ID = gongbasicfieldDB.ID

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
	gongbasicfield.GongEnum = frontRepo.map_ID_GongEnum.get(gongbasicfieldDB.GongBasicFieldPointersEncoding.GongEnumID.Int64)

	// insertion point for slice of pointers fields encoding
}
