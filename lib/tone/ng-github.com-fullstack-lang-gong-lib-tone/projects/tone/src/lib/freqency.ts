// generated code - do not edit

import { FreqencyAPI } from './freqency-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Freqency {

	static GONGSTRUCT_NAME = "Freqency"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFreqencyToFreqencyAPI(freqency: Freqency, freqencyAPI: FreqencyAPI) {

	freqencyAPI.CreatedAt = freqency.CreatedAt
	freqencyAPI.DeletedAt = freqency.DeletedAt
	freqencyAPI.ID = freqency.ID

	// insertion point for basic fields copy operations
	freqencyAPI.Name = freqency.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFreqencyAPIToFreqency update basic, pointers and slice of pointers fields of freqency
// from respectively the basic fields and encoded fields of pointers and slices of pointers of freqencyAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFreqencyAPIToFreqency(freqencyAPI: FreqencyAPI, freqency: Freqency, frontRepo: FrontRepo) {

	freqency.CreatedAt = freqencyAPI.CreatedAt
	freqency.DeletedAt = freqencyAPI.DeletedAt
	freqency.ID = freqencyAPI.ID

	// insertion point for basic fields copy operations
	freqency.Name = freqencyAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
