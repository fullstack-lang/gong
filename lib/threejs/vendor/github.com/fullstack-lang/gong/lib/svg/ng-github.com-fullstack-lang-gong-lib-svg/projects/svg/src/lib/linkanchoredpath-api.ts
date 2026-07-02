// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAnchoredPathAPI {

	static GONGSTRUCT_NAME = "LinkAnchoredPath"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	ScalePropotionnally: boolean = false
	AppliedScaling: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	LinkAnchoredPathPointersEncoding: LinkAnchoredPathPointersEncoding = new LinkAnchoredPathPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class LinkAnchoredPathPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
