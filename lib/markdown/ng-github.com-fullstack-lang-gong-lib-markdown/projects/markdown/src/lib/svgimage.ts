// generated code - do not edit

import { SvgImageAPI } from './svgimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SvgImage {

	static GONGSTRUCT_NAME = "SvgImage"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySvgImageToSvgImageAPI(svgimage: SvgImage, svgimageAPI: SvgImageAPI) {

	svgimageAPI.CreatedAt = svgimage.CreatedAt
	svgimageAPI.DeletedAt = svgimage.DeletedAt
	svgimageAPI.ID = svgimage.ID

	// insertion point for basic fields copy operations
	svgimageAPI.Name = svgimage.Name
	svgimageAPI.Content = svgimage.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySvgImageAPIToSvgImage update basic, pointers and slice of pointers fields of svgimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of svgimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySvgImageAPIToSvgImage(svgimageAPI: SvgImageAPI, svgimage: SvgImage, frontRepo: FrontRepo) {

	svgimage.CreatedAt = svgimageAPI.CreatedAt
	svgimage.DeletedAt = svgimageAPI.DeletedAt
	svgimage.ID = svgimageAPI.ID

	// insertion point for basic fields copy operations
	svgimage.Name = svgimageAPI.Name
	svgimage.Content = svgimageAPI.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
