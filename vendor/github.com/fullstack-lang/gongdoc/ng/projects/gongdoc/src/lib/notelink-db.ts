// insertion point for imports
import { ClassshapeDB } from './classshape-db'
import { LinkDB } from './link-db'
import { VerticeDB } from './vertice-db'
import { NoteShapeDB } from './noteshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteLinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: string = ""

	// insertion point for other declarations
	Classshape?: ClassshapeDB
	ClassshapeID: NullInt64 = new NullInt64 // if pointer is null, Classshape.ID = 0

	Link?: LinkDB
	LinkID: NullInt64 = new NullInt64 // if pointer is null, Link.ID = 0

	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	NoteShape_NoteLinksDBID: NullInt64 = new NullInt64
	NoteShape_NoteLinksDBID_Index: NullInt64  = new NullInt64 // store the index of the notelink instance in NoteShape.NoteLinks
	NoteShape_NoteLinks_reverse?: NoteShapeDB 

}
