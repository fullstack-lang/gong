// generated code - do not edit

import { FormFieldDateTimeAPI } from './formfielddatetime-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDateTime {

	static GONGSTRUCT_NAME = "FormFieldDateTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldDateTimeToFormFieldDateTimeAPI(formfielddatetime: FormFieldDateTime, formfielddatetimeAPI: FormFieldDateTimeAPI) {

	formfielddatetimeAPI.CreatedAt = formfielddatetime.CreatedAt
	formfielddatetimeAPI.DeletedAt = formfielddatetime.DeletedAt
	formfielddatetimeAPI.ID = formfielddatetime.ID

	// insertion point for basic fields copy operations
	formfielddatetimeAPI.Name = formfielddatetime.Name
	formfielddatetimeAPI.Value = formfielddatetime.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldDateTimeAPIToFormFieldDateTime update basic, pointers and slice of pointers fields of formfielddatetime
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfielddatetimeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldDateTimeAPIToFormFieldDateTime(formfielddatetimeAPI: FormFieldDateTimeAPI, formfielddatetime: FormFieldDateTime, frontRepo: FrontRepo) {

	formfielddatetime.CreatedAt = formfielddatetimeAPI.CreatedAt
	formfielddatetime.DeletedAt = formfielddatetimeAPI.DeletedAt
	formfielddatetime.ID = formfielddatetimeAPI.ID

	// insertion point for basic fields copy operations
	formfielddatetime.Name = formfielddatetimeAPI.Name
	formfielddatetime.Value = formfielddatetimeAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
