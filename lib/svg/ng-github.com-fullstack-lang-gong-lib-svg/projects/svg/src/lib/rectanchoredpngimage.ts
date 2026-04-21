// generated code - do not edit

import { RectAnchoredPngImageAPI } from './rectanchoredpngimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredPngImage {

	static GONGSTRUCT_NAME = "RectAnchoredPngImage"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Height: number = 0
	RX: number = 0
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	Base64Content: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyRectAnchoredPngImageToRectAnchoredPngImageAPI(rectanchoredpngimage: RectAnchoredPngImage, rectanchoredpngimageAPI: RectAnchoredPngImageAPI) {

	rectanchoredpngimageAPI.CreatedAt = rectanchoredpngimage.CreatedAt
	rectanchoredpngimageAPI.DeletedAt = rectanchoredpngimage.DeletedAt
	rectanchoredpngimageAPI.ID = rectanchoredpngimage.ID

	// insertion point for basic fields copy operations
	rectanchoredpngimageAPI.Name = rectanchoredpngimage.Name
	rectanchoredpngimageAPI.X = rectanchoredpngimage.X
	rectanchoredpngimageAPI.Y = rectanchoredpngimage.Y
	rectanchoredpngimageAPI.Width = rectanchoredpngimage.Width
	rectanchoredpngimageAPI.Height = rectanchoredpngimage.Height
	rectanchoredpngimageAPI.RX = rectanchoredpngimage.RX
	rectanchoredpngimageAPI.X_Offset = rectanchoredpngimage.X_Offset
	rectanchoredpngimageAPI.Y_Offset = rectanchoredpngimage.Y_Offset
	rectanchoredpngimageAPI.RectAnchorType = rectanchoredpngimage.RectAnchorType
	rectanchoredpngimageAPI.Base64Content = rectanchoredpngimage.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyRectAnchoredPngImageAPIToRectAnchoredPngImage update basic, pointers and slice of pointers fields of rectanchoredpngimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredpngimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredPngImageAPIToRectAnchoredPngImage(rectanchoredpngimageAPI: RectAnchoredPngImageAPI, rectanchoredpngimage: RectAnchoredPngImage, frontRepo: FrontRepo) {

	rectanchoredpngimage.CreatedAt = rectanchoredpngimageAPI.CreatedAt
	rectanchoredpngimage.DeletedAt = rectanchoredpngimageAPI.DeletedAt
	rectanchoredpngimage.ID = rectanchoredpngimageAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredpngimage.Name = rectanchoredpngimageAPI.Name
	rectanchoredpngimage.X = rectanchoredpngimageAPI.X
	rectanchoredpngimage.Y = rectanchoredpngimageAPI.Y
	rectanchoredpngimage.Width = rectanchoredpngimageAPI.Width
	rectanchoredpngimage.Height = rectanchoredpngimageAPI.Height
	rectanchoredpngimage.RX = rectanchoredpngimageAPI.RX
	rectanchoredpngimage.X_Offset = rectanchoredpngimageAPI.X_Offset
	rectanchoredpngimage.Y_Offset = rectanchoredpngimageAPI.Y_Offset
	rectanchoredpngimage.RectAnchorType = rectanchoredpngimageAPI.RectAnchorType
	rectanchoredpngimage.Base64Content = rectanchoredpngimageAPI.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
