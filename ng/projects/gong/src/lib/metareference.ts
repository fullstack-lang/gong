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

export function CopyMetaReferenceDBToMetaReference(metareferenceDB: MetaReferenceDB, metareference: MetaReference, frontRepo: FrontRepo) {

	metareference.CreatedAt = metareferenceDB.CreatedAt
	metareference.DeletedAt = metareferenceDB.DeletedAt
	metareference.ID = metareferenceDB.ID
	
	// insertion point for basic fields copy operations
	metareference.Name = metareferenceDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
