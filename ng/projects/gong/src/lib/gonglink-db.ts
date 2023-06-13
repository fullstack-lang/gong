// insertion point for imports
import { GongNoteDB } from './gongnote-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongLinkDB {

	static GONGSTRUCT_NAME = "GongLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Recv: string = ""
	ImportPath: string = ""

	// insertion point for other declarations
	GongNote_LinksDBID: NullInt64 = new NullInt64
	GongNote_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the gonglink instance in GongNote.Links
	GongNote_Links_reverse?: GongNoteDB 

}
