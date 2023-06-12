// insertion point for imports
import { PositionDB } from './position-db'
import { GongEnumValueEntryDB } from './gongenumvalueentry-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongEnumShapeDB {

	static GONGSTRUCT_NAME = "GongEnumShape"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	Width: number = 0
	Heigth: number = 0

	// insertion point for other declarations
	Position?: PositionDB
	PositionID: NullInt64 = new NullInt64 // if pointer is null, Position.ID = 0

	GongEnumValueEntrys?: Array<GongEnumValueEntryDB>
	Classdiagram_GongEnumShapesDBID: NullInt64 = new NullInt64
	Classdiagram_GongEnumShapesDBID_Index: NullInt64  = new NullInt64 // store the index of the gongenumshape instance in Classdiagram.GongEnumShapes
	Classdiagram_GongEnumShapes_reverse?: ClassdiagramDB 

}
