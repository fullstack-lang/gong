// generated code - do not edit

import { FormFieldTimeAPI } from './formfieldtime-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldTime {

	static GONGSTRUCT_NAME = "FormFieldTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date
	Step: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldTimeToFormFieldTimeAPI(formfieldtime: FormFieldTime, formfieldtimeAPI: FormFieldTimeAPI) {

	formfieldtimeAPI.CreatedAt = formfieldtime.CreatedAt
	formfieldtimeAPI.DeletedAt = formfieldtime.DeletedAt
	formfieldtimeAPI.ID = formfieldtime.ID

	// insertion point for basic fields copy operations
	formfieldtimeAPI.Name = formfieldtime.Name
	formfieldtimeAPI.Value = formfieldtime.Value
	formfieldtimeAPI.Step = formfieldtime.Step

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldTimeAPIToFormFieldTime update basic, pointers and slice of pointers fields of formfieldtime
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldtimeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldTimeAPIToFormFieldTime(formfieldtimeAPI: FormFieldTimeAPI, formfieldtime: FormFieldTime, frontRepo: FrontRepo) {

	formfieldtime.CreatedAt = formfieldtimeAPI.CreatedAt
	formfieldtime.DeletedAt = formfieldtimeAPI.DeletedAt
	formfieldtime.ID = formfieldtimeAPI.ID

	// insertion point for basic fields copy operations
	formfieldtime.Name = formfieldtimeAPI.Name
	formfieldtime.Value = formfieldtimeAPI.Value
	formfieldtime.Step = formfieldtimeAPI.Step

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
