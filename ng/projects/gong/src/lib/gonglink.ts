// generated code - do not edit

import { GongLinkDB } from './gonglink-db'
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

export function CopyGongLinkToGongLinkDB(gonglink: GongLink, gonglinkDB: GongLinkDB) {

	gonglinkDB.CreatedAt = gonglink.CreatedAt
	gonglinkDB.DeletedAt = gonglink.DeletedAt
	gonglinkDB.ID = gonglink.ID

	// insertion point for basic fields copy operations
	gonglinkDB.Name = gonglink.Name
	gonglinkDB.Recv = gonglink.Recv
	gonglinkDB.ImportPath = gonglink.ImportPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongLinkDBToGongLink update basic, pointers and slice of pointers fields of gonglink
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gonglinkDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongLinkDBToGongLink(gonglinkDB: GongLinkDB, gonglink: GongLink, frontRepo: FrontRepo) {

	gonglink.CreatedAt = gonglinkDB.CreatedAt
	gonglink.DeletedAt = gonglinkDB.DeletedAt
	gonglink.ID = gonglinkDB.ID

	// insertion point for basic fields copy operations
	gonglink.Name = gonglinkDB.Name
	gonglink.Recv = gonglinkDB.Recv
	gonglink.ImportPath = gonglinkDB.ImportPath

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
