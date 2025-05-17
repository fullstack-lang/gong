// generated code - do not edit

import { FormEditAssocButtonAPI } from './formeditassocbutton-api'
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
	AssociationStorage: string = ""
	HasChanged: boolean = false
	IsForSavePurpose: boolean = false

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormEditAssocButtonToFormEditAssocButtonAPI(formeditassocbutton: FormEditAssocButton, formeditassocbuttonAPI: FormEditAssocButtonAPI) {

	formeditassocbuttonAPI.CreatedAt = formeditassocbutton.CreatedAt
	formeditassocbuttonAPI.DeletedAt = formeditassocbutton.DeletedAt
	formeditassocbuttonAPI.ID = formeditassocbutton.ID

	// insertion point for basic fields copy operations
	formeditassocbuttonAPI.Name = formeditassocbutton.Name
	formeditassocbuttonAPI.Label = formeditassocbutton.Label
	formeditassocbuttonAPI.AssociationStorage = formeditassocbutton.AssociationStorage
	formeditassocbuttonAPI.HasChanged = formeditassocbutton.HasChanged
	formeditassocbuttonAPI.IsForSavePurpose = formeditassocbutton.IsForSavePurpose

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormEditAssocButtonAPIToFormEditAssocButton update basic, pointers and slice of pointers fields of formeditassocbutton
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formeditassocbuttonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormEditAssocButtonAPIToFormEditAssocButton(formeditassocbuttonAPI: FormEditAssocButtonAPI, formeditassocbutton: FormEditAssocButton, frontRepo: FrontRepo) {

	formeditassocbutton.CreatedAt = formeditassocbuttonAPI.CreatedAt
	formeditassocbutton.DeletedAt = formeditassocbuttonAPI.DeletedAt
	formeditassocbutton.ID = formeditassocbuttonAPI.ID

	// insertion point for basic fields copy operations
	formeditassocbutton.Name = formeditassocbuttonAPI.Name
	formeditassocbutton.Label = formeditassocbuttonAPI.Label
	formeditassocbutton.AssociationStorage = formeditassocbuttonAPI.AssociationStorage
	formeditassocbutton.HasChanged = formeditassocbuttonAPI.HasChanged
	formeditassocbutton.IsForSavePurpose = formeditassocbuttonAPI.IsForSavePurpose

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
