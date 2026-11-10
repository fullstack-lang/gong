// generated code - do not edit

import { LayerAPI } from './layer-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Polygon } from './polygon'
import { Polyline } from './polyline'
import { Marker } from './marker'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Layer {

	static GONGSTRUCT_NAME = "Layer"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsVisible: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Polygons: Array<Polygon> = []
	Polylines: Array<Polyline> = []
	Markers: Array<Marker> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyLayerToLayerAPI(layer: Layer, layerAPI: LayerAPI) {

	layerAPI.CreatedAt = layer.CreatedAt
	layerAPI.DeletedAt = layer.DeletedAt
	layerAPI.ID = layer.ID

	// insertion point for basic fields copy operations
	layerAPI.Name = layer.Name
	layerAPI.IsVisible = layer.IsVisible

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	layerAPI.LayerPointersEncoding.Polygons = []
	for (let _polygon of layer.Polygons) {
		layerAPI.LayerPointersEncoding.Polygons.push(_polygon.ID)
	}

	layerAPI.LayerPointersEncoding.Polylines = []
	for (let _polyline of layer.Polylines) {
		layerAPI.LayerPointersEncoding.Polylines.push(_polyline.ID)
	}

	layerAPI.LayerPointersEncoding.Markers = []
	for (let _marker of layer.Markers) {
		layerAPI.LayerPointersEncoding.Markers.push(_marker.ID)
	}

}

// CopyLayerAPIToLayer update basic, pointers and slice of pointers fields of layer
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerAPIToLayer(layerAPI: LayerAPI, layer: Layer, frontRepo: FrontRepo) {

	layer.CreatedAt = layerAPI.CreatedAt
	layer.DeletedAt = layerAPI.DeletedAt
	layer.ID = layerAPI.ID

	// insertion point for basic fields copy operations
	layer.Name = layerAPI.Name
	layer.IsVisible = layerAPI.IsVisible

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Polygons)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polygons);
		return;
	}

	layer.Polygons = new Array<Polygon>()
	for (let _id of layerAPI.LayerPointersEncoding.Polygons) {
		let _polygon = frontRepo.map_ID_Polygon.get(_id)
		if (_polygon != undefined) {
			layer.Polygons.push(_polygon!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Polylines)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polylines);
		return;
	}

	layer.Polylines = new Array<Polyline>()
	for (let _id of layerAPI.LayerPointersEncoding.Polylines) {
		let _polyline = frontRepo.map_ID_Polyline.get(_id)
		if (_polyline != undefined) {
			layer.Polylines.push(_polyline!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Markers)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Markers);
		return;
	}

	layer.Markers = new Array<Marker>()
	for (let _id of layerAPI.LayerPointersEncoding.Markers) {
		let _marker = frontRepo.map_ID_Marker.get(_id)
		if (_marker != undefined) {
			layer.Markers.push(_marker!)
		}
	}
}
