// generated code - do not edit

import { AsSplitAPI } from './assplit-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { AsSplitArea } from './assplitarea'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplit {

	static GONGSTRUCT_NAME = "AsSplit"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Direction: string = ""

	// insertion point for pointers and slices of pointers declarations
	AsSplitAreas: Array<AsSplitArea> = []
}

export function CopyAsSplitToAsSplitAPI(assplit: AsSplit, assplitAPI: AsSplitAPI) {

	assplitAPI.CreatedAt = assplit.CreatedAt
	assplitAPI.DeletedAt = assplit.DeletedAt
	assplitAPI.ID = assplit.ID

	// insertion point for basic fields copy operations
	assplitAPI.Name = assplit.Name
	assplitAPI.Direction = assplit.Direction

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	assplitAPI.AsSplitPointersEncoding.AsSplitAreas = []
	for (let _assplitarea of assplit.AsSplitAreas) {
		assplitAPI.AsSplitPointersEncoding.AsSplitAreas.push(_assplitarea.ID)
	}

}

// CopyAsSplitAPIToAsSplit update basic, pointers and slice of pointers fields of assplit
// from respectively the basic fields and encoded fields of pointers and slices of pointers of assplitAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAsSplitAPIToAsSplit(assplitAPI: AsSplitAPI, assplit: AsSplit, frontRepo: FrontRepo) {

	assplit.CreatedAt = assplitAPI.CreatedAt
	assplit.DeletedAt = assplitAPI.DeletedAt
	assplit.ID = assplitAPI.ID

	// insertion point for basic fields copy operations
	assplit.Name = assplitAPI.Name
	assplit.Direction = assplitAPI.Direction

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(assplitAPI.AsSplitPointersEncoding.AsSplitAreas)) {
		console.error('Rects is not an array:', assplitAPI.AsSplitPointersEncoding.AsSplitAreas);
		return;
	}

	assplit.AsSplitAreas = new Array<AsSplitArea>()
	for (let _id of assplitAPI.AsSplitPointersEncoding.AsSplitAreas) {
		let _assplitarea = frontRepo.map_ID_AsSplitArea.get(_id)
		if (_assplitarea != undefined) {
			assplit.AsSplitAreas.push(_assplitarea!)
		}
	}
}
