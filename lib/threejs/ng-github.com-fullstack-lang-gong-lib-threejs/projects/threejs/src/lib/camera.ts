// generated code - do not edit

import { CameraAPI } from './camera-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Camera {

	static GONGSTRUCT_NAME = "Camera"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Z: number = 0
	TargetX: number = 0
	TargetY: number = 0
	TargetZ: number = 0
	Fov: number = 0

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCameraToCameraAPI(camera: Camera, cameraAPI: CameraAPI) {

	cameraAPI.CreatedAt = camera.CreatedAt
	cameraAPI.DeletedAt = camera.DeletedAt
	cameraAPI.ID = camera.ID

	// insertion point for basic fields copy operations
	cameraAPI.Name = camera.Name
	cameraAPI.X = camera.X
	cameraAPI.Y = camera.Y
	cameraAPI.Z = camera.Z
	cameraAPI.TargetX = camera.TargetX
	cameraAPI.TargetY = camera.TargetY
	cameraAPI.TargetZ = camera.TargetZ
	cameraAPI.Fov = camera.Fov

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCameraAPIToCamera update basic, pointers and slice of pointers fields of camera
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cameraAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCameraAPIToCamera(cameraAPI: CameraAPI, camera: Camera, frontRepo: FrontRepo) {

	camera.CreatedAt = cameraAPI.CreatedAt
	camera.DeletedAt = cameraAPI.DeletedAt
	camera.ID = cameraAPI.ID

	// insertion point for basic fields copy operations
	camera.Name = cameraAPI.Name
	camera.X = cameraAPI.X
	camera.Y = cameraAPI.Y
	camera.Z = cameraAPI.Z
	camera.TargetX = cameraAPI.TargetX
	camera.TargetY = cameraAPI.TargetY
	camera.TargetZ = cameraAPI.TargetZ
	camera.Fov = cameraAPI.Fov

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
