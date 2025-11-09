// generated code - do not edit

import { LogoOnTheLeftAPI } from './logoontheleft-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LogoOnTheLeft {

	static GONGSTRUCT_NAME = "LogoOnTheLeft"

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

export function CopyLogoOnTheLeftToLogoOnTheLeftAPI(logoontheleft: LogoOnTheLeft, logoontheleftAPI: LogoOnTheLeftAPI) {

	logoontheleftAPI.CreatedAt = logoontheleft.CreatedAt
	logoontheleftAPI.DeletedAt = logoontheleft.DeletedAt
	logoontheleftAPI.ID = logoontheleft.ID

	// insertion point for basic fields copy operations
	logoontheleftAPI.Name = logoontheleft.Name
	logoontheleftAPI.Width = logoontheleft.Width
	logoontheleftAPI.Height = logoontheleft.Height
	logoontheleftAPI.SVG = logoontheleft.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLogoOnTheLeftAPIToLogoOnTheLeft update basic, pointers and slice of pointers fields of logoontheleft
// from respectively the basic fields and encoded fields of pointers and slices of pointers of logoontheleftAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLogoOnTheLeftAPIToLogoOnTheLeft(logoontheleftAPI: LogoOnTheLeftAPI, logoontheleft: LogoOnTheLeft, frontRepo: FrontRepo) {

	logoontheleft.CreatedAt = logoontheleftAPI.CreatedAt
	logoontheleft.DeletedAt = logoontheleftAPI.DeletedAt
	logoontheleft.ID = logoontheleftAPI.ID

	// insertion point for basic fields copy operations
	logoontheleft.Name = logoontheleftAPI.Name
	logoontheleft.Width = logoontheleftAPI.Width
	logoontheleft.Height = logoontheleftAPI.Height
	logoontheleft.SVG = logoontheleftAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
