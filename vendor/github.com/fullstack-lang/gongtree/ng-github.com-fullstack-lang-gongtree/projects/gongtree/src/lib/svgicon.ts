// generated code - do not edit

import { SVGIconAPI } from './svgicon-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGIcon {

	static GONGSTRUCT_NAME = "SVGIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySVGIconToSVGIconAPI(svgicon: SVGIcon, svgiconAPI: SVGIconAPI) {

	svgiconAPI.CreatedAt = svgicon.CreatedAt
	svgiconAPI.DeletedAt = svgicon.DeletedAt
	svgiconAPI.ID = svgicon.ID

	// insertion point for basic fields copy operations
	svgiconAPI.Name = svgicon.Name
	svgiconAPI.SVG = svgicon.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySVGIconAPIToSVGIcon update basic, pointers and slice of pointers fields of svgicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgiconAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySVGIconAPIToSVGIcon(svgiconAPI: SVGIconAPI, svgicon: SVGIcon, frontRepo: FrontRepo) {

	svgicon.CreatedAt = svgiconAPI.CreatedAt
	svgicon.DeletedAt = svgiconAPI.DeletedAt
	svgicon.ID = svgiconAPI.ID

	// insertion point for basic fields copy operations
	svgicon.Name = svgiconAPI.Name
	svgicon.SVG = svgiconAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
