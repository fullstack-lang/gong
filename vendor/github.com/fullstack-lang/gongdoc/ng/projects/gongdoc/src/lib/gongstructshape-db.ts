// insertion point for imports
import { PositionDB } from './position-db'
import { FieldDB } from './field-db'
import { LinkDB } from './link-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongStructShapeDB {

	static GONGSTRUCT_NAME = "GongStructShape"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	ShowNbInstances: boolean = false
	NbInstances: number = 0
	Width: number = 0
	Heigth: number = 0
	IsSelected: boolean = false

	// insertion point for other declarations
	Position?: PositionDB
	PositionID: NullInt64 = new NullInt64 // if pointer is null, Position.ID = 0

	Fields?: Array<FieldDB>
	Links?: Array<LinkDB>
	Classdiagram_GongStructShapesDBID: NullInt64 = new NullInt64
	Classdiagram_GongStructShapesDBID_Index: NullInt64  = new NullInt64 // store the index of the gongstructshape instance in Classdiagram.GongStructShapes
	Classdiagram_GongStructShapes_reverse?: ClassdiagramDB 

}
