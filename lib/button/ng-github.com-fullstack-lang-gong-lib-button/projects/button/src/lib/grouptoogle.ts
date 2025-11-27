// generated code - do not edit

import { GroupToogleAPI } from './grouptoogle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { ButtonToggle } from './buttontoggle'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GroupToogle {

	static GONGSTRUCT_NAME = "GroupToogle"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Percentage: number = 0
	IsSingleSelector: boolean = false

	// insertion point for pointers and slices of pointers declarations
	ButtonToggles: Array<ButtonToggle> = []

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyGroupToogleToGroupToogleAPI(grouptoogle: GroupToogle, grouptoogleAPI: GroupToogleAPI) {

	grouptoogleAPI.CreatedAt = grouptoogle.CreatedAt
	grouptoogleAPI.DeletedAt = grouptoogle.DeletedAt
	grouptoogleAPI.ID = grouptoogle.ID

	// insertion point for basic fields copy operations
	grouptoogleAPI.Name = grouptoogle.Name
	grouptoogleAPI.Percentage = grouptoogle.Percentage
	grouptoogleAPI.IsSingleSelector = grouptoogle.IsSingleSelector

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	grouptoogleAPI.GroupTooglePointersEncoding.ButtonToggles = []
	for (let _buttontoggle of grouptoogle.ButtonToggles) {
		grouptoogleAPI.GroupTooglePointersEncoding.ButtonToggles.push(_buttontoggle.ID)
	}

}

// CopyGroupToogleAPIToGroupToogle update basic, pointers and slice of pointers fields of grouptoogle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of grouptoogleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGroupToogleAPIToGroupToogle(grouptoogleAPI: GroupToogleAPI, grouptoogle: GroupToogle, frontRepo: FrontRepo) {

	grouptoogle.CreatedAt = grouptoogleAPI.CreatedAt
	grouptoogle.DeletedAt = grouptoogleAPI.DeletedAt
	grouptoogle.ID = grouptoogleAPI.ID

	// insertion point for basic fields copy operations
	grouptoogle.Name = grouptoogleAPI.Name
	grouptoogle.Percentage = grouptoogleAPI.Percentage
	grouptoogle.IsSingleSelector = grouptoogleAPI.IsSingleSelector

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(grouptoogleAPI.GroupTooglePointersEncoding.ButtonToggles)) {
		console.error('Rects is not an array:', grouptoogleAPI.GroupTooglePointersEncoding.ButtonToggles);
		return;
	}

	grouptoogle.ButtonToggles = new Array<ButtonToggle>()
	for (let _id of grouptoogleAPI.GroupTooglePointersEncoding.ButtonToggles) {
		let _buttontoggle = frontRepo.map_ID_ButtonToggle.get(_id)
		if (_buttontoggle != undefined) {
			grouptoogle.ButtonToggles.push(_buttontoggle!)
		}
	}
}
