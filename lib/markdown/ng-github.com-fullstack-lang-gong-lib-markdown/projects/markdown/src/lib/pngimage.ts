// generated code - do not edit

import { PngImageAPI } from './pngimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PngImage {

	static GONGSTRUCT_NAME = "PngImage"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyPngImageToPngImageAPI(pngimage: PngImage, pngimageAPI: PngImageAPI) {

	pngimageAPI.CreatedAt = pngimage.CreatedAt
	pngimageAPI.DeletedAt = pngimage.DeletedAt
	pngimageAPI.ID = pngimage.ID

	// insertion point for basic fields copy operations
	pngimageAPI.Name = pngimage.Name
	pngimageAPI.Base64Content = pngimage.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyPngImageAPIToPngImage update basic, pointers and slice of pointers fields of pngimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pngimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPngImageAPIToPngImage(pngimageAPI: PngImageAPI, pngimage: PngImage, frontRepo: FrontRepo) {

	pngimage.CreatedAt = pngimageAPI.CreatedAt
	pngimage.DeletedAt = pngimageAPI.DeletedAt
	pngimage.ID = pngimageAPI.ID

	// insertion point for basic fields copy operations
	pngimage.Name = pngimageAPI.Name
	pngimage.Base64Content = pngimageAPI.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
