// generated code - do not edit

import { GongEnumAPI } from './gongenum-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { GongEnumValue } from './gongenumvalue'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnum {

	static GONGSTRUCT_NAME = "GongEnum"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: number = 0

	// insertion point for pointers and slices of pointers declarations
	Type_string?: string
	GongEnumValues: Array<GongEnumValue> = []
}

export function CopyGongEnumToGongEnumAPI(gongenum: GongEnum, gongenumAPI: GongEnumAPI) {

	gongenumAPI.CreatedAt = gongenum.CreatedAt
	gongenumAPI.DeletedAt = gongenum.DeletedAt
	gongenumAPI.ID = gongenum.ID

	// insertion point for basic fields copy operations
	gongenumAPI.Name = gongenum.Name
	gongenumAPI.Type = gongenum.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	gongenumAPI.GongEnumPointersEncoding.GongEnumValues = []
	for (let _gongenumvalue of gongenum.GongEnumValues) {
		gongenumAPI.GongEnumPointersEncoding.GongEnumValues.push(_gongenumvalue.ID)
	}

}

// CopyGongEnumAPIToGongEnum update basic, pointers and slice of pointers fields of gongenum
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongenumAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongEnumAPIToGongEnum(gongenumAPI: GongEnumAPI, gongenum: GongEnum, frontRepo: FrontRepo) {

	gongenum.CreatedAt = gongenumAPI.CreatedAt
	gongenum.DeletedAt = gongenumAPI.DeletedAt
	gongenum.ID = gongenumAPI.ID

	// insertion point for basic fields copy operations
	gongenum.Name = gongenumAPI.Name
	gongenum.Type = gongenumAPI.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(gongenumAPI.GongEnumPointersEncoding.GongEnumValues)) {
		console.error('Rects is not an array:', gongenumAPI.GongEnumPointersEncoding.GongEnumValues);
		return;
	}

	gongenum.GongEnumValues = new Array<GongEnumValue>()
	for (let _id of gongenumAPI.GongEnumPointersEncoding.GongEnumValues) {
		let _gongenumvalue = frontRepo.map_ID_GongEnumValue.get(_id)
		if (_gongenumvalue != undefined) {
			gongenum.GongEnumValues.push(_gongenumvalue!)
		}
	}
}
