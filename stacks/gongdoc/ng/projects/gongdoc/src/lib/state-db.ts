// insertion point for imports
import { UmlscDB } from './umlsc-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class StateDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	X?: number
	Y?: number

	// insertion point for other declarations
	Umlsc_StatesDBID?: NullInt64
	Umlsc_States_reverse?: UmlscDB

}
