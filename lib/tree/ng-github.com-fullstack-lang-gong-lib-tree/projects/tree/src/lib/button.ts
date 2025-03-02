// generated code - do not edit

import { ButtonAPI } from './button-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { SVGIcon } from './svgicon'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Button {

	static GONGSTRUCT_NAME = "Button"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations
	SVGIcon?: SVGIcon

}

export function CopyButtonToButtonAPI(button: Button, buttonAPI: ButtonAPI) {

	buttonAPI.CreatedAt = button.CreatedAt
	buttonAPI.DeletedAt = button.DeletedAt
	buttonAPI.ID = button.ID

	// insertion point for basic fields copy operations
	buttonAPI.Name = button.Name
	buttonAPI.Icon = button.Icon

	// insertion point for pointer fields encoding
	buttonAPI.ButtonPointersEncoding.SVGIconID.Valid = true
	if (button.SVGIcon != undefined) {
		buttonAPI.ButtonPointersEncoding.SVGIconID.Int64 = button.SVGIcon.ID  
	} else {
		buttonAPI.ButtonPointersEncoding.SVGIconID.Int64 = 0 		
	}


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
	button.Icon = buttonAPI.Icon

	// insertion point for pointer fields encoding
	button.SVGIcon = frontRepo.map_ID_SVGIcon.get(buttonAPI.ButtonPointersEncoding.SVGIconID.Int64)

	// insertion point for slice of pointers fields encoding
}
