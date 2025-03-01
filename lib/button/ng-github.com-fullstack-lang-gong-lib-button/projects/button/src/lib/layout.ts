// generated code - do not edit

import { LayoutAPI } from './layout-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Group } from './group'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Layout {

	static GONGSTRUCT_NAME = "Layout"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Groups: Array<Group> = []
}

export function CopyLayoutToLayoutAPI(layout: Layout, layoutAPI: LayoutAPI) {

	layoutAPI.CreatedAt = layout.CreatedAt
	layoutAPI.DeletedAt = layout.DeletedAt
	layoutAPI.ID = layout.ID

	// insertion point for basic fields copy operations
	layoutAPI.Name = layout.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	layoutAPI.LayoutPointersEncoding.Groups = []
	for (let _group of layout.Groups) {
		layoutAPI.LayoutPointersEncoding.Groups.push(_group.ID)
	}

}

// CopyLayoutAPIToLayout update basic, pointers and slice of pointers fields of layout
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layoutAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayoutAPIToLayout(layoutAPI: LayoutAPI, layout: Layout, frontRepo: FrontRepo) {

	layout.CreatedAt = layoutAPI.CreatedAt
	layout.DeletedAt = layoutAPI.DeletedAt
	layout.ID = layoutAPI.ID

	// insertion point for basic fields copy operations
	layout.Name = layoutAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(layoutAPI.LayoutPointersEncoding.Groups)) {
		console.error('Rects is not an array:', layoutAPI.LayoutPointersEncoding.Groups);
		return;
	}

	layout.Groups = new Array<Group>()
	for (let _id of layoutAPI.LayoutPointersEncoding.Groups) {
		let _group = frontRepo.map_ID_Group.get(_id)
		if (_group != undefined) {
			layout.Groups.push(_group!)
		}
	}
}
