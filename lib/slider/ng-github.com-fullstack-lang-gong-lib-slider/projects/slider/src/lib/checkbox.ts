// generated code - do not edit

import { CheckboxAPI } from './checkbox-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Checkbox {

	static GONGSTRUCT_NAME = "Checkbox"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ValueBool: boolean = false
	LabelForTrue: string = ""
	LabelForFalse: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCheckboxToCheckboxAPI(checkbox: Checkbox, checkboxAPI: CheckboxAPI) {

	checkboxAPI.CreatedAt = checkbox.CreatedAt
	checkboxAPI.DeletedAt = checkbox.DeletedAt
	checkboxAPI.ID = checkbox.ID

	// insertion point for basic fields copy operations
	checkboxAPI.Name = checkbox.Name
	checkboxAPI.ValueBool = checkbox.ValueBool
	checkboxAPI.LabelForTrue = checkbox.LabelForTrue
	checkboxAPI.LabelForFalse = checkbox.LabelForFalse

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCheckboxAPIToCheckbox update basic, pointers and slice of pointers fields of checkbox
// from respectively the basic fields and encoded fields of pointers and slices of pointers of checkboxAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCheckboxAPIToCheckbox(checkboxAPI: CheckboxAPI, checkbox: Checkbox, frontRepo: FrontRepo) {

	checkbox.CreatedAt = checkboxAPI.CreatedAt
	checkbox.DeletedAt = checkboxAPI.DeletedAt
	checkbox.ID = checkboxAPI.ID

	// insertion point for basic fields copy operations
	checkbox.Name = checkboxAPI.Name
	checkbox.ValueBool = checkboxAPI.ValueBool
	checkbox.LabelForTrue = checkboxAPI.LabelForTrue
	checkbox.LabelForFalse = checkboxAPI.LabelForFalse

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
