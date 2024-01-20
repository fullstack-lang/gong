// generated code - do not edit

import { PointDB } from './point-db'
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

export function CopyPointToPointDB(point: Point, pointDB: PointDB) {

	pointDB.CreatedAt = point.CreatedAt
	pointDB.DeletedAt = point.DeletedAt
	pointDB.ID = point.ID
	
	// insertion point for basic fields copy operations
	pointDB.Name = point.Name
	pointDB.X = point.X
	pointDB.Y = point.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyPointDBToPoint(pointDB: PointDB, point: Point, frontRepo: FrontRepo) {

	point.CreatedAt = pointDB.CreatedAt
	point.DeletedAt = pointDB.DeletedAt
	point.ID = pointDB.ID
	
	// insertion point for basic fields copy operations
	point.Name = pointDB.Name
	point.X = pointDB.X
	point.Y = pointDB.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
