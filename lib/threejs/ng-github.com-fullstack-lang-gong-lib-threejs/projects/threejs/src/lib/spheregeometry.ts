// generated code - do not edit

import { SphereGeometryAPI } from './spheregeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SphereGeometry {

	static GONGSTRUCT_NAME = "SphereGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Radius: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0
	PhiStart: number = 0
	PhiLength: number = 0
	ThetaStart: number = 0
	ThetaLength: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopySphereGeometryToSphereGeometryAPI(spheregeometry: SphereGeometry, spheregeometryAPI: SphereGeometryAPI) {

	spheregeometryAPI.CreatedAt = spheregeometry.CreatedAt
	spheregeometryAPI.DeletedAt = spheregeometry.DeletedAt
	spheregeometryAPI.ID = spheregeometry.ID

	// insertion point for basic fields copy operations
	spheregeometryAPI.Name = spheregeometry.Name
	spheregeometryAPI.Radius = spheregeometry.Radius
	spheregeometryAPI.WidthSegments = spheregeometry.WidthSegments
	spheregeometryAPI.HeightSegments = spheregeometry.HeightSegments
	spheregeometryAPI.PhiStart = spheregeometry.PhiStart
	spheregeometryAPI.PhiLength = spheregeometry.PhiLength
	spheregeometryAPI.ThetaStart = spheregeometry.ThetaStart
	spheregeometryAPI.ThetaLength = spheregeometry.ThetaLength

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySphereGeometryAPIToSphereGeometry update basic, pointers and slice of pointers fields of spheregeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of spheregeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySphereGeometryAPIToSphereGeometry(spheregeometryAPI: SphereGeometryAPI, spheregeometry: SphereGeometry, frontRepo: FrontRepo) {

	spheregeometry.CreatedAt = spheregeometryAPI.CreatedAt
	spheregeometry.DeletedAt = spheregeometryAPI.DeletedAt
	spheregeometry.ID = spheregeometryAPI.ID

	// insertion point for basic fields copy operations
	spheregeometry.Name = spheregeometryAPI.Name
	spheregeometry.Radius = spheregeometryAPI.Radius
	spheregeometry.WidthSegments = spheregeometryAPI.WidthSegments
	spheregeometry.HeightSegments = spheregeometryAPI.HeightSegments
	spheregeometry.PhiStart = spheregeometryAPI.PhiStart
	spheregeometry.PhiLength = spheregeometryAPI.PhiLength
	spheregeometry.ThetaStart = spheregeometryAPI.ThetaStart
	spheregeometry.ThetaLength = spheregeometryAPI.ThetaLength

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
