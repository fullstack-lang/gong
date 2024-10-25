// generated code - do not edit

import { ModelPkgAPI } from './modelpkg-api'
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
	DbOrmPkgGenPath: string = ""
	DbLiteOrmPkgGenPath: string = ""
	DbPkgGenPath: string = ""
	ControllersPkgGenPath: string = ""
	FullstackPkgGenPath: string = ""
	StackPkgGenPath: string = ""
	StaticPkgGenPath: string = ""
	ProbePkgGenPath: string = ""
	NgWorkspacePath: string = ""
	NgWorkspaceName: string = ""
	NgDataLibrarySourceCodeDirectory: string = ""
	NgSpecificLibrarySourceCodeDirectory: string = ""
	MaterialLibDatamodelTargetPath: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyModelPkgToModelPkgAPI(modelpkg: ModelPkg, modelpkgAPI: ModelPkgAPI) {

	modelpkgAPI.CreatedAt = modelpkg.CreatedAt
	modelpkgAPI.DeletedAt = modelpkg.DeletedAt
	modelpkgAPI.ID = modelpkg.ID

	// insertion point for basic fields copy operations
	modelpkgAPI.Name = modelpkg.Name
	modelpkgAPI.PkgPath = modelpkg.PkgPath
	modelpkgAPI.PathToGoSubDirectory = modelpkg.PathToGoSubDirectory
	modelpkgAPI.OrmPkgGenPath = modelpkg.OrmPkgGenPath
	modelpkgAPI.DbOrmPkgGenPath = modelpkg.DbOrmPkgGenPath
	modelpkgAPI.DbLiteOrmPkgGenPath = modelpkg.DbLiteOrmPkgGenPath
	modelpkgAPI.DbPkgGenPath = modelpkg.DbPkgGenPath
	modelpkgAPI.ControllersPkgGenPath = modelpkg.ControllersPkgGenPath
	modelpkgAPI.FullstackPkgGenPath = modelpkg.FullstackPkgGenPath
	modelpkgAPI.StackPkgGenPath = modelpkg.StackPkgGenPath
	modelpkgAPI.StaticPkgGenPath = modelpkg.StaticPkgGenPath
	modelpkgAPI.ProbePkgGenPath = modelpkg.ProbePkgGenPath
	modelpkgAPI.NgWorkspacePath = modelpkg.NgWorkspacePath
	modelpkgAPI.NgWorkspaceName = modelpkg.NgWorkspaceName
	modelpkgAPI.NgDataLibrarySourceCodeDirectory = modelpkg.NgDataLibrarySourceCodeDirectory
	modelpkgAPI.NgSpecificLibrarySourceCodeDirectory = modelpkg.NgSpecificLibrarySourceCodeDirectory
	modelpkgAPI.MaterialLibDatamodelTargetPath = modelpkg.MaterialLibDatamodelTargetPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyModelPkgAPIToModelPkg update basic, pointers and slice of pointers fields of modelpkg
// from respectively the basic fields and encoded fields of pointers and slices of pointers of modelpkgAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyModelPkgAPIToModelPkg(modelpkgAPI: ModelPkgAPI, modelpkg: ModelPkg, frontRepo: FrontRepo) {

	modelpkg.CreatedAt = modelpkgAPI.CreatedAt
	modelpkg.DeletedAt = modelpkgAPI.DeletedAt
	modelpkg.ID = modelpkgAPI.ID

	// insertion point for basic fields copy operations
	modelpkg.Name = modelpkgAPI.Name
	modelpkg.PkgPath = modelpkgAPI.PkgPath
	modelpkg.PathToGoSubDirectory = modelpkgAPI.PathToGoSubDirectory
	modelpkg.OrmPkgGenPath = modelpkgAPI.OrmPkgGenPath
	modelpkg.DbOrmPkgGenPath = modelpkgAPI.DbOrmPkgGenPath
	modelpkg.DbLiteOrmPkgGenPath = modelpkgAPI.DbLiteOrmPkgGenPath
	modelpkg.DbPkgGenPath = modelpkgAPI.DbPkgGenPath
	modelpkg.ControllersPkgGenPath = modelpkgAPI.ControllersPkgGenPath
	modelpkg.FullstackPkgGenPath = modelpkgAPI.FullstackPkgGenPath
	modelpkg.StackPkgGenPath = modelpkgAPI.StackPkgGenPath
	modelpkg.StaticPkgGenPath = modelpkgAPI.StaticPkgGenPath
	modelpkg.ProbePkgGenPath = modelpkgAPI.ProbePkgGenPath
	modelpkg.NgWorkspacePath = modelpkgAPI.NgWorkspacePath
	modelpkg.NgWorkspaceName = modelpkgAPI.NgWorkspaceName
	modelpkg.NgDataLibrarySourceCodeDirectory = modelpkgAPI.NgDataLibrarySourceCodeDirectory
	modelpkg.NgSpecificLibrarySourceCodeDirectory = modelpkgAPI.NgSpecificLibrarySourceCodeDirectory
	modelpkg.MaterialLibDatamodelTargetPath = modelpkgAPI.MaterialLibDatamodelTargetPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
