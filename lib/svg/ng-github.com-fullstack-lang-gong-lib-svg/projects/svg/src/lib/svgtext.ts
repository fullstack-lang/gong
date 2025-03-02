// generated code - do not edit

import { SvgTextAPI } from './svgtext-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SvgText {

	static GONGSTRUCT_NAME = "SvgText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySvgTextToSvgTextAPI(svgtext: SvgText, svgtextAPI: SvgTextAPI) {

	svgtextAPI.CreatedAt = svgtext.CreatedAt
	svgtextAPI.DeletedAt = svgtext.DeletedAt
	svgtextAPI.ID = svgtext.ID

	// insertion point for basic fields copy operations
	svgtextAPI.Name = svgtext.Name
	svgtextAPI.Text = svgtext.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySvgTextAPIToSvgText update basic, pointers and slice of pointers fields of svgtext
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgtextAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySvgTextAPIToSvgText(svgtextAPI: SvgTextAPI, svgtext: SvgText, frontRepo: FrontRepo) {

	svgtext.CreatedAt = svgtextAPI.CreatedAt
	svgtext.DeletedAt = svgtextAPI.DeletedAt
	svgtext.ID = svgtextAPI.ID

	// insertion point for basic fields copy operations
	svgtext.Name = svgtextAPI.Name
	svgtext.Text = svgtextAPI.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
