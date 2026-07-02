// generated code - do not edit

import { LogoOnTheRightAPI } from './logoontheright-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LogoOnTheRight {

	static GONGSTRUCT_NAME = "LogoOnTheRight"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyLogoOnTheRightToLogoOnTheRightAPI(logoontheright: LogoOnTheRight, logoontherightAPI: LogoOnTheRightAPI) {

	logoontherightAPI.CreatedAt = logoontheright.CreatedAt
	logoontherightAPI.DeletedAt = logoontheright.DeletedAt
	logoontherightAPI.ID = logoontheright.ID

	// insertion point for basic fields copy operations
	logoontherightAPI.Name = logoontheright.Name
	logoontherightAPI.Width = logoontheright.Width
	logoontherightAPI.Height = logoontheright.Height
	logoontherightAPI.SVG = logoontheright.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLogoOnTheRightAPIToLogoOnTheRight update basic, pointers and slice of pointers fields of logoontheright
// from respectively the basic fields and encoded fields of pointers and slices of pointers of logoontherightAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLogoOnTheRightAPIToLogoOnTheRight(logoontherightAPI: LogoOnTheRightAPI, logoontheright: LogoOnTheRight, frontRepo: FrontRepo) {

	logoontheright.CreatedAt = logoontherightAPI.CreatedAt
	logoontheright.DeletedAt = logoontherightAPI.DeletedAt
	logoontheright.ID = logoontherightAPI.ID

	// insertion point for basic fields copy operations
	logoontheright.Name = logoontherightAPI.Name
	logoontheright.Width = logoontherightAPI.Width
	logoontheright.Height = logoontherightAPI.Height
	logoontheright.SVG = logoontherightAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
