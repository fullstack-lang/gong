// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FormFieldFloat64DB {

	static GONGSTRUCT_NAME = "FormFieldFloat64"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Value: number = 0
	HasMinValidator: boolean = false
	MinValue: number = 0
	HasMaxValidator: boolean = false
	MaxValue: number = 0

	// insertion point for other declarations
}
