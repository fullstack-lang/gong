// generated code - do not edit

import { PolygoneDB } from './polygone-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Polygone {

	static GONGSTRUCT_NAME = "Polygone"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Points: string = ""
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

export function CopyPolygoneToPolygoneDB(polygone: Polygone, polygoneDB: PolygoneDB) {

	polygoneDB.CreatedAt = polygone.CreatedAt
	polygoneDB.DeletedAt = polygone.DeletedAt
	polygoneDB.ID = polygone.ID
	
	// insertion point for basic fields copy operations
	polygoneDB.Name = polygone.Name
	polygoneDB.Points = polygone.Points
	polygoneDB.Color = polygone.Color
	polygoneDB.FillOpacity = polygone.FillOpacity
	polygoneDB.Stroke = polygone.Stroke
	polygoneDB.StrokeWidth = polygone.StrokeWidth
	polygoneDB.StrokeDashArray = polygone.StrokeDashArray
	polygoneDB.StrokeDashArrayWhenSelected = polygone.StrokeDashArrayWhenSelected
	polygoneDB.Transform = polygone.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polygoneDB.PolygonePointersEncoding.Animates = []
    for (let _animate of polygone.Animates) {
		polygoneDB.PolygonePointersEncoding.Animates.push(_animate.ID)
    }
	
}

export function CopyPolygoneDBToPolygone(polygoneDB: PolygoneDB, polygone: Polygone, frontRepo: FrontRepo) {

	polygone.CreatedAt = polygoneDB.CreatedAt
	polygone.DeletedAt = polygoneDB.DeletedAt
	polygone.ID = polygoneDB.ID
	
	// insertion point for basic fields copy operations
	polygone.Name = polygoneDB.Name
	polygone.Points = polygoneDB.Points
	polygone.Color = polygoneDB.Color
	polygone.FillOpacity = polygoneDB.FillOpacity
	polygone.Stroke = polygoneDB.Stroke
	polygone.StrokeWidth = polygoneDB.StrokeWidth
	polygone.StrokeDashArray = polygoneDB.StrokeDashArray
	polygone.StrokeDashArrayWhenSelected = polygoneDB.StrokeDashArrayWhenSelected
	polygone.Transform = polygoneDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	polygone.Animates = new Array<Animate>()
	for (let _id of polygoneDB.PolygonePointersEncoding.Animates) {
	  let _animate = frontRepo.Animates.get(_id)
	  if (_animate != undefined) {
		polygone.Animates.push(_animate!)
	  }
	}
}
