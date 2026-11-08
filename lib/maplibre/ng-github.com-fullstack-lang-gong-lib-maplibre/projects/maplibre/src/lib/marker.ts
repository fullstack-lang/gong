// generated code - do not edit

import { MarkerAPI } from './marker-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Marker {

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

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMarkerToMarkerAPI(marker: Marker, markerAPI: MarkerAPI) {

	markerAPI.CreatedAt = marker.CreatedAt
	markerAPI.DeletedAt = marker.DeletedAt
	markerAPI.ID = marker.ID

	// insertion point for basic fields copy operations
	markerAPI.Name = marker.Name
	markerAPI.Label = marker.Label
	markerAPI.PopupHTML = marker.PopupHTML
	markerAPI.Longitude = marker.Longitude
	markerAPI.Latitude = marker.Latitude
	markerAPI.Altitude = marker.Altitude
	markerAPI.SVG = marker.SVG
	markerAPI.Color = marker.Color
	markerAPI.Scale = marker.Scale
	markerAPI.Rotation = marker.Rotation
	markerAPI.RotationAlignment = marker.RotationAlignment
	markerAPI.PitchAlignment = marker.PitchAlignment
	markerAPI.Anchor = marker.Anchor
	markerAPI.IsDraggable = marker.IsDraggable

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMarkerAPIToMarker update basic, pointers and slice of pointers fields of marker
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMarkerAPIToMarker(markerAPI: MarkerAPI, marker: Marker, frontRepo: FrontRepo) {

	marker.CreatedAt = markerAPI.CreatedAt
	marker.DeletedAt = markerAPI.DeletedAt
	marker.ID = markerAPI.ID

	// insertion point for basic fields copy operations
	marker.Name = markerAPI.Name
	marker.Label = markerAPI.Label
	marker.PopupHTML = markerAPI.PopupHTML
	marker.Longitude = markerAPI.Longitude
	marker.Latitude = markerAPI.Latitude
	marker.Altitude = markerAPI.Altitude
	marker.SVG = markerAPI.SVG
	marker.Color = markerAPI.Color
	marker.Scale = markerAPI.Scale
	marker.Rotation = markerAPI.Rotation
	marker.RotationAlignment = markerAPI.RotationAlignment
	marker.PitchAlignment = markerAPI.PitchAlignment
	marker.Anchor = markerAPI.Anchor
	marker.IsDraggable = markerAPI.IsDraggable

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
