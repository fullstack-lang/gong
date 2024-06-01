// insertion point for imports
import { AnimateAPI } from './animate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAnchoredTextAPI {

	static GONGSTRUCT_NAME = "LinkAnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	AutomaticLayout: boolean = false
	LinkAnchorType: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	FontWeight: string = ""
	FontSize: string = ""
	LetterSpacing: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	LinkAnchoredTextPointersEncoding: LinkAnchoredTextPointersEncoding = new LinkAnchoredTextPointersEncoding
}

export class LinkAnchoredTextPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
