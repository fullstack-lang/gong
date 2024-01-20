// generated code - do not edit

import { SVGDB } from './svg-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Layer } from './layer'
import { Rect } from './rect'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVG {

	static GONGSTRUCT_NAME = "SVG"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DrawingState: string = ""
	IsEditable: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Layers: Array<Layer> = []
	StartRect?: Rect

	EndRect?: Rect

}

export function CopySVGToSVGDB(svg: SVG, svgDB: SVGDB) {

	svgDB.CreatedAt = svg.CreatedAt
	svgDB.DeletedAt = svg.DeletedAt
	svgDB.ID = svg.ID
	
	// insertion point for basic fields copy operations
	svgDB.Name = svg.Name
	svgDB.DrawingState = svg.DrawingState
	svgDB.IsEditable = svg.IsEditable

	// insertion point for pointer fields encoding
    svgDB.SVGPointersEncoding.StartRectID.Valid = true
	if (svg.StartRect != undefined) {
		svgDB.SVGPointersEncoding.StartRectID.Int64 = svg.StartRect.ID  
    } else {
		svgDB.SVGPointersEncoding.StartRectID.Int64 = 0 		
	}

    svgDB.SVGPointersEncoding.EndRectID.Valid = true
	if (svg.EndRect != undefined) {
		svgDB.SVGPointersEncoding.EndRectID.Int64 = svg.EndRect.ID  
    } else {
		svgDB.SVGPointersEncoding.EndRectID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	svgDB.SVGPointersEncoding.Layers = []
    for (let _layer of svg.Layers) {
		svgDB.SVGPointersEncoding.Layers.push(_layer.ID)
    }
	
}

export function CopySVGDBToSVG(svgDB: SVGDB, svg: SVG, frontRepo: FrontRepo) {

	svg.CreatedAt = svgDB.CreatedAt
	svg.DeletedAt = svgDB.DeletedAt
	svg.ID = svgDB.ID
	
	// insertion point for basic fields copy operations
	svg.Name = svgDB.Name
	svg.DrawingState = svgDB.DrawingState
	svg.IsEditable = svgDB.IsEditable

	// insertion point for pointer fields encoding
	svg.StartRect = frontRepo.Rects.get(svgDB.SVGPointersEncoding.StartRectID.Int64)
	svg.EndRect = frontRepo.Rects.get(svgDB.SVGPointersEncoding.EndRectID.Int64)

	// insertion point for slice of pointers fields encoding
	svg.Layers = new Array<Layer>()
	for (let _id of svgDB.SVGPointersEncoding.Layers) {
	  let _layer = frontRepo.Layers.get(_id)
	  if (_layer != undefined) {
		svg.Layers.push(_layer!)
	  }
	}
}
