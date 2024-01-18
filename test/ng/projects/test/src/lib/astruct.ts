// generated code - do not edit

import { AstructDB } from './astruct-db'

// insertion point for imports
import { Bstruct } from './bstruct'
import { Dstruct } from './dstruct'
import { AstructBstructUse } from './astructbstructuse'
import { AstructBstruct2Use } from './astructbstruct2use'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Astruct {

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

	// insertion point for pointers and slices of pointers declarations
	Associationtob?: Bstruct

	Anarrayofb: Array<Bstruct> = []
	Anotherassociationtob_2?: Bstruct

	CEnum_string?: string
	Bstruct?: Bstruct

	Bstruct2?: Bstruct

	Dstruct?: Dstruct

	Dstruct2?: Dstruct

	Dstruct3?: Dstruct

	Dstruct4?: Dstruct

	Duration1_string?: string
	Anarrayofa: Array<Astruct> = []
	Anotherarrayofb: Array<Bstruct> = []
	AnarrayofbUse: Array<AstructBstructUse> = []
	Anarrayofb2Use: Array<AstructBstruct2Use> = []
	AnAstruct?: Astruct

}

export function CopyAstructToAstructDB(astruct: Astruct, astructDB: AstructDB) {

	astructDB.Name = astruct.Name
	astructDB.Date = astruct.Date
}