// generated code - do not edit

import { ExtrudeGeometryAPI } from './extrudegeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Shape } from './shape'
import { Curve } from './curve'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ExtrudeGeometry {

	static GONGSTRUCT_NAME = "ExtrudeGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Steps: number = 0

	// insertion point for pointers and slices of pointers declarations
	Shape?: Shape

	ExtrudePath?: Curve


	CreatedAt?: string
	DeletedAt?: string
}

export function CopyExtrudeGeometryToExtrudeGeometryAPI(extrudegeometry: ExtrudeGeometry, extrudegeometryAPI: ExtrudeGeometryAPI) {

	extrudegeometryAPI.CreatedAt = extrudegeometry.CreatedAt
	extrudegeometryAPI.DeletedAt = extrudegeometry.DeletedAt
	extrudegeometryAPI.ID = extrudegeometry.ID

	// insertion point for basic fields copy operations
	extrudegeometryAPI.Name = extrudegeometry.Name
	extrudegeometryAPI.Steps = extrudegeometry.Steps

	// insertion point for pointer fields encoding
	extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ShapeID.Valid = true
	if (extrudegeometry.Shape != undefined) {
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ShapeID.Int64 = extrudegeometry.Shape.ID  
	} else {
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ShapeID.Int64 = 0 		
	}

	extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ExtrudePathID.Valid = true
	if (extrudegeometry.ExtrudePath != undefined) {
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ExtrudePathID.Int64 = extrudegeometry.ExtrudePath.ID  
	} else {
		extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ExtrudePathID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyExtrudeGeometryAPIToExtrudeGeometry update basic, pointers and slice of pointers fields of extrudegeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of extrudegeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyExtrudeGeometryAPIToExtrudeGeometry(extrudegeometryAPI: ExtrudeGeometryAPI, extrudegeometry: ExtrudeGeometry, frontRepo: FrontRepo) {

	extrudegeometry.CreatedAt = extrudegeometryAPI.CreatedAt
	extrudegeometry.DeletedAt = extrudegeometryAPI.DeletedAt
	extrudegeometry.ID = extrudegeometryAPI.ID

	// insertion point for basic fields copy operations
	extrudegeometry.Name = extrudegeometryAPI.Name
	extrudegeometry.Steps = extrudegeometryAPI.Steps

	// insertion point for pointer fields encoding
	extrudegeometry.Shape = frontRepo.map_ID_Shape.get(extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ShapeID.Int64)
	extrudegeometry.ExtrudePath = frontRepo.map_ID_Curve.get(extrudegeometryAPI.ExtrudeGeometryPointersEncoding.ExtrudePathID.Int64)

	// insertion point for slice of pointers fields encoding
}
