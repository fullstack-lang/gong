// insertion point for imports
import { NoteShapeDB } from './noteshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteShapeLinkDB {

	static GONGSTRUCT_NAME = "NoteShapeLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	Type: string = ""

	// insertion point for other declarations
	NoteShape_NoteShapeLinksDBID: NullInt64 = new NullInt64
	NoteShape_NoteShapeLinksDBID_Index: NullInt64  = new NullInt64 // store the index of the noteshapelink instance in NoteShape.NoteShapeLinks
	NoteShape_NoteShapeLinks_reverse?: NoteShapeDB 

}
