// generated code - do not edit
// insertion point for imports
import { PolygonAPI } from './polygon-api'
import { PolylineAPI } from './polyline-api'
import { MarkerAPI } from './marker-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerAPI {

	static GONGSTRUCT_NAME = "Layer"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsVisible: boolean = false

	// insertion point for other decls

	LayerPointersEncoding: LayerPointersEncoding = new LayerPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class LayerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Polygons: number[] = []
	Polylines: number[] = []
	Markers: number[] = []
}
