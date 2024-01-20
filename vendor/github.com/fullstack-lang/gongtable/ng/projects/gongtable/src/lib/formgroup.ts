// generated code - do not edit

import { FormGroupDB } from './formgroup-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { FormDiv } from './formdiv'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormGroup {

	static GONGSTRUCT_NAME = "FormGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasSuppressButton: boolean = false
	HasSuppressButtonBeenPressed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	FormDivs: Array<FormDiv> = []
}

export function CopyFormGroupToFormGroupDB(formgroup: FormGroup, formgroupDB: FormGroupDB) {

	formgroupDB.CreatedAt = formgroup.CreatedAt
	formgroupDB.DeletedAt = formgroup.DeletedAt
	formgroupDB.ID = formgroup.ID
	
	// insertion point for basic fields copy operations
	formgroupDB.Name = formgroup.Name
	formgroupDB.Label = formgroup.Label
	formgroupDB.HasSuppressButton = formgroup.HasSuppressButton
	formgroupDB.HasSuppressButtonBeenPressed = formgroup.HasSuppressButtonBeenPressed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	formgroupDB.FormGroupPointersEncoding.FormDivs = []
    for (let _formdiv of formgroup.FormDivs) {
		formgroupDB.FormGroupPointersEncoding.FormDivs.push(_formdiv.ID)
    }
	
}

export function CopyFormGroupDBToFormGroup(formgroupDB: FormGroupDB, formgroup: FormGroup, frontRepo: FrontRepo) {

	formgroup.CreatedAt = formgroupDB.CreatedAt
	formgroup.DeletedAt = formgroupDB.DeletedAt
	formgroup.ID = formgroupDB.ID
	
	// insertion point for basic fields copy operations
	formgroup.Name = formgroupDB.Name
	formgroup.Label = formgroupDB.Label
	formgroup.HasSuppressButton = formgroupDB.HasSuppressButton
	formgroup.HasSuppressButtonBeenPressed = formgroupDB.HasSuppressButtonBeenPressed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	formgroup.FormDivs = new Array<FormDiv>()
	for (let _id of formgroupDB.FormGroupPointersEncoding.FormDivs) {
	  let _formdiv = frontRepo.FormDivs.get(_id)
	  if (_formdiv != undefined) {
		formgroup.FormDivs.push(_formdiv!)
	  }
	}
}
