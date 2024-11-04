// generated code - do not edit

import { FormGroupAPI } from './formgroup-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { FormDiv } from './formdiv'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormGroup {

	static GONGSTRUCT_NAME = "FormGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasSuppressButton: boolean = false
	HasSuppressButtonBeenPressed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	FormDivs: Array<FormDiv> = []
}

export function CopyFormGroupToFormGroupAPI(formgroup: FormGroup, formgroupAPI: FormGroupAPI) {

	formgroupAPI.CreatedAt = formgroup.CreatedAt
	formgroupAPI.DeletedAt = formgroup.DeletedAt
	formgroupAPI.ID = formgroup.ID

	// insertion point for basic fields copy operations
	formgroupAPI.Name = formgroup.Name
	formgroupAPI.Label = formgroup.Label
	formgroupAPI.HasSuppressButton = formgroup.HasSuppressButton
	formgroupAPI.HasSuppressButtonBeenPressed = formgroup.HasSuppressButtonBeenPressed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	formgroupAPI.FormGroupPointersEncoding.FormDivs = []
	for (let _formdiv of formgroup.FormDivs) {
		formgroupAPI.FormGroupPointersEncoding.FormDivs.push(_formdiv.ID)
	}

}

// CopyFormGroupAPIToFormGroup update basic, pointers and slice of pointers fields of formgroup
// from respectively the basic fields and encoded fields of pointers and slices of pointers of formgroupAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFormGroupAPIToFormGroup(formgroupAPI: FormGroupAPI, formgroup: FormGroup, frontRepo: FrontRepo) {

	formgroup.CreatedAt = formgroupAPI.CreatedAt
	formgroup.DeletedAt = formgroupAPI.DeletedAt
	formgroup.ID = formgroupAPI.ID

	// insertion point for basic fields copy operations
	formgroup.Name = formgroupAPI.Name
	formgroup.Label = formgroupAPI.Label
	formgroup.HasSuppressButton = formgroupAPI.HasSuppressButton
	formgroup.HasSuppressButtonBeenPressed = formgroupAPI.HasSuppressButtonBeenPressed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(formgroupAPI.FormGroupPointersEncoding.FormDivs)) {
		console.error('Rects is not an array:', formgroupAPI.FormGroupPointersEncoding.FormDivs);
		return;
	}

	formgroup.FormDivs = new Array<FormDiv>()
	for (let _id of formgroupAPI.FormGroupPointersEncoding.FormDivs) {
		let _formdiv = frontRepo.map_ID_FormDiv.get(_id)
		if (_formdiv != undefined) {
			formgroup.FormDivs.push(_formdiv!)
		}
	}
}
