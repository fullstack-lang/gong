// generated code - do not edit

import { CylinderGeometryAPI } from './cylindergeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CylinderGeometry {

	static GONGSTRUCT_NAME = "CylinderGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RadiusTop: number = 0
	RadiusBottom: number = 0
	Height: number = 0
	RadialSegments: number = 0
	HeightSegments: number = 0
	OpenEnded: boolean = false
	ThetaStart: number = 0
	ThetaLength: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCylinderGeometryToCylinderGeometryAPI(cylindergeometry: CylinderGeometry, cylindergeometryAPI: CylinderGeometryAPI) {

	cylindergeometryAPI.CreatedAt = cylindergeometry.CreatedAt
	cylindergeometryAPI.DeletedAt = cylindergeometry.DeletedAt
	cylindergeometryAPI.ID = cylindergeometry.ID

	// insertion point for basic fields copy operations
	cylindergeometryAPI.Name = cylindergeometry.Name
	cylindergeometryAPI.RadiusTop = cylindergeometry.RadiusTop
	cylindergeometryAPI.RadiusBottom = cylindergeometry.RadiusBottom
	cylindergeometryAPI.Height = cylindergeometry.Height
	cylindergeometryAPI.RadialSegments = cylindergeometry.RadialSegments
	cylindergeometryAPI.HeightSegments = cylindergeometry.HeightSegments
	cylindergeometryAPI.OpenEnded = cylindergeometry.OpenEnded
	cylindergeometryAPI.ThetaStart = cylindergeometry.ThetaStart
	cylindergeometryAPI.ThetaLength = cylindergeometry.ThetaLength

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCylinderGeometryAPIToCylinderGeometry update basic, pointers and slice of pointers fields of cylindergeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cylindergeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCylinderGeometryAPIToCylinderGeometry(cylindergeometryAPI: CylinderGeometryAPI, cylindergeometry: CylinderGeometry, frontRepo: FrontRepo) {

	cylindergeometry.CreatedAt = cylindergeometryAPI.CreatedAt
	cylindergeometry.DeletedAt = cylindergeometryAPI.DeletedAt
	cylindergeometry.ID = cylindergeometryAPI.ID

	// insertion point for basic fields copy operations
	cylindergeometry.Name = cylindergeometryAPI.Name
	cylindergeometry.RadiusTop = cylindergeometryAPI.RadiusTop
	cylindergeometry.RadiusBottom = cylindergeometryAPI.RadiusBottom
	cylindergeometry.Height = cylindergeometryAPI.Height
	cylindergeometry.RadialSegments = cylindergeometryAPI.RadialSegments
	cylindergeometry.HeightSegments = cylindergeometryAPI.HeightSegments
	cylindergeometry.OpenEnded = cylindergeometryAPI.OpenEnded
	cylindergeometry.ThetaStart = cylindergeometryAPI.ThetaStart
	cylindergeometry.ThetaLength = cylindergeometryAPI.ThetaLength

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
