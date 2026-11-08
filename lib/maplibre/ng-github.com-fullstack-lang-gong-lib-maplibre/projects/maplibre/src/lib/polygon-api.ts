// generated code - do not edit
// insertion point for imports
import { CoordinateAPI } from './coordinate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolygonAPI {

	static GONGSTRUCT_NAME = "Polygon"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	PopupHTML: string = ""
	FillColor: string = ""
	FillOpacity: number = 0
	StrokeColor: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	IsExtruded: boolean = false
	BaseAltitude: number = 0
	ExtrusionHeight: number = 0

	// insertion point for other decls

	PolygonPointersEncoding: PolygonPointersEncoding = new PolygonPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class PolygonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Coordinates: number[] = []
}
