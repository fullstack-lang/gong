// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MeshMaterialBasicAPI {

	static GONGSTRUCT_NAME = "MeshMaterialBasic"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""

	// insertion point for other decls

	MeshMaterialBasicPointersEncoding: MeshMaterialBasicPointersEncoding = new MeshMaterialBasicPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MeshMaterialBasicPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
