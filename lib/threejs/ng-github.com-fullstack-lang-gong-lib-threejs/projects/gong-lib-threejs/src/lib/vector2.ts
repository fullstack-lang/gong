// generated code - do not edit

import { Vector2API } from './vector2-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Vector2 {

	static GONGSTRUCT_NAME = "Vector2"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyVector2ToVector2API(vector2: Vector2, vector2API: Vector2API) {

	vector2API.CreatedAt = vector2.CreatedAt
	vector2API.DeletedAt = vector2.DeletedAt
	vector2API.ID = vector2.ID

	// insertion point for basic fields copy operations
	vector2API.Name = vector2.Name
	vector2API.X = vector2.X
	vector2API.Y = vector2.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyVector2APIToVector2 update basic, pointers and slice of pointers fields of vector2
// from respectively the basic fields and encoded fields of pointers and slices of pointers of vector2API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVector2APIToVector2(vector2API: Vector2API, vector2: Vector2, frontRepo: FrontRepo) {

	vector2.CreatedAt = vector2API.CreatedAt
	vector2.DeletedAt = vector2API.DeletedAt
	vector2.ID = vector2API.ID

	// insertion point for basic fields copy operations
	vector2.Name = vector2API.Name
	vector2.X = vector2API.X
	vector2.Y = vector2API.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
