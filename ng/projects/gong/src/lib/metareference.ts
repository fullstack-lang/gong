// generated code - do not edit

import { MetaReferenceDB } from './metareference-db'
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

export function CopyMetaReferenceToMetaReferenceDB(metareference: MetaReference, metareferenceDB: MetaReferenceDB) {

	metareferenceDB.CreatedAt = metareference.CreatedAt
	metareferenceDB.DeletedAt = metareference.DeletedAt
	metareferenceDB.ID = metareference.ID

	// insertion point for basic fields copy operations
	metareferenceDB.Name = metareference.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMetaReferenceDBToMetaReference update basic, pointers and slice of pointers fields of metareference
// from respectively the basic fields and encoded fields of pointers and slices of pointers of metareferenceDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMetaReferenceDBToMetaReference(metareferenceDB: MetaReferenceDB, metareference: MetaReference, frontRepo: FrontRepo) {

	metareference.CreatedAt = metareferenceDB.CreatedAt
	metareference.DeletedAt = metareferenceDB.DeletedAt
	metareference.ID = metareferenceDB.ID

	// insertion point for basic fields copy operations
	metareference.Name = metareferenceDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
