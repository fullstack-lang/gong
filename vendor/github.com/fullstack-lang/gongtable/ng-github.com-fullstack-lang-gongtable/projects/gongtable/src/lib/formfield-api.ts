// insertion point for imports
import { FormFieldStringAPI } from './formfieldstring-api'
import { FormFieldFloat64API } from './formfieldfloat64-api'
import { FormFieldIntAPI } from './formfieldint-api'
import { FormFieldDateAPI } from './formfielddate-api'
import { FormFieldTimeAPI } from './formfieldtime-api'
import { FormFieldDateTimeAPI } from './formfielddatetime-api'
import { FormFieldSelectAPI } from './formfieldselect-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldAPI {

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

	// insertion point for other decls

	FormFieldPointersEncoding: FormFieldPointersEncoding = new FormFieldPointersEncoding
}

export class FormFieldPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	FormFieldStringID: NullInt64 = new NullInt64 // if pointer is null, FormFieldString.ID = 0

	FormFieldFloat64ID: NullInt64 = new NullInt64 // if pointer is null, FormFieldFloat64.ID = 0

	FormFieldIntID: NullInt64 = new NullInt64 // if pointer is null, FormFieldInt.ID = 0

	FormFieldDateID: NullInt64 = new NullInt64 // if pointer is null, FormFieldDate.ID = 0

	FormFieldTimeID: NullInt64 = new NullInt64 // if pointer is null, FormFieldTime.ID = 0

	FormFieldDateTimeID: NullInt64 = new NullInt64 // if pointer is null, FormFieldDateTime.ID = 0

	FormFieldSelectID: NullInt64 = new NullInt64 // if pointer is null, FormFieldSelect.ID = 0

}
