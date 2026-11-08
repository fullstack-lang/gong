// generated code - do not edit

import { MaplibreAPI } from './maplibre-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Maplibre {

	static GONGSTRUCT_NAME = "Maplibre"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMaplibreToMaplibreAPI(maplibre: Maplibre, maplibreAPI: MaplibreAPI) {

	maplibreAPI.CreatedAt = maplibre.CreatedAt
	maplibreAPI.DeletedAt = maplibre.DeletedAt
	maplibreAPI.ID = maplibre.ID

	// insertion point for basic fields copy operations
	maplibreAPI.Name = maplibre.Name
	maplibreAPI.StackName = maplibre.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMaplibreAPIToMaplibre update basic, pointers and slice of pointers fields of maplibre
// from respectively the basic fields and encoded fields of pointers and slices of pointers of maplibreAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMaplibreAPIToMaplibre(maplibreAPI: MaplibreAPI, maplibre: Maplibre, frontRepo: FrontRepo) {

	maplibre.CreatedAt = maplibreAPI.CreatedAt
	maplibre.DeletedAt = maplibreAPI.DeletedAt
	maplibre.ID = maplibreAPI.ID

	// insertion point for basic fields copy operations
	maplibre.Name = maplibreAPI.Name
	maplibre.StackName = maplibreAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
