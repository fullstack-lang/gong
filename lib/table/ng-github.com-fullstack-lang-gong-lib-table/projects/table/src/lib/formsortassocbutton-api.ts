// insertion point for imports
import { FormEditAssocButtonAPI } from './formeditassocbutton-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormSortAssocButtonAPI {

	static GONGSTRUCT_NAME = "FormSortAssocButton"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""
	HasToolTip: boolean = false
	ToolTipText: string = ""
	MatTooltipShowDelay: string = ""

	// insertion point for other decls

	FormSortAssocButtonPointersEncoding: FormSortAssocButtonPointersEncoding = new FormSortAssocButtonPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class FormSortAssocButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormEditAssocButtonID: NullInt64 = new NullInt64 // if pointer is null, FormEditAssocButton.ID = 0

}
