// insertion point for imports
import { FormDivDB } from './formdiv-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormGroupDB {

	static GONGSTRUCT_NAME = "FormGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Label: string = ""

	// insertion point for other declarations
	FormDivs?: Array<FormDivDB>
}
