// insertion point for imports
import { BclassDB } from './bclass-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AclassDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Booleanfield?: string
	Aenum?: string
	Aenum_2?: string
	Benum?: string
	CName?: string
	CFloatfield?: number
	Floatfield?: number
	Intfield?: number
	Anotherbooleanfield?: string

	// insertion point for other declarations
	Associationtob?: BclassDB
	AssociationtobID?: NullInt64
	AssociationtobName?: string

	Anotherassociationtob_2?: BclassDB
	Anotherassociationtob_2ID?: NullInt64
	Anotherassociationtob_2Name?: string

	Anarrayofb?: Array<BclassDB>
	Anotherarrayofb?: Array<BclassDB>
	Anarrayofa?: Array<AclassDB>
	Aclass_AnarrayofaDBID?: NullInt64
	Aclass_Anarrayofa_reverse?: AclassDB

}