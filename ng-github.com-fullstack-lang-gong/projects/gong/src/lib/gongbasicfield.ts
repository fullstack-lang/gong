// generated code - do not edit

import { GongBasicFieldAPI } from './gongbasicfield-api'
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

export function CopyGongBasicFieldToGongBasicFieldAPI(gongbasicfield: GongBasicField, gongbasicfieldAPI: GongBasicFieldAPI) {

	gongbasicfieldAPI.CreatedAt = gongbasicfield.CreatedAt
	gongbasicfieldAPI.DeletedAt = gongbasicfield.DeletedAt
	gongbasicfieldAPI.ID = gongbasicfield.ID

	// insertion point for basic fields copy operations
	gongbasicfieldAPI.Name = gongbasicfield.Name
	gongbasicfieldAPI.BasicKindName = gongbasicfield.BasicKindName
	gongbasicfieldAPI.DeclaredType = gongbasicfield.DeclaredType
	gongbasicfieldAPI.CompositeStructName = gongbasicfield.CompositeStructName
	gongbasicfieldAPI.Index = gongbasicfield.Index
	gongbasicfieldAPI.IsDocLink = gongbasicfield.IsDocLink
	gongbasicfieldAPI.IsTextArea = gongbasicfield.IsTextArea
	gongbasicfieldAPI.IsBespokeWidth = gongbasicfield.IsBespokeWidth
	gongbasicfieldAPI.BespokeWidth = gongbasicfield.BespokeWidth
	gongbasicfieldAPI.IsBespokeHeight = gongbasicfield.IsBespokeHeight
	gongbasicfieldAPI.BespokeHeight = gongbasicfield.BespokeHeight

	// insertion point for pointer fields encoding
	gongbasicfieldAPI.GongBasicFieldPointersEncoding.GongEnumID.Valid = true
	if (gongbasicfield.GongEnum != undefined) {
		gongbasicfieldAPI.GongBasicFieldPointersEncoding.GongEnumID.Int64 = gongbasicfield.GongEnum.ID  
	} else {
		gongbasicfieldAPI.GongBasicFieldPointersEncoding.GongEnumID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyGongBasicFieldAPIToGongBasicField update basic, pointers and slice of pointers fields of gongbasicfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongbasicfieldAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongBasicFieldAPIToGongBasicField(gongbasicfieldAPI: GongBasicFieldAPI, gongbasicfield: GongBasicField, frontRepo: FrontRepo) {

	gongbasicfield.CreatedAt = gongbasicfieldAPI.CreatedAt
	gongbasicfield.DeletedAt = gongbasicfieldAPI.DeletedAt
	gongbasicfield.ID = gongbasicfieldAPI.ID

	// insertion point for basic fields copy operations
	gongbasicfield.Name = gongbasicfieldAPI.Name
	gongbasicfield.BasicKindName = gongbasicfieldAPI.BasicKindName
	gongbasicfield.DeclaredType = gongbasicfieldAPI.DeclaredType
	gongbasicfield.CompositeStructName = gongbasicfieldAPI.CompositeStructName
	gongbasicfield.Index = gongbasicfieldAPI.Index
	gongbasicfield.IsDocLink = gongbasicfieldAPI.IsDocLink
	gongbasicfield.IsTextArea = gongbasicfieldAPI.IsTextArea
	gongbasicfield.IsBespokeWidth = gongbasicfieldAPI.IsBespokeWidth
	gongbasicfield.BespokeWidth = gongbasicfieldAPI.BespokeWidth
	gongbasicfield.IsBespokeHeight = gongbasicfieldAPI.IsBespokeHeight
	gongbasicfield.BespokeHeight = gongbasicfieldAPI.BespokeHeight

	// insertion point for pointer fields encoding
	gongbasicfield.GongEnum = frontRepo.map_ID_GongEnum.get(gongbasicfieldAPI.GongBasicFieldPointersEncoding.GongEnumID.Int64)

	// insertion point for slice of pointers fields encoding
}
