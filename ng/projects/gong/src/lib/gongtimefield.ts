// generated code - do not edit

import { GongTimeFieldDB } from './gongtimefield-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongTimeField {

	static GONGSTRUCT_NAME = "GongTimeField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Index: number = 0
	CompositeStructName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongTimeFieldToGongTimeFieldDB(gongtimefield: GongTimeField, gongtimefieldDB: GongTimeFieldDB) {

	gongtimefieldDB.CreatedAt = gongtimefield.CreatedAt
	gongtimefieldDB.DeletedAt = gongtimefield.DeletedAt
	gongtimefieldDB.ID = gongtimefield.ID
	
	// insertion point for basic fields copy operations
	gongtimefieldDB.Name = gongtimefield.Name
	gongtimefieldDB.Index = gongtimefield.Index
	gongtimefieldDB.CompositeStructName = gongtimefield.CompositeStructName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyGongTimeFieldDBToGongTimeField(gongtimefieldDB: GongTimeFieldDB, gongtimefield: GongTimeField, frontRepo: FrontRepo) {

	gongtimefield.CreatedAt = gongtimefieldDB.CreatedAt
	gongtimefield.DeletedAt = gongtimefieldDB.DeletedAt
	gongtimefield.ID = gongtimefieldDB.ID
	
	// insertion point for basic fields copy operations
	gongtimefield.Name = gongtimefieldDB.Name
	gongtimefield.Index = gongtimefieldDB.Index
	gongtimefield.CompositeStructName = gongtimefieldDB.CompositeStructName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
