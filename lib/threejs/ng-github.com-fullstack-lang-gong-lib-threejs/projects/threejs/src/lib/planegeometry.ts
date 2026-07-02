// generated code - do not edit

import { PlaneGeometryAPI } from './planegeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PlaneGeometry {

	static GONGSTRUCT_NAME = "PlaneGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyPlaneGeometryToPlaneGeometryAPI(planegeometry: PlaneGeometry, planegeometryAPI: PlaneGeometryAPI) {

	planegeometryAPI.CreatedAt = planegeometry.CreatedAt
	planegeometryAPI.DeletedAt = planegeometry.DeletedAt
	planegeometryAPI.ID = planegeometry.ID

	// insertion point for basic fields copy operations
	planegeometryAPI.Name = planegeometry.Name
	planegeometryAPI.Width = planegeometry.Width
	planegeometryAPI.Height = planegeometry.Height
	planegeometryAPI.WidthSegments = planegeometry.WidthSegments
	planegeometryAPI.HeightSegments = planegeometry.HeightSegments

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyPlaneGeometryAPIToPlaneGeometry update basic, pointers and slice of pointers fields of planegeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of planegeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPlaneGeometryAPIToPlaneGeometry(planegeometryAPI: PlaneGeometryAPI, planegeometry: PlaneGeometry, frontRepo: FrontRepo) {

	planegeometry.CreatedAt = planegeometryAPI.CreatedAt
	planegeometry.DeletedAt = planegeometryAPI.DeletedAt
	planegeometry.ID = planegeometryAPI.ID

	// insertion point for basic fields copy operations
	planegeometry.Name = planegeometryAPI.Name
	planegeometry.Width = planegeometryAPI.Width
	planegeometry.Height = planegeometryAPI.Height
	planegeometry.WidthSegments = planegeometryAPI.WidthSegments
	planegeometry.HeightSegments = planegeometryAPI.HeightSegments

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
