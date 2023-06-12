// insertion point for imports
import { UmlscDB } from './umlsc-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UmlStateDB {

	static GONGSTRUCT_NAME = "UmlState"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for other declarations
	Umlsc_StatesDBID: NullInt64 = new NullInt64
	Umlsc_StatesDBID_Index: NullInt64  = new NullInt64 // store the index of the umlstate instance in Umlsc.States
	Umlsc_States_reverse?: UmlscDB 

}
