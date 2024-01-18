// generated code - do not edit
// insertion point for imports
import { GongEnumValue } from './gongenumvalue'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnum {

	static GONGSTRUCT_NAME = "GongEnum"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: number = 0

	// insertion point for pointers and slices of pointers declarations
	Type_string?: string
	GongEnumValues: Array<GongEnumValue> = []
}
