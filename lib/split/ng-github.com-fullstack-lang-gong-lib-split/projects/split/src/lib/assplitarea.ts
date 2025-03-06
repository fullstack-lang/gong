// generated code - do not edit

import { AsSplitAreaAPI } from './assplitarea-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { AsSplit } from './assplit'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplitArea {

	static GONGSTRUCT_NAME = "AsSplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Size: number = 0
	IsAny: boolean = false

	// insertion point for pointers and slices of pointers declarations
	AsSplits: Array<AsSplit> = []
}

export function CopyAsSplitAreaToAsSplitAreaAPI(assplitarea: AsSplitArea, assplitareaAPI: AsSplitAreaAPI) {

	assplitareaAPI.CreatedAt = assplitarea.CreatedAt
	assplitareaAPI.DeletedAt = assplitarea.DeletedAt
	assplitareaAPI.ID = assplitarea.ID

	// insertion point for basic fields copy operations
	assplitareaAPI.Name = assplitarea.Name
	assplitareaAPI.Size = assplitarea.Size
	assplitareaAPI.IsAny = assplitarea.IsAny

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	assplitareaAPI.AsSplitAreaPointersEncoding.AsSplits = []
	for (let _assplit of assplitarea.AsSplits) {
		assplitareaAPI.AsSplitAreaPointersEncoding.AsSplits.push(_assplit.ID)
	}

}

// CopyAsSplitAreaAPIToAsSplitArea update basic, pointers and slice of pointers fields of assplitarea
// from respectively the basic fields and encoded fields of pointers and slices of pointers of assplitareaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI: AsSplitAreaAPI, assplitarea: AsSplitArea, frontRepo: FrontRepo) {

	assplitarea.CreatedAt = assplitareaAPI.CreatedAt
	assplitarea.DeletedAt = assplitareaAPI.DeletedAt
	assplitarea.ID = assplitareaAPI.ID

	// insertion point for basic fields copy operations
	assplitarea.Name = assplitareaAPI.Name
	assplitarea.Size = assplitareaAPI.Size
	assplitarea.IsAny = assplitareaAPI.IsAny

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(assplitareaAPI.AsSplitAreaPointersEncoding.AsSplits)) {
		console.error('Rects is not an array:', assplitareaAPI.AsSplitAreaPointersEncoding.AsSplits);
		return;
	}

	assplitarea.AsSplits = new Array<AsSplit>()
	for (let _id of assplitareaAPI.AsSplitAreaPointersEncoding.AsSplits) {
		let _assplit = frontRepo.map_ID_AsSplit.get(_id)
		if (_assplit != undefined) {
			assplitarea.AsSplits.push(_assplit!)
		}
	}
}
