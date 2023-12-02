// insertion point for imports
import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldDateDB } from './formfielddate-db'
import { FormFieldTimeDB } from './formfieldtime-db'
import { FormFieldDateTimeDB } from './formfielddatetime-db'
import { FormFieldSelectDB } from './formfieldselect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldDB {

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
	FormFieldString?: FormFieldStringDB

	FormFieldFloat64?: FormFieldFloat64DB

	FormFieldInt?: FormFieldIntDB

	FormFieldDate?: FormFieldDateDB

	FormFieldTime?: FormFieldTimeDB

	FormFieldDateTime?: FormFieldDateTimeDB

	FormFieldSelect?: FormFieldSelectDB


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
