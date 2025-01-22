// generated code - do not edit

import { GongTimeFieldAPI } from './gongtimefield-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongTimeField {

	static GONGSTRUCT_NAME = "GongTimeField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""
	BespokeTimeFormat: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongTimeFieldToGongTimeFieldAPI(gongtimefield: GongTimeField, gongtimefieldAPI: GongTimeFieldAPI) {

	gongtimefieldAPI.CreatedAt = gongtimefield.CreatedAt
	gongtimefieldAPI.DeletedAt = gongtimefield.DeletedAt
	gongtimefieldAPI.ID = gongtimefield.ID

	// insertion point for basic fields copy operations
	gongtimefieldAPI.Name = gongtimefield.Name
	gongtimefieldAPI.Index = gongtimefield.Index
	gongtimefieldAPI.CompositeStructName = gongtimefield.CompositeStructName
	gongtimefieldAPI.BespokeTimeFormat = gongtimefield.BespokeTimeFormat

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongTimeFieldAPIToGongTimeField update basic, pointers and slice of pointers fields of gongtimefield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongtimefieldAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongTimeFieldAPIToGongTimeField(gongtimefieldAPI: GongTimeFieldAPI, gongtimefield: GongTimeField, frontRepo: FrontRepo) {

	gongtimefield.CreatedAt = gongtimefieldAPI.CreatedAt
	gongtimefield.DeletedAt = gongtimefieldAPI.DeletedAt
	gongtimefield.ID = gongtimefieldAPI.ID

	// insertion point for basic fields copy operations
	gongtimefield.Name = gongtimefieldAPI.Name
	gongtimefield.Index = gongtimefieldAPI.Index
	gongtimefield.CompositeStructName = gongtimefieldAPI.CompositeStructName
	gongtimefield.BespokeTimeFormat = gongtimefieldAPI.BespokeTimeFormat

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
