// generated code - do not edit

import { EllipseAPI } from './ellipse-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Ellipse {

	static GONGSTRUCT_NAME = "Ellipse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	RX: number = 0
	RY: number = 0
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

export function CopyEllipseToEllipseAPI(ellipse: Ellipse, ellipseAPI: EllipseAPI) {

	ellipseAPI.CreatedAt = ellipse.CreatedAt
	ellipseAPI.DeletedAt = ellipse.DeletedAt
	ellipseAPI.ID = ellipse.ID

	// insertion point for basic fields copy operations
	ellipseAPI.Name = ellipse.Name
	ellipseAPI.CX = ellipse.CX
	ellipseAPI.CY = ellipse.CY
	ellipseAPI.RX = ellipse.RX
	ellipseAPI.RY = ellipse.RY
	ellipseAPI.Color = ellipse.Color
	ellipseAPI.FillOpacity = ellipse.FillOpacity
	ellipseAPI.Stroke = ellipse.Stroke
	ellipseAPI.StrokeOpacity = ellipse.StrokeOpacity
	ellipseAPI.StrokeWidth = ellipse.StrokeWidth
	ellipseAPI.StrokeDashArray = ellipse.StrokeDashArray
	ellipseAPI.StrokeDashArrayWhenSelected = ellipse.StrokeDashArrayWhenSelected
	ellipseAPI.Transform = ellipse.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	ellipseAPI.EllipsePointersEncoding.Animates = []
	for (let _animate of ellipse.Animates) {
		ellipseAPI.EllipsePointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyEllipseAPIToEllipse update basic, pointers and slice of pointers fields of ellipse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of ellipseAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyEllipseAPIToEllipse(ellipseAPI: EllipseAPI, ellipse: Ellipse, frontRepo: FrontRepo) {

	ellipse.CreatedAt = ellipseAPI.CreatedAt
	ellipse.DeletedAt = ellipseAPI.DeletedAt
	ellipse.ID = ellipseAPI.ID

	// insertion point for basic fields copy operations
	ellipse.Name = ellipseAPI.Name
	ellipse.CX = ellipseAPI.CX
	ellipse.CY = ellipseAPI.CY
	ellipse.RX = ellipseAPI.RX
	ellipse.RY = ellipseAPI.RY
	ellipse.Color = ellipseAPI.Color
	ellipse.FillOpacity = ellipseAPI.FillOpacity
	ellipse.Stroke = ellipseAPI.Stroke
	ellipse.StrokeOpacity = ellipseAPI.StrokeOpacity
	ellipse.StrokeWidth = ellipseAPI.StrokeWidth
	ellipse.StrokeDashArray = ellipseAPI.StrokeDashArray
	ellipse.StrokeDashArrayWhenSelected = ellipseAPI.StrokeDashArrayWhenSelected
	ellipse.Transform = ellipseAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(ellipseAPI.EllipsePointersEncoding.Animates)) {
		console.error('Rects is not an array:', ellipseAPI.EllipsePointersEncoding.Animates);
		return;
	}

	ellipse.Animates = new Array<Animate>()
	for (let _id of ellipseAPI.EllipsePointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			ellipse.Animates.push(_animate!)
		}
	}
}
