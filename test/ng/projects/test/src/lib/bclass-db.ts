// insertion point for imports
import { AclassDB } from './aclass-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class BclassDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Floatfield?: number
	Intfield?: number

	// insertion point for other declarations
	Aclass_AnarrayofbDBID?: NullInt64
	Aclass_AnarrayofbDBID_Index?: NullInt64 // store the index of the bclass instance in Aclass.Anarrayofb
	Aclass_Anarrayofb_reverse?: AclassDB

	Aclass_AnotherarrayofbDBID?: NullInt64
	Aclass_AnotherarrayofbDBID_Index?: NullInt64 // store the index of the bclass instance in Aclass.Anotherarrayofb
	Aclass_Anotherarrayofb_reverse?: AclassDB

}
