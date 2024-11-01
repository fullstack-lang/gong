// generated code - do not edit

import { SVGAPI } from './svg-api'
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

export function CopySVGToSVGAPI(svg: SVG, svgAPI: SVGAPI) {

	svgAPI.CreatedAt = svg.CreatedAt
	svgAPI.DeletedAt = svg.DeletedAt
	svgAPI.ID = svg.ID

	// insertion point for basic fields copy operations
	svgAPI.Name = svg.Name
	svgAPI.DrawingState = svg.DrawingState
	svgAPI.IsEditable = svg.IsEditable

	// insertion point for pointer fields encoding
	svgAPI.SVGPointersEncoding.StartRectID.Valid = true
	if (svg.StartRect != undefined) {
		svgAPI.SVGPointersEncoding.StartRectID.Int64 = svg.StartRect.ID  
	} else {
		svgAPI.SVGPointersEncoding.StartRectID.Int64 = 0 		
	}

	svgAPI.SVGPointersEncoding.EndRectID.Valid = true
	if (svg.EndRect != undefined) {
		svgAPI.SVGPointersEncoding.EndRectID.Int64 = svg.EndRect.ID  
	} else {
		svgAPI.SVGPointersEncoding.EndRectID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	svgAPI.SVGPointersEncoding.Layers = []
	for (let _layer of svg.Layers) {
		svgAPI.SVGPointersEncoding.Layers.push(_layer.ID)
	}

}

// CopySVGAPIToSVG update basic, pointers and slice of pointers fields of svg
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySVGAPIToSVG(svgAPI: SVGAPI, svg: SVG, frontRepo: FrontRepo) {

	svg.CreatedAt = svgAPI.CreatedAt
	svg.DeletedAt = svgAPI.DeletedAt
	svg.ID = svgAPI.ID

	// insertion point for basic fields copy operations
	svg.Name = svgAPI.Name
	svg.DrawingState = svgAPI.DrawingState
	svg.IsEditable = svgAPI.IsEditable

	// insertion point for pointer fields encoding
	svg.StartRect = frontRepo.map_ID_Rect.get(svgAPI.SVGPointersEncoding.StartRectID.Int64)
	svg.EndRect = frontRepo.map_ID_Rect.get(svgAPI.SVGPointersEncoding.EndRectID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(svgAPI.SVGPointersEncoding.Layers)) {
		console.error('Rects is not an array:', svgAPI.SVGPointersEncoding.Layers);
		return;
	}

	svg.Layers = new Array<Layer>()
	for (let _id of svgAPI.SVGPointersEncoding.Layers) {
		let _layer = frontRepo.map_ID_Layer.get(_id)
		if (_layer != undefined) {
			svg.Layers.push(_layer!)
		}
	}
}
