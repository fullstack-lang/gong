// generated code - do not edit

import { GongLinkAPI } from './gonglink-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongLink {

	static GONGSTRUCT_NAME = "GongLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Recv: string = ""
	ImportPath: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongLinkToGongLinkAPI(gonglink: GongLink, gonglinkAPI: GongLinkAPI) {

	gonglinkAPI.CreatedAt = gonglink.CreatedAt
	gonglinkAPI.DeletedAt = gonglink.DeletedAt
	gonglinkAPI.ID = gonglink.ID

	// insertion point for basic fields copy operations
	gonglinkAPI.Name = gonglink.Name
	gonglinkAPI.Recv = gonglink.Recv
	gonglinkAPI.ImportPath = gonglink.ImportPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongLinkAPIToGongLink update basic, pointers and slice of pointers fields of gonglink
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gonglinkAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongLinkAPIToGongLink(gonglinkAPI: GongLinkAPI, gonglink: GongLink, frontRepo: FrontRepo) {

	gonglink.CreatedAt = gonglinkAPI.CreatedAt
	gonglink.DeletedAt = gonglinkAPI.DeletedAt
	gonglink.ID = gonglinkAPI.ID

	// insertion point for basic fields copy operations
	gonglink.Name = gonglinkAPI.Name
	gonglink.Recv = gonglinkAPI.Recv
	gonglink.ImportPath = gonglinkAPI.ImportPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
