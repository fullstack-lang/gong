// generated code - do not edit

import { ToneAPI } from './tone-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Tone {

	static GONGSTRUCT_NAME = "Tone"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyToneToToneAPI(tone: Tone, toneAPI: ToneAPI) {

	toneAPI.CreatedAt = tone.CreatedAt
	toneAPI.DeletedAt = tone.DeletedAt
	toneAPI.ID = tone.ID

	// insertion point for basic fields copy operations
	toneAPI.Name = tone.Name
	toneAPI.StackName = tone.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyToneAPIToTone update basic, pointers and slice of pointers fields of tone
// from respectively the basic fields and encoded fields of pointers and slices of pointers of toneAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyToneAPIToTone(toneAPI: ToneAPI, tone: Tone, frontRepo: FrontRepo) {

	tone.CreatedAt = toneAPI.CreatedAt
	tone.DeletedAt = toneAPI.DeletedAt
	tone.ID = toneAPI.ID

	// insertion point for basic fields copy operations
	tone.Name = toneAPI.Name
	tone.StackName = toneAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
