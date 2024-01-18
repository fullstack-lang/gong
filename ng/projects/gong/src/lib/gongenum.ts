// generated code - do not edit

import { GongEnumDB } from './gongenum-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongEnumValue } from './gongenumvalue'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnum {

	static GONGSTRUCT_NAME = "GongEnum"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: number = 0

	// insertion point for pointers and slices of pointers declarations
	Type_string?: string
	GongEnumValues: Array<GongEnumValue> = []
}

export function CopyGongEnumToGongEnumDB(gongenum: GongEnum, gongenumDB: GongEnumDB) {

	// insertion point for basic fields copy operations
	gongenumDB.Name = gongenum.Name
	gongenumDB.Type = gongenum.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongenumDB.GongEnumPointersEncoding.GongEnumValues = []
    for (let _gongenumvalue of gongenum.GongEnumValues) {
		gongenumDB.GongEnumPointersEncoding.GongEnumValues.push(_gongenumvalue.ID)
    }
	
}

export function CopyGongEnumDBToGongEnum(gongenumDB: GongEnumDB, gongenum: GongEnum, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations
	gongenum.Name = gongenumDB.Name
	gongenum.Type = gongenumDB.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongenum.GongEnumValues = new Array<GongEnumValue>()
	for (let _id of gongenumDB.GongEnumPointersEncoding.GongEnumValues) {
	  let _gongenumvalue = frontRepo.GongEnumValues.get(_id)
	  if (_gongenumvalue != undefined) {
		gongenum.GongEnumValues.push(_gongenumvalue!)
	  }
	}
}
