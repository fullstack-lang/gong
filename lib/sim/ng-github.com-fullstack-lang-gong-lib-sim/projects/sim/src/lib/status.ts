// generated code - do not edit

import { StatusAPI } from './status-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Status {

	static GONGSTRUCT_NAME = "Status"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CurrentCommand: string = ""
	CompletionDate: string = ""
	CurrentSpeedCommand: string = ""
	SpeedCommandCompletionDate: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyStatusToStatusAPI(status: Status, statusAPI: StatusAPI) {

	statusAPI.CreatedAt = status.CreatedAt
	statusAPI.DeletedAt = status.DeletedAt
	statusAPI.ID = status.ID

	// insertion point for basic fields copy operations
	statusAPI.Name = status.Name
	statusAPI.CurrentCommand = status.CurrentCommand
	statusAPI.CompletionDate = status.CompletionDate
	statusAPI.CurrentSpeedCommand = status.CurrentSpeedCommand
	statusAPI.SpeedCommandCompletionDate = status.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyStatusAPIToStatus update basic, pointers and slice of pointers fields of status
// from respectively the basic fields and encoded fields of pointers and slices of pointers of statusAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyStatusAPIToStatus(statusAPI: StatusAPI, status: Status, frontRepo: FrontRepo) {

	status.CreatedAt = statusAPI.CreatedAt
	status.DeletedAt = statusAPI.DeletedAt
	status.ID = statusAPI.ID

	// insertion point for basic fields copy operations
	status.Name = statusAPI.Name
	status.CurrentCommand = statusAPI.CurrentCommand
	status.CompletionDate = statusAPI.CompletionDate
	status.CurrentSpeedCommand = statusAPI.CurrentSpeedCommand
	status.SpeedCommandCompletionDate = statusAPI.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
