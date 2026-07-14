// generated code - do not edit

import { BufferGeometryAPI } from './buffergeometry-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Vector3 } from './vector3'
import { Triangle } from './triangle'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BufferGeometry {

	static GONGSTRUCT_NAME = "BufferGeometry"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Vertices: Array<Vector3> = []
	Faces: Array<Triangle> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyBufferGeometryToBufferGeometryAPI(buffergeometry: BufferGeometry, buffergeometryAPI: BufferGeometryAPI) {

	buffergeometryAPI.CreatedAt = buffergeometry.CreatedAt
	buffergeometryAPI.DeletedAt = buffergeometry.DeletedAt
	buffergeometryAPI.ID = buffergeometry.ID

	// insertion point for basic fields copy operations
	buffergeometryAPI.Name = buffergeometry.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	buffergeometryAPI.BufferGeometryPointersEncoding.Vertices = []
	for (let _vector3 of buffergeometry.Vertices) {
		buffergeometryAPI.BufferGeometryPointersEncoding.Vertices.push(_vector3.ID)
	}

	buffergeometryAPI.BufferGeometryPointersEncoding.Faces = []
	for (let _triangle of buffergeometry.Faces) {
		buffergeometryAPI.BufferGeometryPointersEncoding.Faces.push(_triangle.ID)
	}

}

// CopyBufferGeometryAPIToBufferGeometry update basic, pointers and slice of pointers fields of buffergeometry
// from respectively the basic fields and encoded fields of pointers and slices of pointers of buffergeometryAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBufferGeometryAPIToBufferGeometry(buffergeometryAPI: BufferGeometryAPI, buffergeometry: BufferGeometry, frontRepo: FrontRepo) {

	buffergeometry.CreatedAt = buffergeometryAPI.CreatedAt
	buffergeometry.DeletedAt = buffergeometryAPI.DeletedAt
	buffergeometry.ID = buffergeometryAPI.ID

	// insertion point for basic fields copy operations
	buffergeometry.Name = buffergeometryAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(buffergeometryAPI.BufferGeometryPointersEncoding.Vertices)) {
		console.error('Rects is not an array:', buffergeometryAPI.BufferGeometryPointersEncoding.Vertices);
		return;
	}

	buffergeometry.Vertices = new Array<Vector3>()
	for (let _id of buffergeometryAPI.BufferGeometryPointersEncoding.Vertices) {
		let _vector3 = frontRepo.map_ID_Vector3.get(_id)
		if (_vector3 != undefined) {
			buffergeometry.Vertices.push(_vector3!)
		}
	}
	if (!Array.isArray(buffergeometryAPI.BufferGeometryPointersEncoding.Faces)) {
		console.error('Rects is not an array:', buffergeometryAPI.BufferGeometryPointersEncoding.Faces);
		return;
	}

	buffergeometry.Faces = new Array<Triangle>()
	for (let _id of buffergeometryAPI.BufferGeometryPointersEncoding.Faces) {
		let _triangle = frontRepo.map_ID_Triangle.get(_id)
		if (_triangle != undefined) {
			buffergeometry.Faces.push(_triangle!)
		}
	}
}
