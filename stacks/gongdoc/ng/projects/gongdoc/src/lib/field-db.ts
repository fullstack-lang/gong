// insertion point for imports
import { ClassshapeDB } from './classshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class FieldDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Fieldname?: string
	FieldTypeAsString?: string
	Structname?: string
	Fieldtypename?: string

	// insertion point for other declarations
	Classshape_FieldsDBID?: NullInt64
	Classshape_Fields_reverse?: ClassshapeDB

}
