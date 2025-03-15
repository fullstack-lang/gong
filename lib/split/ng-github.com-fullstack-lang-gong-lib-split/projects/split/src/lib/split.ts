// generated code - do not edit

import { SplitAPI } from './split-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Split {

	static GONGSTRUCT_NAME = "Split"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySplitToSplitAPI(split: Split, splitAPI: SplitAPI) {

	splitAPI.CreatedAt = split.CreatedAt
	splitAPI.DeletedAt = split.DeletedAt
	splitAPI.ID = split.ID

	// insertion point for basic fields copy operations
	splitAPI.Name = split.Name
	splitAPI.StackName = split.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySplitAPIToSplit update basic, pointers and slice of pointers fields of split
// from respectively the basic fields and encoded fields of pointers and slices of pointers of splitAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySplitAPIToSplit(splitAPI: SplitAPI, split: Split, frontRepo: FrontRepo) {

	split.CreatedAt = splitAPI.CreatedAt
	split.DeletedAt = splitAPI.DeletedAt
	split.ID = splitAPI.ID

	// insertion point for basic fields copy operations
	split.Name = splitAPI.Name
	split.StackName = splitAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
