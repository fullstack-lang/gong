// insertion point for imports
import { BstructAPI } from './bstruct-api'
import { DstructAPI } from './dstruct-api'
import { AstructBstructUseAPI } from './astructbstructuse-api'
import { AstructBstruct2UseAPI } from './astructbstruct2use-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AstructAPI {

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
	TextFieldBespokeSize: string = ""
	TextArea: string = ""

	// insertion point for other decls
	CEnum_string?: string
	Duration1_string?: string

	AstructPointersEncoding: AstructPointersEncoding = new AstructPointersEncoding
}

export class AstructPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	AssociationtobID: NullInt64 = new NullInt64 // if pointer is null, Associationtob.ID = 0

	Anarrayofb: number[] = []
	Anotherassociationtob_2ID: NullInt64 = new NullInt64 // if pointer is null, Anotherassociationtob_2.ID = 0

	BstructID: NullInt64 = new NullInt64 // if pointer is null, Bstruct.ID = 0

	Bstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Bstruct2.ID = 0

	DstructID: NullInt64 = new NullInt64 // if pointer is null, Dstruct.ID = 0

	Dstruct2ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct2.ID = 0

	Dstruct3ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct3.ID = 0

	Dstruct4ID: NullInt64 = new NullInt64 // if pointer is null, Dstruct4.ID = 0

	Anarrayofa: number[] = []
	Anotherarrayofb: number[] = []
	AnarrayofbUse: number[] = []
	Anarrayofb2Use: number[] = []
	AnAstructID: NullInt64 = new NullInt64 // if pointer is null, AnAstruct.ID = 0

}
