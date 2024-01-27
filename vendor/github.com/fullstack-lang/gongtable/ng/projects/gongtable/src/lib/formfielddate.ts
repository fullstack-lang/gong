// generated code - do not edit

import { FormFieldDateDB } from './formfielddate-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDate {

	static GONGSTRUCT_NAME = "FormFieldDate"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldDateToFormFieldDateDB(formfielddate: FormFieldDate, formfielddateDB: FormFieldDateDB) {

	formfielddateDB.CreatedAt = formfielddate.CreatedAt
	formfielddateDB.DeletedAt = formfielddate.DeletedAt
	formfielddateDB.ID = formfielddate.ID

	// insertion point for basic fields copy operations
	formfielddateDB.Name = formfielddate.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldDateDBToFormFieldDate update basic, pointers and slice of pointers fields of formfielddate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfielddateDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldDateDBToFormFieldDate(formfielddateDB: FormFieldDateDB, formfielddate: FormFieldDate, frontRepo: FrontRepo) {

	formfielddate.CreatedAt = formfielddateDB.CreatedAt
	formfielddate.DeletedAt = formfielddateDB.DeletedAt
	formfielddate.ID = formfielddateDB.ID

	// insertion point for basic fields copy operations
	formfielddate.Name = formfielddateDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
