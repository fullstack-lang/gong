// insertion point for imports
import { ClassdiagramDB } from './classdiagram-db'
import { UmlscDB } from './umlsc-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DiagramPackageDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Path: string = ""
	GongModelPath: string = ""
	IsEditable: boolean = false
	IsReloaded: boolean = false
	AbsolutePathToDiagramPackage: string = ""

	// insertion point for other declarations
	Classdiagrams?: Array<ClassdiagramDB>
	Umlscs?: Array<UmlscDB>
}
