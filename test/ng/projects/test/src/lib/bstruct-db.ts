// insertion point for imports
import { AstructDB } from './astruct-db'
import { DstructDB } from './dstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BstructDB {

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

	BstructPointersEncoding: BstructPointersEncoding = new BstructPointersEncoding
}

export class BstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	// reverse pointers encoding (to be removed)
	Astruct_AnarrayofbDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofbDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Astruct.Anarrayofb
	Astruct_Anarrayofb_reverse?: AstructDB 

	// reverse pointers encoding (to be removed)
	Astruct_AnotherarrayofbDBID: NullInt64 = new NullInt64
	Astruct_AnotherarrayofbDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Astruct.Anotherarrayofb
	Astruct_Anotherarrayofb_reverse?: AstructDB 

	// reverse pointers encoding (to be removed)
	Dstruct_AnarrayofbDBID: NullInt64 = new NullInt64
	Dstruct_AnarrayofbDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Dstruct.Anarrayofb
	Dstruct_Anarrayofb_reverse?: DstructDB 

}
