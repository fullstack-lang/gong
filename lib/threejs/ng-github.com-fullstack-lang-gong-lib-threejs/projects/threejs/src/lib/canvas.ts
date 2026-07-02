// generated code - do not edit

import { CanvasAPI } from './canvas-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { DirectionalLight } from './directionallight'
import { AmbiantLight } from './ambiantlight'
import { Mesh } from './mesh'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Canvas {

	static GONGSTRUCT_NAME = "Canvas"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	DirectionalLights: Array<DirectionalLight> = []
	AmbiantLight?: AmbiantLight

	Meshs: Array<Mesh> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCanvasToCanvasAPI(canvas: Canvas, canvasAPI: CanvasAPI) {

	canvasAPI.CreatedAt = canvas.CreatedAt
	canvasAPI.DeletedAt = canvas.DeletedAt
	canvasAPI.ID = canvas.ID

	// insertion point for basic fields copy operations
	canvasAPI.Name = canvas.Name

	// insertion point for pointer fields encoding
	canvasAPI.CanvasPointersEncoding.AmbiantLightID.Valid = true
	if (canvas.AmbiantLight != undefined) {
		canvasAPI.CanvasPointersEncoding.AmbiantLightID.Int64 = canvas.AmbiantLight.ID  
	} else {
		canvasAPI.CanvasPointersEncoding.AmbiantLightID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	canvasAPI.CanvasPointersEncoding.DirectionalLights = []
	for (let _directionallight of canvas.DirectionalLights) {
		canvasAPI.CanvasPointersEncoding.DirectionalLights.push(_directionallight.ID)
	}

	canvasAPI.CanvasPointersEncoding.Meshs = []
	for (let _mesh of canvas.Meshs) {
		canvasAPI.CanvasPointersEncoding.Meshs.push(_mesh.ID)
	}

}

// CopyCanvasAPIToCanvas update basic, pointers and slice of pointers fields of canvas
// from respectively the basic fields and encoded fields of pointers and slices of pointers of canvasAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCanvasAPIToCanvas(canvasAPI: CanvasAPI, canvas: Canvas, frontRepo: FrontRepo) {

	canvas.CreatedAt = canvasAPI.CreatedAt
	canvas.DeletedAt = canvasAPI.DeletedAt
	canvas.ID = canvasAPI.ID

	// insertion point for basic fields copy operations
	canvas.Name = canvasAPI.Name

	// insertion point for pointer fields encoding
	canvas.AmbiantLight = frontRepo.map_ID_AmbiantLight.get(canvasAPI.CanvasPointersEncoding.AmbiantLightID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(canvasAPI.CanvasPointersEncoding.DirectionalLights)) {
		console.error('Rects is not an array:', canvasAPI.CanvasPointersEncoding.DirectionalLights);
		return;
	}

	canvas.DirectionalLights = new Array<DirectionalLight>()
	for (let _id of canvasAPI.CanvasPointersEncoding.DirectionalLights) {
		let _directionallight = frontRepo.map_ID_DirectionalLight.get(_id)
		if (_directionallight != undefined) {
			canvas.DirectionalLights.push(_directionallight!)
		}
	}
	if (!Array.isArray(canvasAPI.CanvasPointersEncoding.Meshs)) {
		console.error('Rects is not an array:', canvasAPI.CanvasPointersEncoding.Meshs);
		return;
	}

	canvas.Meshs = new Array<Mesh>()
	for (let _id of canvasAPI.CanvasPointersEncoding.Meshs) {
		let _mesh = frontRepo.map_ID_Mesh.get(_id)
		if (_mesh != undefined) {
			canvas.Meshs.push(_mesh!)
		}
	}
}
