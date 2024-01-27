// generated code - do not edit

import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldFloat64 {

	static GONGSTRUCT_NAME = "FormFieldFloat64"

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

export function CopyFormFieldFloat64ToFormFieldFloat64DB(formfieldfloat64: FormFieldFloat64, formfieldfloat64DB: FormFieldFloat64DB) {

	formfieldfloat64DB.CreatedAt = formfieldfloat64.CreatedAt
	formfieldfloat64DB.DeletedAt = formfieldfloat64.DeletedAt
	formfieldfloat64DB.ID = formfieldfloat64.ID

	// insertion point for basic fields copy operations
	formfieldfloat64DB.Name = formfieldfloat64.Name
	formfieldfloat64DB.Value = formfieldfloat64.Value
	formfieldfloat64DB.HasMinValidator = formfieldfloat64.HasMinValidator
	formfieldfloat64DB.MinValue = formfieldfloat64.MinValue
	formfieldfloat64DB.HasMaxValidator = formfieldfloat64.HasMaxValidator
	formfieldfloat64DB.MaxValue = formfieldfloat64.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldFloat64DBToFormFieldFloat64 update basic, pointers and slice of pointers fields of formfieldfloat64
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldfloat64DB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldFloat64DBToFormFieldFloat64(formfieldfloat64DB: FormFieldFloat64DB, formfieldfloat64: FormFieldFloat64, frontRepo: FrontRepo) {

	formfieldfloat64.CreatedAt = formfieldfloat64DB.CreatedAt
	formfieldfloat64.DeletedAt = formfieldfloat64DB.DeletedAt
	formfieldfloat64.ID = formfieldfloat64DB.ID

	// insertion point for basic fields copy operations
	formfieldfloat64.Name = formfieldfloat64DB.Name
	formfieldfloat64.Value = formfieldfloat64DB.Value
	formfieldfloat64.HasMinValidator = formfieldfloat64DB.HasMinValidator
	formfieldfloat64.MinValue = formfieldfloat64DB.MinValue
	formfieldfloat64.HasMaxValidator = formfieldfloat64DB.HasMaxValidator
	formfieldfloat64.MaxValue = formfieldfloat64DB.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
