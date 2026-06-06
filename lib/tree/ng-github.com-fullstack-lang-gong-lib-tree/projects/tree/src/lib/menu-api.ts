// insertion point for imports
import { ButtonAPI } from './button-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MenuAPI {

	static GONGSTRUCT_NAME = "Menu"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	MenuPointersEncoding: MenuPointersEncoding = new MenuPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MenuPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Buttons: number[] = []
}
