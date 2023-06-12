// insertion point for imports
import { UmlStateDB } from './umlstate-db'
import { DiagramPackageDB } from './diagrampackage-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UmlscDB {

	static GONGSTRUCT_NAME = "Umlsc"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Activestate: string = ""
	IsInDrawMode: boolean = false

	// insertion point for other declarations
	States?: Array<UmlStateDB>
	DiagramPackage_UmlscsDBID: NullInt64 = new NullInt64
	DiagramPackage_UmlscsDBID_Index: NullInt64  = new NullInt64 // store the index of the umlsc instance in DiagramPackage.Umlscs
	DiagramPackage_Umlscs_reverse?: DiagramPackageDB 

}
