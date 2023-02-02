// insertion point for imports
import { GongStructShapeDB } from './gongstructshape-db'
import { LinkDB } from './link-db'
import { VerticeDB } from './vertice-db'
import { NoteShapeDB } from './noteshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteShapeLinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""

	// insertion point for other declarations
	Classshape?: GongStructShapeDB
	ClassshapeID: NullInt64 = new NullInt64 // if pointer is null, Classshape.ID = 0

	Link?: LinkDB
	LinkID: NullInt64 = new NullInt64 // if pointer is null, Link.ID = 0

	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	NoteShape_NoteShapeLinksDBID: NullInt64 = new NullInt64
	NoteShape_NoteShapeLinksDBID_Index: NullInt64  = new NullInt64 // store the index of the noteshapelink instance in NoteShape.NoteShapeLinks
	NoteShape_NoteShapeLinks_reverse?: NoteShapeDB 

}
