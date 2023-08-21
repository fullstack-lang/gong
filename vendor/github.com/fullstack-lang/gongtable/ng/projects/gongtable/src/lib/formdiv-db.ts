// insertion point for imports
import { FormFieldDB } from './formfield-db'
import { CheckBoxDB } from './checkbox-db'
import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormSortAssocButtonDB } from './formsortassocbutton-db'
import { FormGroupDB } from './formgroup-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormDivDB {

	static GONGSTRUCT_NAME = "FormDiv"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	FormFields?: Array<FormFieldDB>
	CheckBoxs?: Array<CheckBoxDB>
	FormEditAssocButton?: FormEditAssocButtonDB
	FormEditAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormEditAssocButton.ID = 0

	FormSortAssocButton?: FormSortAssocButtonDB
	FormSortAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormSortAssocButton.ID = 0

	FormGroup_FormDivsDBID: NullInt64 = new NullInt64
	FormGroup_FormDivsDBID_Index: NullInt64  = new NullInt64 // store the index of the formdiv instance in FormGroup.FormDivs
	FormGroup_FormDivs_reverse?: FormGroupDB 

}
