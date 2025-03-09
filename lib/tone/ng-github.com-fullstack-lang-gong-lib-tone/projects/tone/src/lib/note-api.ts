// insertion point for imports
import { FreqencyAPI } from './freqency-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteAPI {

	static GONGSTRUCT_NAME = "Note"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Start: number = 0
	Duration: number = 0
	Velocity: number = 0
	Info: string = ""

	// insertion point for other decls

	NotePointersEncoding: NotePointersEncoding = new NotePointersEncoding
}

export class NotePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Frequencies: number[] = []
}
