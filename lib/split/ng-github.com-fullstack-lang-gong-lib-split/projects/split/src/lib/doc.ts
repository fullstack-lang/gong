// generated code - do not edit

import { DocAPI } from './doc-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Doc {

	static GONGSTRUCT_NAME = "Doc"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDocToDocAPI(doc: Doc, docAPI: DocAPI) {

	docAPI.CreatedAt = doc.CreatedAt
	docAPI.DeletedAt = doc.DeletedAt
	docAPI.ID = doc.ID

	// insertion point for basic fields copy operations
	docAPI.Name = doc.Name
	docAPI.StackName = doc.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDocAPIToDoc update basic, pointers and slice of pointers fields of doc
// from respectively the basic fields and encoded fields of pointers and slices of pointers of docAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDocAPIToDoc(docAPI: DocAPI, doc: Doc, frontRepo: FrontRepo) {

	doc.CreatedAt = docAPI.CreatedAt
	doc.DeletedAt = docAPI.DeletedAt
	doc.ID = docAPI.ID

	// insertion point for basic fields copy operations
	doc.Name = docAPI.Name
	doc.StackName = docAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
