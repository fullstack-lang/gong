// generated code - do not edit

import { CheckBoxAPI } from './checkbox-api'
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

export function CopyCheckBoxToCheckBoxAPI(checkbox: CheckBox, checkboxAPI: CheckBoxAPI) {

	checkboxAPI.CreatedAt = checkbox.CreatedAt
	checkboxAPI.DeletedAt = checkbox.DeletedAt
	checkboxAPI.ID = checkbox.ID

	// insertion point for basic fields copy operations
	checkboxAPI.Name = checkbox.Name
	checkboxAPI.Value = checkbox.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCheckBoxAPIToCheckBox update basic, pointers and slice of pointers fields of checkbox
// from respectively the basic fields and encoded fields of pointers and slices of pointers of checkboxAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCheckBoxAPIToCheckBox(checkboxAPI: CheckBoxAPI, checkbox: CheckBox, frontRepo: FrontRepo) {

	checkbox.CreatedAt = checkboxAPI.CreatedAt
	checkbox.DeletedAt = checkboxAPI.DeletedAt
	checkbox.ID = checkboxAPI.ID

	// insertion point for basic fields copy operations
	checkbox.Name = checkboxAPI.Name
	checkbox.Value = checkboxAPI.Value

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
