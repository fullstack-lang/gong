// generated code - do not edit

//insertion point for imports
import { AmbiantLightAPI } from './ambiantlight-api'

import { BoxGeometryAPI } from './boxgeometry-api'

import { BufferGeometryAPI } from './buffergeometry-api'

import { CameraAPI } from './camera-api'

import { CanvasAPI } from './canvas-api'

import { CurveAPI } from './curve-api'

import { CylinderGeometryAPI } from './cylindergeometry-api'

import { DirectionalLightAPI } from './directionallight-api'

import { ExtrudeGeometryAPI } from './extrudegeometry-api'

import { MeshAPI } from './mesh-api'

import { MeshMaterialBasicAPI } from './meshmaterialbasic-api'

import { MeshPhysicalMaterialAPI } from './meshphysicalmaterial-api'

import { PlaneGeometryAPI } from './planegeometry-api'

import { ShapeAPI } from './shape-api'

import { SphereGeometryAPI } from './spheregeometry-api'

import { TorusGeometryAPI } from './torusgeometry-api'

import { TriangleAPI } from './triangle-api'

import { TubeGeometryAPI } from './tubegeometry-api'

import { Vector2API } from './vector2-api'

import { Vector3API } from './vector3-api'


export class BackRepoData {
	// insertion point for declarations
	AmbiantLightAPIs = new Array<AmbiantLightAPI>()

	BoxGeometryAPIs = new Array<BoxGeometryAPI>()

	BufferGeometryAPIs = new Array<BufferGeometryAPI>()

	CameraAPIs = new Array<CameraAPI>()

	CanvasAPIs = new Array<CanvasAPI>()

	CurveAPIs = new Array<CurveAPI>()

	CylinderGeometryAPIs = new Array<CylinderGeometryAPI>()

	DirectionalLightAPIs = new Array<DirectionalLightAPI>()

	ExtrudeGeometryAPIs = new Array<ExtrudeGeometryAPI>()

	MeshAPIs = new Array<MeshAPI>()

	MeshMaterialBasicAPIs = new Array<MeshMaterialBasicAPI>()

	MeshPhysicalMaterialAPIs = new Array<MeshPhysicalMaterialAPI>()

	PlaneGeometryAPIs = new Array<PlaneGeometryAPI>()

	ShapeAPIs = new Array<ShapeAPI>()

	SphereGeometryAPIs = new Array<SphereGeometryAPI>()

	TorusGeometryAPIs = new Array<TorusGeometryAPI>()

	TriangleAPIs = new Array<TriangleAPI>()

	TubeGeometryAPIs = new Array<TubeGeometryAPI>()

	Vector2APIs = new Array<Vector2API>()

	Vector3APIs = new Array<Vector3API>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AmbiantLightAPIs = data?.AmbiantLightAPIs || [];

		this.BoxGeometryAPIs = data?.BoxGeometryAPIs || [];

		this.BufferGeometryAPIs = data?.BufferGeometryAPIs || [];

		this.CameraAPIs = data?.CameraAPIs || [];

		this.CanvasAPIs = data?.CanvasAPIs || [];

		this.CurveAPIs = data?.CurveAPIs || [];

		this.CylinderGeometryAPIs = data?.CylinderGeometryAPIs || [];

		this.DirectionalLightAPIs = data?.DirectionalLightAPIs || [];

		this.ExtrudeGeometryAPIs = data?.ExtrudeGeometryAPIs || [];

		this.MeshAPIs = data?.MeshAPIs || [];

		this.MeshMaterialBasicAPIs = data?.MeshMaterialBasicAPIs || [];

		this.MeshPhysicalMaterialAPIs = data?.MeshPhysicalMaterialAPIs || [];

		this.PlaneGeometryAPIs = data?.PlaneGeometryAPIs || [];

		this.ShapeAPIs = data?.ShapeAPIs || [];

		this.SphereGeometryAPIs = data?.SphereGeometryAPIs || [];

		this.TorusGeometryAPIs = data?.TorusGeometryAPIs || [];

		this.TriangleAPIs = data?.TriangleAPIs || [];

		this.TubeGeometryAPIs = data?.TubeGeometryAPIs || [];

		this.Vector2APIs = data?.Vector2APIs || [];

		this.Vector3APIs = data?.Vector3APIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}