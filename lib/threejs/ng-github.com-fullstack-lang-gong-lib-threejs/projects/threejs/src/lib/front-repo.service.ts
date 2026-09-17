// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AmbiantLightAPI } from './ambiantlight-api'
import { AmbiantLight, CopyAmbiantLightAPIToAmbiantLight } from './ambiantlight'

import { BoxGeometryAPI } from './boxgeometry-api'
import { BoxGeometry, CopyBoxGeometryAPIToBoxGeometry } from './boxgeometry'

import { BufferGeometryAPI } from './buffergeometry-api'
import { BufferGeometry, CopyBufferGeometryAPIToBufferGeometry } from './buffergeometry'

import { CameraAPI } from './camera-api'
import { Camera, CopyCameraAPIToCamera } from './camera'

import { CanvasAPI } from './canvas-api'
import { Canvas, CopyCanvasAPIToCanvas } from './canvas'

import { CurveAPI } from './curve-api'
import { Curve, CopyCurveAPIToCurve } from './curve'

import { CylinderGeometryAPI } from './cylindergeometry-api'
import { CylinderGeometry, CopyCylinderGeometryAPIToCylinderGeometry } from './cylindergeometry'

import { DirectionalLightAPI } from './directionallight-api'
import { DirectionalLight, CopyDirectionalLightAPIToDirectionalLight } from './directionallight'

import { ExtrudeGeometryAPI } from './extrudegeometry-api'
import { ExtrudeGeometry, CopyExtrudeGeometryAPIToExtrudeGeometry } from './extrudegeometry'

import { MeshAPI } from './mesh-api'
import { Mesh, CopyMeshAPIToMesh } from './mesh'

import { MeshMaterialBasicAPI } from './meshmaterialbasic-api'
import { MeshMaterialBasic, CopyMeshMaterialBasicAPIToMeshMaterialBasic } from './meshmaterialbasic'

import { MeshPhysicalMaterialAPI } from './meshphysicalmaterial-api'
import { MeshPhysicalMaterial, CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial } from './meshphysicalmaterial'

import { PlaneGeometryAPI } from './planegeometry-api'
import { PlaneGeometry, CopyPlaneGeometryAPIToPlaneGeometry } from './planegeometry'

import { ShapeAPI } from './shape-api'
import { Shape, CopyShapeAPIToShape } from './shape'

import { SphereGeometryAPI } from './spheregeometry-api'
import { SphereGeometry, CopySphereGeometryAPIToSphereGeometry } from './spheregeometry'

import { TorusGeometryAPI } from './torusgeometry-api'
import { TorusGeometry, CopyTorusGeometryAPIToTorusGeometry } from './torusgeometry'

import { TriangleAPI } from './triangle-api'
import { Triangle, CopyTriangleAPIToTriangle } from './triangle'

import { TubeGeometryAPI } from './tubegeometry-api'
import { TubeGeometry, CopyTubeGeometryAPIToTubeGeometry } from './tubegeometry'

import { Vector2API } from './vector2-api'
import { Vector2, CopyVector2APIToVector2 } from './vector2'

import { Vector3API } from './vector3-api'
import { Vector3, CopyVector3APIToVector3 } from './vector3'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/threejs/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_AmbiantLights = new Array<AmbiantLight>() // array of front instances
	map_ID_AmbiantLight = new Map<number, AmbiantLight>() // map of front instances

	array_BoxGeometrys = new Array<BoxGeometry>() // array of front instances
	map_ID_BoxGeometry = new Map<number, BoxGeometry>() // map of front instances

	array_BufferGeometrys = new Array<BufferGeometry>() // array of front instances
	map_ID_BufferGeometry = new Map<number, BufferGeometry>() // map of front instances

	array_Cameras = new Array<Camera>() // array of front instances
	map_ID_Camera = new Map<number, Camera>() // map of front instances

	array_Canvass = new Array<Canvas>() // array of front instances
	map_ID_Canvas = new Map<number, Canvas>() // map of front instances

	array_Curves = new Array<Curve>() // array of front instances
	map_ID_Curve = new Map<number, Curve>() // map of front instances

	array_CylinderGeometrys = new Array<CylinderGeometry>() // array of front instances
	map_ID_CylinderGeometry = new Map<number, CylinderGeometry>() // map of front instances

	array_DirectionalLights = new Array<DirectionalLight>() // array of front instances
	map_ID_DirectionalLight = new Map<number, DirectionalLight>() // map of front instances

	array_ExtrudeGeometrys = new Array<ExtrudeGeometry>() // array of front instances
	map_ID_ExtrudeGeometry = new Map<number, ExtrudeGeometry>() // map of front instances

	array_Meshs = new Array<Mesh>() // array of front instances
	map_ID_Mesh = new Map<number, Mesh>() // map of front instances

	array_MeshMaterialBasics = new Array<MeshMaterialBasic>() // array of front instances
	map_ID_MeshMaterialBasic = new Map<number, MeshMaterialBasic>() // map of front instances

	array_MeshPhysicalMaterials = new Array<MeshPhysicalMaterial>() // array of front instances
	map_ID_MeshPhysicalMaterial = new Map<number, MeshPhysicalMaterial>() // map of front instances

	array_PlaneGeometrys = new Array<PlaneGeometry>() // array of front instances
	map_ID_PlaneGeometry = new Map<number, PlaneGeometry>() // map of front instances

	array_Shapes = new Array<Shape>() // array of front instances
	map_ID_Shape = new Map<number, Shape>() // map of front instances

	array_SphereGeometrys = new Array<SphereGeometry>() // array of front instances
	map_ID_SphereGeometry = new Map<number, SphereGeometry>() // map of front instances

	array_TorusGeometrys = new Array<TorusGeometry>() // array of front instances
	map_ID_TorusGeometry = new Map<number, TorusGeometry>() // map of front instances

	array_Triangles = new Array<Triangle>() // array of front instances
	map_ID_Triangle = new Map<number, Triangle>() // map of front instances

	array_TubeGeometrys = new Array<TubeGeometry>() // array of front instances
	map_ID_TubeGeometry = new Map<number, TubeGeometry>() // map of front instances

	array_Vector2s = new Array<Vector2>() // array of front instances
	map_ID_Vector2 = new Map<number, Vector2>() // map of front instances

	array_Vector3s = new Array<Vector3>() // array of front instances
	map_ID_Vector3 = new Map<number, Vector3>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'AmbiantLight':
				return this.array_AmbiantLights as unknown as Array<Type>
			case 'BoxGeometry':
				return this.array_BoxGeometrys as unknown as Array<Type>
			case 'BufferGeometry':
				return this.array_BufferGeometrys as unknown as Array<Type>
			case 'Camera':
				return this.array_Cameras as unknown as Array<Type>
			case 'Canvas':
				return this.array_Canvass as unknown as Array<Type>
			case 'Curve':
				return this.array_Curves as unknown as Array<Type>
			case 'CylinderGeometry':
				return this.array_CylinderGeometrys as unknown as Array<Type>
			case 'DirectionalLight':
				return this.array_DirectionalLights as unknown as Array<Type>
			case 'ExtrudeGeometry':
				return this.array_ExtrudeGeometrys as unknown as Array<Type>
			case 'Mesh':
				return this.array_Meshs as unknown as Array<Type>
			case 'MeshMaterialBasic':
				return this.array_MeshMaterialBasics as unknown as Array<Type>
			case 'MeshPhysicalMaterial':
				return this.array_MeshPhysicalMaterials as unknown as Array<Type>
			case 'PlaneGeometry':
				return this.array_PlaneGeometrys as unknown as Array<Type>
			case 'Shape':
				return this.array_Shapes as unknown as Array<Type>
			case 'SphereGeometry':
				return this.array_SphereGeometrys as unknown as Array<Type>
			case 'TorusGeometry':
				return this.array_TorusGeometrys as unknown as Array<Type>
			case 'Triangle':
				return this.array_Triangles as unknown as Array<Type>
			case 'TubeGeometry':
				return this.array_TubeGeometrys as unknown as Array<Type>
			case 'Vector2':
				return this.array_Vector2s as unknown as Array<Type>
			case 'Vector3':
				return this.array_Vector3s as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'AmbiantLight':
				return this.map_ID_AmbiantLight as unknown as Map<number, Type>
			case 'BoxGeometry':
				return this.map_ID_BoxGeometry as unknown as Map<number, Type>
			case 'BufferGeometry':
				return this.map_ID_BufferGeometry as unknown as Map<number, Type>
			case 'Camera':
				return this.map_ID_Camera as unknown as Map<number, Type>
			case 'Canvas':
				return this.map_ID_Canvas as unknown as Map<number, Type>
			case 'Curve':
				return this.map_ID_Curve as unknown as Map<number, Type>
			case 'CylinderGeometry':
				return this.map_ID_CylinderGeometry as unknown as Map<number, Type>
			case 'DirectionalLight':
				return this.map_ID_DirectionalLight as unknown as Map<number, Type>
			case 'ExtrudeGeometry':
				return this.map_ID_ExtrudeGeometry as unknown as Map<number, Type>
			case 'Mesh':
				return this.map_ID_Mesh as unknown as Map<number, Type>
			case 'MeshMaterialBasic':
				return this.map_ID_MeshMaterialBasic as unknown as Map<number, Type>
			case 'MeshPhysicalMaterial':
				return this.map_ID_MeshPhysicalMaterial as unknown as Map<number, Type>
			case 'PlaneGeometry':
				return this.map_ID_PlaneGeometry as unknown as Map<number, Type>
			case 'Shape':
				return this.map_ID_Shape as unknown as Map<number, Type>
			case 'SphereGeometry':
				return this.map_ID_SphereGeometry as unknown as Map<number, Type>
			case 'TorusGeometry':
				return this.map_ID_TorusGeometry as unknown as Map<number, Type>
			case 'Triangle':
				return this.map_ID_Triangle as unknown as Map<number, Type>
			case 'TubeGeometry':
				return this.map_ID_TubeGeometry as unknown as Map<number, Type>
			case 'Vector2':
				return this.map_ID_Vector2 as unknown as Map<number, Type>
			case 'Vector3':
				return this.map_ID_Vector3 as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized")
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	Name: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	Name: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient,
		private ngZone: NgZone,
	) { }

	//
	// pull returns the FrontRepo Observable
	//
	pull(Name: string = ""): Observable<FrontRepo> {
		this.Name = Name
		return of(this.frontRepo)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		// console.log("github.com/fullstack-lang/gong/lib/threejs/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/threejs/go; connectToWebSocket: returning existing connection")
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/threejs/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/threejs/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/threejs/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_AmbiantLights = []
				frontRepo.map_ID_AmbiantLight.clear()

				backRepoData.AmbiantLightAPIs.forEach(
					ambiantlightAPI => {
						let ambiantlight = new AmbiantLight
						frontRepo.array_AmbiantLights.push(ambiantlight)
						frontRepo.map_ID_AmbiantLight.set(ambiantlightAPI.ID, ambiantlight)
					}
				)

				// init the arrays
				frontRepo.array_BoxGeometrys = []
				frontRepo.map_ID_BoxGeometry.clear()

				backRepoData.BoxGeometryAPIs.forEach(
					boxgeometryAPI => {
						let boxgeometry = new BoxGeometry
						frontRepo.array_BoxGeometrys.push(boxgeometry)
						frontRepo.map_ID_BoxGeometry.set(boxgeometryAPI.ID, boxgeometry)
					}
				)

				// init the arrays
				frontRepo.array_BufferGeometrys = []
				frontRepo.map_ID_BufferGeometry.clear()

				backRepoData.BufferGeometryAPIs.forEach(
					buffergeometryAPI => {
						let buffergeometry = new BufferGeometry
						frontRepo.array_BufferGeometrys.push(buffergeometry)
						frontRepo.map_ID_BufferGeometry.set(buffergeometryAPI.ID, buffergeometry)
					}
				)

				// init the arrays
				frontRepo.array_Cameras = []
				frontRepo.map_ID_Camera.clear()

				backRepoData.CameraAPIs.forEach(
					cameraAPI => {
						let camera = new Camera
						frontRepo.array_Cameras.push(camera)
						frontRepo.map_ID_Camera.set(cameraAPI.ID, camera)
					}
				)

				// init the arrays
				frontRepo.array_Canvass = []
				frontRepo.map_ID_Canvas.clear()

				backRepoData.CanvasAPIs.forEach(
					canvasAPI => {
						let canvas = new Canvas
						frontRepo.array_Canvass.push(canvas)
						frontRepo.map_ID_Canvas.set(canvasAPI.ID, canvas)
					}
				)

				// init the arrays
				frontRepo.array_Curves = []
				frontRepo.map_ID_Curve.clear()

				backRepoData.CurveAPIs.forEach(
					curveAPI => {
						let curve = new Curve
						frontRepo.array_Curves.push(curve)
						frontRepo.map_ID_Curve.set(curveAPI.ID, curve)
					}
				)

				// init the arrays
				frontRepo.array_CylinderGeometrys = []
				frontRepo.map_ID_CylinderGeometry.clear()

				backRepoData.CylinderGeometryAPIs.forEach(
					cylindergeometryAPI => {
						let cylindergeometry = new CylinderGeometry
						frontRepo.array_CylinderGeometrys.push(cylindergeometry)
						frontRepo.map_ID_CylinderGeometry.set(cylindergeometryAPI.ID, cylindergeometry)
					}
				)

				// init the arrays
				frontRepo.array_DirectionalLights = []
				frontRepo.map_ID_DirectionalLight.clear()

				backRepoData.DirectionalLightAPIs.forEach(
					directionallightAPI => {
						let directionallight = new DirectionalLight
						frontRepo.array_DirectionalLights.push(directionallight)
						frontRepo.map_ID_DirectionalLight.set(directionallightAPI.ID, directionallight)
					}
				)

				// init the arrays
				frontRepo.array_ExtrudeGeometrys = []
				frontRepo.map_ID_ExtrudeGeometry.clear()

				backRepoData.ExtrudeGeometryAPIs.forEach(
					extrudegeometryAPI => {
						let extrudegeometry = new ExtrudeGeometry
						frontRepo.array_ExtrudeGeometrys.push(extrudegeometry)
						frontRepo.map_ID_ExtrudeGeometry.set(extrudegeometryAPI.ID, extrudegeometry)
					}
				)

				// init the arrays
				frontRepo.array_Meshs = []
				frontRepo.map_ID_Mesh.clear()

				backRepoData.MeshAPIs.forEach(
					meshAPI => {
						let mesh = new Mesh
						frontRepo.array_Meshs.push(mesh)
						frontRepo.map_ID_Mesh.set(meshAPI.ID, mesh)
					}
				)

				// init the arrays
				frontRepo.array_MeshMaterialBasics = []
				frontRepo.map_ID_MeshMaterialBasic.clear()

				backRepoData.MeshMaterialBasicAPIs.forEach(
					meshmaterialbasicAPI => {
						let meshmaterialbasic = new MeshMaterialBasic
						frontRepo.array_MeshMaterialBasics.push(meshmaterialbasic)
						frontRepo.map_ID_MeshMaterialBasic.set(meshmaterialbasicAPI.ID, meshmaterialbasic)
					}
				)

				// init the arrays
				frontRepo.array_MeshPhysicalMaterials = []
				frontRepo.map_ID_MeshPhysicalMaterial.clear()

				backRepoData.MeshPhysicalMaterialAPIs.forEach(
					meshphysicalmaterialAPI => {
						let meshphysicalmaterial = new MeshPhysicalMaterial
						frontRepo.array_MeshPhysicalMaterials.push(meshphysicalmaterial)
						frontRepo.map_ID_MeshPhysicalMaterial.set(meshphysicalmaterialAPI.ID, meshphysicalmaterial)
					}
				)

				// init the arrays
				frontRepo.array_PlaneGeometrys = []
				frontRepo.map_ID_PlaneGeometry.clear()

				backRepoData.PlaneGeometryAPIs.forEach(
					planegeometryAPI => {
						let planegeometry = new PlaneGeometry
						frontRepo.array_PlaneGeometrys.push(planegeometry)
						frontRepo.map_ID_PlaneGeometry.set(planegeometryAPI.ID, planegeometry)
					}
				)

				// init the arrays
				frontRepo.array_Shapes = []
				frontRepo.map_ID_Shape.clear()

				backRepoData.ShapeAPIs.forEach(
					shapeAPI => {
						let shape = new Shape
						frontRepo.array_Shapes.push(shape)
						frontRepo.map_ID_Shape.set(shapeAPI.ID, shape)
					}
				)

				// init the arrays
				frontRepo.array_SphereGeometrys = []
				frontRepo.map_ID_SphereGeometry.clear()

				backRepoData.SphereGeometryAPIs.forEach(
					spheregeometryAPI => {
						let spheregeometry = new SphereGeometry
						frontRepo.array_SphereGeometrys.push(spheregeometry)
						frontRepo.map_ID_SphereGeometry.set(spheregeometryAPI.ID, spheregeometry)
					}
				)

				// init the arrays
				frontRepo.array_TorusGeometrys = []
				frontRepo.map_ID_TorusGeometry.clear()

				backRepoData.TorusGeometryAPIs.forEach(
					torusgeometryAPI => {
						let torusgeometry = new TorusGeometry
						frontRepo.array_TorusGeometrys.push(torusgeometry)
						frontRepo.map_ID_TorusGeometry.set(torusgeometryAPI.ID, torusgeometry)
					}
				)

				// init the arrays
				frontRepo.array_Triangles = []
				frontRepo.map_ID_Triangle.clear()

				backRepoData.TriangleAPIs.forEach(
					triangleAPI => {
						let triangle = new Triangle
						frontRepo.array_Triangles.push(triangle)
						frontRepo.map_ID_Triangle.set(triangleAPI.ID, triangle)
					}
				)

				// init the arrays
				frontRepo.array_TubeGeometrys = []
				frontRepo.map_ID_TubeGeometry.clear()

				backRepoData.TubeGeometryAPIs.forEach(
					tubegeometryAPI => {
						let tubegeometry = new TubeGeometry
						frontRepo.array_TubeGeometrys.push(tubegeometry)
						frontRepo.map_ID_TubeGeometry.set(tubegeometryAPI.ID, tubegeometry)
					}
				)

				// init the arrays
				frontRepo.array_Vector2s = []
				frontRepo.map_ID_Vector2.clear()

				backRepoData.Vector2APIs.forEach(
					vector2API => {
						let vector2 = new Vector2
						frontRepo.array_Vector2s.push(vector2)
						frontRepo.map_ID_Vector2.set(vector2API.ID, vector2)
					}
				)

				// init the arrays
				frontRepo.array_Vector3s = []
				frontRepo.map_ID_Vector3.clear()

				backRepoData.Vector3APIs.forEach(
					vector3API => {
						let vector3 = new Vector3
						frontRepo.array_Vector3s.push(vector3)
						frontRepo.map_ID_Vector3.set(vector3API.ID, vector3)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AmbiantLightAPIs.forEach(
					ambiantlightAPI => {
						let ambiantlight = frontRepo.map_ID_AmbiantLight.get(ambiantlightAPI.ID)
						CopyAmbiantLightAPIToAmbiantLight(ambiantlightAPI, ambiantlight!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BoxGeometryAPIs.forEach(
					boxgeometryAPI => {
						let boxgeometry = frontRepo.map_ID_BoxGeometry.get(boxgeometryAPI.ID)
						CopyBoxGeometryAPIToBoxGeometry(boxgeometryAPI, boxgeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BufferGeometryAPIs.forEach(
					buffergeometryAPI => {
						let buffergeometry = frontRepo.map_ID_BufferGeometry.get(buffergeometryAPI.ID)
						CopyBufferGeometryAPIToBufferGeometry(buffergeometryAPI, buffergeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CameraAPIs.forEach(
					cameraAPI => {
						let camera = frontRepo.map_ID_Camera.get(cameraAPI.ID)
						CopyCameraAPIToCamera(cameraAPI, camera!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CanvasAPIs.forEach(
					canvasAPI => {
						let canvas = frontRepo.map_ID_Canvas.get(canvasAPI.ID)
						CopyCanvasAPIToCanvas(canvasAPI, canvas!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CurveAPIs.forEach(
					curveAPI => {
						let curve = frontRepo.map_ID_Curve.get(curveAPI.ID)
						CopyCurveAPIToCurve(curveAPI, curve!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CylinderGeometryAPIs.forEach(
					cylindergeometryAPI => {
						let cylindergeometry = frontRepo.map_ID_CylinderGeometry.get(cylindergeometryAPI.ID)
						CopyCylinderGeometryAPIToCylinderGeometry(cylindergeometryAPI, cylindergeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DirectionalLightAPIs.forEach(
					directionallightAPI => {
						let directionallight = frontRepo.map_ID_DirectionalLight.get(directionallightAPI.ID)
						CopyDirectionalLightAPIToDirectionalLight(directionallightAPI, directionallight!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ExtrudeGeometryAPIs.forEach(
					extrudegeometryAPI => {
						let extrudegeometry = frontRepo.map_ID_ExtrudeGeometry.get(extrudegeometryAPI.ID)
						CopyExtrudeGeometryAPIToExtrudeGeometry(extrudegeometryAPI, extrudegeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MeshAPIs.forEach(
					meshAPI => {
						let mesh = frontRepo.map_ID_Mesh.get(meshAPI.ID)
						CopyMeshAPIToMesh(meshAPI, mesh!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MeshMaterialBasicAPIs.forEach(
					meshmaterialbasicAPI => {
						let meshmaterialbasic = frontRepo.map_ID_MeshMaterialBasic.get(meshmaterialbasicAPI.ID)
						CopyMeshMaterialBasicAPIToMeshMaterialBasic(meshmaterialbasicAPI, meshmaterialbasic!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MeshPhysicalMaterialAPIs.forEach(
					meshphysicalmaterialAPI => {
						let meshphysicalmaterial = frontRepo.map_ID_MeshPhysicalMaterial.get(meshphysicalmaterialAPI.ID)
						CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial(meshphysicalmaterialAPI, meshphysicalmaterial!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PlaneGeometryAPIs.forEach(
					planegeometryAPI => {
						let planegeometry = frontRepo.map_ID_PlaneGeometry.get(planegeometryAPI.ID)
						CopyPlaneGeometryAPIToPlaneGeometry(planegeometryAPI, planegeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ShapeAPIs.forEach(
					shapeAPI => {
						let shape = frontRepo.map_ID_Shape.get(shapeAPI.ID)
						CopyShapeAPIToShape(shapeAPI, shape!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SphereGeometryAPIs.forEach(
					spheregeometryAPI => {
						let spheregeometry = frontRepo.map_ID_SphereGeometry.get(spheregeometryAPI.ID)
						CopySphereGeometryAPIToSphereGeometry(spheregeometryAPI, spheregeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TorusGeometryAPIs.forEach(
					torusgeometryAPI => {
						let torusgeometry = frontRepo.map_ID_TorusGeometry.get(torusgeometryAPI.ID)
						CopyTorusGeometryAPIToTorusGeometry(torusgeometryAPI, torusgeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TriangleAPIs.forEach(
					triangleAPI => {
						let triangle = frontRepo.map_ID_Triangle.get(triangleAPI.ID)
						CopyTriangleAPIToTriangle(triangleAPI, triangle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TubeGeometryAPIs.forEach(
					tubegeometryAPI => {
						let tubegeometry = frontRepo.map_ID_TubeGeometry.get(tubegeometryAPI.ID)
						CopyTubeGeometryAPIToTubeGeometry(tubegeometryAPI, tubegeometry!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.Vector2APIs.forEach(
					vector2API => {
						let vector2 = frontRepo.map_ID_Vector2.get(vector2API.ID)
						CopyVector2APIToVector2(vector2API, vector2!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.Vector3APIs.forEach(
					vector3API => {
						let vector3 = frontRepo.map_ID_Vector3.get(vector3API.ID)
						CopyVector3APIToVector3(vector3API, vector3!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				// console.log("github.com/fullstack-lang/gong/lib/threejs/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					// console.log("github.com/fullstack-lang/gong/lib/threejs/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/lib/threejs/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					// console.log("github.com/fullstack-lang/gong/lib/threejs/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					// console.log("github.com/fullstack-lang/gong/lib/threejs/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/threejs/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						// console.log("github.com/fullstack-lang/gong/lib/threejs/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/threejs/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/threejs/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/threejs/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.");
				}
			};

			attemptConnection(50);

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				if (socket) {
					socket.close()
				}
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getAmbiantLightUniqueID(id: number): number {
	return 31 * id
}
export function getBoxGeometryUniqueID(id: number): number {
	return 37 * id
}
export function getBufferGeometryUniqueID(id: number): number {
	return 41 * id
}
export function getCameraUniqueID(id: number): number {
	return 43 * id
}
export function getCanvasUniqueID(id: number): number {
	return 47 * id
}
export function getCurveUniqueID(id: number): number {
	return 53 * id
}
export function getCylinderGeometryUniqueID(id: number): number {
	return 59 * id
}
export function getDirectionalLightUniqueID(id: number): number {
	return 61 * id
}
export function getExtrudeGeometryUniqueID(id: number): number {
	return 67 * id
}
export function getMeshUniqueID(id: number): number {
	return 71 * id
}
export function getMeshMaterialBasicUniqueID(id: number): number {
	return 73 * id
}
export function getMeshPhysicalMaterialUniqueID(id: number): number {
	return 79 * id
}
export function getPlaneGeometryUniqueID(id: number): number {
	return 83 * id
}
export function getShapeUniqueID(id: number): number {
	return 89 * id
}
export function getSphereGeometryUniqueID(id: number): number {
	return 97 * id
}
export function getTorusGeometryUniqueID(id: number): number {
	return 101 * id
}
export function getTriangleUniqueID(id: number): number {
	return 103 * id
}
export function getTubeGeometryUniqueID(id: number): number {
	return 107 * id
}
export function getVector2UniqueID(id: number): number {
	return 109 * id
}
export function getVector3UniqueID(id: number): number {
	return 113 * id
}
