// generated code - do not edit

import { AstructAPI } from './astruct-api'
import { FrontRepo } from './front-repo.service';

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
	Field:  = 
	Date: Date = new Date
	Date2: Date = new Date
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

	Dstruct4s: Array<Dstruct> = []
	Duration1_string?: string
	Anarrayofa: Array<Astruct> = []
	Anotherarrayofb: Array<Bstruct> = []
	AnarrayofbUse: Array<AstructBstructUse> = []
	Anarrayofb2Use: Array<AstructBstruct2Use> = []
	AnAstruct?: Astruct

}

export function CopyAstructToAstructAPI(astruct: Astruct, astructAPI: AstructAPI) {

	astructAPI.CreatedAt = astruct.CreatedAt
	astructAPI.DeletedAt = astruct.DeletedAt
	astructAPI.ID = astruct.ID

	// insertion point for basic fields copy operations
	astructAPI.Name = astruct.Name
	astructAPI.Field = astruct.Field
	astructAPI.Date = astruct.Date
	astructAPI.Date2 = astruct.Date2
	astructAPI.Booleanfield = astruct.Booleanfield
	astructAPI.Aenum = astruct.Aenum
	astructAPI.Aenum_2 = astruct.Aenum_2
	astructAPI.Benum = astruct.Benum
	astructAPI.CEnum = astruct.CEnum
	astructAPI.CName = astruct.CName
	astructAPI.CFloatfield = astruct.CFloatfield
	astructAPI.Floatfield = astruct.Floatfield
	astructAPI.Intfield = astruct.Intfield
	astructAPI.Anotherbooleanfield = astruct.Anotherbooleanfield
	astructAPI.Duration1 = astruct.Duration1
	astructAPI.TextFieldBespokeSize = astruct.TextFieldBespokeSize
	astructAPI.TextArea = astruct.TextArea

	// insertion point for pointer fields encoding
	astructAPI.AstructPointersEncoding.AssociationtobID.Valid = true
	if (astruct.Associationtob != undefined) {
		astructAPI.AstructPointersEncoding.AssociationtobID.Int64 = astruct.Associationtob.ID  
	} else {
		astructAPI.AstructPointersEncoding.AssociationtobID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.Anotherassociationtob_2ID.Valid = true
	if (astruct.Anotherassociationtob_2 != undefined) {
		astructAPI.AstructPointersEncoding.Anotherassociationtob_2ID.Int64 = astruct.Anotherassociationtob_2.ID  
	} else {
		astructAPI.AstructPointersEncoding.Anotherassociationtob_2ID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.BstructID.Valid = true
	if (astruct.Bstruct != undefined) {
		astructAPI.AstructPointersEncoding.BstructID.Int64 = astruct.Bstruct.ID  
	} else {
		astructAPI.AstructPointersEncoding.BstructID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.Bstruct2ID.Valid = true
	if (astruct.Bstruct2 != undefined) {
		astructAPI.AstructPointersEncoding.Bstruct2ID.Int64 = astruct.Bstruct2.ID  
	} else {
		astructAPI.AstructPointersEncoding.Bstruct2ID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.DstructID.Valid = true
	if (astruct.Dstruct != undefined) {
		astructAPI.AstructPointersEncoding.DstructID.Int64 = astruct.Dstruct.ID  
	} else {
		astructAPI.AstructPointersEncoding.DstructID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.Dstruct2ID.Valid = true
	if (astruct.Dstruct2 != undefined) {
		astructAPI.AstructPointersEncoding.Dstruct2ID.Int64 = astruct.Dstruct2.ID  
	} else {
		astructAPI.AstructPointersEncoding.Dstruct2ID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.Dstruct3ID.Valid = true
	if (astruct.Dstruct3 != undefined) {
		astructAPI.AstructPointersEncoding.Dstruct3ID.Int64 = astruct.Dstruct3.ID  
	} else {
		astructAPI.AstructPointersEncoding.Dstruct3ID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.Dstruct4ID.Valid = true
	if (astruct.Dstruct4 != undefined) {
		astructAPI.AstructPointersEncoding.Dstruct4ID.Int64 = astruct.Dstruct4.ID  
	} else {
		astructAPI.AstructPointersEncoding.Dstruct4ID.Int64 = 0 		
	}

	astructAPI.AstructPointersEncoding.AnAstructID.Valid = true
	if (astruct.AnAstruct != undefined) {
		astructAPI.AstructPointersEncoding.AnAstructID.Int64 = astruct.AnAstruct.ID  
	} else {
		astructAPI.AstructPointersEncoding.AnAstructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	astructAPI.AstructPointersEncoding.Anarrayofb = []
	for (let _bstruct of astruct.Anarrayofb) {
		astructAPI.AstructPointersEncoding.Anarrayofb.push(_bstruct.ID)
	}

	astructAPI.AstructPointersEncoding.Dstruct4s = []
	for (let _dstruct of astruct.Dstruct4s) {
		astructAPI.AstructPointersEncoding.Dstruct4s.push(_dstruct.ID)
	}

	astructAPI.AstructPointersEncoding.Anarrayofa = []
	for (let _astruct of astruct.Anarrayofa) {
		astructAPI.AstructPointersEncoding.Anarrayofa.push(_astruct.ID)
	}

	astructAPI.AstructPointersEncoding.Anotherarrayofb = []
	for (let _bstruct of astruct.Anotherarrayofb) {
		astructAPI.AstructPointersEncoding.Anotherarrayofb.push(_bstruct.ID)
	}

	astructAPI.AstructPointersEncoding.AnarrayofbUse = []
	for (let _astructbstructuse of astruct.AnarrayofbUse) {
		astructAPI.AstructPointersEncoding.AnarrayofbUse.push(_astructbstructuse.ID)
	}

	astructAPI.AstructPointersEncoding.Anarrayofb2Use = []
	for (let _astructbstruct2use of astruct.Anarrayofb2Use) {
		astructAPI.AstructPointersEncoding.Anarrayofb2Use.push(_astructbstruct2use.ID)
	}

}

// CopyAstructAPIToAstruct update basic, pointers and slice of pointers fields of astruct
// from respectively the basic fields and encoded fields of pointers and slices of pointers of astructAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAstructAPIToAstruct(astructAPI: AstructAPI, astruct: Astruct, frontRepo: FrontRepo) {

	astruct.CreatedAt = astructAPI.CreatedAt
	astruct.DeletedAt = astructAPI.DeletedAt
	astruct.ID = astructAPI.ID

	// insertion point for basic fields copy operations
	astruct.Name = astructAPI.Name
	astruct.Field = astructAPI.Field
	astruct.Date = astructAPI.Date
	astruct.Date2 = astructAPI.Date2
	astruct.Booleanfield = astructAPI.Booleanfield
	astruct.Aenum = astructAPI.Aenum
	astruct.Aenum_2 = astructAPI.Aenum_2
	astruct.Benum = astructAPI.Benum
	astruct.CEnum = astructAPI.CEnum
	astruct.CName = astructAPI.CName
	astruct.CFloatfield = astructAPI.CFloatfield
	astruct.Floatfield = astructAPI.Floatfield
	astruct.Intfield = astructAPI.Intfield
	astruct.Anotherbooleanfield = astructAPI.Anotherbooleanfield
	astruct.Duration1 = astructAPI.Duration1
	astruct.TextFieldBespokeSize = astructAPI.TextFieldBespokeSize
	astruct.TextArea = astructAPI.TextArea

	// insertion point for pointer fields encoding
	astruct.Associationtob = frontRepo.map_ID_Bstruct.get(astructAPI.AstructPointersEncoding.AssociationtobID.Int64)
	astruct.Anotherassociationtob_2 = frontRepo.map_ID_Bstruct.get(astructAPI.AstructPointersEncoding.Anotherassociationtob_2ID.Int64)
	astruct.Bstruct = frontRepo.map_ID_Bstruct.get(astructAPI.AstructPointersEncoding.BstructID.Int64)
	astruct.Bstruct2 = frontRepo.map_ID_Bstruct.get(astructAPI.AstructPointersEncoding.Bstruct2ID.Int64)
	astruct.Dstruct = frontRepo.map_ID_Dstruct.get(astructAPI.AstructPointersEncoding.DstructID.Int64)
	astruct.Dstruct2 = frontRepo.map_ID_Dstruct.get(astructAPI.AstructPointersEncoding.Dstruct2ID.Int64)
	astruct.Dstruct3 = frontRepo.map_ID_Dstruct.get(astructAPI.AstructPointersEncoding.Dstruct3ID.Int64)
	astruct.Dstruct4 = frontRepo.map_ID_Dstruct.get(astructAPI.AstructPointersEncoding.Dstruct4ID.Int64)
	astruct.AnAstruct = frontRepo.map_ID_Astruct.get(astructAPI.AstructPointersEncoding.AnAstructID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(astructAPI.AstructPointersEncoding.Anarrayofb)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.Anarrayofb);
		return;
	}

	astruct.Anarrayofb = new Array<Bstruct>()
	for (let _id of astructAPI.AstructPointersEncoding.Anarrayofb) {
		let _bstruct = frontRepo.map_ID_Bstruct.get(_id)
		if (_bstruct != undefined) {
			astruct.Anarrayofb.push(_bstruct!)
		}
	}
	if (!Array.isArray(astructAPI.AstructPointersEncoding.Dstruct4s)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.Dstruct4s);
		return;
	}

	astruct.Dstruct4s = new Array<Dstruct>()
	for (let _id of astructAPI.AstructPointersEncoding.Dstruct4s) {
		let _dstruct = frontRepo.map_ID_Dstruct.get(_id)
		if (_dstruct != undefined) {
			astruct.Dstruct4s.push(_dstruct!)
		}
	}
	if (!Array.isArray(astructAPI.AstructPointersEncoding.Anarrayofa)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.Anarrayofa);
		return;
	}

	astruct.Anarrayofa = new Array<Astruct>()
	for (let _id of astructAPI.AstructPointersEncoding.Anarrayofa) {
		let _astruct = frontRepo.map_ID_Astruct.get(_id)
		if (_astruct != undefined) {
			astruct.Anarrayofa.push(_astruct!)
		}
	}
	if (!Array.isArray(astructAPI.AstructPointersEncoding.Anotherarrayofb)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.Anotherarrayofb);
		return;
	}

	astruct.Anotherarrayofb = new Array<Bstruct>()
	for (let _id of astructAPI.AstructPointersEncoding.Anotherarrayofb) {
		let _bstruct = frontRepo.map_ID_Bstruct.get(_id)
		if (_bstruct != undefined) {
			astruct.Anotherarrayofb.push(_bstruct!)
		}
	}
	if (!Array.isArray(astructAPI.AstructPointersEncoding.AnarrayofbUse)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.AnarrayofbUse);
		return;
	}

	astruct.AnarrayofbUse = new Array<AstructBstructUse>()
	for (let _id of astructAPI.AstructPointersEncoding.AnarrayofbUse) {
		let _astructbstructuse = frontRepo.map_ID_AstructBstructUse.get(_id)
		if (_astructbstructuse != undefined) {
			astruct.AnarrayofbUse.push(_astructbstructuse!)
		}
	}
	if (!Array.isArray(astructAPI.AstructPointersEncoding.Anarrayofb2Use)) {
		console.error('Rects is not an array:', astructAPI.AstructPointersEncoding.Anarrayofb2Use);
		return;
	}

	astruct.Anarrayofb2Use = new Array<AstructBstruct2Use>()
	for (let _id of astructAPI.AstructPointersEncoding.Anarrayofb2Use) {
		let _astructbstruct2use = frontRepo.map_ID_AstructBstruct2Use.get(_id)
		if (_astructbstruct2use != undefined) {
			astruct.Anarrayofb2Use.push(_astructbstruct2use!)
		}
	}
}
