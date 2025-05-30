// generated code - do not edit

import { MessageAPI } from './message-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Message {

	static GONGSTRUCT_NAME = "Message"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyMessageToMessageAPI(message: Message, messageAPI: MessageAPI) {

	messageAPI.CreatedAt = message.CreatedAt
	messageAPI.DeletedAt = message.DeletedAt
	messageAPI.ID = message.ID

	// insertion point for basic fields copy operations
	messageAPI.Name = message.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMessageAPIToMessage update basic, pointers and slice of pointers fields of message
// from respectively the basic fields and encoded fields of pointers and slices of pointers of messageAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMessageAPIToMessage(messageAPI: MessageAPI, message: Message, frontRepo: FrontRepo) {

	message.CreatedAt = messageAPI.CreatedAt
	message.DeletedAt = messageAPI.DeletedAt
	message.ID = messageAPI.ID

	// insertion point for basic fields copy operations
	message.Name = messageAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
