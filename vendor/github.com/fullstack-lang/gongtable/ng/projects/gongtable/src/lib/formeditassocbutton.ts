// generated code - do not edit

import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormEditAssocButton {

	static GONGSTRUCT_NAME = "FormEditAssocButton"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormEditAssocButtonToFormEditAssocButtonDB(formeditassocbutton: FormEditAssocButton, formeditassocbuttonDB: FormEditAssocButtonDB) {

	formeditassocbuttonDB.CreatedAt = formeditassocbutton.CreatedAt
	formeditassocbuttonDB.DeletedAt = formeditassocbutton.DeletedAt
	formeditassocbuttonDB.ID = formeditassocbutton.ID

	// insertion point for basic fields copy operations
	formeditassocbuttonDB.Name = formeditassocbutton.Name
	formeditassocbuttonDB.Label = formeditassocbutton.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormEditAssocButtonDBToFormEditAssocButton update basic, pointers and slice of pointers fields of formeditassocbutton
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formeditassocbuttonDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormEditAssocButtonDBToFormEditAssocButton(formeditassocbuttonDB: FormEditAssocButtonDB, formeditassocbutton: FormEditAssocButton, frontRepo: FrontRepo) {

	formeditassocbutton.CreatedAt = formeditassocbuttonDB.CreatedAt
	formeditassocbutton.DeletedAt = formeditassocbuttonDB.DeletedAt
	formeditassocbutton.ID = formeditassocbuttonDB.ID

	// insertion point for basic fields copy operations
	formeditassocbutton.Name = formeditassocbuttonDB.Name
	formeditassocbutton.Label = formeditassocbuttonDB.Label

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
