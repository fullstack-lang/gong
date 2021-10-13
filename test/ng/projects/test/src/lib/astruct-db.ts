// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstruct2UseDB } from './astructbstruct2use-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Date: Date = new Date
	Booleanfield: boolean = false
	Aenum: string = ""
	Aenum_2: string = ""
	Benum: string = ""
	CName: string = ""
	CFloatfield: number = 0
	Floatfield: number = 0
	Intfield: number = 0
	Anotherbooleanfield: boolean = false
	Duration1: number = 0

	// insertion point for other declarations
	Duration1_string?: string
	Associationtob?: BstructDB
	AssociationtobID: NullInt64 = new NullInt64 // if pointer is null, Associationtob.ID = 0

	Anotherassociationtob_2?: BstructDB
	Anotherassociationtob_2ID: NullInt64 = new NullInt64 // if pointer is null, Anotherassociationtob_2.ID = 0

	Anarrayofb?: Array<BstructDB>
	Anotherarrayofb?: Array<BstructDB>
	Anarrayofa?: Array<AstructDB>
	AnarrayofbUse?: Array<AstructBstructUseDB>
	Anarrayofb2Use?: Array<AstructBstruct2UseDB>
	Astruct_AnarrayofaDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofaDBID_Index: NullInt64  = new NullInt64 // store the index of the astruct instance in Astruct.Anarrayofa
	Astruct_Anarrayofa_reverse: AstructDB = new AstructDB

}
