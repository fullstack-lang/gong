// generated code - do not edit

import { PolylineDB } from './polyline-db'
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
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyPolylineToPolylineDB(polyline: Polyline, polylineDB: PolylineDB) {

	polylineDB.CreatedAt = polyline.CreatedAt
	polylineDB.DeletedAt = polyline.DeletedAt
	polylineDB.ID = polyline.ID

	// insertion point for basic fields copy operations
	polylineDB.Name = polyline.Name
	polylineDB.Points = polyline.Points
	polylineDB.Color = polyline.Color
	polylineDB.FillOpacity = polyline.FillOpacity
	polylineDB.Stroke = polyline.Stroke
	polylineDB.StrokeWidth = polyline.StrokeWidth
	polylineDB.StrokeDashArray = polyline.StrokeDashArray
	polylineDB.StrokeDashArrayWhenSelected = polyline.StrokeDashArrayWhenSelected
	polylineDB.Transform = polyline.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polylineDB.PolylinePointersEncoding.Animates = []
	for (let _animate of polyline.Animates) {
		polylineDB.PolylinePointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyPolylineDBToPolyline update basic, pointers and slice of pointers fields of polyline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polylineDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPolylineDBToPolyline(polylineDB: PolylineDB, polyline: Polyline, frontRepo: FrontRepo) {

	polyline.CreatedAt = polylineDB.CreatedAt
	polyline.DeletedAt = polylineDB.DeletedAt
	polyline.ID = polylineDB.ID

	// insertion point for basic fields copy operations
	polyline.Name = polylineDB.Name
	polyline.Points = polylineDB.Points
	polyline.Color = polylineDB.Color
	polyline.FillOpacity = polylineDB.FillOpacity
	polyline.Stroke = polylineDB.Stroke
	polyline.StrokeWidth = polylineDB.StrokeWidth
	polyline.StrokeDashArray = polylineDB.StrokeDashArray
	polyline.StrokeDashArrayWhenSelected = polylineDB.StrokeDashArrayWhenSelected
	polyline.Transform = polylineDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polyline.Animates = new Array<Animate>()
	for (let _id of polylineDB.PolylinePointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			polyline.Animates.push(_animate!)
		}
	}
}
