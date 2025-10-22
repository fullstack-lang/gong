// generated code - do not edit

import { PolylineAPI } from './polyline-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Polyline {

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
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyPolylineToPolylineAPI(polyline: Polyline, polylineAPI: PolylineAPI) {

	polylineAPI.CreatedAt = polyline.CreatedAt
	polylineAPI.DeletedAt = polyline.DeletedAt
	polylineAPI.ID = polyline.ID

	// insertion point for basic fields copy operations
	polylineAPI.Name = polyline.Name
	polylineAPI.Points = polyline.Points
	polylineAPI.Color = polyline.Color
	polylineAPI.FillOpacity = polyline.FillOpacity
	polylineAPI.Stroke = polyline.Stroke
	polylineAPI.StrokeOpacity = polyline.StrokeOpacity
	polylineAPI.StrokeWidth = polyline.StrokeWidth
	polylineAPI.StrokeDashArray = polyline.StrokeDashArray
	polylineAPI.StrokeDashArrayWhenSelected = polyline.StrokeDashArrayWhenSelected
	polylineAPI.Transform = polyline.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polylineAPI.PolylinePointersEncoding.Animates = []
	for (let _animate of polyline.Animates) {
		polylineAPI.PolylinePointersEncoding.Animates.push(_animate.ID)
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
	polyline.Points = polylineAPI.Points
	polyline.Color = polylineAPI.Color
	polyline.FillOpacity = polylineAPI.FillOpacity
	polyline.Stroke = polylineAPI.Stroke
	polyline.StrokeOpacity = polylineAPI.StrokeOpacity
	polyline.StrokeWidth = polylineAPI.StrokeWidth
	polyline.StrokeDashArray = polylineAPI.StrokeDashArray
	polyline.StrokeDashArrayWhenSelected = polylineAPI.StrokeDashArrayWhenSelected
	polyline.Transform = polylineAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(polylineAPI.PolylinePointersEncoding.Animates)) {
		console.error('Rects is not an array:', polylineAPI.PolylinePointersEncoding.Animates);
		return;
	}

	polyline.Animates = new Array<Animate>()
	for (let _id of polylineAPI.PolylinePointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			polyline.Animates.push(_animate!)
		}
	}
}
