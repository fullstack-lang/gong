// insertion point for imports
import { PositionDB } from './position-db'
import { FieldDB } from './field-db'
import { LinkDB } from './link-db'
import { ClassdiagramDB } from './classdiagram-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class ClassshapeDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Structname?: string
	Width?: number
	Heigth?: number
	ClassshapeTargetType?: string

	// insertion point for other declarations
	Position?: PositionDB
	PositionID?: NullInt64
	PositionName?: string

	Fields?: Array<FieldDB>
	Links?: Array<LinkDB>
	Classdiagram_ClassshapesDBID?: NullInt64
	Classdiagram_Classshapes_reverse?: ClassdiagramDB

}
