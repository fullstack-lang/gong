// generated code - do not edit

import { FormFieldSelectAPI } from './formfieldselect-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Option } from './option'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldSelect {

	static GONGSTRUCT_NAME = "FormFieldSelect"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CanBeEmpty: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Value?: Option

	Options: Array<Option> = []
}

export function CopyFormFieldSelectToFormFieldSelectAPI(formfieldselect: FormFieldSelect, formfieldselectAPI: FormFieldSelectAPI) {

	formfieldselectAPI.CreatedAt = formfieldselect.CreatedAt
	formfieldselectAPI.DeletedAt = formfieldselect.DeletedAt
	formfieldselectAPI.ID = formfieldselect.ID

	// insertion point for basic fields copy operations
	formfieldselectAPI.Name = formfieldselect.Name
	formfieldselectAPI.CanBeEmpty = formfieldselect.CanBeEmpty

	// insertion point for pointer fields encoding
	formfieldselectAPI.FormFieldSelectPointersEncoding.ValueID.Valid = true
	if (formfieldselect.Value != undefined) {
		formfieldselectAPI.FormFieldSelectPointersEncoding.ValueID.Int64 = formfieldselect.Value.ID  
	} else {
		formfieldselectAPI.FormFieldSelectPointersEncoding.ValueID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	formfieldselectAPI.FormFieldSelectPointersEncoding.Options = []
	for (let _option of formfieldselect.Options) {
		formfieldselectAPI.FormFieldSelectPointersEncoding.Options.push(_option.ID)
	}

}

// CopyFormFieldSelectAPIToFormFieldSelect update basic, pointers and slice of pointers fields of formfieldselect
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldselectAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldSelectAPIToFormFieldSelect(formfieldselectAPI: FormFieldSelectAPI, formfieldselect: FormFieldSelect, frontRepo: FrontRepo) {

	formfieldselect.CreatedAt = formfieldselectAPI.CreatedAt
	formfieldselect.DeletedAt = formfieldselectAPI.DeletedAt
	formfieldselect.ID = formfieldselectAPI.ID

	// insertion point for basic fields copy operations
	formfieldselect.Name = formfieldselectAPI.Name
	formfieldselect.CanBeEmpty = formfieldselectAPI.CanBeEmpty

	// insertion point for pointer fields encoding
	formfieldselect.Value = frontRepo.map_ID_Option.get(formfieldselectAPI.FormFieldSelectPointersEncoding.ValueID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(formfieldselectAPI.FormFieldSelectPointersEncoding.Options)) {
		console.error('Rects is not an array:', formfieldselectAPI.FormFieldSelectPointersEncoding.Options);
		return;
	}

	formfieldselect.Options = new Array<Option>()
	for (let _id of formfieldselectAPI.FormFieldSelectPointersEncoding.Options) {
		let _option = frontRepo.map_ID_Option.get(_id)
		if (_option != undefined) {
			formfieldselect.Options.push(_option!)
		}
	}
}
