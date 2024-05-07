// generated code - do not edit

import { OptionAPI } from './option-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Option {

	static GONGSTRUCT_NAME = "Option"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyOptionToOptionAPI(option: Option, optionAPI: OptionAPI) {

	optionAPI.CreatedAt = option.CreatedAt
	optionAPI.DeletedAt = option.DeletedAt
	optionAPI.ID = option.ID

	// insertion point for basic fields copy operations
	optionAPI.Name = option.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyOptionAPIToOption update basic, pointers and slice of pointers fields of option
// from respectively the basic fields and encoded fields of pointers and slices of pointers of optionAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyOptionAPIToOption(optionAPI: OptionAPI, option: Option, frontRepo: FrontRepo) {

	option.CreatedAt = optionAPI.CreatedAt
	option.DeletedAt = optionAPI.DeletedAt
	option.ID = optionAPI.ID

	// insertion point for basic fields copy operations
	option.Name = optionAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
