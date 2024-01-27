// generated code - do not edit

import { MetaDB } from './meta-db'
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

export function CopyMetaToMetaDB(meta: Meta, metaDB: MetaDB) {

	metaDB.CreatedAt = meta.CreatedAt
	metaDB.DeletedAt = meta.DeletedAt
	metaDB.ID = meta.ID

	// insertion point for basic fields copy operations
	metaDB.Name = meta.Name
	metaDB.Text = meta.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	metaDB.MetaPointersEncoding.MetaReferences = []
	for (let _metareference of meta.MetaReferences) {
		metaDB.MetaPointersEncoding.MetaReferences.push(_metareference.ID)
	}

}

// CopyMetaDBToMeta update basic, pointers and slice of pointers fields of meta
// from respectively the basic fields and encoded fields of pointers and slices of pointers of metaDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMetaDBToMeta(metaDB: MetaDB, meta: Meta, frontRepo: FrontRepo) {

	meta.CreatedAt = metaDB.CreatedAt
	meta.DeletedAt = metaDB.DeletedAt
	meta.ID = metaDB.ID

	// insertion point for basic fields copy operations
	meta.Name = metaDB.Name
	meta.Text = metaDB.Text

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	meta.MetaReferences = new Array<MetaReference>()
	for (let _id of metaDB.MetaPointersEncoding.MetaReferences) {
		let _metareference = frontRepo.map_ID_MetaReference.get(_id)
		if (_metareference != undefined) {
			meta.MetaReferences.push(_metareference!)
		}
	}
}
