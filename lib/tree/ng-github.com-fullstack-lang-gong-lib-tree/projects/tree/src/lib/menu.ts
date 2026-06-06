// generated code - do not edit

import { MenuAPI } from './menu-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Button } from './button'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Menu {

	static GONGSTRUCT_NAME = "Menu"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Buttons: Array<Button> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMenuToMenuAPI(menu: Menu, menuAPI: MenuAPI) {

	menuAPI.CreatedAt = menu.CreatedAt
	menuAPI.DeletedAt = menu.DeletedAt
	menuAPI.ID = menu.ID

	// insertion point for basic fields copy operations
	menuAPI.Name = menu.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	menuAPI.MenuPointersEncoding.Buttons = []
	for (let _button of menu.Buttons) {
		menuAPI.MenuPointersEncoding.Buttons.push(_button.ID)
	}

}

// CopyMenuAPIToMenu update basic, pointers and slice of pointers fields of menu
// from respectively the basic fields and encoded fields of pointers and slices of pointers of menuAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMenuAPIToMenu(menuAPI: MenuAPI, menu: Menu, frontRepo: FrontRepo) {

	menu.CreatedAt = menuAPI.CreatedAt
	menu.DeletedAt = menuAPI.DeletedAt
	menu.ID = menuAPI.ID

	// insertion point for basic fields copy operations
	menu.Name = menuAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(menuAPI.MenuPointersEncoding.Buttons)) {
		console.error('Rects is not an array:', menuAPI.MenuPointersEncoding.Buttons);
		return;
	}

	menu.Buttons = new Array<Button>()
	for (let _id of menuAPI.MenuPointersEncoding.Buttons) {
		let _button = frontRepo.map_ID_Button.get(_id)
		if (_button != undefined) {
			menu.Buttons.push(_button!)
		}
	}
}
