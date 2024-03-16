// insertion point for imports
import { FormFieldAPI } from './formfield-api'
import { CheckBoxAPI } from './checkbox-api'
import { FormEditAssocButtonAPI } from './formeditassocbutton-api'
import { FormSortAssocButtonAPI } from './formsortassocbutton-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormDivAPI {

	static GONGSTRUCT_NAME = "FormDiv"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	FormDivPointersEncoding: FormDivPointersEncoding = new FormDivPointersEncoding
}

export class FormDivPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormFields: number[] = []
	CheckBoxs: number[] = []
	FormEditAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormEditAssocButton.ID = 0

	FormSortAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormSortAssocButton.ID = 0

}
