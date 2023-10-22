// insertion point for imports
import { FormFieldDB } from './formfield-db'
import { CheckBoxDB } from './checkbox-db'
import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormSortAssocButtonDB } from './formsortassocbutton-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormDivDB {

	static GONGSTRUCT_NAME = "FormDiv"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	FormFields: Array<FormFieldDB> = []
	CheckBoxs: Array<CheckBoxDB> = []
	FormEditAssocButton?: FormEditAssocButtonDB

	FormSortAssocButton?: FormSortAssocButtonDB


	FormDivPointersEncoding: FormDivPointersEncoding = new FormDivPointersEncoding
}

export class FormDivPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormFields: number[] = []
	CheckBoxs: number[] = []
	FormEditAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormEditAssocButton.ID = 0

	FormSortAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormSortAssocButton.ID = 0

}
