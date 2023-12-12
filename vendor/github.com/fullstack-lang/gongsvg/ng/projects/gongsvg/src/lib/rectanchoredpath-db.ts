// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredPathDB {

	static GONGSTRUCT_NAME = "RectAnchoredPath"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	WidthFollowRect: boolean = false
	HeightFollowRect: boolean = false
	ScalePropotionnally: boolean = false
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations

	RectAnchoredPathPointersEncoding: RectAnchoredPathPointersEncoding = new RectAnchoredPathPointersEncoding
}

export class RectAnchoredPathPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
