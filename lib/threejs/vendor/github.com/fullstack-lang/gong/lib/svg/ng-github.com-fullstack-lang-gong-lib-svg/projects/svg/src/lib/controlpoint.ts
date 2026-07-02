// generated code - do not edit

import { ControlPointAPI } from './controlpoint-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rect } from './rect'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ControlPoint {

	static GONGSTRUCT_NAME = "ControlPoint"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X_Relative: number = 0
	Y_Relative: number = 0

	// insertion point for pointers and slices of pointers declarations
	ClosestRect?: Rect


	CreatedAt?: string
	DeletedAt?: string
}

export function CopyControlPointToControlPointAPI(controlpoint: ControlPoint, controlpointAPI: ControlPointAPI) {

	controlpointAPI.CreatedAt = controlpoint.CreatedAt
	controlpointAPI.DeletedAt = controlpoint.DeletedAt
	controlpointAPI.ID = controlpoint.ID

	// insertion point for basic fields copy operations
	controlpointAPI.Name = controlpoint.Name
	controlpointAPI.X_Relative = controlpoint.X_Relative
	controlpointAPI.Y_Relative = controlpoint.Y_Relative

	// insertion point for pointer fields encoding
	controlpointAPI.ControlPointPointersEncoding.ClosestRectID.Valid = true
	if (controlpoint.ClosestRect != undefined) {
		controlpointAPI.ControlPointPointersEncoding.ClosestRectID.Int64 = controlpoint.ClosestRect.ID  
	} else {
		controlpointAPI.ControlPointPointersEncoding.ClosestRectID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyControlPointAPIToControlPoint update basic, pointers and slice of pointers fields of controlpoint
// from respectively the basic fields and encoded fields of pointers and slices of pointers of controlpointAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyControlPointAPIToControlPoint(controlpointAPI: ControlPointAPI, controlpoint: ControlPoint, frontRepo: FrontRepo) {

	controlpoint.CreatedAt = controlpointAPI.CreatedAt
	controlpoint.DeletedAt = controlpointAPI.DeletedAt
	controlpoint.ID = controlpointAPI.ID

	// insertion point for basic fields copy operations
	controlpoint.Name = controlpointAPI.Name
	controlpoint.X_Relative = controlpointAPI.X_Relative
	controlpoint.Y_Relative = controlpointAPI.Y_Relative

	// insertion point for pointer fields encoding
	controlpoint.ClosestRect = frontRepo.map_ID_Rect.get(controlpointAPI.ControlPointPointersEncoding.ClosestRectID.Int64)

	// insertion point for slice of pointers fields encoding
}
