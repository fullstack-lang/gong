// insertion point for imports
import { GongEnumDB } from './gongenum-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GongEnumValueDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Value?: string

	// insertion point for other declarations
	GongEnum_GongEnumValuesDBID?: NullInt64
	GongEnum_GongEnumValues_reverse?: GongEnumDB

}
