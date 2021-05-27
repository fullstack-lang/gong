// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AclassDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Date?: Date
	Booleanfield?: string
	Floatfield?: number
	Intfield?: number
	Anotherbooleanfield?: string
	Duration1?: number

	// insertion point for other declarations
	Duration1_string?: string
	Anarrayofa?: Array<AclassDB>
	Aclass_AnarrayofaDBID?: NullInt64
	Aclass_AnarrayofaDBID_Index?: NullInt64 // store the index of the aclass instance in Aclass.Anarrayofa
	Aclass_Anarrayofa_reverse?: AclassDB

}
