// generated code - do not edit

import { CircleAPI } from './circle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Circle {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	Radius: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animations: Array<Animate> = []
}

export function CopyCircleToCircleAPI(circle: Circle, circleAPI: CircleAPI) {

	circleAPI.CreatedAt = circle.CreatedAt
	circleAPI.DeletedAt = circle.DeletedAt
	circleAPI.ID = circle.ID

	// insertion point for basic fields copy operations
	circleAPI.Name = circle.Name
	circleAPI.CX = circle.CX
	circleAPI.CY = circle.CY
	circleAPI.Radius = circle.Radius
	circleAPI.Color = circle.Color
	circleAPI.FillOpacity = circle.FillOpacity
	circleAPI.Stroke = circle.Stroke
	circleAPI.StrokeOpacity = circle.StrokeOpacity
	circleAPI.StrokeWidth = circle.StrokeWidth
	circleAPI.StrokeDashArray = circle.StrokeDashArray
	circleAPI.StrokeDashArrayWhenSelected = circle.StrokeDashArrayWhenSelected
	circleAPI.Transform = circle.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	circleAPI.CirclePointersEncoding.Animations = []
	for (let _animate of circle.Animations) {
		circleAPI.CirclePointersEncoding.Animations.push(_animate.ID)
	}

}

// CopyCircleAPIToCircle update basic, pointers and slice of pointers fields of circle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleAPIToCircle(circleAPI: CircleAPI, circle: Circle, frontRepo: FrontRepo) {

	circle.CreatedAt = circleAPI.CreatedAt
	circle.DeletedAt = circleAPI.DeletedAt
	circle.ID = circleAPI.ID

	// insertion point for basic fields copy operations
	circle.Name = circleAPI.Name
	circle.CX = circleAPI.CX
	circle.CY = circleAPI.CY
	circle.Radius = circleAPI.Radius
	circle.Color = circleAPI.Color
	circle.FillOpacity = circleAPI.FillOpacity
	circle.Stroke = circleAPI.Stroke
	circle.StrokeOpacity = circleAPI.StrokeOpacity
	circle.StrokeWidth = circleAPI.StrokeWidth
	circle.StrokeDashArray = circleAPI.StrokeDashArray
	circle.StrokeDashArrayWhenSelected = circleAPI.StrokeDashArrayWhenSelected
	circle.Transform = circleAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(circleAPI.CirclePointersEncoding.Animations)) {
		console.error('Rects is not an array:', circleAPI.CirclePointersEncoding.Animations);
		return;
	}

	circle.Animations = new Array<Animate>()
	for (let _id of circleAPI.CirclePointersEncoding.Animations) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			circle.Animations.push(_animate!)
		}
	}
}
