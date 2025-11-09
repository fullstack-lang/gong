// generated code - do not edit

import { FormSortAssocButtonAPI } from './formsortassocbutton-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { FormEditAssocButton } from './formeditassocbutton'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormSortAssocButton {

	static GONGSTRUCT_NAME = "FormSortAssocButton"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasToolTip: boolean = false
	ToolTipText: string = ""

	// insertion point for pointers and slices of pointers declarations
	FormEditAssocButton?: FormEditAssocButton


	CreatedAt?: string
	DeletedAt?: string
}

export function CopyFormSortAssocButtonToFormSortAssocButtonAPI(formsortassocbutton: FormSortAssocButton, formsortassocbuttonAPI: FormSortAssocButtonAPI) {

	formsortassocbuttonAPI.CreatedAt = formsortassocbutton.CreatedAt
	formsortassocbuttonAPI.DeletedAt = formsortassocbutton.DeletedAt
	formsortassocbuttonAPI.ID = formsortassocbutton.ID

	// insertion point for basic fields copy operations
	formsortassocbuttonAPI.Name = formsortassocbutton.Name
	formsortassocbuttonAPI.Label = formsortassocbutton.Label
	formsortassocbuttonAPI.HasToolTip = formsortassocbutton.HasToolTip
	formsortassocbuttonAPI.ToolTipText = formsortassocbutton.ToolTipText

	// insertion point for pointer fields encoding
	formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding.FormEditAssocButtonID.Valid = true
	if (formsortassocbutton.FormEditAssocButton != undefined) {
		formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding.FormEditAssocButtonID.Int64 = formsortassocbutton.FormEditAssocButton.ID  
	} else {
		formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding.FormEditAssocButtonID.Int64 = 0 		
	}


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
	formsortassocbutton.HasToolTip = formsortassocbuttonAPI.HasToolTip
	formsortassocbutton.ToolTipText = formsortassocbuttonAPI.ToolTipText

	// insertion point for pointer fields encoding
	formsortassocbutton.FormEditAssocButton = frontRepo.map_ID_FormEditAssocButton.get(formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding.FormEditAssocButtonID.Int64)

	// insertion point for slice of pointers fields encoding
}
