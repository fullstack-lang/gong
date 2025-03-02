// generated code - do not edit

import { FormFieldDateAPI } from './formfielddate-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDate {

	static GONGSTRUCT_NAME = "FormFieldDate"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldDateToFormFieldDateAPI(formfielddate: FormFieldDate, formfielddateAPI: FormFieldDateAPI) {

	formfielddateAPI.CreatedAt = formfielddate.CreatedAt
	formfielddateAPI.DeletedAt = formfielddate.DeletedAt
	formfielddateAPI.ID = formfielddate.ID

	// insertion point for basic fields copy operations
	formfielddateAPI.Name = formfielddate.Name
	formfielddateAPI.Value = formfielddate.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldDateAPIToFormFieldDate update basic, pointers and slice of pointers fields of formfielddate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfielddateAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldDateAPIToFormFieldDate(formfielddateAPI: FormFieldDateAPI, formfielddate: FormFieldDate, frontRepo: FrontRepo) {

	formfielddate.CreatedAt = formfielddateAPI.CreatedAt
	formfielddate.DeletedAt = formfielddateAPI.DeletedAt
	formfielddate.ID = formfielddateAPI.ID

	// insertion point for basic fields copy operations
	formfielddate.Name = formfielddateAPI.Name
	formfielddate.Value = formfielddateAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
