// generated code - do not edit

import { FormFieldSelectDB } from './formfieldselect-db'
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

export function CopyFormFieldSelectToFormFieldSelectDB(formfieldselect: FormFieldSelect, formfieldselectDB: FormFieldSelectDB) {

	formfieldselectDB.CreatedAt = formfieldselect.CreatedAt
	formfieldselectDB.DeletedAt = formfieldselect.DeletedAt
	formfieldselectDB.ID = formfieldselect.ID

	// insertion point for basic fields copy operations
	formfieldselectDB.Name = formfieldselect.Name
	formfieldselectDB.CanBeEmpty = formfieldselect.CanBeEmpty

	// insertion point for pointer fields encoding
	formfieldselectDB.FormFieldSelectPointersEncoding.ValueID.Valid = true
	if (formfieldselect.Value != undefined) {
		formfieldselectDB.FormFieldSelectPointersEncoding.ValueID.Int64 = formfieldselect.Value.ID  
	} else {
		formfieldselectDB.FormFieldSelectPointersEncoding.ValueID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	formfieldselectDB.FormFieldSelectPointersEncoding.Options = []
	for (let _option of formfieldselect.Options) {
		formfieldselectDB.FormFieldSelectPointersEncoding.Options.push(_option.ID)
	}

}

// CopyFormFieldSelectDBToFormFieldSelect update basic, pointers and slice of pointers fields of formfieldselect
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldselectDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldSelectDBToFormFieldSelect(formfieldselectDB: FormFieldSelectDB, formfieldselect: FormFieldSelect, frontRepo: FrontRepo) {

	formfieldselect.CreatedAt = formfieldselectDB.CreatedAt
	formfieldselect.DeletedAt = formfieldselectDB.DeletedAt
	formfieldselect.ID = formfieldselectDB.ID

	// insertion point for basic fields copy operations
	formfieldselect.Name = formfieldselectDB.Name
	formfieldselect.CanBeEmpty = formfieldselectDB.CanBeEmpty

	// insertion point for pointer fields encoding
	formfieldselect.Value = frontRepo.map_ID_Option.get(formfieldselectDB.FormFieldSelectPointersEncoding.ValueID.Int64)

	// insertion point for slice of pointers fields encoding
	formfieldselect.Options = new Array<Option>()
	for (let _id of formfieldselectDB.FormFieldSelectPointersEncoding.Options) {
		let _option = frontRepo.map_ID_Option.get(_id)
		if (_option != undefined) {
			formfieldselect.Options.push(_option!)
		}
	}
}
