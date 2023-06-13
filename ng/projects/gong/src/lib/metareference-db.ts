// insertion point for imports
import { MetaDB } from './meta-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MetaReferenceDB {

	static GONGSTRUCT_NAME = "MetaReference"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Meta_MetaReferencesDBID: NullInt64 = new NullInt64
	Meta_MetaReferencesDBID_Index: NullInt64  = new NullInt64 // store the index of the metareference instance in Meta.MetaReferences
	Meta_MetaReferences_reverse?: MetaDB 

}
