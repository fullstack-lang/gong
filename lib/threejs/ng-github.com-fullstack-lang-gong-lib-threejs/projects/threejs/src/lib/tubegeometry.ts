// generated code - do not edit

import { TubeGeometryAPI } from './tubegeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Curve } from './curve'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TubeGeometry {

	static GONGSTRUCT_NAME = "TubeGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	TubularSegments: number = 0
	Radius: number = 0
	RadialSegments: number = 0
	Closed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Path?: Curve


	CreatedAt?: string
	DeletedAt?: string
}

export function CopyTubeGeometryToTubeGeometryAPI(tubegeometry: TubeGeometry, tubegeometryAPI: TubeGeometryAPI) {

	tubegeometryAPI.CreatedAt = tubegeometry.CreatedAt
	tubegeometryAPI.DeletedAt = tubegeometry.DeletedAt
	tubegeometryAPI.ID = tubegeometry.ID

	// insertion point for basic fields copy operations
	tubegeometryAPI.Name = tubegeometry.Name
	tubegeometryAPI.TubularSegments = tubegeometry.TubularSegments
	tubegeometryAPI.Radius = tubegeometry.Radius
	tubegeometryAPI.RadialSegments = tubegeometry.RadialSegments
	tubegeometryAPI.Closed = tubegeometry.Closed

	// insertion point for pointer fields encoding
	tubegeometryAPI.TubeGeometryPointersEncoding.PathID.Valid = true
	if (tubegeometry.Path != undefined) {
		tubegeometryAPI.TubeGeometryPointersEncoding.PathID.Int64 = tubegeometry.Path.ID  
	} else {
		tubegeometryAPI.TubeGeometryPointersEncoding.PathID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyTubeGeometryAPIToTubeGeometry update basic, pointers and slice of pointers fields of tubegeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of tubegeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTubeGeometryAPIToTubeGeometry(tubegeometryAPI: TubeGeometryAPI, tubegeometry: TubeGeometry, frontRepo: FrontRepo) {

	tubegeometry.CreatedAt = tubegeometryAPI.CreatedAt
	tubegeometry.DeletedAt = tubegeometryAPI.DeletedAt
	tubegeometry.ID = tubegeometryAPI.ID

	// insertion point for basic fields copy operations
	tubegeometry.Name = tubegeometryAPI.Name
	tubegeometry.TubularSegments = tubegeometryAPI.TubularSegments
	tubegeometry.Radius = tubegeometryAPI.Radius
	tubegeometry.RadialSegments = tubegeometryAPI.RadialSegments
	tubegeometry.Closed = tubegeometryAPI.Closed

	// insertion point for pointer fields encoding
	tubegeometry.Path = frontRepo.map_ID_Curve.get(tubegeometryAPI.TubeGeometryPointersEncoding.PathID.Int64)

	// insertion point for slice of pointers fields encoding
}
