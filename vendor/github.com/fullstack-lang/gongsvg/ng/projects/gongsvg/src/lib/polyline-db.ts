// insertion point for imports
import { AnimateDB } from './animate-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolylineDB {

	static GONGSTRUCT_NAME = "Polyline"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Points: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<AnimateDB> = []

	PolylinePointersEncoding: PolylinePointersEncoding = new PolylinePointersEncoding
}

export class PolylinePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
