// generated code - do not edit

import { FormFieldTimeDB } from './formfieldtime-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldTime {

	static GONGSTRUCT_NAME = "FormFieldTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date
	Step: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldTimeToFormFieldTimeDB(formfieldtime: FormFieldTime, formfieldtimeDB: FormFieldTimeDB) {

	formfieldtimeDB.CreatedAt = formfieldtime.CreatedAt
	formfieldtimeDB.DeletedAt = formfieldtime.DeletedAt
	formfieldtimeDB.ID = formfieldtime.ID
	
	// insertion point for basic fields copy operations
	formfieldtimeDB.Name = formfieldtime.Name
	formfieldtimeDB.Step = formfieldtime.Step

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyFormFieldTimeDBToFormFieldTime(formfieldtimeDB: FormFieldTimeDB, formfieldtime: FormFieldTime, frontRepo: FrontRepo) {

	formfieldtime.CreatedAt = formfieldtimeDB.CreatedAt
	formfieldtime.DeletedAt = formfieldtimeDB.DeletedAt
	formfieldtime.ID = formfieldtimeDB.ID
	
	// insertion point for basic fields copy operations
	formfieldtime.Name = formfieldtimeDB.Name
	formfieldtime.Step = formfieldtimeDB.Step

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
