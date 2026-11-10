// generated code - do not edit

import { MapViewAPI } from './mapview-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Layer } from './layer'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MapView {

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

	// insertion point for pointers and slices of pointers declarations
	Layers: Array<Layer> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMapViewToMapViewAPI(mapview: MapView, mapviewAPI: MapViewAPI) {

	mapviewAPI.CreatedAt = mapview.CreatedAt
	mapviewAPI.DeletedAt = mapview.DeletedAt
	mapviewAPI.ID = mapview.ID

	// insertion point for basic fields copy operations
	mapviewAPI.Name = mapview.Name
	mapviewAPI.StyleURL = mapview.StyleURL
	mapviewAPI.Longitude = mapview.Longitude
	mapviewAPI.Latitude = mapview.Latitude
	mapviewAPI.Zoom = mapview.Zoom
	mapviewAPI.Pitch = mapview.Pitch
	mapviewAPI.Bearing = mapview.Bearing

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	mapviewAPI.MapViewPointersEncoding.Layers = []
	for (let _layer of mapview.Layers) {
		mapviewAPI.MapViewPointersEncoding.Layers.push(_layer.ID)
	}

}

// CopyMapViewAPIToMapView update basic, pointers and slice of pointers fields of mapview
// from respectively the basic fields and encoded fields of pointers and slices of pointers of mapviewAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMapViewAPIToMapView(mapviewAPI: MapViewAPI, mapview: MapView, frontRepo: FrontRepo) {

	mapview.CreatedAt = mapviewAPI.CreatedAt
	mapview.DeletedAt = mapviewAPI.DeletedAt
	mapview.ID = mapviewAPI.ID

	// insertion point for basic fields copy operations
	mapview.Name = mapviewAPI.Name
	mapview.StyleURL = mapviewAPI.StyleURL
	mapview.Longitude = mapviewAPI.Longitude
	mapview.Latitude = mapviewAPI.Latitude
	mapview.Zoom = mapviewAPI.Zoom
	mapview.Pitch = mapviewAPI.Pitch
	mapview.Bearing = mapviewAPI.Bearing

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(mapviewAPI.MapViewPointersEncoding.Layers)) {
		console.error('Rects is not an array:', mapviewAPI.MapViewPointersEncoding.Layers);
		return;
	}

	mapview.Layers = new Array<Layer>()
	for (let _id of mapviewAPI.MapViewPointersEncoding.Layers) {
		let _layer = frontRepo.map_ID_Layer.get(_id)
		if (_layer != undefined) {
			mapview.Layers.push(_layer!)
		}
	}
}
