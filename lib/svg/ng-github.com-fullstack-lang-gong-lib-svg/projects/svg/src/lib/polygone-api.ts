// insertion point for imports
import { AnimateAPI } from './animate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolygoneAPI {

	static GONGSTRUCT_NAME = "Polygone"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Points: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	PolygonePointersEncoding: PolygonePointersEncoding = new PolygonePointersEncoding
}

export class PolygonePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
