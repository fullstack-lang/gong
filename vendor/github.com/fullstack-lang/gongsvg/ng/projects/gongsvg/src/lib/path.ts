// generated code - do not edit

import { PathDB } from './path-db'
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
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyPathToPathDB(path: Path, pathDB: PathDB) {

	pathDB.CreatedAt = path.CreatedAt
	pathDB.DeletedAt = path.DeletedAt
	pathDB.ID = path.ID
	
	// insertion point for basic fields copy operations
	pathDB.Name = path.Name
	pathDB.Definition = path.Definition
	pathDB.Color = path.Color
	pathDB.FillOpacity = path.FillOpacity
	pathDB.Stroke = path.Stroke
	pathDB.StrokeWidth = path.StrokeWidth
	pathDB.StrokeDashArray = path.StrokeDashArray
	pathDB.StrokeDashArrayWhenSelected = path.StrokeDashArrayWhenSelected
	pathDB.Transform = path.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	pathDB.PathPointersEncoding.Animates = []
    for (let _animate of path.Animates) {
		pathDB.PathPointersEncoding.Animates.push(_animate.ID)
    }
	
}

export function CopyPathDBToPath(pathDB: PathDB, path: Path, frontRepo: FrontRepo) {

	path.CreatedAt = pathDB.CreatedAt
	path.DeletedAt = pathDB.DeletedAt
	path.ID = pathDB.ID
	
	// insertion point for basic fields copy operations
	path.Name = pathDB.Name
	path.Definition = pathDB.Definition
	path.Color = pathDB.Color
	path.FillOpacity = pathDB.FillOpacity
	path.Stroke = pathDB.Stroke
	path.StrokeWidth = pathDB.StrokeWidth
	path.StrokeDashArray = pathDB.StrokeDashArray
	path.StrokeDashArrayWhenSelected = pathDB.StrokeDashArrayWhenSelected
	path.Transform = pathDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	path.Animates = new Array<Animate>()
	for (let _id of pathDB.PathPointersEncoding.Animates) {
	  let _animate = frontRepo.Animates.get(_id)
	  if (_animate != undefined) {
		path.Animates.push(_animate!)
	  }
	}
}
