// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Anarrayofa?: Array<AstructDB>
	Astruct_AnarrayofaDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofaDBID_Index: NullInt64  = new NullInt64 // store the index of the astruct instance in Astruct.Anarrayofa
	Astruct_Anarrayofa_reverse?: AstructDB 

}
