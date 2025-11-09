// generated code - do not edit

import { ConditionAPI } from './condition-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Condition {

	static GONGSTRUCT_NAME = "Condition"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyConditionToConditionAPI(condition: Condition, conditionAPI: ConditionAPI) {

	conditionAPI.CreatedAt = condition.CreatedAt
	conditionAPI.DeletedAt = condition.DeletedAt
	conditionAPI.ID = condition.ID

	// insertion point for basic fields copy operations
	conditionAPI.Name = condition.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyConditionAPIToCondition update basic, pointers and slice of pointers fields of condition
// from respectively the basic fields and encoded fields of pointers and slices of pointers of conditionAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyConditionAPIToCondition(conditionAPI: ConditionAPI, condition: Condition, frontRepo: FrontRepo) {

	condition.CreatedAt = conditionAPI.CreatedAt
	condition.DeletedAt = conditionAPI.DeletedAt
	condition.ID = conditionAPI.ID

	// insertion point for basic fields copy operations
	condition.Name = conditionAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
