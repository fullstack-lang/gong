// generated code - do not edit

import { FormFieldStringDB } from './formfieldstring-db'
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

export function CopyFormFieldStringToFormFieldStringDB(formfieldstring: FormFieldString, formfieldstringDB: FormFieldStringDB) {

	formfieldstringDB.CreatedAt = formfieldstring.CreatedAt
	formfieldstringDB.DeletedAt = formfieldstring.DeletedAt
	formfieldstringDB.ID = formfieldstring.ID

	// insertion point for basic fields copy operations
	formfieldstringDB.Name = formfieldstring.Name
	formfieldstringDB.Value = formfieldstring.Value
	formfieldstringDB.IsTextArea = formfieldstring.IsTextArea

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldStringDBToFormFieldString update basic, pointers and slice of pointers fields of formfieldstring
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldstringDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldStringDBToFormFieldString(formfieldstringDB: FormFieldStringDB, formfieldstring: FormFieldString, frontRepo: FrontRepo) {

	formfieldstring.CreatedAt = formfieldstringDB.CreatedAt
	formfieldstring.DeletedAt = formfieldstringDB.DeletedAt
	formfieldstring.ID = formfieldstringDB.ID

	// insertion point for basic fields copy operations
	formfieldstring.Name = formfieldstringDB.Name
	formfieldstring.Value = formfieldstringDB.Value
	formfieldstring.IsTextArea = formfieldstringDB.IsTextArea

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
