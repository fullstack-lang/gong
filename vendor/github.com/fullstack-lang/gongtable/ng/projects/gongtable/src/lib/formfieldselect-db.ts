// insertion point for imports
import { OptionDB } from './option-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldSelectDB {

	static GONGSTRUCT_NAME = "FormFieldSelect"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Value?: OptionDB
	ValueID: NullInt64 = new NullInt64 // if pointer is null, Value.ID = 0

	Options?: Array<OptionDB>
}
