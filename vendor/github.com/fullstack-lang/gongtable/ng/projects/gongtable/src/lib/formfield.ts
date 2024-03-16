// generated code - do not edit

import { FormFieldAPI } from './formfield-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { FormFieldString } from './formfieldstring'
import { FormFieldFloat64 } from './formfieldfloat64'
import { FormFieldInt } from './formfieldint'
import { FormFieldDate } from './formfielddate'
import { FormFieldTime } from './formfieldtime'
import { FormFieldDateTime } from './formfielddatetime'
import { FormFieldSelect } from './formfieldselect'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormField {

	static GONGSTRUCT_NAME = "FormField"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	InputTypeEnum: string = ""
	Label: string = ""
	Placeholder: string = ""
	HasBespokeWidth: boolean = false
	BespokeWidthPx: number = 0
	HasBespokeHeight: boolean = false
	BespokeHeightPx: number = 0

	// insertion point for pointers and slices of pointers declarations
	FormFieldString?: FormFieldString

	FormFieldFloat64?: FormFieldFloat64

	FormFieldInt?: FormFieldInt

	FormFieldDate?: FormFieldDate

	FormFieldTime?: FormFieldTime

	FormFieldDateTime?: FormFieldDateTime

	FormFieldSelect?: FormFieldSelect

}

export function CopyFormFieldToFormFieldAPI(formfield: FormField, formfieldAPI: FormFieldAPI) {

	formfieldAPI.CreatedAt = formfield.CreatedAt
	formfieldAPI.DeletedAt = formfield.DeletedAt
	formfieldAPI.ID = formfield.ID

	// insertion point for basic fields copy operations
	formfieldAPI.Name = formfield.Name
	formfieldAPI.InputTypeEnum = formfield.InputTypeEnum
	formfieldAPI.Label = formfield.Label
	formfieldAPI.Placeholder = formfield.Placeholder
	formfieldAPI.HasBespokeWidth = formfield.HasBespokeWidth
	formfieldAPI.BespokeWidthPx = formfield.BespokeWidthPx
	formfieldAPI.HasBespokeHeight = formfield.HasBespokeHeight
	formfieldAPI.BespokeHeightPx = formfield.BespokeHeightPx

	// insertion point for pointer fields encoding
	formfieldAPI.FormFieldPointersEncoding.FormFieldStringID.Valid = true
	if (formfield.FormFieldString != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldStringID.Int64 = formfield.FormFieldString.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldStringID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldFloat64ID.Valid = true
	if (formfield.FormFieldFloat64 != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldFloat64ID.Int64 = formfield.FormFieldFloat64.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldFloat64ID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldIntID.Valid = true
	if (formfield.FormFieldInt != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldIntID.Int64 = formfield.FormFieldInt.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldIntID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldDateID.Valid = true
	if (formfield.FormFieldDate != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldDateID.Int64 = formfield.FormFieldDate.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldDateID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldTimeID.Valid = true
	if (formfield.FormFieldTime != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldTimeID.Int64 = formfield.FormFieldTime.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldTimeID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldDateTimeID.Valid = true
	if (formfield.FormFieldDateTime != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldDateTimeID.Int64 = formfield.FormFieldDateTime.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldDateTimeID.Int64 = 0 		
	}

	formfieldAPI.FormFieldPointersEncoding.FormFieldSelectID.Valid = true
	if (formfield.FormFieldSelect != undefined) {
		formfieldAPI.FormFieldPointersEncoding.FormFieldSelectID.Int64 = formfield.FormFieldSelect.ID  
	} else {
		formfieldAPI.FormFieldPointersEncoding.FormFieldSelectID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyFormFieldAPIToFormField update basic, pointers and slice of pointers fields of formfield
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formfieldAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormFieldAPIToFormField(formfieldAPI: FormFieldAPI, formfield: FormField, frontRepo: FrontRepo) {

	formfield.CreatedAt = formfieldAPI.CreatedAt
	formfield.DeletedAt = formfieldAPI.DeletedAt
	formfield.ID = formfieldAPI.ID

	// insertion point for basic fields copy operations
	formfield.Name = formfieldAPI.Name
	formfield.InputTypeEnum = formfieldAPI.InputTypeEnum
	formfield.Label = formfieldAPI.Label
	formfield.Placeholder = formfieldAPI.Placeholder
	formfield.HasBespokeWidth = formfieldAPI.HasBespokeWidth
	formfield.BespokeWidthPx = formfieldAPI.BespokeWidthPx
	formfield.HasBespokeHeight = formfieldAPI.HasBespokeHeight
	formfield.BespokeHeightPx = formfieldAPI.BespokeHeightPx

	// insertion point for pointer fields encoding
	formfield.FormFieldString = frontRepo.map_ID_FormFieldString.get(formfieldAPI.FormFieldPointersEncoding.FormFieldStringID.Int64)
	formfield.FormFieldFloat64 = frontRepo.map_ID_FormFieldFloat64.get(formfieldAPI.FormFieldPointersEncoding.FormFieldFloat64ID.Int64)
	formfield.FormFieldInt = frontRepo.map_ID_FormFieldInt.get(formfieldAPI.FormFieldPointersEncoding.FormFieldIntID.Int64)
	formfield.FormFieldDate = frontRepo.map_ID_FormFieldDate.get(formfieldAPI.FormFieldPointersEncoding.FormFieldDateID.Int64)
	formfield.FormFieldTime = frontRepo.map_ID_FormFieldTime.get(formfieldAPI.FormFieldPointersEncoding.FormFieldTimeID.Int64)
	formfield.FormFieldDateTime = frontRepo.map_ID_FormFieldDateTime.get(formfieldAPI.FormFieldPointersEncoding.FormFieldDateTimeID.Int64)
	formfield.FormFieldSelect = frontRepo.map_ID_FormFieldSelect.get(formfieldAPI.FormFieldPointersEncoding.FormFieldSelectID.Int64)

	// insertion point for slice of pointers fields encoding
}
