// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongdocCommandDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Command: string = ""
	DiagramName: string = ""
	Date: string = ""
	GongdocNodeType: string = ""
	StructName: string = ""
	FieldName: string = ""
	FieldTypeName: string = ""
	PositionX: number = 0
	PositionY: number = 0
	NoteName: string = ""

	// insertion point for other declarations
}
