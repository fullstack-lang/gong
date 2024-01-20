// generated code - do not edit

import { ButtonDB } from './button-db'
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

export function CopyButtonToButtonDB(button: Button, buttonDB: ButtonDB) {

	buttonDB.CreatedAt = button.CreatedAt
	buttonDB.DeletedAt = button.DeletedAt
	buttonDB.ID = button.ID
	
	// insertion point for basic fields copy operations
	buttonDB.Name = button.Name
	buttonDB.Icon = button.Icon

	// insertion point for pointer fields encoding
    buttonDB.ButtonPointersEncoding.SVGIconID.Valid = true
	if (button.SVGIcon != undefined) {
		buttonDB.ButtonPointersEncoding.SVGIconID.Int64 = button.SVGIcon.ID  
    } else {
		buttonDB.ButtonPointersEncoding.SVGIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyButtonDBToButton(buttonDB: ButtonDB, button: Button, frontRepo: FrontRepo) {

	button.CreatedAt = buttonDB.CreatedAt
	button.DeletedAt = buttonDB.DeletedAt
	button.ID = buttonDB.ID
	
	// insertion point for basic fields copy operations
	button.Name = buttonDB.Name
	button.Icon = buttonDB.Icon

	// insertion point for pointer fields encoding
	button.SVGIcon = frontRepo.SVGIcons.get(buttonDB.ButtonPointersEncoding.SVGIconID.Int64)

	// insertion point for slice of pointers fields encoding
}
