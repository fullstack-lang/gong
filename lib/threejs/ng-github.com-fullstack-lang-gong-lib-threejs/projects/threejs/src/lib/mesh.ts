// generated code - do not edit

import { MeshAPI } from './mesh-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { MeshMaterialBasic } from './meshmaterialbasic'
import { MeshPhysicalMaterial } from './meshphysicalmaterial'
import { CylinderGeometry } from './cylindergeometry'
import { BoxGeometry } from './boxgeometry'
import { SphereGeometry } from './spheregeometry'
import { TorusGeometry } from './torusgeometry'
import { PlaneGeometry } from './planegeometry'
import { TubeGeometry } from './tubegeometry'
import { ExtrudeGeometry } from './extrudegeometry'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Mesh {

	static GONGSTRUCT_NAME = "Mesh"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0

	// insertion point for pointers and slices of pointers declarations
	MeshMaterialBasic?: MeshMaterialBasic

	MeshPhysicalMaterial?: MeshPhysicalMaterial

	CylinderGeometry?: CylinderGeometry

	BoxGeometry?: BoxGeometry

	SphereGeometry?: SphereGeometry

	TorusGeometry?: TorusGeometry

	PlaneGeometry?: PlaneGeometry

	TubeGeometry?: TubeGeometry

	ExtrudeGeometry?: ExtrudeGeometry


	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMeshToMeshAPI(mesh: Mesh, meshAPI: MeshAPI) {

	meshAPI.CreatedAt = mesh.CreatedAt
	meshAPI.DeletedAt = mesh.DeletedAt
	meshAPI.ID = mesh.ID

	// insertion point for basic fields copy operations
	meshAPI.Name = mesh.Name
	meshAPI.X = mesh.X
	meshAPI.Y = mesh.Y
	meshAPI.Z = mesh.Z

	// insertion point for pointer fields encoding
	meshAPI.MeshPointersEncoding.MeshMaterialBasicID.Valid = true
	if (mesh.MeshMaterialBasic != undefined) {
		meshAPI.MeshPointersEncoding.MeshMaterialBasicID.Int64 = mesh.MeshMaterialBasic.ID  
	} else {
		meshAPI.MeshPointersEncoding.MeshMaterialBasicID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.MeshPhysicalMaterialID.Valid = true
	if (mesh.MeshPhysicalMaterial != undefined) {
		meshAPI.MeshPointersEncoding.MeshPhysicalMaterialID.Int64 = mesh.MeshPhysicalMaterial.ID  
	} else {
		meshAPI.MeshPointersEncoding.MeshPhysicalMaterialID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.CylinderGeometryID.Valid = true
	if (mesh.CylinderGeometry != undefined) {
		meshAPI.MeshPointersEncoding.CylinderGeometryID.Int64 = mesh.CylinderGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.CylinderGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.BoxGeometryID.Valid = true
	if (mesh.BoxGeometry != undefined) {
		meshAPI.MeshPointersEncoding.BoxGeometryID.Int64 = mesh.BoxGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.BoxGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.SphereGeometryID.Valid = true
	if (mesh.SphereGeometry != undefined) {
		meshAPI.MeshPointersEncoding.SphereGeometryID.Int64 = mesh.SphereGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.SphereGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.TorusGeometryID.Valid = true
	if (mesh.TorusGeometry != undefined) {
		meshAPI.MeshPointersEncoding.TorusGeometryID.Int64 = mesh.TorusGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.TorusGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.PlaneGeometryID.Valid = true
	if (mesh.PlaneGeometry != undefined) {
		meshAPI.MeshPointersEncoding.PlaneGeometryID.Int64 = mesh.PlaneGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.PlaneGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.TubeGeometryID.Valid = true
	if (mesh.TubeGeometry != undefined) {
		meshAPI.MeshPointersEncoding.TubeGeometryID.Int64 = mesh.TubeGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.TubeGeometryID.Int64 = 0 		
	}

	meshAPI.MeshPointersEncoding.ExtrudeGeometryID.Valid = true
	if (mesh.ExtrudeGeometry != undefined) {
		meshAPI.MeshPointersEncoding.ExtrudeGeometryID.Int64 = mesh.ExtrudeGeometry.ID  
	} else {
		meshAPI.MeshPointersEncoding.ExtrudeGeometryID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyMeshAPIToMesh update basic, pointers and slice of pointers fields of mesh
// from respectively the basic fields and encoded fields of pointers and slices of pointers of meshAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMeshAPIToMesh(meshAPI: MeshAPI, mesh: Mesh, frontRepo: FrontRepo) {

	mesh.CreatedAt = meshAPI.CreatedAt
	mesh.DeletedAt = meshAPI.DeletedAt
	mesh.ID = meshAPI.ID

	// insertion point for basic fields copy operations
	mesh.Name = meshAPI.Name
	mesh.X = meshAPI.X
	mesh.Y = meshAPI.Y
	mesh.Z = meshAPI.Z

	// insertion point for pointer fields encoding
	mesh.MeshMaterialBasic = frontRepo.map_ID_MeshMaterialBasic.get(meshAPI.MeshPointersEncoding.MeshMaterialBasicID.Int64)
	mesh.MeshPhysicalMaterial = frontRepo.map_ID_MeshPhysicalMaterial.get(meshAPI.MeshPointersEncoding.MeshPhysicalMaterialID.Int64)
	mesh.CylinderGeometry = frontRepo.map_ID_CylinderGeometry.get(meshAPI.MeshPointersEncoding.CylinderGeometryID.Int64)
	mesh.BoxGeometry = frontRepo.map_ID_BoxGeometry.get(meshAPI.MeshPointersEncoding.BoxGeometryID.Int64)
	mesh.SphereGeometry = frontRepo.map_ID_SphereGeometry.get(meshAPI.MeshPointersEncoding.SphereGeometryID.Int64)
	mesh.TorusGeometry = frontRepo.map_ID_TorusGeometry.get(meshAPI.MeshPointersEncoding.TorusGeometryID.Int64)
	mesh.PlaneGeometry = frontRepo.map_ID_PlaneGeometry.get(meshAPI.MeshPointersEncoding.PlaneGeometryID.Int64)
	mesh.TubeGeometry = frontRepo.map_ID_TubeGeometry.get(meshAPI.MeshPointersEncoding.TubeGeometryID.Int64)
	mesh.ExtrudeGeometry = frontRepo.map_ID_ExtrudeGeometry.get(meshAPI.MeshPointersEncoding.ExtrudeGeometryID.Int64)

	// insertion point for slice of pointers fields encoding
}
