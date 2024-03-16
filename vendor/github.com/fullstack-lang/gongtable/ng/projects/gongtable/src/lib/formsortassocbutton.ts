// generated code - do not edit

import { FormSortAssocButtonAPI } from './formsortassocbutton-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormSortAssocButton {

	static GONGSTRUCT_NAME = "FormSortAssocButton"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormSortAssocButtonToFormSortAssocButtonAPI(formsortassocbutton: FormSortAssocButton, formsortassocbuttonAPI: FormSortAssocButtonAPI) {

	formsortassocbuttonAPI.CreatedAt = formsortassocbutton.CreatedAt
	formsortassocbuttonAPI.DeletedAt = formsortassocbutton.DeletedAt
	formsortassocbuttonAPI.ID = formsortassocbutton.ID

	// insertion point for basic fields copy operations
	formsortassocbuttonAPI.Name = formsortassocbutton.Name
	formsortassocbuttonAPI.Label = formsortassocbutton.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormSortAssocButtonAPIToFormSortAssocButton update basic, pointers and slice of pointers fields of formsortassocbutton
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formsortassocbuttonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormSortAssocButtonAPIToFormSortAssocButton(formsortassocbuttonAPI: FormSortAssocButtonAPI, formsortassocbutton: FormSortAssocButton, frontRepo: FrontRepo) {

	formsortassocbutton.CreatedAt = formsortassocbuttonAPI.CreatedAt
	formsortassocbutton.DeletedAt = formsortassocbuttonAPI.DeletedAt
	formsortassocbutton.ID = formsortassocbuttonAPI.ID

	// insertion point for basic fields copy operations
	formsortassocbutton.Name = formsortassocbuttonAPI.Name
	formsortassocbutton.Label = formsortassocbuttonAPI.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
