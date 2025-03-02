// generated code - do not edit

import { PathAPI } from './path-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Path {

	static GONGSTRUCT_NAME = "Path"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyPathToPathAPI(path: Path, pathAPI: PathAPI) {

	pathAPI.CreatedAt = path.CreatedAt
	pathAPI.DeletedAt = path.DeletedAt
	pathAPI.ID = path.ID

	// insertion point for basic fields copy operations
	pathAPI.Name = path.Name
	pathAPI.Definition = path.Definition
	pathAPI.Color = path.Color
	pathAPI.FillOpacity = path.FillOpacity
	pathAPI.Stroke = path.Stroke
	pathAPI.StrokeOpacity = path.StrokeOpacity
	pathAPI.StrokeWidth = path.StrokeWidth
	pathAPI.StrokeDashArray = path.StrokeDashArray
	pathAPI.StrokeDashArrayWhenSelected = path.StrokeDashArrayWhenSelected
	pathAPI.Transform = path.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	pathAPI.PathPointersEncoding.Animates = []
	for (let _animate of path.Animates) {
		pathAPI.PathPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyPathAPIToPath update basic, pointers and slice of pointers fields of path
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pathAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPathAPIToPath(pathAPI: PathAPI, path: Path, frontRepo: FrontRepo) {

	path.CreatedAt = pathAPI.CreatedAt
	path.DeletedAt = pathAPI.DeletedAt
	path.ID = pathAPI.ID

	// insertion point for basic fields copy operations
	path.Name = pathAPI.Name
	path.Definition = pathAPI.Definition
	path.Color = pathAPI.Color
	path.FillOpacity = pathAPI.FillOpacity
	path.Stroke = pathAPI.Stroke
	path.StrokeOpacity = pathAPI.StrokeOpacity
	path.StrokeWidth = pathAPI.StrokeWidth
	path.StrokeDashArray = pathAPI.StrokeDashArray
	path.StrokeDashArrayWhenSelected = pathAPI.StrokeDashArrayWhenSelected
	path.Transform = pathAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(pathAPI.PathPointersEncoding.Animates)) {
		console.error('Rects is not an array:', pathAPI.PathPointersEncoding.Animates);
		return;
	}

	path.Animates = new Array<Animate>()
	for (let _id of pathAPI.PathPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			path.Animates.push(_animate!)
		}
	}
}
