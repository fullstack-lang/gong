// generated code - do not edit

import { FormFieldStringAPI } from './formfieldstring-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldString {

	static GONGSTRUCT_NAME = "FormFieldString"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""
	IsTextArea: boolean = false

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldStringToFormFieldStringAPI(formfieldstring: FormFieldString, formfieldstringAPI: FormFieldStringAPI) {

	formfieldstringAPI.CreatedAt = formfieldstring.CreatedAt
	formfieldstringAPI.DeletedAt = formfieldstring.DeletedAt
	formfieldstringAPI.ID = formfieldstring.ID

	// insertion point for basic fields copy operations
	formfieldstringAPI.Name = formfieldstring.Name
	formfieldstringAPI.Value = formfieldstring.Value
	formfieldstringAPI.IsTextArea = formfieldstring.IsTextArea

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldStringAPIToFormFieldString update basic, pointers and slice of pointers fields of formfieldstring
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldstringAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldStringAPIToFormFieldString(formfieldstringAPI: FormFieldStringAPI, formfieldstring: FormFieldString, frontRepo: FrontRepo) {

	formfieldstring.CreatedAt = formfieldstringAPI.CreatedAt
	formfieldstring.DeletedAt = formfieldstringAPI.DeletedAt
	formfieldstring.ID = formfieldstringAPI.ID

	// insertion point for basic fields copy operations
	formfieldstring.Name = formfieldstringAPI.Name
	formfieldstring.Value = formfieldstringAPI.Value
	formfieldstring.IsTextArea = formfieldstringAPI.IsTextArea

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
