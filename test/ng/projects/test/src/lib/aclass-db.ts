// insertion point for imports
import { BclassDB } from './bclass-db'
import { AclassBclassUseDB } from './aclassbclassuse-db'
import { AclassBclass2UseDB } from './aclassbclass2use-db'

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
	Aenum?: string
	Aenum_2?: string
	Benum?: string
	CName?: string
	CFloatfield?: number
	Floatfield?: number
	Intfield?: number
	Anotherbooleanfield?: string
	Duration1?: number

	// insertion point for other declarations
	Duration1_string?: string
	Associationtob?: BclassDB
	AssociationtobID?: NullInt64

	Anotherassociationtob_2?: BclassDB
	Anotherassociationtob_2ID?: NullInt64

	Anarrayofb?: Array<BclassDB>
	Anotherarrayofb?: Array<BclassDB>
	Anarrayofa?: Array<AclassDB>
	AnarrayofbUse?: Array<AclassBclassUseDB>
	Anarrayofb2Use?: Array<AclassBclass2UseDB>
	Aclass_AnarrayofaDBID?: NullInt64
	Aclass_AnarrayofaDBID_Index?: NullInt64 // store the index of the aclass instance in Aclass.Anarrayofa
	Aclass_Anarrayofa_reverse?: AclassDB

}
