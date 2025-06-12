// generated code - do not edit

import { LogoAPI } from './logo-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Logo {

	static GONGSTRUCT_NAME = "Logo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyLogoToLogoAPI(logo: Logo, logoAPI: LogoAPI) {

	logoAPI.CreatedAt = logo.CreatedAt
	logoAPI.DeletedAt = logo.DeletedAt
	logoAPI.ID = logo.ID

	// insertion point for basic fields copy operations
	logoAPI.Name = logo.Name
	logoAPI.SVG = logo.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLogoAPIToLogo update basic, pointers and slice of pointers fields of logo
// from respectively the basic fields and encoded fields of pointers and slices of pointers of logoAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLogoAPIToLogo(logoAPI: LogoAPI, logo: Logo, frontRepo: FrontRepo) {

	logo.CreatedAt = logoAPI.CreatedAt
	logo.DeletedAt = logoAPI.DeletedAt
	logo.ID = logoAPI.ID

	// insertion point for basic fields copy operations
	logo.Name = logoAPI.Name
	logo.SVG = logoAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
