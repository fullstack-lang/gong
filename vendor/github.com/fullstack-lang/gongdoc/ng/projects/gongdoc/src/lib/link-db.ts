// insertion point for imports
import { VerticeDB } from './vertice-db'
import { GongStructShapeDB } from './gongstructshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	Fieldtypename: string = ""
	FieldOffsetX: number = 0
	FieldOffsetY: number = 0
	TargetMultiplicity: string = ""
	TargetMultiplicityOffsetX: number = 0
	TargetMultiplicityOffsetY: number = 0
	SourceMultiplicity: string = ""
	SourceMultiplicityOffsetX: number = 0
	SourceMultiplicityOffsetY: number = 0
	StartOrientation: string = ""
	StartRatio: number = 0
	EndOrientation: string = ""
	EndRatio: number = 0
	CornerOffsetRatio: number = 0

	// insertion point for other declarations
	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	GongStructShape_LinksDBID: NullInt64 = new NullInt64
	GongStructShape_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the link instance in GongStructShape.Links
	GongStructShape_Links_reverse?: GongStructShapeDB 

}
