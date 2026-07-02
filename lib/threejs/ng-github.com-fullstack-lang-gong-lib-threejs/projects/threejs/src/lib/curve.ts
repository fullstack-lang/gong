// generated code - do not edit

import { CurveAPI } from './curve-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Vector3 } from './vector3'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Curve {

	static GONGSTRUCT_NAME = "Curve"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Points: Array<Vector3> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyCurveToCurveAPI(curve: Curve, curveAPI: CurveAPI) {

	curveAPI.CreatedAt = curve.CreatedAt
	curveAPI.DeletedAt = curve.DeletedAt
	curveAPI.ID = curve.ID

	// insertion point for basic fields copy operations
	curveAPI.Name = curve.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	curveAPI.CurvePointersEncoding.Points = []
	for (let _vector3 of curve.Points) {
		curveAPI.CurvePointersEncoding.Points.push(_vector3.ID)
	}

}

// CopyCurveAPIToCurve update basic, pointers and slice of pointers fields of curve
// from respectively the basic fields and encoded fields of pointers and slices of pointers of curveAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCurveAPIToCurve(curveAPI: CurveAPI, curve: Curve, frontRepo: FrontRepo) {

	curve.CreatedAt = curveAPI.CreatedAt
	curve.DeletedAt = curveAPI.DeletedAt
	curve.ID = curveAPI.ID

	// insertion point for basic fields copy operations
	curve.Name = curveAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(curveAPI.CurvePointersEncoding.Points)) {
		console.error('Rects is not an array:', curveAPI.CurvePointersEncoding.Points);
		return;
	}

	curve.Points = new Array<Vector3>()
	for (let _id of curveAPI.CurvePointersEncoding.Points) {
		let _vector3 = frontRepo.map_ID_Vector3.get(_id)
		if (_vector3 != undefined) {
			curve.Points.push(_vector3!)
		}
	}
}
