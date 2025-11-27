// generated code - do not edit

import { ButtonToggleAPI } from './buttontoggle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonToggle {

	static GONGSTRUCT_NAME = "ButtonToggle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	Icon: string = ""
	IsDisabled: boolean = false

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyButtonToggleToButtonToggleAPI(buttontoggle: ButtonToggle, buttontoggleAPI: ButtonToggleAPI) {

	buttontoggleAPI.CreatedAt = buttontoggle.CreatedAt
	buttontoggleAPI.DeletedAt = buttontoggle.DeletedAt
	buttontoggleAPI.ID = buttontoggle.ID

	// insertion point for basic fields copy operations
	buttontoggleAPI.Name = buttontoggle.Name
	buttontoggleAPI.Label = buttontoggle.Label
	buttontoggleAPI.Icon = buttontoggle.Icon
	buttontoggleAPI.IsDisabled = buttontoggle.IsDisabled

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyButtonToggleAPIToButtonToggle update basic, pointers and slice of pointers fields of buttontoggle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of buttontoggleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyButtonToggleAPIToButtonToggle(buttontoggleAPI: ButtonToggleAPI, buttontoggle: ButtonToggle, frontRepo: FrontRepo) {

	buttontoggle.CreatedAt = buttontoggleAPI.CreatedAt
	buttontoggle.DeletedAt = buttontoggleAPI.DeletedAt
	buttontoggle.ID = buttontoggleAPI.ID

	// insertion point for basic fields copy operations
	buttontoggle.Name = buttontoggleAPI.Name
	buttontoggle.Label = buttontoggleAPI.Label
	buttontoggle.Icon = buttontoggleAPI.Icon
	buttontoggle.IsDisabled = buttontoggleAPI.IsDisabled

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
