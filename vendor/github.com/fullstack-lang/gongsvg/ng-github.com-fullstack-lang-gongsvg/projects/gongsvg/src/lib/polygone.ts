// generated code - do not edit

import { PolygoneAPI } from './polygone-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Polygone {

	static GONGSTRUCT_NAME = "Polygone"

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

export function CopyPolygoneToPolygoneAPI(polygone: Polygone, polygoneAPI: PolygoneAPI) {

	polygoneAPI.CreatedAt = polygone.CreatedAt
	polygoneAPI.DeletedAt = polygone.DeletedAt
	polygoneAPI.ID = polygone.ID

	// insertion point for basic fields copy operations
	polygoneAPI.Name = polygone.Name
	polygoneAPI.Points = polygone.Points
	polygoneAPI.Color = polygone.Color
	polygoneAPI.FillOpacity = polygone.FillOpacity
	polygoneAPI.Stroke = polygone.Stroke
	polygoneAPI.StrokeOpacity = polygone.StrokeOpacity
	polygoneAPI.StrokeWidth = polygone.StrokeWidth
	polygoneAPI.StrokeDashArray = polygone.StrokeDashArray
	polygoneAPI.StrokeDashArrayWhenSelected = polygone.StrokeDashArrayWhenSelected
	polygoneAPI.Transform = polygone.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polygoneAPI.PolygonePointersEncoding.Animates = []
	for (let _animate of polygone.Animates) {
		polygoneAPI.PolygonePointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyPolygoneAPIToPolygone update basic, pointers and slice of pointers fields of polygone
// from respectively the basic fields and encoded fields of pointers and slices of pointers of polygoneAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPolygoneAPIToPolygone(polygoneAPI: PolygoneAPI, polygone: Polygone, frontRepo: FrontRepo) {

	polygone.CreatedAt = polygoneAPI.CreatedAt
	polygone.DeletedAt = polygoneAPI.DeletedAt
	polygone.ID = polygoneAPI.ID

	// insertion point for basic fields copy operations
	polygone.Name = polygoneAPI.Name
	polygone.Points = polygoneAPI.Points
	polygone.Color = polygoneAPI.Color
	polygone.FillOpacity = polygoneAPI.FillOpacity
	polygone.Stroke = polygoneAPI.Stroke
	polygone.StrokeOpacity = polygoneAPI.StrokeOpacity
	polygone.StrokeWidth = polygoneAPI.StrokeWidth
	polygone.StrokeDashArray = polygoneAPI.StrokeDashArray
	polygone.StrokeDashArrayWhenSelected = polygoneAPI.StrokeDashArrayWhenSelected
	polygone.Transform = polygoneAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(polygoneAPI.PolygonePointersEncoding.Animates)) {
		console.error('Rects is not an array:', polygoneAPI.PolygonePointersEncoding.Animates);
		return;
	}

	polygone.Animates = new Array<Animate>()
	for (let _id of polygoneAPI.PolygonePointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			polygone.Animates.push(_animate!)
		}
	}
}
