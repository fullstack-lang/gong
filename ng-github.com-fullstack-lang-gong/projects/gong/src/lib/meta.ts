// generated code - do not edit

import { MetaAPI } from './meta-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { MetaReference } from './metareference'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Meta {

	static GONGSTRUCT_NAME = "Meta"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""

	// insertion point for pointers and slices of pointers declarations
	MetaReferences: Array<MetaReference> = []
}

export function CopyMetaToMetaAPI(meta: Meta, metaAPI: MetaAPI) {

	metaAPI.CreatedAt = meta.CreatedAt
	metaAPI.DeletedAt = meta.DeletedAt
	metaAPI.ID = meta.ID

	// insertion point for basic fields copy operations
	metaAPI.Name = meta.Name
	metaAPI.Text = meta.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	metaAPI.MetaPointersEncoding.MetaReferences = []
	for (let _metareference of meta.MetaReferences) {
		metaAPI.MetaPointersEncoding.MetaReferences.push(_metareference.ID)
	}

}

// CopyMetaAPIToMeta update basic, pointers and slice of pointers fields of meta
// from respectively the basic fields and encoded fields of pointers and slices of pointers of metaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMetaAPIToMeta(metaAPI: MetaAPI, meta: Meta, frontRepo: FrontRepo) {

	meta.CreatedAt = metaAPI.CreatedAt
	meta.DeletedAt = metaAPI.DeletedAt
	meta.ID = metaAPI.ID

	// insertion point for basic fields copy operations
	meta.Name = metaAPI.Name
	meta.Text = metaAPI.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(metaAPI.MetaPointersEncoding.MetaReferences)) {
		console.error('Rects is not an array:', metaAPI.MetaPointersEncoding.MetaReferences);
		return;
	}

	meta.MetaReferences = new Array<MetaReference>()
	for (let _id of metaAPI.MetaPointersEncoding.MetaReferences) {
		let _metareference = frontRepo.map_ID_MetaReference.get(_id)
		if (_metareference != undefined) {
			meta.MetaReferences.push(_metareference!)
		}
	}
}
