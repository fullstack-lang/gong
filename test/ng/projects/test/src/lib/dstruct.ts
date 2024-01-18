// generated code - do not edit
// insertion point for imports
import { Bstruct } from './bstruct'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Dstruct {

	static GONGSTRUCT_NAME = "Dstruct"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Anarrayofb: Array<Bstruct> = []
}
