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

	// insertion point for basic fields copy operations
	astructDB.Name = astruct.Name
	astructDB.Booleanfield = astruct.Booleanfield
	astructDB.Aenum = astruct.Aenum
	astructDB.Aenum_2 = astruct.Aenum_2
	astructDB.Benum = astruct.Benum
	astructDB.CEnum = astruct.CEnum
	astructDB.CName = astruct.CName
	astructDB.CFloatfield = astruct.CFloatfield
	astructDB.Floatfield = astruct.Floatfield
	astructDB.Intfield = astruct.Intfield
	astructDB.Anotherbooleanfield = astruct.Anotherbooleanfield
	astructDB.Duration1 = astruct.Duration1
	astructDB.StructRef = astruct.StructRef
	astructDB.FieldRef = astruct.FieldRef
	astructDB.EnumIntRef = astruct.EnumIntRef
	astructDB.EnumStringRef = astruct.EnumStringRef
	astructDB.EnumValue = astruct.EnumValue
	astructDB.ConstIdentifierValue = astruct.ConstIdentifierValue
	astructDB.TextFieldBespokeSize = astruct.TextFieldBespokeSize
	astructDB.TextArea = astruct.TextArea

	// insertion point for pointer fields encoding
	astructDB.AstructPointersEncoding.AssociationtobID.Valid = true
	if (astruct.Associationtob != undefined) {
		astructDB.AstructPointersEncoding.AssociationtobID.Int64 = astruct.Associationtob.ID
	} else {
		astructDB.AstructPointersEncoding.AssociationtobID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.Anotherassociationtob_2ID.Valid = true
	if (astruct.Anotherassociationtob_2 != undefined) {
		astructDB.AstructPointersEncoding.Anotherassociationtob_2ID.Int64 = astruct.Anotherassociationtob_2.ID
	} else {
		astructDB.AstructPointersEncoding.Anotherassociationtob_2ID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.BstructID.Valid = true
	if (astruct.Bstruct != undefined) {
		astructDB.AstructPointersEncoding.BstructID.Int64 = astruct.Bstruct.ID
	} else {
		astructDB.AstructPointersEncoding.BstructID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.Bstruct2ID.Valid = true
	if (astruct.Bstruct2 != undefined) {
		astructDB.AstructPointersEncoding.Bstruct2ID.Int64 = astruct.Bstruct2.ID
	} else {
		astructDB.AstructPointersEncoding.Bstruct2ID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.DstructID.Valid = true
	if (astruct.Dstruct != undefined) {
		astructDB.AstructPointersEncoding.DstructID.Int64 = astruct.Dstruct.ID
	} else {
		astructDB.AstructPointersEncoding.DstructID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.Dstruct2ID.Valid = true
	if (astruct.Dstruct2 != undefined) {
		astructDB.AstructPointersEncoding.Dstruct2ID.Int64 = astruct.Dstruct2.ID
	} else {
		astructDB.AstructPointersEncoding.Dstruct2ID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.Dstruct3ID.Valid = true
	if (astruct.Dstruct3 != undefined) {
		astructDB.AstructPointersEncoding.Dstruct3ID.Int64 = astruct.Dstruct3.ID
	} else {
		astructDB.AstructPointersEncoding.Dstruct3ID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.Dstruct4ID.Valid = true
	if (astruct.Dstruct4 != undefined) {
		astructDB.AstructPointersEncoding.Dstruct4ID.Int64 = astruct.Dstruct4.ID
	} else {
		astructDB.AstructPointersEncoding.Dstruct4ID.Int64 = 0
	}

	astructDB.AstructPointersEncoding.AnAstructID.Valid = true
	if (astruct.AnAstruct != undefined) {
		astructDB.AstructPointersEncoding.AnAstructID.Int64 = astruct.AnAstruct.ID
	} else {
		astructDB.AstructPointersEncoding.AnAstructID.Int64 = 0
	}


	// insertion point for slice of pointers fields encoding
	astructDB.AstructPointersEncoding.Anarrayofb = []
	for (let _bstruct of astruct.Anarrayofb) {
		astructDB.AstructPointersEncoding.Anarrayofb.push(_bstruct.ID)
	}

	astructDB.AstructPointersEncoding.Anarrayofa = []
	for (let _astruct of astruct.Anarrayofa) {
		astructDB.AstructPointersEncoding.Anarrayofa.push(_astruct.ID)
	}

	astructDB.AstructPointersEncoding.Anotherarrayofb = []
	for (let _bstruct of astruct.Anotherarrayofb) {
		astructDB.AstructPointersEncoding.Anotherarrayofb.push(_bstruct.ID)
	}

	astructDB.AstructPointersEncoding.AnarrayofbUse = []
	for (let _astructbstructuse of astruct.AnarrayofbUse) {
		astructDB.AstructPointersEncoding.AnarrayofbUse.push(_astructbstructuse.ID)
	}

	astructDB.AstructPointersEncoding.Anarrayofb2Use = []
	for (let _astructbstruct2use of astruct.Anarrayofb2Use) {
		astructDB.AstructPointersEncoding.Anarrayofb2Use.push(_astructbstruct2use.ID)
	}

}
