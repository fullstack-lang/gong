// generated code - do not edit

import { ButtonAPI } from './button-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Button {

	static GONGSTRUCT_NAME = "Button"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyButtonToButtonAPI(button: Button, buttonAPI: ButtonAPI) {

	buttonAPI.CreatedAt = button.CreatedAt
	buttonAPI.DeletedAt = button.DeletedAt
	buttonAPI.ID = button.ID

	// insertion point for basic fields copy operations
	buttonAPI.Name = button.Name
	buttonAPI.Label = button.Label
	buttonAPI.Icon = button.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyButtonAPIToButton update basic, pointers and slice of pointers fields of button
// from respectively the basic fields and encoded fields of pointers and slices of pointers of buttonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyButtonAPIToButton(buttonAPI: ButtonAPI, button: Button, frontRepo: FrontRepo) {

	button.CreatedAt = buttonAPI.CreatedAt
	button.DeletedAt = buttonAPI.DeletedAt
	button.ID = buttonAPI.ID

	// insertion point for basic fields copy operations
	button.Name = buttonAPI.Name
	button.Label = buttonAPI.Label
	button.Icon = buttonAPI.Icon

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
