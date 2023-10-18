// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructBstruct2UseDB {

	static GONGSTRUCT_NAME = "AstructBstruct2Use"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Bstrcut2?: BstructDB


	AstructBstruct2UsePointersEncoding: AstructBstruct2UsePointersEncoding = new AstructBstruct2UsePointersEncoding
}

export class AstructBstruct2UsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Bstrcut2ID: NullInt64 = new NullInt64 // if pointer is null, Bstrcut2.ID = 0

	// reverse pointers encoding (to be removed)
	Astruct_Anarrayofb2UseDBID: NullInt64 = new NullInt64
	Astruct_Anarrayofb2UseDBID_Index: NullInt64  = new NullInt64 // store the index of the astructbstruct2use instance in Astruct.Anarrayofb2Use
	Astruct_Anarrayofb2Use_reverse?: AstructDB 

}
