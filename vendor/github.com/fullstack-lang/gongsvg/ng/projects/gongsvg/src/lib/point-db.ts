// insertion point for imports
import { LinkDB } from './link-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PointDB {

	static GONGSTRUCT_NAME = "Point"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for other declarations
	Link_ControlPointsDBID: NullInt64 = new NullInt64
	Link_ControlPointsDBID_Index: NullInt64  = new NullInt64 // store the index of the point instance in Link.ControlPoints
	Link_ControlPoints_reverse?: LinkDB 

}
