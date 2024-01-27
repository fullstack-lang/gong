// generated code - do not edit

import { ModelPkgDB } from './modelpkg-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ModelPkg {

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
}

export function CopyModelPkgToModelPkgDB(modelpkg: ModelPkg, modelpkgDB: ModelPkgDB) {

	modelpkgDB.CreatedAt = modelpkg.CreatedAt
	modelpkgDB.DeletedAt = modelpkg.DeletedAt
	modelpkgDB.ID = modelpkg.ID

	// insertion point for basic fields copy operations
	modelpkgDB.Name = modelpkg.Name
	modelpkgDB.PkgPath = modelpkg.PkgPath
	modelpkgDB.PathToGoSubDirectory = modelpkg.PathToGoSubDirectory
	modelpkgDB.OrmPkgGenPath = modelpkg.OrmPkgGenPath
	modelpkgDB.ControllersPkgGenPath = modelpkg.ControllersPkgGenPath
	modelpkgDB.FullstackPkgGenPath = modelpkg.FullstackPkgGenPath
	modelpkgDB.StackPkgGenPath = modelpkg.StackPkgGenPath
	modelpkgDB.StaticPkgGenPath = modelpkg.StaticPkgGenPath
	modelpkgDB.ProbePkgGenPath = modelpkg.ProbePkgGenPath
	modelpkgDB.NgWorkspacePath = modelpkg.NgWorkspacePath
	modelpkgDB.NgDataLibrarySourceCodeDirectory = modelpkg.NgDataLibrarySourceCodeDirectory
	modelpkgDB.NgSpecificLibrarySourceCodeDirectory = modelpkg.NgSpecificLibrarySourceCodeDirectory
	modelpkgDB.MaterialLibDatamodelTargetPath = modelpkg.MaterialLibDatamodelTargetPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyModelPkgDBToModelPkg update basic, pointers and slice of pointers fields of modelpkg
// from respectively the basic fields and encoded fields of pointers and slices of pointers of modelpkgDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyModelPkgDBToModelPkg(modelpkgDB: ModelPkgDB, modelpkg: ModelPkg, frontRepo: FrontRepo) {

	modelpkg.CreatedAt = modelpkgDB.CreatedAt
	modelpkg.DeletedAt = modelpkgDB.DeletedAt
	modelpkg.ID = modelpkgDB.ID

	// insertion point for basic fields copy operations
	modelpkg.Name = modelpkgDB.Name
	modelpkg.PkgPath = modelpkgDB.PkgPath
	modelpkg.PathToGoSubDirectory = modelpkgDB.PathToGoSubDirectory
	modelpkg.OrmPkgGenPath = modelpkgDB.OrmPkgGenPath
	modelpkg.ControllersPkgGenPath = modelpkgDB.ControllersPkgGenPath
	modelpkg.FullstackPkgGenPath = modelpkgDB.FullstackPkgGenPath
	modelpkg.StackPkgGenPath = modelpkgDB.StackPkgGenPath
	modelpkg.StaticPkgGenPath = modelpkgDB.StaticPkgGenPath
	modelpkg.ProbePkgGenPath = modelpkgDB.ProbePkgGenPath
	modelpkg.NgWorkspacePath = modelpkgDB.NgWorkspacePath
	modelpkg.NgDataLibrarySourceCodeDirectory = modelpkgDB.NgDataLibrarySourceCodeDirectory
	modelpkg.NgSpecificLibrarySourceCodeDirectory = modelpkgDB.NgSpecificLibrarySourceCodeDirectory
	modelpkg.MaterialLibDatamodelTargetPath = modelpkgDB.MaterialLibDatamodelTargetPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
