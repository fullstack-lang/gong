// generated code - do not edit

import { ShapeAPI } from './shape-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Vector2 } from './vector2'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Shape {

	static GONGSTRUCT_NAME = "Shape"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Points: Array<Vector2> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyShapeToShapeAPI(shape: Shape, shapeAPI: ShapeAPI) {

	shapeAPI.CreatedAt = shape.CreatedAt
	shapeAPI.DeletedAt = shape.DeletedAt
	shapeAPI.ID = shape.ID

	// insertion point for basic fields copy operations
	shapeAPI.Name = shape.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	shapeAPI.ShapePointersEncoding.Points = []
	for (let _vector2 of shape.Points) {
		shapeAPI.ShapePointersEncoding.Points.push(_vector2.ID)
	}

}

// CopyShapeAPIToShape update basic, pointers and slice of pointers fields of shape
// from respectively the basic fields and encoded fields of pointers and slices of pointers of shapeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyShapeAPIToShape(shapeAPI: ShapeAPI, shape: Shape, frontRepo: FrontRepo) {

	shape.CreatedAt = shapeAPI.CreatedAt
	shape.DeletedAt = shapeAPI.DeletedAt
	shape.ID = shapeAPI.ID

	// insertion point for basic fields copy operations
	shape.Name = shapeAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(shapeAPI.ShapePointersEncoding.Points)) {
		console.error('Rects is not an array:', shapeAPI.ShapePointersEncoding.Points);
		return;
	}

	shape.Points = new Array<Vector2>()
	for (let _id of shapeAPI.ShapePointersEncoding.Points) {
		let _vector2 = frontRepo.map_ID_Vector2.get(_id)
		if (_vector2 != undefined) {
			shape.Points.push(_vector2!)
		}
	}
}
