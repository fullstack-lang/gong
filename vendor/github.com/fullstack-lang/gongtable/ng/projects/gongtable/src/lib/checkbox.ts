// generated code - do not edit

import { CheckBoxDB } from './checkbox-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CheckBox {

	static GONGSTRUCT_NAME = "CheckBox"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: boolean = false

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCheckBoxToCheckBoxDB(checkbox: CheckBox, checkboxDB: CheckBoxDB) {

	checkboxDB.CreatedAt = checkbox.CreatedAt
	checkboxDB.DeletedAt = checkbox.DeletedAt
	checkboxDB.ID = checkbox.ID

	// insertion point for basic fields copy operations
	checkboxDB.Name = checkbox.Name
	checkboxDB.Value = checkbox.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCheckBoxDBToCheckBox update basic, pointers and slice of pointers fields of checkbox
// from respectively the basic fields and encoded fields of pointers and slices of pointers of checkboxDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCheckBoxDBToCheckBox(checkboxDB: CheckBoxDB, checkbox: CheckBox, frontRepo: FrontRepo) {

	checkbox.CreatedAt = checkboxDB.CreatedAt
	checkbox.DeletedAt = checkboxDB.DeletedAt
	checkbox.ID = checkboxDB.ID

	// insertion point for basic fields copy operations
	checkbox.Name = checkboxDB.Name
	checkbox.Value = checkboxDB.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
