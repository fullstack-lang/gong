// generated code - do not edit

import { FormSortAssocButtonDB } from './formsortassocbutton-db'
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

export function CopyFormSortAssocButtonToFormSortAssocButtonDB(formsortassocbutton: FormSortAssocButton, formsortassocbuttonDB: FormSortAssocButtonDB) {

	formsortassocbuttonDB.CreatedAt = formsortassocbutton.CreatedAt
	formsortassocbuttonDB.DeletedAt = formsortassocbutton.DeletedAt
	formsortassocbuttonDB.ID = formsortassocbutton.ID

	// insertion point for basic fields copy operations
	formsortassocbuttonDB.Name = formsortassocbutton.Name
	formsortassocbuttonDB.Label = formsortassocbutton.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormSortAssocButtonDBToFormSortAssocButton update basic, pointers and slice of pointers fields of formsortassocbutton
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formsortassocbuttonDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormSortAssocButtonDBToFormSortAssocButton(formsortassocbuttonDB: FormSortAssocButtonDB, formsortassocbutton: FormSortAssocButton, frontRepo: FrontRepo) {

	formsortassocbutton.CreatedAt = formsortassocbuttonDB.CreatedAt
	formsortassocbutton.DeletedAt = formsortassocbuttonDB.DeletedAt
	formsortassocbutton.ID = formsortassocbuttonDB.ID

	// insertion point for basic fields copy operations
	formsortassocbutton.Name = formsortassocbuttonDB.Name
	formsortassocbutton.Label = formsortassocbuttonDB.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
