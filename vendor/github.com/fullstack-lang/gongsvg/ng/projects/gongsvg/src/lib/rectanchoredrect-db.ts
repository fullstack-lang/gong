// insertion point for imports
import { RectDB } from './rect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredRectDB {

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
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Rect_RectAnchoredRectsDBID: NullInt64 = new NullInt64
	Rect_RectAnchoredRectsDBID_Index: NullInt64  = new NullInt64 // store the index of the rectanchoredrect instance in Rect.RectAnchoredRects
	Rect_RectAnchoredRects_reverse?: RectDB 

}
