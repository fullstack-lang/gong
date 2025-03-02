// generated code - do not edit

import { PointAPI } from './point-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Point {

	static GONGSTRUCT_NAME = "Point"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyPointToPointAPI(point: Point, pointAPI: PointAPI) {

	pointAPI.CreatedAt = point.CreatedAt
	pointAPI.DeletedAt = point.DeletedAt
	pointAPI.ID = point.ID

	// insertion point for basic fields copy operations
	pointAPI.Name = point.Name
	pointAPI.X = point.X
	pointAPI.Y = point.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyPointAPIToPoint update basic, pointers and slice of pointers fields of point
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pointAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPointAPIToPoint(pointAPI: PointAPI, point: Point, frontRepo: FrontRepo) {

	point.CreatedAt = pointAPI.CreatedAt
	point.DeletedAt = pointAPI.DeletedAt
	point.ID = pointAPI.ID

	// insertion point for basic fields copy operations
	point.Name = pointAPI.Name
	point.X = pointAPI.X
	point.Y = pointAPI.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
