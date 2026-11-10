// generated code - do not edit

import { CesiumAPI } from './cesium-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Cesium {

	static GONGSTRUCT_NAME = "Cesium"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCesiumToCesiumAPI(cesium: Cesium, cesiumAPI: CesiumAPI) {

	cesiumAPI.CreatedAt = cesium.CreatedAt
	cesiumAPI.DeletedAt = cesium.DeletedAt
	cesiumAPI.ID = cesium.ID

	// insertion point for basic fields copy operations
	cesiumAPI.Name = cesium.Name
	cesiumAPI.StackName = cesium.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCesiumAPIToCesium update basic, pointers and slice of pointers fields of cesium
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cesiumAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCesiumAPIToCesium(cesiumAPI: CesiumAPI, cesium: Cesium, frontRepo: FrontRepo) {

	cesium.CreatedAt = cesiumAPI.CreatedAt
	cesium.DeletedAt = cesiumAPI.DeletedAt
	cesium.ID = cesiumAPI.ID

	// insertion point for basic fields copy operations
	cesium.Name = cesiumAPI.Name
	cesium.StackName = cesiumAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
