// generated code - do not edit

import { DummyAgentAPI } from './dummyagent-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DummyAgent {

	static GONGSTRUCT_NAME = "DummyAgent"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDummyAgentToDummyAgentAPI(dummyagent: DummyAgent, dummyagentAPI: DummyAgentAPI) {

	dummyagentAPI.CreatedAt = dummyagent.CreatedAt
	dummyagentAPI.DeletedAt = dummyagent.DeletedAt
	dummyagentAPI.ID = dummyagent.ID

	// insertion point for basic fields copy operations
	dummyagentAPI.TechName = dummyagent.TechName
	dummyagentAPI.Name = dummyagent.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDummyAgentAPIToDummyAgent update basic, pointers and slice of pointers fields of dummyagent
// from respectively the basic fields and encoded fields of pointers and slices of pointers of dummyagentAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDummyAgentAPIToDummyAgent(dummyagentAPI: DummyAgentAPI, dummyagent: DummyAgent, frontRepo: FrontRepo) {

	dummyagent.CreatedAt = dummyagentAPI.CreatedAt
	dummyagent.DeletedAt = dummyagentAPI.DeletedAt
	dummyagent.ID = dummyagentAPI.ID

	// insertion point for basic fields copy operations
	dummyagent.TechName = dummyagentAPI.TechName
	dummyagent.Name = dummyagentAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
