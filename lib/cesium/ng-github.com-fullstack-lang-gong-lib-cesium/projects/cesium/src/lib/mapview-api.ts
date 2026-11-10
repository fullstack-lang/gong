// generated code - do not edit
// insertion point for imports
import { LayerAPI } from './layer-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MapViewAPI {

	static GONGSTRUCT_NAME = "MapView"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StyleURL: string = ""
	Longitude: number = 0
	Latitude: number = 0
	Zoom: number = 0
	Pitch: number = 0
	Bearing: number = 0

	// insertion point for other decls

	MapViewPointersEncoding: MapViewPointersEncoding = new MapViewPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MapViewPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Layers: number[] = []
}
