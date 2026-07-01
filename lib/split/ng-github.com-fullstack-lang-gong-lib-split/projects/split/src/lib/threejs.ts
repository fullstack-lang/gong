// generated code - do not edit

import { ThreejsAPI } from './threejs-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Threejs {

	static GONGSTRUCT_NAME = "Threejs"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyThreejsToThreejsAPI(threejs: Threejs, threejsAPI: ThreejsAPI) {

	threejsAPI.CreatedAt = threejs.CreatedAt
	threejsAPI.DeletedAt = threejs.DeletedAt
	threejsAPI.ID = threejs.ID

	// insertion point for basic fields copy operations
	threejsAPI.Name = threejs.Name
	threejsAPI.StackName = threejs.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyThreejsAPIToThreejs update basic, pointers and slice of pointers fields of threejs
// from respectively the basic fields and encoded fields of pointers and slices of pointers of threejsAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyThreejsAPIToThreejs(threejsAPI: ThreejsAPI, threejs: Threejs, frontRepo: FrontRepo) {

	threejs.CreatedAt = threejsAPI.CreatedAt
	threejs.DeletedAt = threejsAPI.DeletedAt
	threejs.ID = threejsAPI.ID

	// insertion point for basic fields copy operations
	threejs.Name = threejsAPI.Name
	threejs.StackName = threejsAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
