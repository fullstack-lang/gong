// generated code - do not edit

import { FormFieldIntDB } from './formfieldint-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldInt {

	static GONGSTRUCT_NAME = "FormFieldInt"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0
	HasMinValidator: boolean = false
	MinValue: number = 0
	HasMaxValidator: boolean = false
	MaxValue: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFormFieldIntToFormFieldIntDB(formfieldint: FormFieldInt, formfieldintDB: FormFieldIntDB) {

	formfieldintDB.CreatedAt = formfieldint.CreatedAt
	formfieldintDB.DeletedAt = formfieldint.DeletedAt
	formfieldintDB.ID = formfieldint.ID
	
	// insertion point for basic fields copy operations
	formfieldintDB.Name = formfieldint.Name
	formfieldintDB.Value = formfieldint.Value
	formfieldintDB.HasMinValidator = formfieldint.HasMinValidator
	formfieldintDB.MinValue = formfieldint.MinValue
	formfieldintDB.HasMaxValidator = formfieldint.HasMaxValidator
	formfieldintDB.MaxValue = formfieldint.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyFormFieldIntDBToFormFieldInt(formfieldintDB: FormFieldIntDB, formfieldint: FormFieldInt, frontRepo: FrontRepo) {

	formfieldint.CreatedAt = formfieldintDB.CreatedAt
	formfieldint.DeletedAt = formfieldintDB.DeletedAt
	formfieldint.ID = formfieldintDB.ID
	
	// insertion point for basic fields copy operations
	formfieldint.Name = formfieldintDB.Name
	formfieldint.Value = formfieldintDB.Value
	formfieldint.HasMinValidator = formfieldintDB.HasMinValidator
	formfieldint.MinValue = formfieldintDB.MinValue
	formfieldint.HasMaxValidator = formfieldintDB.HasMaxValidator
	formfieldint.MaxValue = formfieldintDB.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
