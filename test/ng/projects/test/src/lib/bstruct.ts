// generated code - do not edit

import { BstructDB } from './bstruct-db'

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Bstruct {

	static GONGSTRUCT_NAME = "Bstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyBstructToBstructDB(bstruct: Bstruct, bstructDB: BstructDB) {

	// insertion point for basic fields copy operations
	bstructDB.Name = bstruct.Name
	bstructDB.Floatfield = bstruct.Floatfield
	bstructDB.Floatfield2 = bstruct.Floatfield2
	bstructDB.Intfield = bstruct.Intfield
	
	// insertion point for pointer fields encoding
}
