// insertion point for imports
import { ClassshapeDB } from './classshape-db'
import { NoteDB } from './note-db'
import { DiagramPackageDB } from './diagrampackage-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ClassdiagramDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsInDrawMode: boolean = false

	// insertion point for other declarations
	Classshapes?: Array<ClassshapeDB>
	Notes?: Array<NoteDB>
	DiagramPackage_ClassdiagramsDBID: NullInt64 = new NullInt64
	DiagramPackage_ClassdiagramsDBID_Index: NullInt64  = new NullInt64 // store the index of the classdiagram instance in DiagramPackage.Classdiagrams
	DiagramPackage_Classdiagrams_reverse?: DiagramPackageDB 

}
