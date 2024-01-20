// generated code - do not edit

import { FormFieldDB } from './formfield-db'
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

export function CopyFormFieldToFormFieldDB(formfield: FormField, formfieldDB: FormFieldDB) {

	formfieldDB.CreatedAt = formfield.CreatedAt
	formfieldDB.DeletedAt = formfield.DeletedAt
	formfieldDB.ID = formfield.ID
	
	// insertion point for basic fields copy operations
	formfieldDB.Name = formfield.Name
	formfieldDB.InputTypeEnum = formfield.InputTypeEnum
	formfieldDB.Label = formfield.Label
	formfieldDB.Placeholder = formfield.Placeholder
	formfieldDB.HasBespokeWidth = formfield.HasBespokeWidth
	formfieldDB.BespokeWidthPx = formfield.BespokeWidthPx
	formfieldDB.HasBespokeHeight = formfield.HasBespokeHeight
	formfieldDB.BespokeHeightPx = formfield.BespokeHeightPx

	// insertion point for pointer fields encoding
    formfieldDB.FormFieldPointersEncoding.FormFieldStringID.Valid = true
	if (formfield.FormFieldString != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldStringID.Int64 = formfield.FormFieldString.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldStringID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldFloat64ID.Valid = true
	if (formfield.FormFieldFloat64 != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldFloat64ID.Int64 = formfield.FormFieldFloat64.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldFloat64ID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldIntID.Valid = true
	if (formfield.FormFieldInt != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldIntID.Int64 = formfield.FormFieldInt.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldIntID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldDateID.Valid = true
	if (formfield.FormFieldDate != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldDateID.Int64 = formfield.FormFieldDate.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldDateID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldTimeID.Valid = true
	if (formfield.FormFieldTime != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldTimeID.Int64 = formfield.FormFieldTime.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldTimeID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldDateTimeID.Valid = true
	if (formfield.FormFieldDateTime != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldDateTimeID.Int64 = formfield.FormFieldDateTime.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldDateTimeID.Int64 = 0 		
	}

    formfieldDB.FormFieldPointersEncoding.FormFieldSelectID.Valid = true
	if (formfield.FormFieldSelect != undefined) {
		formfieldDB.FormFieldPointersEncoding.FormFieldSelectID.Int64 = formfield.FormFieldSelect.ID  
    } else {
		formfieldDB.FormFieldPointersEncoding.FormFieldSelectID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyFormFieldDBToFormField(formfieldDB: FormFieldDB, formfield: FormField, frontRepo: FrontRepo) {

	formfield.CreatedAt = formfieldDB.CreatedAt
	formfield.DeletedAt = formfieldDB.DeletedAt
	formfield.ID = formfieldDB.ID
	
	// insertion point for basic fields copy operations
	formfield.Name = formfieldDB.Name
	formfield.InputTypeEnum = formfieldDB.InputTypeEnum
	formfield.Label = formfieldDB.Label
	formfield.Placeholder = formfieldDB.Placeholder
	formfield.HasBespokeWidth = formfieldDB.HasBespokeWidth
	formfield.BespokeWidthPx = formfieldDB.BespokeWidthPx
	formfield.HasBespokeHeight = formfieldDB.HasBespokeHeight
	formfield.BespokeHeightPx = formfieldDB.BespokeHeightPx

	// insertion point for pointer fields encoding
	formfield.FormFieldString = frontRepo.FormFieldStrings.get(formfieldDB.FormFieldPointersEncoding.FormFieldStringID.Int64)
	formfield.FormFieldFloat64 = frontRepo.FormFieldFloat64s.get(formfieldDB.FormFieldPointersEncoding.FormFieldFloat64ID.Int64)
	formfield.FormFieldInt = frontRepo.FormFieldInts.get(formfieldDB.FormFieldPointersEncoding.FormFieldIntID.Int64)
	formfield.FormFieldDate = frontRepo.FormFieldDates.get(formfieldDB.FormFieldPointersEncoding.FormFieldDateID.Int64)
	formfield.FormFieldTime = frontRepo.FormFieldTimes.get(formfieldDB.FormFieldPointersEncoding.FormFieldTimeID.Int64)
	formfield.FormFieldDateTime = frontRepo.FormFieldDateTimes.get(formfieldDB.FormFieldPointersEncoding.FormFieldDateTimeID.Int64)
	formfield.FormFieldSelect = frontRepo.FormFieldSelects.get(formfieldDB.FormFieldPointersEncoding.FormFieldSelectID.Int64)

	// insertion point for slice of pointers fields encoding
}
