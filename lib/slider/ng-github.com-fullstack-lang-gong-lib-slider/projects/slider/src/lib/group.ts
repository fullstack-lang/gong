// generated code - do not edit

import { GroupAPI } from './group-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Slider } from './slider'
import { Checkbox } from './checkbox'

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
	Sliders: Array<Slider> = []
	Checkboxes: Array<Checkbox> = []
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
	groupAPI.GroupPointersEncoding.Sliders = []
	for (let _slider of group.Sliders) {
		groupAPI.GroupPointersEncoding.Sliders.push(_slider.ID)
	}

	groupAPI.GroupPointersEncoding.Checkboxes = []
	for (let _checkbox of group.Checkboxes) {
		groupAPI.GroupPointersEncoding.Checkboxes.push(_checkbox.ID)
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
	if (!Array.isArray(groupAPI.GroupPointersEncoding.Sliders)) {
		console.error('Rects is not an array:', groupAPI.GroupPointersEncoding.Sliders);
		return;
	}

	group.Sliders = new Array<Slider>()
	for (let _id of groupAPI.GroupPointersEncoding.Sliders) {
		let _slider = frontRepo.map_ID_Slider.get(_id)
		if (_slider != undefined) {
			group.Sliders.push(_slider!)
		}
	}
	if (!Array.isArray(groupAPI.GroupPointersEncoding.Checkboxes)) {
		console.error('Rects is not an array:', groupAPI.GroupPointersEncoding.Checkboxes);
		return;
	}

	group.Checkboxes = new Array<Checkbox>()
	for (let _id of groupAPI.GroupPointersEncoding.Checkboxes) {
		let _checkbox = frontRepo.map_ID_Checkbox.get(_id)
		if (_checkbox != undefined) {
			group.Checkboxes.push(_checkbox!)
		}
	}
}
