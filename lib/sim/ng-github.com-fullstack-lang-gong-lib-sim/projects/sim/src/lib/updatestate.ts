// generated code - do not edit

import { UpdateStateAPI } from './updatestate-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UpdateState {

	static GONGSTRUCT_NAME = "UpdateState"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0
	Period: number = 0

	// insertion point for pointers and slices of pointers declarations
	Duration_string?: string
	Period_string?: string
}

export function CopyUpdateStateToUpdateStateAPI(updatestate: UpdateState, updatestateAPI: UpdateStateAPI) {

	updatestateAPI.CreatedAt = updatestate.CreatedAt
	updatestateAPI.DeletedAt = updatestate.DeletedAt
	updatestateAPI.ID = updatestate.ID

	// insertion point for basic fields copy operations
	updatestateAPI.Name = updatestate.Name
	updatestateAPI.Duration = updatestate.Duration
	updatestateAPI.Period = updatestate.Period

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyUpdateStateAPIToUpdateState update basic, pointers and slice of pointers fields of updatestate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of updatestateAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyUpdateStateAPIToUpdateState(updatestateAPI: UpdateStateAPI, updatestate: UpdateState, frontRepo: FrontRepo) {

	updatestate.CreatedAt = updatestateAPI.CreatedAt
	updatestate.DeletedAt = updatestateAPI.DeletedAt
	updatestate.ID = updatestateAPI.ID

	// insertion point for basic fields copy operations
	updatestate.Name = updatestateAPI.Name
	updatestate.Duration = updatestateAPI.Duration
	updatestate.Period = updatestateAPI.Period

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
