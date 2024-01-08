// insertion point for imports
import { AnimateDB } from './animate-db'
import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { RectAnchoredPathDB } from './rectanchoredpath-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectDB {

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

	// insertion point for pointers and slices of pointers declarations
	Animations: Array<AnimateDB> = []
	RectAnchoredTexts: Array<RectAnchoredTextDB> = []
	RectAnchoredRects: Array<RectAnchoredRectDB> = []
	RectAnchoredPaths: Array<RectAnchoredPathDB> = []

	RectPointersEncoding: RectPointersEncoding = new RectPointersEncoding
}

export class RectPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animations: number[] = []
	RectAnchoredTexts: number[] = []
	RectAnchoredRects: number[] = []
	RectAnchoredPaths: number[] = []
}
