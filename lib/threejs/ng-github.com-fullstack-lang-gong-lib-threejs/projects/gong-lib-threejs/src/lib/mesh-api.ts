// insertion point for imports
import { MeshMaterialBasicAPI } from './meshmaterialbasic-api'
import { MeshPhysicalMaterialAPI } from './meshphysicalmaterial-api'
import { CylinderGeometryAPI } from './cylindergeometry-api'
import { BoxGeometryAPI } from './boxgeometry-api'
import { SphereGeometryAPI } from './spheregeometry-api'
import { TorusGeometryAPI } from './torusgeometry-api'
import { PlaneGeometryAPI } from './planegeometry-api'
import { TubeGeometryAPI } from './tubegeometry-api'
import { ExtrudeGeometryAPI } from './extrudegeometry-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MeshAPI {

	static GONGSTRUCT_NAME = "Mesh"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0

	// insertion point for other decls

	MeshPointersEncoding: MeshPointersEncoding = new MeshPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MeshPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	MeshMaterialBasicID: NullInt64 = new NullInt64 // if pointer is null, MeshMaterialBasic.ID = 0

	MeshPhysicalMaterialID: NullInt64 = new NullInt64 // if pointer is null, MeshPhysicalMaterial.ID = 0

	CylinderGeometryID: NullInt64 = new NullInt64 // if pointer is null, CylinderGeometry.ID = 0

	BoxGeometryID: NullInt64 = new NullInt64 // if pointer is null, BoxGeometry.ID = 0

	SphereGeometryID: NullInt64 = new NullInt64 // if pointer is null, SphereGeometry.ID = 0

	TorusGeometryID: NullInt64 = new NullInt64 // if pointer is null, TorusGeometry.ID = 0

	PlaneGeometryID: NullInt64 = new NullInt64 // if pointer is null, PlaneGeometry.ID = 0

	TubeGeometryID: NullInt64 = new NullInt64 // if pointer is null, TubeGeometry.ID = 0

	ExtrudeGeometryID: NullInt64 = new NullInt64 // if pointer is null, ExtrudeGeometry.ID = 0

}
