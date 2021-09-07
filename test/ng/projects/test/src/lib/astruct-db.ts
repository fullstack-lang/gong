// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstruct2UseDB } from './astructbstruct2use-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AstructDB {
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
	Associationtob?: BstructDB
	AssociationtobID?: NullInt64

	Anotherassociationtob_2?: BstructDB
	Anotherassociationtob_2ID?: NullInt64

	Anarrayofb?: Array<BstructDB>
	Anotherarrayofb?: Array<BstructDB>
	Anarrayofa?: Array<AstructDB>
	AnarrayofbUse?: Array<AstructBstructUseDB>
	Anarrayofb2Use?: Array<AstructBstruct2UseDB>
	Astruct_AnarrayofaDBID?: NullInt64
	Astruct_AnarrayofaDBID_Index?: NullInt64 // store the index of the astruct instance in Astruct.Anarrayofa
	Astruct_Anarrayofa_reverse?: AstructDB

}
