// insertion point for imports
import { ADB } from './a-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BDB {

	static GONGSTRUCT_NAME = "B"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	A_BsDBID: NullInt64 = new NullInt64
	A_BsDBID_Index: NullInt64  = new NullInt64 // store the index of the b instance in A.Bs
	A_Bs_reverse?: ADB 


	BPointersEncoding: BPointersEncoding = new BPointersEncoding
}

export class BPointersEncoding {
	// insertion point for other declarations
}
