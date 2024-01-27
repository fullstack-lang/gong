// generated code - do not edit

import { CircleDB } from './circle-db'
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
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animations: Array<Animate> = []
}

export function CopyCircleToCircleDB(circle: Circle, circleDB: CircleDB) {

	circleDB.CreatedAt = circle.CreatedAt
	circleDB.DeletedAt = circle.DeletedAt
	circleDB.ID = circle.ID

	// insertion point for basic fields copy operations
	circleDB.Name = circle.Name
	circleDB.CX = circle.CX
	circleDB.CY = circle.CY
	circleDB.Radius = circle.Radius
	circleDB.Color = circle.Color
	circleDB.FillOpacity = circle.FillOpacity
	circleDB.Stroke = circle.Stroke
	circleDB.StrokeWidth = circle.StrokeWidth
	circleDB.StrokeDashArray = circle.StrokeDashArray
	circleDB.StrokeDashArrayWhenSelected = circle.StrokeDashArrayWhenSelected
	circleDB.Transform = circle.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	circleDB.CirclePointersEncoding.Animations = []
	for (let _animate of circle.Animations) {
		circleDB.CirclePointersEncoding.Animations.push(_animate.ID)
	}

}

// CopyCircleDBToCircle update basic, pointers and slice of pointers fields of circle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circleDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleDBToCircle(circleDB: CircleDB, circle: Circle, frontRepo: FrontRepo) {

	circle.CreatedAt = circleDB.CreatedAt
	circle.DeletedAt = circleDB.DeletedAt
	circle.ID = circleDB.ID

	// insertion point for basic fields copy operations
	circle.Name = circleDB.Name
	circle.CX = circleDB.CX
	circle.CY = circleDB.CY
	circle.Radius = circleDB.Radius
	circle.Color = circleDB.Color
	circle.FillOpacity = circleDB.FillOpacity
	circle.Stroke = circleDB.Stroke
	circle.StrokeWidth = circleDB.StrokeWidth
	circle.StrokeDashArray = circleDB.StrokeDashArray
	circle.StrokeDashArrayWhenSelected = circleDB.StrokeDashArrayWhenSelected
	circle.Transform = circleDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	circle.Animations = new Array<Animate>()
	for (let _id of circleDB.CirclePointersEncoding.Animations) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			circle.Animations.push(_animate!)
		}
	}
}
