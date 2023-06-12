// insertion point for imports
import { NoteShapeLinkDB } from './noteshapelink-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteShapeDB {

	static GONGSTRUCT_NAME = "NoteShape"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	Body: string = ""
	BodyHTML: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Heigth: number = 0
	Matched: boolean = false

	// insertion point for other declarations
	NoteShapeLinks?: Array<NoteShapeLinkDB>
	Classdiagram_NoteShapesDBID: NullInt64 = new NullInt64
	Classdiagram_NoteShapesDBID_Index: NullInt64  = new NullInt64 // store the index of the noteshape instance in Classdiagram.NoteShapes
	Classdiagram_NoteShapes_reverse?: ClassdiagramDB 

}
