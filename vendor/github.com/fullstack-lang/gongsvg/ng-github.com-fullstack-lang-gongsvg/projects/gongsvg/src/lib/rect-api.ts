// insertion point for imports
import { AnimateAPI } from './animate-api'
import { RectAnchoredTextAPI } from './rectanchoredtext-api'
import { RectAnchoredRectAPI } from './rectanchoredrect-api'
import { RectAnchoredPathAPI } from './rectanchoredpath-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAPI {

	static GONGSTRUCT_NAME = "Rect"

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
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	IsSelectable: boolean = false
	IsSelected: boolean = false
	CanHaveLeftHandle: boolean = false
	HasLeftHandle: boolean = false
	CanHaveRightHandle: boolean = false
	HasRightHandle: boolean = false
	CanHaveTopHandle: boolean = false
	HasTopHandle: boolean = false
	IsScalingProportionally: boolean = false
	CanHaveBottomHandle: boolean = false
	HasBottomHandle: boolean = false
	CanMoveHorizontaly: boolean = false
	CanMoveVerticaly: boolean = false

	// insertion point for other decls

	RectPointersEncoding: RectPointersEncoding = new RectPointersEncoding
}

export class RectPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animations: number[] = []
	RectAnchoredTexts: number[] = []
	RectAnchoredRects: number[] = []
	RectAnchoredPaths: number[] = []
}
