// generated code - do not edit

import { FormFieldFloat64API } from './formfieldfloat64-api'
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

export function CopyFormFieldFloat64ToFormFieldFloat64API(formfieldfloat64: FormFieldFloat64, formfieldfloat64API: FormFieldFloat64API) {

	formfieldfloat64API.CreatedAt = formfieldfloat64.CreatedAt
	formfieldfloat64API.DeletedAt = formfieldfloat64.DeletedAt
	formfieldfloat64API.ID = formfieldfloat64.ID

	// insertion point for basic fields copy operations
	formfieldfloat64API.Name = formfieldfloat64.Name
	formfieldfloat64API.Value = formfieldfloat64.Value
	formfieldfloat64API.HasMinValidator = formfieldfloat64.HasMinValidator
	formfieldfloat64API.MinValue = formfieldfloat64.MinValue
	formfieldfloat64API.HasMaxValidator = formfieldfloat64.HasMaxValidator
	formfieldfloat64API.MaxValue = formfieldfloat64.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldFloat64APIToFormFieldFloat64 update basic, pointers and slice of pointers fields of formfieldfloat64
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldfloat64API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldFloat64APIToFormFieldFloat64(formfieldfloat64API: FormFieldFloat64API, formfieldfloat64: FormFieldFloat64, frontRepo: FrontRepo) {

	formfieldfloat64.CreatedAt = formfieldfloat64API.CreatedAt
	formfieldfloat64.DeletedAt = formfieldfloat64API.DeletedAt
	formfieldfloat64.ID = formfieldfloat64API.ID

	// insertion point for basic fields copy operations
	formfieldfloat64.Name = formfieldfloat64API.Name
	formfieldfloat64.Value = formfieldfloat64API.Value
	formfieldfloat64.HasMinValidator = formfieldfloat64API.HasMinValidator
	formfieldfloat64.MinValue = formfieldfloat64API.MinValue
	formfieldfloat64.HasMaxValidator = formfieldfloat64API.HasMaxValidator
	formfieldfloat64.MaxValue = formfieldfloat64API.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
