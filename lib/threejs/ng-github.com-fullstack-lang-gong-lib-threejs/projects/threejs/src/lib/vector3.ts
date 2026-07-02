// generated code - do not edit

import { Vector3API } from './vector3-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Vector3 {

	static GONGSTRUCT_NAME = "Vector3"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyVector3ToVector3API(vector3: Vector3, vector3API: Vector3API) {

	vector3API.CreatedAt = vector3.CreatedAt
	vector3API.DeletedAt = vector3.DeletedAt
	vector3API.ID = vector3.ID

	// insertion point for basic fields copy operations
	vector3API.Name = vector3.Name
	vector3API.X = vector3.X
	vector3API.Y = vector3.Y
	vector3API.Z = vector3.Z

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyVector3APIToVector3 update basic, pointers and slice of pointers fields of vector3
// from respectively the basic fields and encoded fields of pointers and slices of pointers of vector3API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVector3APIToVector3(vector3API: Vector3API, vector3: Vector3, frontRepo: FrontRepo) {

	vector3.CreatedAt = vector3API.CreatedAt
	vector3.DeletedAt = vector3API.DeletedAt
	vector3.ID = vector3API.ID

	// insertion point for basic fields copy operations
	vector3.Name = vector3API.Name
	vector3.X = vector3API.X
	vector3.Y = vector3API.Y
	vector3.Z = vector3API.Z

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
