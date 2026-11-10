// generated code - do not edit
// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkerAPI {

	static GONGSTRUCT_NAME = "Marker"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	PopupHTML: string = ""
	Longitude: number = 0
	Latitude: number = 0
	Altitude: number = 0
	SVG: string = ""
	Color: string = ""
	Scale: number = 0
	Rotation: number = 0
	RotationAlignment: string = ""
	PitchAlignment: string = ""
	Anchor: string = ""
	IsDraggable: boolean = false

	// insertion point for other decls

	MarkerPointersEncoding: MarkerPointersEncoding = new MarkerPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MarkerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
