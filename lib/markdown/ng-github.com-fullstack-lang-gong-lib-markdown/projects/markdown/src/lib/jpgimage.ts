// generated code - do not edit

import { JpgImageAPI } from './jpgimage-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class JpgImage {

	static GONGSTRUCT_NAME = "JpgImage"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64Content: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyJpgImageToJpgImageAPI(jpgimage: JpgImage, jpgimageAPI: JpgImageAPI) {

	jpgimageAPI.CreatedAt = jpgimage.CreatedAt
	jpgimageAPI.DeletedAt = jpgimage.DeletedAt
	jpgimageAPI.ID = jpgimage.ID

	// insertion point for basic fields copy operations
	jpgimageAPI.Name = jpgimage.Name
	jpgimageAPI.Base64Content = jpgimage.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyJpgImageAPIToJpgImage update basic, pointers and slice of pointers fields of jpgimage
// from respectively the basic fields and encoded fields of pointers and slices of pointers of jpgimageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyJpgImageAPIToJpgImage(jpgimageAPI: JpgImageAPI, jpgimage: JpgImage, frontRepo: FrontRepo) {

	jpgimage.CreatedAt = jpgimageAPI.CreatedAt
	jpgimage.DeletedAt = jpgimageAPI.DeletedAt
	jpgimage.ID = jpgimageAPI.ID

	// insertion point for basic fields copy operations
	jpgimage.Name = jpgimageAPI.Name
	jpgimage.Base64Content = jpgimageAPI.Base64Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
