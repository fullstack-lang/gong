// insertion point for imports
import { VerticeDB } from './vertice-db'
import { ClassshapeDB } from './classshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Fieldname: string = ""
	Structname: string = ""
	Fieldtypename: string = ""
	TargetMultiplicity: string = ""
	SourceMultiplicity: string = ""

	// insertion point for other declarations
	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	Classshape_LinksDBID: NullInt64 = new NullInt64
	Classshape_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the link instance in Classshape.Links
	Classshape_Links_reverse?: ClassshapeDB 

}
