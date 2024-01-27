// generated code - do not edit

import { GongEnumValueDB } from './gongenumvalue-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumValue {

	static GONGSTRUCT_NAME = "GongEnumValue"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongEnumValueToGongEnumValueDB(gongenumvalue: GongEnumValue, gongenumvalueDB: GongEnumValueDB) {

	gongenumvalueDB.CreatedAt = gongenumvalue.CreatedAt
	gongenumvalueDB.DeletedAt = gongenumvalue.DeletedAt
	gongenumvalueDB.ID = gongenumvalue.ID

	// insertion point for basic fields copy operations
	gongenumvalueDB.Name = gongenumvalue.Name
	gongenumvalueDB.Value = gongenumvalue.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongEnumValueDBToGongEnumValue update basic, pointers and slice of pointers fields of gongenumvalue
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongenumvalueDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongEnumValueDBToGongEnumValue(gongenumvalueDB: GongEnumValueDB, gongenumvalue: GongEnumValue, frontRepo: FrontRepo) {

	gongenumvalue.CreatedAt = gongenumvalueDB.CreatedAt
	gongenumvalue.DeletedAt = gongenumvalueDB.DeletedAt
	gongenumvalue.ID = gongenumvalueDB.ID

	// insertion point for basic fields copy operations
	gongenumvalue.Name = gongenumvalueDB.Name
	gongenumvalue.Value = gongenumvalueDB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
