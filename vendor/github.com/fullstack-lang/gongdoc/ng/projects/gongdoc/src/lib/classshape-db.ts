// insertion point for imports
import { PositionDB } from './position-db'
import { GongStructDB } from './gongstruct-db'
import { FieldDB } from './field-db'
import { LinkDB } from './link-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ClassshapeDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Structname: string = ""
	ShowNbInstances: boolean = false
	NbInstances: number = 0
	Width: number = 0
	Heigth: number = 0
	ClassshapeTargetType: string = ""

	// insertion point for other declarations
	Position?: PositionDB
	PositionID: NullInt64 = new NullInt64 // if pointer is null, Position.ID = 0

	GongStruct?: GongStructDB
	GongStructID: NullInt64 = new NullInt64 // if pointer is null, GongStruct.ID = 0

	Fields?: Array<FieldDB>
	Links?: Array<LinkDB>
	Classdiagram_ClassshapesDBID: NullInt64 = new NullInt64
	Classdiagram_ClassshapesDBID_Index: NullInt64  = new NullInt64 // store the index of the classshape instance in Classdiagram.Classshapes
	Classdiagram_Classshapes_reverse?: ClassdiagramDB 

}
