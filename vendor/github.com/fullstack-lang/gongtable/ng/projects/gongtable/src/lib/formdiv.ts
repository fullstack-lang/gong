// generated code - do not edit

import { FormDivDB } from './formdiv-db'
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

export function CopyFormDivToFormDivDB(formdiv: FormDiv, formdivDB: FormDivDB) {

	formdivDB.CreatedAt = formdiv.CreatedAt
	formdivDB.DeletedAt = formdiv.DeletedAt
	formdivDB.ID = formdiv.ID

	// insertion point for basic fields copy operations
	formdivDB.Name = formdiv.Name

	// insertion point for pointer fields encoding
	formdivDB.FormDivPointersEncoding.FormEditAssocButtonID.Valid = true
	if (formdiv.FormEditAssocButton != undefined) {
		formdivDB.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = formdiv.FormEditAssocButton.ID  
	} else {
		formdivDB.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = 0 		
	}

	formdivDB.FormDivPointersEncoding.FormSortAssocButtonID.Valid = true
	if (formdiv.FormSortAssocButton != undefined) {
		formdivDB.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = formdiv.FormSortAssocButton.ID  
	} else {
		formdivDB.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	formdivDB.FormDivPointersEncoding.FormFields = []
	for (let _formfield of formdiv.FormFields) {
		formdivDB.FormDivPointersEncoding.FormFields.push(_formfield.ID)
	}

	formdivDB.FormDivPointersEncoding.CheckBoxs = []
	for (let _checkbox of formdiv.CheckBoxs) {
		formdivDB.FormDivPointersEncoding.CheckBoxs.push(_checkbox.ID)
	}

}

// CopyFormDivDBToFormDiv update basic, pointers and slice of pointers fields of formdiv
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formdivDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormDivDBToFormDiv(formdivDB: FormDivDB, formdiv: FormDiv, frontRepo: FrontRepo) {

	formdiv.CreatedAt = formdivDB.CreatedAt
	formdiv.DeletedAt = formdivDB.DeletedAt
	formdiv.ID = formdivDB.ID

	// insertion point for basic fields copy operations
	formdiv.Name = formdivDB.Name

	// insertion point for pointer fields encoding
	formdiv.FormEditAssocButton = frontRepo.map_ID_FormEditAssocButton.get(formdivDB.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
	formdiv.FormSortAssocButton = frontRepo.map_ID_FormSortAssocButton.get(formdivDB.FormDivPointersEncoding.FormSortAssocButtonID.Int64)

	// insertion point for slice of pointers fields encoding
	formdiv.FormFields = new Array<FormField>()
	for (let _id of formdivDB.FormDivPointersEncoding.FormFields) {
		let _formfield = frontRepo.map_ID_FormField.get(_id)
		if (_formfield != undefined) {
			formdiv.FormFields.push(_formfield!)
		}
	}
	formdiv.CheckBoxs = new Array<CheckBox>()
	for (let _id of formdivDB.FormDivPointersEncoding.CheckBoxs) {
		let _checkbox = frontRepo.map_ID_CheckBox.get(_id)
		if (_checkbox != undefined) {
			formdiv.CheckBoxs.push(_checkbox!)
		}
	}
}
