// generated code - do not edit

import { AmbiantLightAPI } from './ambiantlight-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AmbiantLight {

	static GONGSTRUCT_NAME = "AmbiantLight"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Intensity: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyAmbiantLightToAmbiantLightAPI(ambiantlight: AmbiantLight, ambiantlightAPI: AmbiantLightAPI) {

	ambiantlightAPI.CreatedAt = ambiantlight.CreatedAt
	ambiantlightAPI.DeletedAt = ambiantlight.DeletedAt
	ambiantlightAPI.ID = ambiantlight.ID

	// insertion point for basic fields copy operations
	ambiantlightAPI.Name = ambiantlight.Name
	ambiantlightAPI.Intensity = ambiantlight.Intensity

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAmbiantLightAPIToAmbiantLight update basic, pointers and slice of pointers fields of ambiantlight
// from respectively the basic fields and encoded fields of pointers and slices of pointers of ambiantlightAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAmbiantLightAPIToAmbiantLight(ambiantlightAPI: AmbiantLightAPI, ambiantlight: AmbiantLight, frontRepo: FrontRepo) {

	ambiantlight.CreatedAt = ambiantlightAPI.CreatedAt
	ambiantlight.DeletedAt = ambiantlightAPI.DeletedAt
	ambiantlight.ID = ambiantlightAPI.ID

	// insertion point for basic fields copy operations
	ambiantlight.Name = ambiantlightAPI.Name
	ambiantlight.Intensity = ambiantlightAPI.Intensity

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
