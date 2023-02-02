// insertion point for imports
import { VerticeDB } from './vertice-db'
import { GongStructShapeDB } from './gongstructshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Structname: string = ""
	Identifier: string = ""
	Fieldtypename: string = ""
	TargetMultiplicity: string = ""
	SourceMultiplicity: string = ""

	// insertion point for other declarations
	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	GongStructShape_LinksDBID: NullInt64 = new NullInt64
	GongStructShape_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the link instance in GongStructShape.Links
	GongStructShape_Links_reverse?: GongStructShapeDB 

}
