// generated code - do not edit

import { RectAnchoredJpgImageAPI } from './rectanchoredjpgimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredJpgImage {

	static GONGSTRUCT_NAME = "RectAnchoredJpgImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyRectAnchoredJpgImageToRectAnchoredJpgImageAPI(rectanchoredjpgimage: RectAnchoredJpgImage, rectanchoredjpgimageAPI: RectAnchoredJpgImageAPI) {

	rectanchoredjpgimageAPI.CreatedAt = rectanchoredjpgimage.CreatedAt
	rectanchoredjpgimageAPI.DeletedAt = rectanchoredjpgimage.DeletedAt
	rectanchoredjpgimageAPI.ID = rectanchoredjpgimage.ID

	// insertion point for basic fields copy operations
	rectanchoredjpgimageAPI.Name = rectanchoredjpgimage.Name
	rectanchoredjpgimageAPI.Base64Content = rectanchoredjpgimage.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyRectAnchoredJpgImageAPIToRectAnchoredJpgImage update basic, pointers and slice of pointers fields of rectanchoredjpgimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredjpgimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredJpgImageAPIToRectAnchoredJpgImage(rectanchoredjpgimageAPI: RectAnchoredJpgImageAPI, rectanchoredjpgimage: RectAnchoredJpgImage, frontRepo: FrontRepo) {

	rectanchoredjpgimage.CreatedAt = rectanchoredjpgimageAPI.CreatedAt
	rectanchoredjpgimage.DeletedAt = rectanchoredjpgimageAPI.DeletedAt
	rectanchoredjpgimage.ID = rectanchoredjpgimageAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredjpgimage.Name = rectanchoredjpgimageAPI.Name
	rectanchoredjpgimage.Base64Content = rectanchoredjpgimageAPI.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
