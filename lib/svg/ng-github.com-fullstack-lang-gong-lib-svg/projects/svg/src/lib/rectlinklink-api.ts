// insertion point for imports
import { RectAPI } from './rect-api'
import { LinkAPI } from './link-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectLinkLinkAPI {

	static GONGSTRUCT_NAME = "RectLinkLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	TargetAnchorPosition: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	RectLinkLinkPointersEncoding: RectLinkLinkPointersEncoding = new RectLinkLinkPointersEncoding
}

export class RectLinkLinkPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	StartID: NullInt64 = new NullInt64 // if pointer is null, Start.ID = 0

	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

}
