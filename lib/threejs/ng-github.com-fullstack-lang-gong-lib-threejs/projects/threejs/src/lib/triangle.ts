// generated code - do not edit

import { TriangleAPI } from './triangle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Triangle {

	static GONGSTRUCT_NAME = "Triangle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	V1: number = 0
	V2: number = 0
	V3: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyTriangleToTriangleAPI(triangle: Triangle, triangleAPI: TriangleAPI) {

	triangleAPI.CreatedAt = triangle.CreatedAt
	triangleAPI.DeletedAt = triangle.DeletedAt
	triangleAPI.ID = triangle.ID

	// insertion point for basic fields copy operations
	triangleAPI.Name = triangle.Name
	triangleAPI.V1 = triangle.V1
	triangleAPI.V2 = triangle.V2
	triangleAPI.V3 = triangle.V3

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyTriangleAPIToTriangle update basic, pointers and slice of pointers fields of triangle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of triangleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTriangleAPIToTriangle(triangleAPI: TriangleAPI, triangle: Triangle, frontRepo: FrontRepo) {

	triangle.CreatedAt = triangleAPI.CreatedAt
	triangle.DeletedAt = triangleAPI.DeletedAt
	triangle.ID = triangleAPI.ID

	// insertion point for basic fields copy operations
	triangle.Name = triangleAPI.Name
	triangle.V1 = triangleAPI.V1
	triangle.V2 = triangleAPI.V2
	triangle.V3 = triangleAPI.V3

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
