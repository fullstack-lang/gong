// generated code - do not edit

import { OptionDB } from './option-db'
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

export function CopyOptionToOptionDB(option: Option, optionDB: OptionDB) {

	optionDB.CreatedAt = option.CreatedAt
	optionDB.DeletedAt = option.DeletedAt
	optionDB.ID = option.ID
	
	// insertion point for basic fields copy operations
	optionDB.Name = option.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyOptionDBToOption(optionDB: OptionDB, option: Option, frontRepo: FrontRepo) {

	option.CreatedAt = optionDB.CreatedAt
	option.DeletedAt = optionDB.DeletedAt
	option.ID = optionDB.ID
	
	// insertion point for basic fields copy operations
	option.Name = optionDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
