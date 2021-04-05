// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GongdocCommandDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Command?: string
	DiagramName?: string
	Date?: string
	GongdocNodeType?: string
	StructName?: string
	FieldName?: string
	FieldTypeName?: string
	PositionX?: number
	PositionY?: number

	// insertion point for other declarations
}
