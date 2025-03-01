// generated code - do not edit

import { GroupAPI } from './group-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Button } from './button'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Group {

	static GONGSTRUCT_NAME = "Group"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Percentage: number = 0

	// insertion point for pointers and slices of pointers declarations
	Buttons: Array<Button> = []
}

export function CopyGroupToGroupAPI(group: Group, groupAPI: GroupAPI) {

	groupAPI.CreatedAt = group.CreatedAt
	groupAPI.DeletedAt = group.DeletedAt
	groupAPI.ID = group.ID

	// insertion point for basic fields copy operations
	groupAPI.Name = group.Name
	groupAPI.Percentage = group.Percentage

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	groupAPI.GroupPointersEncoding.Buttons = []
	for (let _button of group.Buttons) {
		groupAPI.GroupPointersEncoding.Buttons.push(_button.ID)
	}

}

// CopyGroupAPIToGroup update basic, pointers and slice of pointers fields of group
// from respectively the basic fields and encoded fields of pointers and slices of pointers of groupAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGroupAPIToGroup(groupAPI: GroupAPI, group: Group, frontRepo: FrontRepo) {

	group.CreatedAt = groupAPI.CreatedAt
	group.DeletedAt = groupAPI.DeletedAt
	group.ID = groupAPI.ID

	// insertion point for basic fields copy operations
	group.Name = groupAPI.Name
	group.Percentage = groupAPI.Percentage

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(groupAPI.GroupPointersEncoding.Buttons)) {
		console.error('Rects is not an array:', groupAPI.GroupPointersEncoding.Buttons);
		return;
	}

	group.Buttons = new Array<Button>()
	for (let _id of groupAPI.GroupPointersEncoding.Buttons) {
		let _button = frontRepo.map_ID_Button.get(_id)
		if (_button != undefined) {
			group.Buttons.push(_button!)
		}
	}
}
