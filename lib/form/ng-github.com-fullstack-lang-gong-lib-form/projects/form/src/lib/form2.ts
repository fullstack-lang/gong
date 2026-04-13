// generated code - do not edit

import { Form2API } from './form2-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Form2 {

	static GONGSTRUCT_NAME = "Form2"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyForm2ToForm2API(form2: Form2, form2API: Form2API) {

	form2API.CreatedAt = form2.CreatedAt
	form2API.DeletedAt = form2.DeletedAt
	form2API.ID = form2.ID

	// insertion point for basic fields copy operations
	form2API.Name = form2.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyForm2APIToForm2 update basic, pointers and slice of pointers fields of form2
// from respectively the basic fields and encoded fields of pointers and slices of pointers of form2API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyForm2APIToForm2(form2API: Form2API, form2: Form2, frontRepo: FrontRepo) {

	form2.CreatedAt = form2API.CreatedAt
	form2.DeletedAt = form2API.DeletedAt
	form2.ID = form2API.ID

	// insertion point for basic fields copy operations
	form2.Name = form2API.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
