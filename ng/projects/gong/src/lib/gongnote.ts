// generated code - do not edit

import { GongNoteDB } from './gongnote-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongLink } from './gonglink'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongNote {

	static GONGSTRUCT_NAME = "GongNote"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Body: string = ""
	BodyHTML: string = ""

	// insertion point for pointers and slices of pointers declarations
	Links: Array<GongLink> = []
}

export function CopyGongNoteToGongNoteDB(gongnote: GongNote, gongnoteDB: GongNoteDB) {

	gongnoteDB.CreatedAt = gongnote.CreatedAt
	gongnoteDB.DeletedAt = gongnote.DeletedAt
	gongnoteDB.ID = gongnote.ID
	
	// insertion point for basic fields copy operations
	gongnoteDB.Name = gongnote.Name
	gongnoteDB.Body = gongnote.Body
	gongnoteDB.BodyHTML = gongnote.BodyHTML

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongnoteDB.GongNotePointersEncoding.Links = []
    for (let _gonglink of gongnote.Links) {
		gongnoteDB.GongNotePointersEncoding.Links.push(_gonglink.ID)
    }
	
}

export function CopyGongNoteDBToGongNote(gongnoteDB: GongNoteDB, gongnote: GongNote, frontRepo: FrontRepo) {

	gongnote.CreatedAt = gongnoteDB.CreatedAt
	gongnote.DeletedAt = gongnoteDB.DeletedAt
	gongnote.ID = gongnoteDB.ID
	
	// insertion point for basic fields copy operations
	gongnote.Name = gongnoteDB.Name
	gongnote.Body = gongnoteDB.Body
	gongnote.BodyHTML = gongnoteDB.BodyHTML

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongnote.Links = new Array<GongLink>()
	for (let _id of gongnoteDB.GongNotePointersEncoding.Links) {
	  let _gonglink = frontRepo.GongLinks.get(_id)
	  if (_gonglink != undefined) {
		gongnote.Links.push(_gonglink!)
	  }
	}
}
