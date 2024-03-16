// generated code - do not edit

import { MetaReferenceAPI } from './metareference-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MetaReference {

	static GONGSTRUCT_NAME = "MetaReference"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyMetaReferenceToMetaReferenceAPI(metareference: MetaReference, metareferenceAPI: MetaReferenceAPI) {

	metareferenceAPI.CreatedAt = metareference.CreatedAt
	metareferenceAPI.DeletedAt = metareference.DeletedAt
	metareferenceAPI.ID = metareference.ID

	// insertion point for basic fields copy operations
	metareferenceAPI.Name = metareference.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMetaReferenceAPIToMetaReference update basic, pointers and slice of pointers fields of metareference
// from respectively the basic fields and encoded fields of pointers and slices of pointers of metareferenceAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMetaReferenceAPIToMetaReference(metareferenceAPI: MetaReferenceAPI, metareference: MetaReference, frontRepo: FrontRepo) {

	metareference.CreatedAt = metareferenceAPI.CreatedAt
	metareference.DeletedAt = metareferenceAPI.DeletedAt
	metareference.ID = metareferenceAPI.ID

	// insertion point for basic fields copy operations
	metareference.Name = metareferenceAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
