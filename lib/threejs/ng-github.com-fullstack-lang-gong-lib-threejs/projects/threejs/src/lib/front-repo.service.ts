import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AmbiantLightAPI } from './ambiantlight-api'
import { AmbiantLight, CopyAmbiantLightAPIToAmbiantLight } from './ambiantlight'
import { AmbiantLightService } from './ambiantlight.service'

import { BoxGeometryAPI } from './boxgeometry-api'
import { BoxGeometry, CopyBoxGeometryAPIToBoxGeometry } from './boxgeometry'
import { BoxGeometryService } from './boxgeometry.service'

import { CameraAPI } from './camera-api'
import { Camera, CopyCameraAPIToCamera } from './camera'
import { CameraService } from './camera.service'

import { CanvasAPI } from './canvas-api'
import { Canvas, CopyCanvasAPIToCanvas } from './canvas'
import { CanvasService } from './canvas.service'

import { CurveAPI } from './curve-api'
import { Curve, CopyCurveAPIToCurve } from './curve'
import { CurveService } from './curve.service'

import { CylinderGeometryAPI } from './cylindergeometry-api'
import { CylinderGeometry, CopyCylinderGeometryAPIToCylinderGeometry } from './cylindergeometry'
import { CylinderGeometryService } from './cylindergeometry.service'

import { DirectionalLightAPI } from './directionallight-api'
import { DirectionalLight, CopyDirectionalLightAPIToDirectionalLight } from './directionallight'
import { DirectionalLightService } from './directionallight.service'

import { ExtrudeGeometryAPI } from './extrudegeometry-api'
import { ExtrudeGeometry, CopyExtrudeGeometryAPIToExtrudeGeometry } from './extrudegeometry'
import { ExtrudeGeometryService } from './extrudegeometry.service'

import { MeshAPI } from './mesh-api'
import { Mesh, CopyMeshAPIToMesh } from './mesh'
import { MeshService } from './mesh.service'

import { MeshMaterialBasicAPI } from './meshmaterialbasic-api'
import { MeshMaterialBasic, CopyMeshMaterialBasicAPIToMeshMaterialBasic } from './meshmaterialbasic'
import { MeshMaterialBasicService } from './meshmaterialbasic.service'

import { MeshPhysicalMaterialAPI } from './meshphysicalmaterial-api'
import { MeshPhysicalMaterial, CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial } from './meshphysicalmaterial'
import { MeshPhysicalMaterialService } from './meshphysicalmaterial.service'

import { PlaneGeometryAPI } from './planegeometry-api'
import { PlaneGeometry, CopyPlaneGeometryAPIToPlaneGeometry } from './planegeometry'
import { PlaneGeometryService } from './planegeometry.service'

import { ShapeAPI } from './shape-api'
import { Shape, CopyShapeAPIToShape } from './shape'
import { ShapeService } from './shape.service'

import { SphereGeometryAPI } from './spheregeometry-api'
import { SphereGeometry, CopySphereGeometryAPIToSphereGeometry } from './spheregeometry'
import { SphereGeometryService } from './spheregeometry.service'

import { TorusGeometryAPI } from './torusgeometry-api'
import { TorusGeometry, CopyTorusGeometryAPIToTorusGeometry } from './torusgeometry'
import { TorusGeometryService } from './torusgeometry.service'

import { TubeGeometryAPI } from './tubegeometry-api'
import { TubeGeometry, CopyTubeGeometryAPIToTubeGeometry } from './tubegeometry'
import { TubeGeometryService } from './tubegeometry.service'

import { Vector2API } from './vector2-api'
import { Vector2, CopyVector2APIToVector2 } from './vector2'
import { Vector2Service } from './vector2.service'

import { Vector3API } from './vector3-api'
import { Vector3, CopyVector3APIToVector3 } from './vector3'
import { Vector3Service } from './vector3.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/threejs/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_AmbiantLights = new Array<AmbiantLight>() // array of front instances
	map_ID_AmbiantLight = new Map<number, AmbiantLight>() // map of front instances

	array_BoxGeometrys = new Array<BoxGeometry>() // array of front instances
	map_ID_BoxGeometry = new Map<number, BoxGeometry>() // map of front instances

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
		private http: HttpClient, // insertion point sub template 
		private ambiantlightService: AmbiantLightService,
		private boxgeometryService: BoxGeometryService,
		private cameraService: CameraService,
		private canvasService: CanvasService,
		private curveService: CurveService,
		private cylindergeometryService: CylinderGeometryService,
		private directionallightService: DirectionalLightService,
		private extrudegeometryService: ExtrudeGeometryService,
		private meshService: MeshService,
		private meshmaterialbasicService: MeshMaterialBasicService,
		private meshphysicalmaterialService: MeshPhysicalMaterialService,
		private planegeometryService: PlaneGeometryService,
		private shapeService: ShapeService,
		private spheregeometryService: SphereGeometryService,
		private torusgeometryService: TorusGeometryService,
		private tubegeometryService: TubeGeometryService,
		private vector2Service: Vector2Service,
		private vector3Service: Vector3Service,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		)
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		)
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AmbiantLightAPI[]>,
		Observable<BoxGeometryAPI[]>,
		Observable<CameraAPI[]>,
		Observable<CanvasAPI[]>,
		Observable<CurveAPI[]>,
		Observable<CylinderGeometryAPI[]>,
		Observable<DirectionalLightAPI[]>,
		Observable<ExtrudeGeometryAPI[]>,
		Observable<MeshAPI[]>,
		Observable<MeshMaterialBasicAPI[]>,
		Observable<MeshPhysicalMaterialAPI[]>,
		Observable<PlaneGeometryAPI[]>,
		Observable<ShapeAPI[]>,
		Observable<SphereGeometryAPI[]>,
		Observable<TorusGeometryAPI[]>,
		Observable<TubeGeometryAPI[]>,
		Observable<Vector2API[]>,
		Observable<Vector3API[]>,
	]

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(Name: string = ""): Observable<FrontRepo> {

		this.Name = Name

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.ambiantlightService.getAmbiantLights(this.Name, this.frontRepo),
			this.boxgeometryService.getBoxGeometrys(this.Name, this.frontRepo),
			this.cameraService.getCameras(this.Name, this.frontRepo),
			this.canvasService.getCanvass(this.Name, this.frontRepo),
			this.curveService.getCurves(this.Name, this.frontRepo),
			this.cylindergeometryService.getCylinderGeometrys(this.Name, this.frontRepo),
			this.directionallightService.getDirectionalLights(this.Name, this.frontRepo),
			this.extrudegeometryService.getExtrudeGeometrys(this.Name, this.frontRepo),
			this.meshService.getMeshs(this.Name, this.frontRepo),
			this.meshmaterialbasicService.getMeshMaterialBasics(this.Name, this.frontRepo),
			this.meshphysicalmaterialService.getMeshPhysicalMaterials(this.Name, this.frontRepo),
			this.planegeometryService.getPlaneGeometrys(this.Name, this.frontRepo),
			this.shapeService.getShapes(this.Name, this.frontRepo),
			this.spheregeometryService.getSphereGeometrys(this.Name, this.frontRepo),
			this.torusgeometryService.getTorusGeometrys(this.Name, this.frontRepo),
			this.tubegeometryService.getTubeGeometrys(this.Name, this.frontRepo),
			this.vector2Service.getVector2s(this.Name, this.frontRepo),
			this.vector3Service.getVector3s(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						ambiantlights_,
						boxgeometrys_,
						cameras_,
						canvass_,
						curves_,
						cylindergeometrys_,
						directionallights_,
						extrudegeometrys_,
						meshs_,
						meshmaterialbasics_,
						meshphysicalmaterials_,
						planegeometrys_,
						shapes_,
						spheregeometrys_,
						torusgeometrys_,
						tubegeometrys_,
						vector2s_,
						vector3s_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var ambiantlights: AmbiantLightAPI[]
						ambiantlights = ambiantlights_ as AmbiantLightAPI[]
						var boxgeometrys: BoxGeometryAPI[]
						boxgeometrys = boxgeometrys_ as BoxGeometryAPI[]
						var cameras: CameraAPI[]
						cameras = cameras_ as CameraAPI[]
						var canvass: CanvasAPI[]
						canvass = canvass_ as CanvasAPI[]
						var curves: CurveAPI[]
						curves = curves_ as CurveAPI[]
						var cylindergeometrys: CylinderGeometryAPI[]
						cylindergeometrys = cylindergeometrys_ as CylinderGeometryAPI[]
						var directionallights: DirectionalLightAPI[]
						directionallights = directionallights_ as DirectionalLightAPI[]
						var extrudegeometrys: ExtrudeGeometryAPI[]
						extrudegeometrys = extrudegeometrys_ as ExtrudeGeometryAPI[]
						var meshs: MeshAPI[]
						meshs = meshs_ as MeshAPI[]
						var meshmaterialbasics: MeshMaterialBasicAPI[]
						meshmaterialbasics = meshmaterialbasics_ as MeshMaterialBasicAPI[]
						var meshphysicalmaterials: MeshPhysicalMaterialAPI[]
						meshphysicalmaterials = meshphysicalmaterials_ as MeshPhysicalMaterialAPI[]
						var planegeometrys: PlaneGeometryAPI[]
						planegeometrys = planegeometrys_ as PlaneGeometryAPI[]
						var shapes: ShapeAPI[]
						shapes = shapes_ as ShapeAPI[]
						var spheregeometrys: SphereGeometryAPI[]
						spheregeometrys = spheregeometrys_ as SphereGeometryAPI[]
						var torusgeometrys: TorusGeometryAPI[]
						torusgeometrys = torusgeometrys_ as TorusGeometryAPI[]
						var tubegeometrys: TubeGeometryAPI[]
						tubegeometrys = tubegeometrys_ as TubeGeometryAPI[]
						var vector2s: Vector2API[]
						vector2s = vector2s_ as Vector2API[]
						var vector3s: Vector3API[]
						vector3s = vector3s_ as Vector3API[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_AmbiantLights = []
						this.frontRepo.map_ID_AmbiantLight.clear()

						ambiantlights.forEach(
							ambiantlightAPI => {
								let ambiantlight = new AmbiantLight
								this.frontRepo.array_AmbiantLights.push(ambiantlight)
								this.frontRepo.map_ID_AmbiantLight.set(ambiantlightAPI.ID, ambiantlight)
							}
						)

						// init the arrays
						this.frontRepo.array_BoxGeometrys = []
						this.frontRepo.map_ID_BoxGeometry.clear()

						boxgeometrys.forEach(
							boxgeometryAPI => {
								let boxgeometry = new BoxGeometry
								this.frontRepo.array_BoxGeometrys.push(boxgeometry)
								this.frontRepo.map_ID_BoxGeometry.set(boxgeometryAPI.ID, boxgeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_Cameras = []
						this.frontRepo.map_ID_Camera.clear()

						cameras.forEach(
							cameraAPI => {
								let camera = new Camera
								this.frontRepo.array_Cameras.push(camera)
								this.frontRepo.map_ID_Camera.set(cameraAPI.ID, camera)
							}
						)

						// init the arrays
						this.frontRepo.array_Canvass = []
						this.frontRepo.map_ID_Canvas.clear()

						canvass.forEach(
							canvasAPI => {
								let canvas = new Canvas
								this.frontRepo.array_Canvass.push(canvas)
								this.frontRepo.map_ID_Canvas.set(canvasAPI.ID, canvas)
							}
						)

						// init the arrays
						this.frontRepo.array_Curves = []
						this.frontRepo.map_ID_Curve.clear()

						curves.forEach(
							curveAPI => {
								let curve = new Curve
								this.frontRepo.array_Curves.push(curve)
								this.frontRepo.map_ID_Curve.set(curveAPI.ID, curve)
							}
						)

						// init the arrays
						this.frontRepo.array_CylinderGeometrys = []
						this.frontRepo.map_ID_CylinderGeometry.clear()

						cylindergeometrys.forEach(
							cylindergeometryAPI => {
								let cylindergeometry = new CylinderGeometry
								this.frontRepo.array_CylinderGeometrys.push(cylindergeometry)
								this.frontRepo.map_ID_CylinderGeometry.set(cylindergeometryAPI.ID, cylindergeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_DirectionalLights = []
						this.frontRepo.map_ID_DirectionalLight.clear()

						directionallights.forEach(
							directionallightAPI => {
								let directionallight = new DirectionalLight
								this.frontRepo.array_DirectionalLights.push(directionallight)
								this.frontRepo.map_ID_DirectionalLight.set(directionallightAPI.ID, directionallight)
							}
						)

						// init the arrays
						this.frontRepo.array_ExtrudeGeometrys = []
						this.frontRepo.map_ID_ExtrudeGeometry.clear()

						extrudegeometrys.forEach(
							extrudegeometryAPI => {
								let extrudegeometry = new ExtrudeGeometry
								this.frontRepo.array_ExtrudeGeometrys.push(extrudegeometry)
								this.frontRepo.map_ID_ExtrudeGeometry.set(extrudegeometryAPI.ID, extrudegeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_Meshs = []
						this.frontRepo.map_ID_Mesh.clear()

						meshs.forEach(
							meshAPI => {
								let mesh = new Mesh
								this.frontRepo.array_Meshs.push(mesh)
								this.frontRepo.map_ID_Mesh.set(meshAPI.ID, mesh)
							}
						)

						// init the arrays
						this.frontRepo.array_MeshMaterialBasics = []
						this.frontRepo.map_ID_MeshMaterialBasic.clear()

						meshmaterialbasics.forEach(
							meshmaterialbasicAPI => {
								let meshmaterialbasic = new MeshMaterialBasic
								this.frontRepo.array_MeshMaterialBasics.push(meshmaterialbasic)
								this.frontRepo.map_ID_MeshMaterialBasic.set(meshmaterialbasicAPI.ID, meshmaterialbasic)
							}
						)

						// init the arrays
						this.frontRepo.array_MeshPhysicalMaterials = []
						this.frontRepo.map_ID_MeshPhysicalMaterial.clear()

						meshphysicalmaterials.forEach(
							meshphysicalmaterialAPI => {
								let meshphysicalmaterial = new MeshPhysicalMaterial
								this.frontRepo.array_MeshPhysicalMaterials.push(meshphysicalmaterial)
								this.frontRepo.map_ID_MeshPhysicalMaterial.set(meshphysicalmaterialAPI.ID, meshphysicalmaterial)
							}
						)

						// init the arrays
						this.frontRepo.array_PlaneGeometrys = []
						this.frontRepo.map_ID_PlaneGeometry.clear()

						planegeometrys.forEach(
							planegeometryAPI => {
								let planegeometry = new PlaneGeometry
								this.frontRepo.array_PlaneGeometrys.push(planegeometry)
								this.frontRepo.map_ID_PlaneGeometry.set(planegeometryAPI.ID, planegeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_Shapes = []
						this.frontRepo.map_ID_Shape.clear()

						shapes.forEach(
							shapeAPI => {
								let shape = new Shape
								this.frontRepo.array_Shapes.push(shape)
								this.frontRepo.map_ID_Shape.set(shapeAPI.ID, shape)
							}
						)

						// init the arrays
						this.frontRepo.array_SphereGeometrys = []
						this.frontRepo.map_ID_SphereGeometry.clear()

						spheregeometrys.forEach(
							spheregeometryAPI => {
								let spheregeometry = new SphereGeometry
								this.frontRepo.array_SphereGeometrys.push(spheregeometry)
								this.frontRepo.map_ID_SphereGeometry.set(spheregeometryAPI.ID, spheregeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_TorusGeometrys = []
						this.frontRepo.map_ID_TorusGeometry.clear()

						torusgeometrys.forEach(
							torusgeometryAPI => {
								let torusgeometry = new TorusGeometry
								this.frontRepo.array_TorusGeometrys.push(torusgeometry)
								this.frontRepo.map_ID_TorusGeometry.set(torusgeometryAPI.ID, torusgeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_TubeGeometrys = []
						this.frontRepo.map_ID_TubeGeometry.clear()

						tubegeometrys.forEach(
							tubegeometryAPI => {
								let tubegeometry = new TubeGeometry
								this.frontRepo.array_TubeGeometrys.push(tubegeometry)
								this.frontRepo.map_ID_TubeGeometry.set(tubegeometryAPI.ID, tubegeometry)
							}
						)

						// init the arrays
						this.frontRepo.array_Vector2s = []
						this.frontRepo.map_ID_Vector2.clear()

						vector2s.forEach(
							vector2API => {
								let vector2 = new Vector2
								this.frontRepo.array_Vector2s.push(vector2)
								this.frontRepo.map_ID_Vector2.set(vector2API.ID, vector2)
							}
						)

						// init the arrays
						this.frontRepo.array_Vector3s = []
						this.frontRepo.map_ID_Vector3.clear()

						vector3s.forEach(
							vector3API => {
								let vector3 = new Vector3
								this.frontRepo.array_Vector3s.push(vector3)
								this.frontRepo.map_ID_Vector3.set(vector3API.ID, vector3)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						ambiantlights.forEach(
							ambiantlightAPI => {
								let ambiantlight = this.frontRepo.map_ID_AmbiantLight.get(ambiantlightAPI.ID)
								CopyAmbiantLightAPIToAmbiantLight(ambiantlightAPI, ambiantlight!, this.frontRepo)
							}
						)

						// fill up front objects
						boxgeometrys.forEach(
							boxgeometryAPI => {
								let boxgeometry = this.frontRepo.map_ID_BoxGeometry.get(boxgeometryAPI.ID)
								CopyBoxGeometryAPIToBoxGeometry(boxgeometryAPI, boxgeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						cameras.forEach(
							cameraAPI => {
								let camera = this.frontRepo.map_ID_Camera.get(cameraAPI.ID)
								CopyCameraAPIToCamera(cameraAPI, camera!, this.frontRepo)
							}
						)

						// fill up front objects
						canvass.forEach(
							canvasAPI => {
								let canvas = this.frontRepo.map_ID_Canvas.get(canvasAPI.ID)
								CopyCanvasAPIToCanvas(canvasAPI, canvas!, this.frontRepo)
							}
						)

						// fill up front objects
						curves.forEach(
							curveAPI => {
								let curve = this.frontRepo.map_ID_Curve.get(curveAPI.ID)
								CopyCurveAPIToCurve(curveAPI, curve!, this.frontRepo)
							}
						)

						// fill up front objects
						cylindergeometrys.forEach(
							cylindergeometryAPI => {
								let cylindergeometry = this.frontRepo.map_ID_CylinderGeometry.get(cylindergeometryAPI.ID)
								CopyCylinderGeometryAPIToCylinderGeometry(cylindergeometryAPI, cylindergeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						directionallights.forEach(
							directionallightAPI => {
								let directionallight = this.frontRepo.map_ID_DirectionalLight.get(directionallightAPI.ID)
								CopyDirectionalLightAPIToDirectionalLight(directionallightAPI, directionallight!, this.frontRepo)
							}
						)

						// fill up front objects
						extrudegeometrys.forEach(
							extrudegeometryAPI => {
								let extrudegeometry = this.frontRepo.map_ID_ExtrudeGeometry.get(extrudegeometryAPI.ID)
								CopyExtrudeGeometryAPIToExtrudeGeometry(extrudegeometryAPI, extrudegeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						meshs.forEach(
							meshAPI => {
								let mesh = this.frontRepo.map_ID_Mesh.get(meshAPI.ID)
								CopyMeshAPIToMesh(meshAPI, mesh!, this.frontRepo)
							}
						)

						// fill up front objects
						meshmaterialbasics.forEach(
							meshmaterialbasicAPI => {
								let meshmaterialbasic = this.frontRepo.map_ID_MeshMaterialBasic.get(meshmaterialbasicAPI.ID)
								CopyMeshMaterialBasicAPIToMeshMaterialBasic(meshmaterialbasicAPI, meshmaterialbasic!, this.frontRepo)
							}
						)

						// fill up front objects
						meshphysicalmaterials.forEach(
							meshphysicalmaterialAPI => {
								let meshphysicalmaterial = this.frontRepo.map_ID_MeshPhysicalMaterial.get(meshphysicalmaterialAPI.ID)
								CopyMeshPhysicalMaterialAPIToMeshPhysicalMaterial(meshphysicalmaterialAPI, meshphysicalmaterial!, this.frontRepo)
							}
						)

						// fill up front objects
						planegeometrys.forEach(
							planegeometryAPI => {
								let planegeometry = this.frontRepo.map_ID_PlaneGeometry.get(planegeometryAPI.ID)
								CopyPlaneGeometryAPIToPlaneGeometry(planegeometryAPI, planegeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						shapes.forEach(
							shapeAPI => {
								let shape = this.frontRepo.map_ID_Shape.get(shapeAPI.ID)
								CopyShapeAPIToShape(shapeAPI, shape!, this.frontRepo)
							}
						)

						// fill up front objects
						spheregeometrys.forEach(
							spheregeometryAPI => {
								let spheregeometry = this.frontRepo.map_ID_SphereGeometry.get(spheregeometryAPI.ID)
								CopySphereGeometryAPIToSphereGeometry(spheregeometryAPI, spheregeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						torusgeometrys.forEach(
							torusgeometryAPI => {
								let torusgeometry = this.frontRepo.map_ID_TorusGeometry.get(torusgeometryAPI.ID)
								CopyTorusGeometryAPIToTorusGeometry(torusgeometryAPI, torusgeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						tubegeometrys.forEach(
							tubegeometryAPI => {
								let tubegeometry = this.frontRepo.map_ID_TubeGeometry.get(tubegeometryAPI.ID)
								CopyTubeGeometryAPIToTubeGeometry(tubegeometryAPI, tubegeometry!, this.frontRepo)
							}
						)

						// fill up front objects
						vector2s.forEach(
							vector2API => {
								let vector2 = this.frontRepo.map_ID_Vector2.get(vector2API.ID)
								CopyVector2APIToVector2(vector2API, vector2!, this.frontRepo)
							}
						)

						// fill up front objects
						vector3s.forEach(
							vector3API => {
								let vector3 = this.frontRepo.map_ID_Vector3.get(vector3API.ID)
								CopyVector3APIToVector3(vector3API, vector3!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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


				observer.next(frontRepo)
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
export function getCameraUniqueID(id: number): number {
	return 41 * id
}
export function getCanvasUniqueID(id: number): number {
	return 43 * id
}
export function getCurveUniqueID(id: number): number {
	return 47 * id
}
export function getCylinderGeometryUniqueID(id: number): number {
	return 53 * id
}
export function getDirectionalLightUniqueID(id: number): number {
	return 59 * id
}
export function getExtrudeGeometryUniqueID(id: number): number {
	return 61 * id
}
export function getMeshUniqueID(id: number): number {
	return 67 * id
}
export function getMeshMaterialBasicUniqueID(id: number): number {
	return 71 * id
}
export function getMeshPhysicalMaterialUniqueID(id: number): number {
	return 73 * id
}
export function getPlaneGeometryUniqueID(id: number): number {
	return 79 * id
}
export function getShapeUniqueID(id: number): number {
	return 83 * id
}
export function getSphereGeometryUniqueID(id: number): number {
	return 89 * id
}
export function getTorusGeometryUniqueID(id: number): number {
	return 97 * id
}
export function getTubeGeometryUniqueID(id: number): number {
	return 101 * id
}
export function getVector2UniqueID(id: number): number {
	return 103 * id
}
export function getVector3UniqueID(id: number): number {
	return 107 * id
}
