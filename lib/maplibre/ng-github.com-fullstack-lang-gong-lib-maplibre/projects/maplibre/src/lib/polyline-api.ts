// generated code - do not edit
// insertion point for imports
import { CoordinateAPI } from './coordinate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolylineAPI {

	static GONGSTRUCT_NAME = "Polyline"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	PopupHTML: string = ""
	StrokeColor: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""

	// insertion point for other decls

	PolylinePointersEncoding: PolylinePointersEncoding = new PolylinePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class PolylinePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Coordinates: number[] = []
}
