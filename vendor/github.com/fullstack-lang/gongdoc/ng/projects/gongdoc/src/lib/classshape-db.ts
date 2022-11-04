// insertion point for imports
import { PositionDB } from './position-db'
import { ReferenceDB } from './reference-db'
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
	ReferenceName: string = ""
	ShowNbInstances: boolean = false
	NbInstances: number = 0
	Width: number = 0
	Heigth: number = 0
	IsSelected: boolean = false

	// insertion point for other declarations
	Position?: PositionDB
	PositionID: NullInt64 = new NullInt64 // if pointer is null, Position.ID = 0

	Reference?: ReferenceDB
	ReferenceID: NullInt64 = new NullInt64 // if pointer is null, Reference.ID = 0

	Fields?: Array<FieldDB>
	Links?: Array<LinkDB>
	Classdiagram_ClassshapesDBID: NullInt64 = new NullInt64
	Classdiagram_ClassshapesDBID_Index: NullInt64  = new NullInt64 // store the index of the classshape instance in Classdiagram.Classshapes
	Classdiagram_Classshapes_reverse?: ClassdiagramDB 

}
