// generated code - do not edit

import { TorusGeometryAPI } from './torusgeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TorusGeometry {

	static GONGSTRUCT_NAME = "TorusGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Radius: number = 0
	Tube: number = 0
	RadialSegments: number = 0
	TubularSegments: number = 0
	Arc: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyTorusGeometryToTorusGeometryAPI(torusgeometry: TorusGeometry, torusgeometryAPI: TorusGeometryAPI) {

	torusgeometryAPI.CreatedAt = torusgeometry.CreatedAt
	torusgeometryAPI.DeletedAt = torusgeometry.DeletedAt
	torusgeometryAPI.ID = torusgeometry.ID

	// insertion point for basic fields copy operations
	torusgeometryAPI.Name = torusgeometry.Name
	torusgeometryAPI.Radius = torusgeometry.Radius
	torusgeometryAPI.Tube = torusgeometry.Tube
	torusgeometryAPI.RadialSegments = torusgeometry.RadialSegments
	torusgeometryAPI.TubularSegments = torusgeometry.TubularSegments
	torusgeometryAPI.Arc = torusgeometry.Arc

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyTorusGeometryAPIToTorusGeometry update basic, pointers and slice of pointers fields of torusgeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of torusgeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTorusGeometryAPIToTorusGeometry(torusgeometryAPI: TorusGeometryAPI, torusgeometry: TorusGeometry, frontRepo: FrontRepo) {

	torusgeometry.CreatedAt = torusgeometryAPI.CreatedAt
	torusgeometry.DeletedAt = torusgeometryAPI.DeletedAt
	torusgeometry.ID = torusgeometryAPI.ID

	// insertion point for basic fields copy operations
	torusgeometry.Name = torusgeometryAPI.Name
	torusgeometry.Radius = torusgeometryAPI.Radius
	torusgeometry.Tube = torusgeometryAPI.Tube
	torusgeometry.RadialSegments = torusgeometryAPI.RadialSegments
	torusgeometry.TubularSegments = torusgeometryAPI.TubularSegments
	torusgeometry.Arc = torusgeometryAPI.Arc

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
