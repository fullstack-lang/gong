// generated code - do not edit

import { EventAPI } from './event-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Event {

	static GONGSTRUCT_NAME = "Event"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0

	// insertion point for pointers and slices of pointers declarations
	Duration_string?: string
}

export function CopyEventToEventAPI(event: Event, eventAPI: EventAPI) {

	eventAPI.CreatedAt = event.CreatedAt
	eventAPI.DeletedAt = event.DeletedAt
	eventAPI.ID = event.ID

	// insertion point for basic fields copy operations
	eventAPI.Name = event.Name
	eventAPI.Duration = event.Duration

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyEventAPIToEvent update basic, pointers and slice of pointers fields of event
// from respectively the basic fields and encoded fields of pointers and slices of pointers of eventAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyEventAPIToEvent(eventAPI: EventAPI, event: Event, frontRepo: FrontRepo) {

	event.CreatedAt = eventAPI.CreatedAt
	event.DeletedAt = eventAPI.DeletedAt
	event.ID = eventAPI.ID

	// insertion point for basic fields copy operations
	event.Name = eventAPI.Name
	event.Duration = eventAPI.Duration

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
