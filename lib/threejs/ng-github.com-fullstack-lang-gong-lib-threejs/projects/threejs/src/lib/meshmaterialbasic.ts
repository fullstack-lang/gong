// generated code - do not edit

import { MeshMaterialBasicAPI } from './meshmaterialbasic-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MeshMaterialBasic {

	static GONGSTRUCT_NAME = "MeshMaterialBasic"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMeshMaterialBasicToMeshMaterialBasicAPI(meshmaterialbasic: MeshMaterialBasic, meshmaterialbasicAPI: MeshMaterialBasicAPI) {

	meshmaterialbasicAPI.CreatedAt = meshmaterialbasic.CreatedAt
	meshmaterialbasicAPI.DeletedAt = meshmaterialbasic.DeletedAt
	meshmaterialbasicAPI.ID = meshmaterialbasic.ID

	// insertion point for basic fields copy operations
	meshmaterialbasicAPI.Name = meshmaterialbasic.Name
	meshmaterialbasicAPI.Color = meshmaterialbasic.Color

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMeshMaterialBasicAPIToMeshMaterialBasic update basic, pointers and slice of pointers fields of meshmaterialbasic
// from respectively the basic fields and encoded fields of pointers and slices of pointers of meshmaterialbasicAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMeshMaterialBasicAPIToMeshMaterialBasic(meshmaterialbasicAPI: MeshMaterialBasicAPI, meshmaterialbasic: MeshMaterialBasic, frontRepo: FrontRepo) {

	meshmaterialbasic.CreatedAt = meshmaterialbasicAPI.CreatedAt
	meshmaterialbasic.DeletedAt = meshmaterialbasicAPI.DeletedAt
	meshmaterialbasic.ID = meshmaterialbasicAPI.ID

	// insertion point for basic fields copy operations
	meshmaterialbasic.Name = meshmaterialbasicAPI.Name
	meshmaterialbasic.Color = meshmaterialbasicAPI.Color

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
