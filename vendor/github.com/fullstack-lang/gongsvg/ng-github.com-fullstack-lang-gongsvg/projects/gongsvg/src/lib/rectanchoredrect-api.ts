// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredRectAPI {

	static GONGSTRUCT_NAME = "RectAnchoredRect"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Height: number = 0
	RX: number = 0
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	WidthFollowRect: boolean = false
	HeightFollowRect: boolean = false
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	RectAnchoredRectPointersEncoding: RectAnchoredRectPointersEncoding = new RectAnchoredRectPointersEncoding
}

export class RectAnchoredRectPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
