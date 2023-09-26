// insertion point for imports
import { BstructDB } from './bstruct-db'
import { DstructDB } from './dstruct-db'
import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstruct2UseDB } from './astructbstruct2use-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructDB {

	static GONGSTRUCT_NAME = "Astruct"

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
	CEnum: number = 0
	CName: string = ""
	CFloatfield: number = 0
	Floatfield: number = 0
	Intfield: number = 0
	Anotherbooleanfield: boolean = false
	Duration1: number = 0
	StructRef: string = ""
	FieldRef: string = ""
	EnumIntRef: string = ""
	EnumStringRef: string = ""
	EnumValue: string = ""
	ConstIdentifierValue: string = ""
	TextArea: string = ""

	// insertion point for other declarations
	Associationtob?: BstructDB
	AssociationtobID: NullInt64 = new NullInt64 // if pointer is null, Associationtob.ID = 0

	Anarrayofb?: Array<BstructDB>
	Anotherassociationtob_2?: BstructDB
	Anotherassociationtob_2ID: NullInt64 = new NullInt64 // if pointer is null, Anotherassociationtob_2.ID = 0

	CEnum_string?: string
	Bstruct?: BstructDB
	BstructID: NullInt64 = new NullInt64 // if pointer is null, Bstruct.ID = 0

	Bstruct2?: BstructDB
	Bstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Bstruct2.ID = 0

	Dstruct?: DstructDB
	DstructID: NullInt64 = new NullInt64 // if pointer is null, Dstruct.ID = 0

	Dstruct2?: DstructDB
	Dstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct2.ID = 0

	Dstruct3?: DstructDB
	Dstruct3ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct3.ID = 0

	Dstruct4?: DstructDB
	Dstruct4ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct4.ID = 0

	Duration1_string?: string
	Anarrayofa?: Array<AstructDB>
	Anotherarrayofb?: Array<BstructDB>
	AnarrayofbUse?: Array<AstructBstructUseDB>
	Anarrayofb2Use?: Array<AstructBstruct2UseDB>
	AnAstruct?: AstructDB
	AnAstructID: NullInt64 = new NullInt64 // if pointer is null, AnAstruct.ID = 0

	Astruct_AnarrayofaDBID: NullInt64 = new NullInt64
	Astruct_AnarrayofaDBID_Index: NullInt64  = new NullInt64 // store the index of the astruct instance in Astruct.Anarrayofa
	Astruct_Anarrayofa_reverse?: AstructDB 

}
