// insertion point for imports
import { VerticeDB } from './vertice-db'
import { ClassshapeDB } from './classshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LinkDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Fieldname?: string
	Structname?: string
	Fieldtypename?: string
	Multiplicity?: string

	// insertion point for other declarations
	Middlevertice?: VerticeDB
	MiddleverticeID?: NullInt64
	MiddleverticeName?: string

	Classshape_LinksDBID?: NullInt64
	Classshape_Links_reverse?: ClassshapeDB

}
