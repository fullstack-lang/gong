// insertion point for imports
import { AnimateDB } from './animate-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredTextDB {

	static GONGSTRUCT_NAME = "RectAnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	FontWeight: string = ""
	FontSize: number = 0
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	TextAnchorType: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<AnimateDB> = []

	RectAnchoredTextPointersEncoding: RectAnchoredTextPointersEncoding = new RectAnchoredTextPointersEncoding
}

export class RectAnchoredTextPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
