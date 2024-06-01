// generated code - do not edit

import { GongEnumValueAPI } from './gongenumvalue-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValue {

	static GONGSTRUCT_NAME = "GongEnumValue"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongEnumValueToGongEnumValueAPI(gongenumvalue: GongEnumValue, gongenumvalueAPI: GongEnumValueAPI) {

	gongenumvalueAPI.CreatedAt = gongenumvalue.CreatedAt
	gongenumvalueAPI.DeletedAt = gongenumvalue.DeletedAt
	gongenumvalueAPI.ID = gongenumvalue.ID

	// insertion point for basic fields copy operations
	gongenumvalueAPI.Name = gongenumvalue.Name
	gongenumvalueAPI.Value = gongenumvalue.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongEnumValueAPIToGongEnumValue update basic, pointers and slice of pointers fields of gongenumvalue
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongenumvalueAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongEnumValueAPIToGongEnumValue(gongenumvalueAPI: GongEnumValueAPI, gongenumvalue: GongEnumValue, frontRepo: FrontRepo) {

	gongenumvalue.CreatedAt = gongenumvalueAPI.CreatedAt
	gongenumvalue.DeletedAt = gongenumvalueAPI.DeletedAt
	gongenumvalue.ID = gongenumvalueAPI.ID

	// insertion point for basic fields copy operations
	gongenumvalue.Name = gongenumvalueAPI.Name
	gongenumvalue.Value = gongenumvalueAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
