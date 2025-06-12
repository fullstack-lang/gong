// generated code - do not edit

import { TitleAPI } from './title-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Title {

	static GONGSTRUCT_NAME = "Title"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyTitleToTitleAPI(title: Title, titleAPI: TitleAPI) {

	titleAPI.CreatedAt = title.CreatedAt
	titleAPI.DeletedAt = title.DeletedAt
	titleAPI.ID = title.ID

	// insertion point for basic fields copy operations
	titleAPI.Name = title.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyTitleAPIToTitle update basic, pointers and slice of pointers fields of title
// from respectively the basic fields and encoded fields of pointers and slices of pointers of titleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTitleAPIToTitle(titleAPI: TitleAPI, title: Title, frontRepo: FrontRepo) {

	title.CreatedAt = titleAPI.CreatedAt
	title.DeletedAt = titleAPI.DeletedAt
	title.ID = titleAPI.ID

	// insertion point for basic fields copy operations
	title.Name = titleAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
