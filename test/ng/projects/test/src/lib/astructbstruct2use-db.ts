// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AstructBstruct2UseDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	Bstrcut2?: BstructDB
	Bstrcut2ID?: NullInt64

	Astruct_Anarrayofb2UseDBID?: NullInt64
	Astruct_Anarrayofb2UseDBID_Index?: NullInt64 // store the index of the astructbstruct2use instance in Astruct.Anarrayofb2Use
	Astruct_Anarrayofb2Use_reverse?: AstructDB

}
