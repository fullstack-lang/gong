// generated code - do not edit

import { BoxGeometryAPI } from './boxgeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BoxGeometry {

	static GONGSTRUCT_NAME = "BoxGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	Depth: number = 0
	WidthSegments: number = 0
	HeightSegments: number = 0
	DepthSegments: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyBoxGeometryToBoxGeometryAPI(boxgeometry: BoxGeometry, boxgeometryAPI: BoxGeometryAPI) {

	boxgeometryAPI.CreatedAt = boxgeometry.CreatedAt
	boxgeometryAPI.DeletedAt = boxgeometry.DeletedAt
	boxgeometryAPI.ID = boxgeometry.ID

	// insertion point for basic fields copy operations
	boxgeometryAPI.Name = boxgeometry.Name
	boxgeometryAPI.Width = boxgeometry.Width
	boxgeometryAPI.Height = boxgeometry.Height
	boxgeometryAPI.Depth = boxgeometry.Depth
	boxgeometryAPI.WidthSegments = boxgeometry.WidthSegments
	boxgeometryAPI.HeightSegments = boxgeometry.HeightSegments
	boxgeometryAPI.DepthSegments = boxgeometry.DepthSegments

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBoxGeometryAPIToBoxGeometry update basic, pointers and slice of pointers fields of boxgeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of boxgeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBoxGeometryAPIToBoxGeometry(boxgeometryAPI: BoxGeometryAPI, boxgeometry: BoxGeometry, frontRepo: FrontRepo) {

	boxgeometry.CreatedAt = boxgeometryAPI.CreatedAt
	boxgeometry.DeletedAt = boxgeometryAPI.DeletedAt
	boxgeometry.ID = boxgeometryAPI.ID

	// insertion point for basic fields copy operations
	boxgeometry.Name = boxgeometryAPI.Name
	boxgeometry.Width = boxgeometryAPI.Width
	boxgeometry.Height = boxgeometryAPI.Height
	boxgeometry.Depth = boxgeometryAPI.Depth
	boxgeometry.WidthSegments = boxgeometryAPI.WidthSegments
	boxgeometry.HeightSegments = boxgeometryAPI.HeightSegments
	boxgeometry.DepthSegments = boxgeometryAPI.DepthSegments

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
