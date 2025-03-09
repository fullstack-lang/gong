// generated code - do not edit

import { NoteAPI } from './note-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Freqency } from './freqency'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Note {

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

	// insertion point for pointers and slices of pointers declarations
	Frequencies: Array<Freqency> = []
}

export function CopyNoteToNoteAPI(note: Note, noteAPI: NoteAPI) {

	noteAPI.CreatedAt = note.CreatedAt
	noteAPI.DeletedAt = note.DeletedAt
	noteAPI.ID = note.ID

	// insertion point for basic fields copy operations
	noteAPI.Name = note.Name
	noteAPI.Start = note.Start
	noteAPI.Duration = note.Duration
	noteAPI.Velocity = note.Velocity
	noteAPI.Info = note.Info

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	noteAPI.NotePointersEncoding.Frequencies = []
	for (let _freqency of note.Frequencies) {
		noteAPI.NotePointersEncoding.Frequencies.push(_freqency.ID)
	}

}

// CopyNoteAPIToNote update basic, pointers and slice of pointers fields of note
// from respectively the basic fields and encoded fields of pointers and slices of pointers of noteAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyNoteAPIToNote(noteAPI: NoteAPI, note: Note, frontRepo: FrontRepo) {

	note.CreatedAt = noteAPI.CreatedAt
	note.DeletedAt = noteAPI.DeletedAt
	note.ID = noteAPI.ID

	// insertion point for basic fields copy operations
	note.Name = noteAPI.Name
	note.Start = noteAPI.Start
	note.Duration = noteAPI.Duration
	note.Velocity = noteAPI.Velocity
	note.Info = noteAPI.Info

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(noteAPI.NotePointersEncoding.Frequencies)) {
		console.error('Rects is not an array:', noteAPI.NotePointersEncoding.Frequencies);
		return;
	}

	note.Frequencies = new Array<Freqency>()
	for (let _id of noteAPI.NotePointersEncoding.Frequencies) {
		let _freqency = frontRepo.map_ID_Freqency.get(_id)
		if (_freqency != undefined) {
			note.Frequencies.push(_freqency!)
		}
	}
}
