// generated code - do not edit

import { FormFieldIntAPI } from './formfieldint-api'
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

export function CopyFormFieldIntToFormFieldIntAPI(formfieldint: FormFieldInt, formfieldintAPI: FormFieldIntAPI) {

	formfieldintAPI.CreatedAt = formfieldint.CreatedAt
	formfieldintAPI.DeletedAt = formfieldint.DeletedAt
	formfieldintAPI.ID = formfieldint.ID

	// insertion point for basic fields copy operations
	formfieldintAPI.Name = formfieldint.Name
	formfieldintAPI.Value = formfieldint.Value
	formfieldintAPI.HasMinValidator = formfieldint.HasMinValidator
	formfieldintAPI.MinValue = formfieldint.MinValue
	formfieldintAPI.HasMaxValidator = formfieldint.HasMaxValidator
	formfieldintAPI.MaxValue = formfieldint.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldIntAPIToFormFieldInt update basic, pointers and slice of pointers fields of formfieldint
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldintAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldIntAPIToFormFieldInt(formfieldintAPI: FormFieldIntAPI, formfieldint: FormFieldInt, frontRepo: FrontRepo) {

	formfieldint.CreatedAt = formfieldintAPI.CreatedAt
	formfieldint.DeletedAt = formfieldintAPI.DeletedAt
	formfieldint.ID = formfieldintAPI.ID

	// insertion point for basic fields copy operations
	formfieldint.Name = formfieldintAPI.Name
	formfieldint.Value = formfieldintAPI.Value
	formfieldint.HasMinValidator = formfieldintAPI.HasMinValidator
	formfieldint.MinValue = formfieldintAPI.MinValue
	formfieldint.HasMaxValidator = formfieldintAPI.HasMaxValidator
	formfieldint.MaxValue = formfieldintAPI.MaxValue

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
