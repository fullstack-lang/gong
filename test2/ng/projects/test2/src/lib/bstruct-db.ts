// insertion point for imports
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class BstructDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Astruct_AnarrayofbstructDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofbstructDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Astruct.Anarrayofbstruct
	Astruct_Anarrayofbstruct_reverse: AstructDB = new AstructDB

}
