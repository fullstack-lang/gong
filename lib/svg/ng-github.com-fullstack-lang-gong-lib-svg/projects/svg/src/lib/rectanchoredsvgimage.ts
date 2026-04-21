// generated code - do not edit

import { RectAnchoredSvgImageAPI } from './rectanchoredsvgimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredSvgImage {

	static GONGSTRUCT_NAME = "RectAnchoredSvgImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyRectAnchoredSvgImageToRectAnchoredSvgImageAPI(rectanchoredsvgimage: RectAnchoredSvgImage, rectanchoredsvgimageAPI: RectAnchoredSvgImageAPI) {

	rectanchoredsvgimageAPI.CreatedAt = rectanchoredsvgimage.CreatedAt
	rectanchoredsvgimageAPI.DeletedAt = rectanchoredsvgimage.DeletedAt
	rectanchoredsvgimageAPI.ID = rectanchoredsvgimage.ID

	// insertion point for basic fields copy operations
	rectanchoredsvgimageAPI.Name = rectanchoredsvgimage.Name
	rectanchoredsvgimageAPI.Content = rectanchoredsvgimage.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyRectAnchoredSvgImageAPIToRectAnchoredSvgImage update basic, pointers and slice of pointers fields of rectanchoredsvgimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredsvgimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredSvgImageAPIToRectAnchoredSvgImage(rectanchoredsvgimageAPI: RectAnchoredSvgImageAPI, rectanchoredsvgimage: RectAnchoredSvgImage, frontRepo: FrontRepo) {

	rectanchoredsvgimage.CreatedAt = rectanchoredsvgimageAPI.CreatedAt
	rectanchoredsvgimage.DeletedAt = rectanchoredsvgimageAPI.DeletedAt
	rectanchoredsvgimage.ID = rectanchoredsvgimageAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredsvgimage.Name = rectanchoredsvgimageAPI.Name
	rectanchoredsvgimage.Content = rectanchoredsvgimageAPI.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
