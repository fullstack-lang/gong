// generated code - do not edit

import { FavIconAPI } from './favicon-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FavIcon {

	static GONGSTRUCT_NAME = "FavIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFavIconToFavIconAPI(favicon: FavIcon, faviconAPI: FavIconAPI) {

	faviconAPI.CreatedAt = favicon.CreatedAt
	faviconAPI.DeletedAt = favicon.DeletedAt
	faviconAPI.ID = favicon.ID

	// insertion point for basic fields copy operations
	faviconAPI.Name = favicon.Name
	faviconAPI.SVG = favicon.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFavIconAPIToFavIcon update basic, pointers and slice of pointers fields of favicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of faviconAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFavIconAPIToFavIcon(faviconAPI: FavIconAPI, favicon: FavIcon, frontRepo: FrontRepo) {

	favicon.CreatedAt = faviconAPI.CreatedAt
	favicon.DeletedAt = faviconAPI.DeletedAt
	favicon.ID = faviconAPI.ID

	// insertion point for basic fields copy operations
	favicon.Name = faviconAPI.Name
	favicon.SVG = faviconAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
