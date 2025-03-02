// insertion point for imports
import { RectAPI } from './rect-api'
import { LinkAnchoredTextAPI } from './linkanchoredtext-api'
import { PointAPI } from './point-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAPI {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: string = ""
	IsBezierCurve: boolean = false
	StartAnchorType: string = ""
	EndAnchorType: string = ""
	StartOrientation: string = ""
	StartRatio: number = 0
	EndOrientation: string = ""
	EndRatio: number = 0
	CornerOffsetRatio: number = 0
	CornerRadius: number = 0
	HasEndArrow: boolean = false
	EndArrowSize: number = 0
	HasStartArrow: boolean = false
	StartArrowSize: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	LinkPointersEncoding: LinkPointersEncoding = new LinkPointersEncoding
}

export class LinkPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	StartID: NullInt64 = new NullInt64 // if pointer is null, Start.ID = 0

	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

	TextAtArrowEnd: number[] = []
	TextAtArrowStart: number[] = []
	ControlPoints: number[] = []
}
