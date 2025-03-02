// generated code - do not edit

import { FormDivAPI } from './formdiv-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { FormField } from './formfield'
import { CheckBox } from './checkbox'
import { FormEditAssocButton } from './formeditassocbutton'
import { FormSortAssocButton } from './formsortassocbutton'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormDiv {

	static GONGSTRUCT_NAME = "FormDiv"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	FormFields: Array<FormField> = []
	CheckBoxs: Array<CheckBox> = []
	FormEditAssocButton?: FormEditAssocButton

	FormSortAssocButton?: FormSortAssocButton

}

export function CopyFormDivToFormDivAPI(formdiv: FormDiv, formdivAPI: FormDivAPI) {

	formdivAPI.CreatedAt = formdiv.CreatedAt
	formdivAPI.DeletedAt = formdiv.DeletedAt
	formdivAPI.ID = formdiv.ID

	// insertion point for basic fields copy operations
	formdivAPI.Name = formdiv.Name

	// insertion point for pointer fields encoding
	formdivAPI.FormDivPointersEncoding.FormEditAssocButtonID.Valid = true
	if (formdiv.FormEditAssocButton != undefined) {
		formdivAPI.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = formdiv.FormEditAssocButton.ID  
	} else {
		formdivAPI.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = 0 		
	}

	formdivAPI.FormDivPointersEncoding.FormSortAssocButtonID.Valid = true
	if (formdiv.FormSortAssocButton != undefined) {
		formdivAPI.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = formdiv.FormSortAssocButton.ID  
	} else {
		formdivAPI.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	formdivAPI.FormDivPointersEncoding.FormFields = []
	for (let _formfield of formdiv.FormFields) {
		formdivAPI.FormDivPointersEncoding.FormFields.push(_formfield.ID)
	}

	formdivAPI.FormDivPointersEncoding.CheckBoxs = []
	for (let _checkbox of formdiv.CheckBoxs) {
		formdivAPI.FormDivPointersEncoding.CheckBoxs.push(_checkbox.ID)
	}

}

// CopyFormDivAPIToFormDiv update basic, pointers and slice of pointers fields of formdiv
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formdivAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormDivAPIToFormDiv(formdivAPI: FormDivAPI, formdiv: FormDiv, frontRepo: FrontRepo) {

	formdiv.CreatedAt = formdivAPI.CreatedAt
	formdiv.DeletedAt = formdivAPI.DeletedAt
	formdiv.ID = formdivAPI.ID

	// insertion point for basic fields copy operations
	formdiv.Name = formdivAPI.Name

	// insertion point for pointer fields encoding
	formdiv.FormEditAssocButton = frontRepo.map_ID_FormEditAssocButton.get(formdivAPI.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
	formdiv.FormSortAssocButton = frontRepo.map_ID_FormSortAssocButton.get(formdivAPI.FormDivPointersEncoding.FormSortAssocButtonID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(formdivAPI.FormDivPointersEncoding.FormFields)) {
		console.error('Rects is not an array:', formdivAPI.FormDivPointersEncoding.FormFields);
		return;
	}

	formdiv.FormFields = new Array<FormField>()
	for (let _id of formdivAPI.FormDivPointersEncoding.FormFields) {
		let _formfield = frontRepo.map_ID_FormField.get(_id)
		if (_formfield != undefined) {
			formdiv.FormFields.push(_formfield!)
		}
	}
	if (!Array.isArray(formdivAPI.FormDivPointersEncoding.CheckBoxs)) {
		console.error('Rects is not an array:', formdivAPI.FormDivPointersEncoding.CheckBoxs);
		return;
	}

	formdiv.CheckBoxs = new Array<CheckBox>()
	for (let _id of formdivAPI.FormDivPointersEncoding.CheckBoxs) {
		let _checkbox = frontRepo.map_ID_CheckBox.get(_id)
		if (_checkbox != undefined) {
			formdiv.CheckBoxs.push(_checkbox!)
		}
	}
}
