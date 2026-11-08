// generated code - do not edit

import { PolygonAPI } from './polygon-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Coordinate } from './coordinate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Polygon {

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

	// insertion point for pointers and slices of pointers declarations
	Coordinates: Array<Coordinate> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyPolygonToPolygonAPI(polygon: Polygon, polygonAPI: PolygonAPI) {

	polygonAPI.CreatedAt = polygon.CreatedAt
	polygonAPI.DeletedAt = polygon.DeletedAt
	polygonAPI.ID = polygon.ID

	// insertion point for basic fields copy operations
	polygonAPI.Name = polygon.Name
	polygonAPI.Label = polygon.Label
	polygonAPI.PopupHTML = polygon.PopupHTML
	polygonAPI.FillColor = polygon.FillColor
	polygonAPI.FillOpacity = polygon.FillOpacity
	polygonAPI.StrokeColor = polygon.StrokeColor
	polygonAPI.StrokeOpacity = polygon.StrokeOpacity
	polygonAPI.StrokeWidth = polygon.StrokeWidth
	polygonAPI.IsExtruded = polygon.IsExtruded
	polygonAPI.BaseAltitude = polygon.BaseAltitude
	polygonAPI.ExtrusionHeight = polygon.ExtrusionHeight

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polygonAPI.PolygonPointersEncoding.Coordinates = []
	for (let _coordinate of polygon.Coordinates) {
		polygonAPI.PolygonPointersEncoding.Coordinates.push(_coordinate.ID)
	}

}

// CopyPolygonAPIToPolygon update basic, pointers and slice of pointers fields of polygon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polygonAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPolygonAPIToPolygon(polygonAPI: PolygonAPI, polygon: Polygon, frontRepo: FrontRepo) {

	polygon.CreatedAt = polygonAPI.CreatedAt
	polygon.DeletedAt = polygonAPI.DeletedAt
	polygon.ID = polygonAPI.ID

	// insertion point for basic fields copy operations
	polygon.Name = polygonAPI.Name
	polygon.Label = polygonAPI.Label
	polygon.PopupHTML = polygonAPI.PopupHTML
	polygon.FillColor = polygonAPI.FillColor
	polygon.FillOpacity = polygonAPI.FillOpacity
	polygon.StrokeColor = polygonAPI.StrokeColor
	polygon.StrokeOpacity = polygonAPI.StrokeOpacity
	polygon.StrokeWidth = polygonAPI.StrokeWidth
	polygon.IsExtruded = polygonAPI.IsExtruded
	polygon.BaseAltitude = polygonAPI.BaseAltitude
	polygon.ExtrusionHeight = polygonAPI.ExtrusionHeight

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(polygonAPI.PolygonPointersEncoding.Coordinates)) {
		console.error('Rects is not an array:', polygonAPI.PolygonPointersEncoding.Coordinates);
		return;
	}

	polygon.Coordinates = new Array<Coordinate>()
	for (let _id of polygonAPI.PolygonPointersEncoding.Coordinates) {
		let _coordinate = frontRepo.map_ID_Coordinate.get(_id)
		if (_coordinate != undefined) {
			polygon.Coordinates.push(_coordinate!)
		}
	}
}
