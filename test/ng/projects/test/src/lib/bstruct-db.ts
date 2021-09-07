// insertion point for imports
import { AstructDB } from './astruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class BstructDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Floatfield?: number
	Intfield?: number

	// insertion point for other declarations
	Astruct_AnarrayofbDBID?: NullInt64
	Astruct_AnarrayofbDBID_Index?: NullInt64 // store the index of the bstruct instance in Astruct.Anarrayofb
	Astruct_Anarrayofb_reverse?: AstructDB

	Astruct_AnotherarrayofbDBID?: NullInt64
	Astruct_AnotherarrayofbDBID_Index?: NullInt64 // store the index of the bstruct instance in Astruct.Anotherarrayofb
	Astruct_Anotherarrayofb_reverse?: AstructDB

}
