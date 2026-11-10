// generated code - do not edit

import { PolylineAPI } from './polyline-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Coordinate } from './coordinate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Polyline {

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

	// insertion point for pointers and slices of pointers declarations
	Coordinates: Array<Coordinate> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyPolylineToPolylineAPI(polyline: Polyline, polylineAPI: PolylineAPI) {

	polylineAPI.CreatedAt = polyline.CreatedAt
	polylineAPI.DeletedAt = polyline.DeletedAt
	polylineAPI.ID = polyline.ID

	// insertion point for basic fields copy operations
	polylineAPI.Name = polyline.Name
	polylineAPI.Label = polyline.Label
	polylineAPI.PopupHTML = polyline.PopupHTML
	polylineAPI.StrokeColor = polyline.StrokeColor
	polylineAPI.StrokeOpacity = polyline.StrokeOpacity
	polylineAPI.StrokeWidth = polyline.StrokeWidth
	polylineAPI.StrokeDashArray = polyline.StrokeDashArray

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polylineAPI.PolylinePointersEncoding.Coordinates = []
	for (let _coordinate of polyline.Coordinates) {
		polylineAPI.PolylinePointersEncoding.Coordinates.push(_coordinate.ID)
	}

}

// CopyPolylineAPIToPolyline update basic, pointers and slice of pointers fields of polyline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polylineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPolylineAPIToPolyline(polylineAPI: PolylineAPI, polyline: Polyline, frontRepo: FrontRepo) {

	polyline.CreatedAt = polylineAPI.CreatedAt
	polyline.DeletedAt = polylineAPI.DeletedAt
	polyline.ID = polylineAPI.ID

	// insertion point for basic fields copy operations
	polyline.Name = polylineAPI.Name
	polyline.Label = polylineAPI.Label
	polyline.PopupHTML = polylineAPI.PopupHTML
	polyline.StrokeColor = polylineAPI.StrokeColor
	polyline.StrokeOpacity = polylineAPI.StrokeOpacity
	polyline.StrokeWidth = polylineAPI.StrokeWidth
	polyline.StrokeDashArray = polylineAPI.StrokeDashArray

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(polylineAPI.PolylinePointersEncoding.Coordinates)) {
		console.error('Rects is not an array:', polylineAPI.PolylinePointersEncoding.Coordinates);
		return;
	}

	polyline.Coordinates = new Array<Coordinate>()
	for (let _id of polylineAPI.PolylinePointersEncoding.Coordinates) {
		let _coordinate = frontRepo.map_ID_Coordinate.get(_id)
		if (_coordinate != undefined) {
			polyline.Coordinates.push(_coordinate!)
		}
	}
}
