// generated code - do not edit

import { SvgAPI } from './svg-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Svg {

	static GONGSTRUCT_NAME = "Svg"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySvgToSvgAPI(svg: Svg, svgAPI: SvgAPI) {

	svgAPI.CreatedAt = svg.CreatedAt
	svgAPI.DeletedAt = svg.DeletedAt
	svgAPI.ID = svg.ID

	// insertion point for basic fields copy operations
	svgAPI.Name = svg.Name
	svgAPI.StackName = svg.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySvgAPIToSvg update basic, pointers and slice of pointers fields of svg
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySvgAPIToSvg(svgAPI: SvgAPI, svg: Svg, frontRepo: FrontRepo) {

	svg.CreatedAt = svgAPI.CreatedAt
	svg.DeletedAt = svgAPI.DeletedAt
	svg.ID = svgAPI.ID

	// insertion point for basic fields copy operations
	svg.Name = svgAPI.Name
	svg.StackName = svgAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
