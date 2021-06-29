// insertion point for imports
import { BclassDB } from './bclass-db'
import { AclassDB } from './aclass-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AclassBclassUseDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	Bclass2?: BclassDB
	Bclass2ID?: NullInt64

	Aclass_AnarrayofbUseDBID?: NullInt64
	Aclass_AnarrayofbUseDBID_Index?: NullInt64 // store the index of the aclassbclassuse instance in Aclass.AnarrayofbUse
	Aclass_AnarrayofbUse_reverse?: AclassDB

}
