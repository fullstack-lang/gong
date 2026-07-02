// generated code - do not edit

import { MeshPhysicalMaterialAPI } from './meshphysicalmaterial-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MeshPhysicalMaterial {

	static GONGSTRUCT_NAME = "MeshPhysicalMaterial"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Color: string = ""
	Wireframe: boolean = false
	Opacity: number = 0
	Transparent: boolean = false
	Visible: boolean = false

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMeshPhysicalMaterialToMeshPhysicalMaterialAPI(meshphysicalmaterial: MeshPhysicalMaterial, meshphysicalmaterialAPI: MeshPhysicalMaterialAPI) {

	meshphysicalmaterialAPI.CreatedAt = meshphysicalmaterial.CreatedAt
	meshphysicalmaterialAPI.DeletedAt = meshphysicalmaterial.DeletedAt
	meshphysicalmaterialAPI.ID = meshphysicalmaterial.ID

	// insertion point for basic fields copy operations
	meshphysicalmaterialAPI.Name = meshphysicalmaterial.Name
	meshphysicalmaterialAPI.Color = meshphysicalmaterial.Color
	meshphysicalmaterialAPI.Wireframe = meshphysicalmaterial.Wireframe
	meshphysicalmaterialAPI.Opacity = meshphysicalmaterial.Opacity
	meshphysicalmaterialAPI.Transparent = meshphysicalmaterial.Transparent
	meshphysicalmaterialAPI.Visible = meshphysicalmaterial.Visible

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial update basic, pointers and slice of pointers fields of meshphysicalmaterial
// from respectively the basic fields and encoded fields of pointers and slices of pointers of meshphysicalmaterialAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial(meshphysicalmaterialAPI: MeshPhysicalMaterialAPI, meshphysicalmaterial: MeshPhysicalMaterial, frontRepo: FrontRepo) {

	meshphysicalmaterial.CreatedAt = meshphysicalmaterialAPI.CreatedAt
	meshphysicalmaterial.DeletedAt = meshphysicalmaterialAPI.DeletedAt
	meshphysicalmaterial.ID = meshphysicalmaterialAPI.ID

	// insertion point for basic fields copy operations
	meshphysicalmaterial.Name = meshphysicalmaterialAPI.Name
	meshphysicalmaterial.Color = meshphysicalmaterialAPI.Color
	meshphysicalmaterial.Wireframe = meshphysicalmaterialAPI.Wireframe
	meshphysicalmaterial.Opacity = meshphysicalmaterialAPI.Opacity
	meshphysicalmaterial.Transparent = meshphysicalmaterialAPI.Transparent
	meshphysicalmaterial.Visible = meshphysicalmaterialAPI.Visible

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
