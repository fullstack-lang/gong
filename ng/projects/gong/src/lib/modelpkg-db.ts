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
	PathToGoSubDirectory: string = ""
	OrmPkgGenPath: string = ""
	ControllersPkgGenPath: string = ""
	FullstackPkgGenPath: string = ""
	StackPkgGenPath: string = ""
	StaticPkgGenPath: string = ""
	ProbePkgGenPath: string = ""
	NgWorkspacePath: string = ""
	NgDataLibrarySourceCodeDirectory: string = ""
	NgSpecificLibrarySourceCodeDirectory: string = ""
	MaterialLibDatamodelTargetPath: string = ""

	// insertion point for pointers and slices of pointers declarations

	ModelPkgPointersEncoding: ModelPkgPointersEncoding = new ModelPkgPointersEncoding
}

export class ModelPkgPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
