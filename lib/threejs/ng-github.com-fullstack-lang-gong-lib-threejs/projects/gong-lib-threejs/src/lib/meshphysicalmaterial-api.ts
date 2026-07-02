// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MeshPhysicalMaterialAPI {

	static GONGSTRUCT_NAME = "MeshPhysicalMaterial"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""
	Wireframe: boolean = false
	Opacity: number = 0
	Transparent: boolean = false
	Visible: boolean = false

	// insertion point for other decls

	MeshPhysicalMaterialPointersEncoding: MeshPhysicalMaterialPointersEncoding = new MeshPhysicalMaterialPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MeshPhysicalMaterialPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
