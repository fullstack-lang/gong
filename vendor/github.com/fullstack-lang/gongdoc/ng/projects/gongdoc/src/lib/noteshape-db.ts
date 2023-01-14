// insertion point for imports
import { NoteLinkDB } from './notelink-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteShapeDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Body: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Heigth: number = 0
	Matched: boolean = false

	// insertion point for other declarations
	NoteLinks?: Array<NoteLinkDB>
	Classdiagram_NotesDBID: NullInt64 = new NullInt64
	Classdiagram_NotesDBID_Index: NullInt64  = new NullInt64 // store the index of the noteshape instance in Classdiagram.Notes
	Classdiagram_Notes_reverse?: ClassdiagramDB 

}
