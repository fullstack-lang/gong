// generated code - do not edit

import { ClippingPlaneAPI } from './clippingplane-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ClippingPlane {

	static GONGSTRUCT_NAME = "ClippingPlane"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsEnabled: boolean = false
	NormalX: number = 0
	NormalY: number = 0
	NormalZ: number = 0
	Constant: number = 0
	IsHelperVisible: boolean = false
	HelperSize: number = 0
	HelperCenterX: number = 0
	HelperCenterY: number = 0
	HelperCenterZ: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyClippingPlaneToClippingPlaneAPI(clippingplane: ClippingPlane, clippingplaneAPI: ClippingPlaneAPI) {

	clippingplaneAPI.CreatedAt = clippingplane.CreatedAt
	clippingplaneAPI.DeletedAt = clippingplane.DeletedAt
	clippingplaneAPI.ID = clippingplane.ID

	// insertion point for basic fields copy operations
	clippingplaneAPI.Name = clippingplane.Name
	clippingplaneAPI.IsEnabled = clippingplane.IsEnabled
	clippingplaneAPI.NormalX = clippingplane.NormalX
	clippingplaneAPI.NormalY = clippingplane.NormalY
	clippingplaneAPI.NormalZ = clippingplane.NormalZ
	clippingplaneAPI.Constant = clippingplane.Constant
	clippingplaneAPI.IsHelperVisible = clippingplane.IsHelperVisible
	clippingplaneAPI.HelperSize = clippingplane.HelperSize
	clippingplaneAPI.HelperCenterX = clippingplane.HelperCenterX
	clippingplaneAPI.HelperCenterY = clippingplane.HelperCenterY
	clippingplaneAPI.HelperCenterZ = clippingplane.HelperCenterZ

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyClippingPlaneAPIToClippingPlane update basic, pointers and slice of pointers fields of clippingplane
// from respectively the basic fields and encoded fields of pointers and slices of pointers of clippingplaneAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyClippingPlaneAPIToClippingPlane(clippingplaneAPI: ClippingPlaneAPI, clippingplane: ClippingPlane, frontRepo: FrontRepo) {

	clippingplane.CreatedAt = clippingplaneAPI.CreatedAt
	clippingplane.DeletedAt = clippingplaneAPI.DeletedAt
	clippingplane.ID = clippingplaneAPI.ID

	// insertion point for basic fields copy operations
	clippingplane.Name = clippingplaneAPI.Name
	clippingplane.IsEnabled = clippingplaneAPI.IsEnabled
	clippingplane.NormalX = clippingplaneAPI.NormalX
	clippingplane.NormalY = clippingplaneAPI.NormalY
	clippingplane.NormalZ = clippingplaneAPI.NormalZ
	clippingplane.Constant = clippingplaneAPI.Constant
	clippingplane.IsHelperVisible = clippingplaneAPI.IsHelperVisible
	clippingplane.HelperSize = clippingplaneAPI.HelperSize
	clippingplane.HelperCenterX = clippingplaneAPI.HelperCenterX
	clippingplane.HelperCenterY = clippingplaneAPI.HelperCenterY
	clippingplane.HelperCenterZ = clippingplaneAPI.HelperCenterZ

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
