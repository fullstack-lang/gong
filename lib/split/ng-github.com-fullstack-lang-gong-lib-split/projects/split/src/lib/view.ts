// generated code - do not edit

import { ViewAPI } from './view-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { AsSplitArea } from './assplitarea'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class View {

	static GONGSTRUCT_NAME = "View"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ShowViewName: boolean = false

	// insertion point for pointers and slices of pointers declarations
	RootAsSplitAreas: Array<AsSplitArea> = []
}

export function CopyViewToViewAPI(view: View, viewAPI: ViewAPI) {

	viewAPI.CreatedAt = view.CreatedAt
	viewAPI.DeletedAt = view.DeletedAt
	viewAPI.ID = view.ID

	// insertion point for basic fields copy operations
	viewAPI.Name = view.Name
	viewAPI.ShowViewName = view.ShowViewName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	viewAPI.ViewPointersEncoding.RootAsSplitAreas = []
	for (let _assplitarea of view.RootAsSplitAreas) {
		viewAPI.ViewPointersEncoding.RootAsSplitAreas.push(_assplitarea.ID)
	}

}

// CopyViewAPIToView update basic, pointers and slice of pointers fields of view
// from respectively the basic fields and encoded fields of pointers and slices of pointers of viewAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyViewAPIToView(viewAPI: ViewAPI, view: View, frontRepo: FrontRepo) {

	view.CreatedAt = viewAPI.CreatedAt
	view.DeletedAt = viewAPI.DeletedAt
	view.ID = viewAPI.ID

	// insertion point for basic fields copy operations
	view.Name = viewAPI.Name
	view.ShowViewName = viewAPI.ShowViewName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(viewAPI.ViewPointersEncoding.RootAsSplitAreas)) {
		console.error('Rects is not an array:', viewAPI.ViewPointersEncoding.RootAsSplitAreas);
		return;
	}

	view.RootAsSplitAreas = new Array<AsSplitArea>()
	for (let _id of viewAPI.ViewPointersEncoding.RootAsSplitAreas) {
		let _assplitarea = frontRepo.map_ID_AsSplitArea.get(_id)
		if (_assplitarea != undefined) {
			view.RootAsSplitAreas.push(_assplitarea!)
		}
	}
}
