// generated code - do not edit

import { ContentAPI } from './content-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Content {

	static GONGSTRUCT_NAME = "Content"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyContentToContentAPI(content: Content, contentAPI: ContentAPI) {

	contentAPI.CreatedAt = content.CreatedAt
	contentAPI.DeletedAt = content.DeletedAt
	contentAPI.ID = content.ID

	// insertion point for basic fields copy operations
	contentAPI.Name = content.Name
	contentAPI.Content = content.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyContentAPIToContent update basic, pointers and slice of pointers fields of content
// from respectively the basic fields and encoded fields of pointers and slices of pointers of contentAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyContentAPIToContent(contentAPI: ContentAPI, content: Content, frontRepo: FrontRepo) {

	content.CreatedAt = contentAPI.CreatedAt
	content.DeletedAt = contentAPI.DeletedAt
	content.ID = contentAPI.ID

	// insertion point for basic fields copy operations
	content.Name = contentAPI.Name
	content.Content = contentAPI.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
