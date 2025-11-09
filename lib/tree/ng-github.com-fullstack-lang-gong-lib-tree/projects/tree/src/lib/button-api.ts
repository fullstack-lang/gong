// insertion point for imports
import { SVGIconAPI } from './svgicon-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonAPI {

	static GONGSTRUCT_NAME = "Button"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""
	IsDisabled: boolean = false
	HasToolTip: boolean = false
	ToolTipText: string = ""
	ToolTipPosition: string = ""

	// insertion point for other decls

	ButtonPointersEncoding: ButtonPointersEncoding = new ButtonPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	SVGIconID: NullInt64 = new NullInt64 // if pointer is null, SVGIcon.ID = 0

}
