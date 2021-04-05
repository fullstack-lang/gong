// insertion point for imports
import { StateDB } from './state-db'
import { PkgeltDB } from './pkgelt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class UmlscDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Activestate?: string

	// insertion point for other declarations
	States?: Array<StateDB>
	Pkgelt_UmlscsDBID?: NullInt64
	Pkgelt_Umlscs_reverse?: PkgeltDB

}
