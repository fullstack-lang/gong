// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonToggleAPI {

	static GONGSTRUCT_NAME = "ButtonToggle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	Icon: string = ""
	IsDisabled: boolean = false

	// insertion point for other decls

	ButtonTogglePointersEncoding: ButtonTogglePointersEncoding = new ButtonTogglePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ButtonTogglePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
