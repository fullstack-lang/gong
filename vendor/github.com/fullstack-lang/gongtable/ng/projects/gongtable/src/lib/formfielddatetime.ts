// generated code - do not edit

import { FormFieldDateTimeDB } from './formfielddatetime-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDateTime {

	static GONGSTRUCT_NAME = "FormFieldDateTime"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: Date = new Date

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldDateTimeToFormFieldDateTimeDB(formfielddatetime: FormFieldDateTime, formfielddatetimeDB: FormFieldDateTimeDB) {

	formfielddatetimeDB.CreatedAt = formfielddatetime.CreatedAt
	formfielddatetimeDB.DeletedAt = formfielddatetime.DeletedAt
	formfielddatetimeDB.ID = formfielddatetime.ID
	
	// insertion point for basic fields copy operations
	formfielddatetimeDB.Name = formfielddatetime.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyFormFieldDateTimeDBToFormFieldDateTime(formfielddatetimeDB: FormFieldDateTimeDB, formfielddatetime: FormFieldDateTime, frontRepo: FrontRepo) {

	formfielddatetime.CreatedAt = formfielddatetimeDB.CreatedAt
	formfielddatetime.DeletedAt = formfielddatetimeDB.DeletedAt
	formfielddatetime.ID = formfielddatetimeDB.ID
	
	// insertion point for basic fields copy operations
	formfielddatetime.Name = formfielddatetimeDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
