// insertion point for imports
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BstructDB {
	CreatedAt?: string
	DeletedAt?: string

	// With the multiple stacks, each sack has its own DB, therefore, 
	// the unique identifier of an instance is a combination of the ID and the stack path
	ID: number = 0
	GONG__StackPath: string = ""

	// insertion point for basic fields declarations
	Name: string = ""
	Floatfield: number = 0
	Floatfield2: number = 0
	Intfield: number = 0

	// insertion point for other declarations
	Astruct_AnarrayofbDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofbDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Astruct.Anarrayofb
	Astruct_Anarrayofb_reverse?: AstructDB 

	Astruct_AnotherarrayofbDBID: NullInt64 = new NullInt64
	Astruct_AnotherarrayofbDBID_Index: NullInt64  = new NullInt64 // store the index of the bstruct instance in Astruct.Anotherarrayofb
	Astruct_Anotherarrayofb_reverse?: AstructDB 

}
