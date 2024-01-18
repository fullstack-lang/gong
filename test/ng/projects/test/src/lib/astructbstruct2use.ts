// generated code - do not edit

import { AstructBstruct2UseDB } from './astructbstruct2use-db'

// insertion point for imports
import { Bstruct } from './bstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstruct2Use {

	static GONGSTRUCT_NAME = "AstructBstruct2Use"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Bstrcut2?: Bstruct

}

export function CopyAstructBstruct2UseToAstructBstruct2UseDB(astructbstruct2use: AstructBstruct2Use, astructbstruct2useDB: AstructBstruct2UseDB) {

	// insertion point for basic fields declarations
	astructbstruct2useDB.Name = astructbstruct2use.Name

}
