// generated code - do not edit

import { FormAPI } from './form-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Form {

	static GONGSTRUCT_NAME = "Form"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""
	FormName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormToFormAPI(form: Form, formAPI: FormAPI) {

	formAPI.CreatedAt = form.CreatedAt
	formAPI.DeletedAt = form.DeletedAt
	formAPI.ID = form.ID

	// insertion point for basic fields copy operations
	formAPI.Name = form.Name
	formAPI.StackName = form.StackName
	formAPI.FormName = form.FormName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormAPIToForm update basic, pointers and slice of pointers fields of form
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormAPIToForm(formAPI: FormAPI, form: Form, frontRepo: FrontRepo) {

	form.CreatedAt = formAPI.CreatedAt
	form.DeletedAt = formAPI.DeletedAt
	form.ID = formAPI.ID

	// insertion point for basic fields copy operations
	form.Name = formAPI.Name
	form.StackName = formAPI.StackName
	form.FormName = formAPI.FormName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
