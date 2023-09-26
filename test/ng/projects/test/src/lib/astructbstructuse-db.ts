// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstructUseDB {

	static GONGSTRUCT_NAME = "AstructBstructUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Bstruct2?: BstructDB
	Bstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Bstruct2.ID = 0

	Astruct_AnarrayofbUseDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofbUseDBID_Index: NullInt64  = new NullInt64 // store the index of the astructbstructuse instance in Astruct.AnarrayofbUse
	Astruct_AnarrayofbUse_reverse?: AstructDB 

}
