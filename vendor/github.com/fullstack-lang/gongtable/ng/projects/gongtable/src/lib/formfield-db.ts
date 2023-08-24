// insertion point for imports
import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldDateDB } from './formfielddate-db'
import { FormFieldTimeDB } from './formfieldtime-db'
import { FormFieldDateTimeDB } from './formfielddatetime-db'
import { FormFieldSelectDB } from './formfieldselect-db'
import { FormDivDB } from './formdiv-db'

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

	// insertion point for other declarations
	FormFieldString?: FormFieldStringDB
	FormFieldStringID: NullInt64 = new NullInt64 // if pointer is null, FormFieldString.ID = 0

	FormFieldFloat64?: FormFieldFloat64DB
	FormFieldFloat64ID: NullInt64 = new NullInt64 // if pointer is null, FormFieldFloat64.ID = 0

	FormFieldInt?: FormFieldIntDB
	FormFieldIntID: NullInt64 = new NullInt64 // if pointer is null, FormFieldInt.ID = 0

	FormFieldDate?: FormFieldDateDB
	FormFieldDateID: NullInt64 = new NullInt64 // if pointer is null, FormFieldDate.ID = 0

	FormFieldTime?: FormFieldTimeDB
	FormFieldTimeID: NullInt64 = new NullInt64 // if pointer is null, FormFieldTime.ID = 0

	FormFieldDateTime?: FormFieldDateTimeDB
	FormFieldDateTimeID: NullInt64 = new NullInt64 // if pointer is null, FormFieldDateTime.ID = 0

	FormFieldSelect?: FormFieldSelectDB
	FormFieldSelectID: NullInt64 = new NullInt64 // if pointer is null, FormFieldSelect.ID = 0

	FormDiv_FormFieldsDBID: NullInt64 = new NullInt64
	FormDiv_FormFieldsDBID_Index: NullInt64  = new NullInt64 // store the index of the formfield instance in FormDiv.FormFields
	FormDiv_FormFields_reverse?: FormDivDB 

}
