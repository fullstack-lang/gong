// generated code - do not edit

import { DirectionalLightAPI } from './directionallight-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DirectionalLight {

	static GONGSTRUCT_NAME = "DirectionalLight"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0
	Intensity: number = 0
	IsWithCastShadow: boolean = false

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyDirectionalLightToDirectionalLightAPI(directionallight: DirectionalLight, directionallightAPI: DirectionalLightAPI) {

	directionallightAPI.CreatedAt = directionallight.CreatedAt
	directionallightAPI.DeletedAt = directionallight.DeletedAt
	directionallightAPI.ID = directionallight.ID

	// insertion point for basic fields copy operations
	directionallightAPI.Name = directionallight.Name
	directionallightAPI.X = directionallight.X
	directionallightAPI.Y = directionallight.Y
	directionallightAPI.Z = directionallight.Z
	directionallightAPI.Intensity = directionallight.Intensity
	directionallightAPI.IsWithCastShadow = directionallight.IsWithCastShadow

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDirectionalLightAPIToDirectionalLight update basic, pointers and slice of pointers fields of directionallight
// from respectively the basic fields and encoded fields of pointers and slices of pointers of directionallightAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDirectionalLightAPIToDirectionalLight(directionallightAPI: DirectionalLightAPI, directionallight: DirectionalLight, frontRepo: FrontRepo) {

	directionallight.CreatedAt = directionallightAPI.CreatedAt
	directionallight.DeletedAt = directionallightAPI.DeletedAt
	directionallight.ID = directionallightAPI.ID

	// insertion point for basic fields copy operations
	directionallight.Name = directionallightAPI.Name
	directionallight.X = directionallightAPI.X
	directionallight.Y = directionallightAPI.Y
	directionallight.Z = directionallightAPI.Z
	directionallight.Intensity = directionallightAPI.Intensity
	directionallight.IsWithCastShadow = directionallightAPI.IsWithCastShadow

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
