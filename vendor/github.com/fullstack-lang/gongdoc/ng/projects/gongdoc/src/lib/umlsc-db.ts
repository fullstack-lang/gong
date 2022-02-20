// insertion point for imports
import { UmlStateDB } from './umlstate-db'
import { PkgeltDB } from './pkgelt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UmlscDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Activestate: string = ""

	// insertion point for other declarations
	States?: Array<UmlStateDB>
	Pkgelt_UmlscsDBID: NullInt64 = new NullInt64
	Pkgelt_UmlscsDBID_Index: NullInt64  = new NullInt64 // store the index of the umlsc instance in Pkgelt.Umlscs
	Pkgelt_Umlscs_reverse?: PkgeltDB 

}
