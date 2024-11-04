// generated code - do not edit

import { GongNoteAPI } from './gongnote-api'
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

export function CopyGongNoteToGongNoteAPI(gongnote: GongNote, gongnoteAPI: GongNoteAPI) {

	gongnoteAPI.CreatedAt = gongnote.CreatedAt
	gongnoteAPI.DeletedAt = gongnote.DeletedAt
	gongnoteAPI.ID = gongnote.ID

	// insertion point for basic fields copy operations
	gongnoteAPI.Name = gongnote.Name
	gongnoteAPI.Body = gongnote.Body
	gongnoteAPI.BodyHTML = gongnote.BodyHTML

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongnoteAPI.GongNotePointersEncoding.Links = []
	for (let _gonglink of gongnote.Links) {
		gongnoteAPI.GongNotePointersEncoding.Links.push(_gonglink.ID)
	}

}

// CopyGongNoteAPIToGongNote update basic, pointers and slice of pointers fields of gongnote
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongnoteAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongNoteAPIToGongNote(gongnoteAPI: GongNoteAPI, gongnote: GongNote, frontRepo: FrontRepo) {

	gongnote.CreatedAt = gongnoteAPI.CreatedAt
	gongnote.DeletedAt = gongnoteAPI.DeletedAt
	gongnote.ID = gongnoteAPI.ID

	// insertion point for basic fields copy operations
	gongnote.Name = gongnoteAPI.Name
	gongnote.Body = gongnoteAPI.Body
	gongnote.BodyHTML = gongnoteAPI.BodyHTML

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(gongnoteAPI.GongNotePointersEncoding.Links)) {
		console.error('Rects is not an array:', gongnoteAPI.GongNotePointersEncoding.Links);
		return;
	}

	gongnote.Links = new Array<GongLink>()
	for (let _id of gongnoteAPI.GongNotePointersEncoding.Links) {
		let _gonglink = frontRepo.map_ID_GongLink.get(_id)
		if (_gonglink != undefined) {
			gongnote.Links.push(_gonglink!)
		}
	}
}
