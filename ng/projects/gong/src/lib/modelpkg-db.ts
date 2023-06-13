// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ModelPkgDB {

	static GONGSTRUCT_NAME = "ModelPkg"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	PkgPath: string = ""

	// insertion point for other declarations
}
