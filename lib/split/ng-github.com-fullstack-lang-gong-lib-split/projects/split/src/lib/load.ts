// generated code - do not edit

import { LoadAPI } from './load-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Load {

	static GONGSTRUCT_NAME = "Load"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyLoadToLoadAPI(load: Load, loadAPI: LoadAPI) {

	loadAPI.CreatedAt = load.CreatedAt
	loadAPI.DeletedAt = load.DeletedAt
	loadAPI.ID = load.ID

	// insertion point for basic fields copy operations
	loadAPI.Name = load.Name
	loadAPI.StackName = load.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLoadAPIToLoad update basic, pointers and slice of pointers fields of load
// from respectively the basic fields and encoded fields of pointers and slices of pointers of loadAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLoadAPIToLoad(loadAPI: LoadAPI, load: Load, frontRepo: FrontRepo) {

	load.CreatedAt = loadAPI.CreatedAt
	load.DeletedAt = loadAPI.DeletedAt
	load.ID = loadAPI.ID

	// insertion point for basic fields copy operations
	load.Name = loadAPI.Name
	load.StackName = loadAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
