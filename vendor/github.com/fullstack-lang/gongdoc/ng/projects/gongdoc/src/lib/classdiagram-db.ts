// insertion point for imports
import { ClassshapeDB } from './classshape-db'
import { NoteDB } from './note-db'
import { PkgeltDB } from './pkgelt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ClassdiagramDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsEditable: boolean = false

	// insertion point for other declarations
	Classshapes?: Array<ClassshapeDB>
	Notes?: Array<NoteDB>
	Pkgelt_ClassdiagramsDBID: NullInt64 = new NullInt64
	Pkgelt_ClassdiagramsDBID_Index: NullInt64  = new NullInt64 // store the index of the classdiagram instance in Pkgelt.Classdiagrams
	Pkgelt_Classdiagrams_reverse?: PkgeltDB 

}
