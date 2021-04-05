// insertion point for imports
import { ClassshapeDB } from './classshape-db'
import { PkgeltDB } from './pkgelt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class ClassdiagramDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	Classshapes?: Array<ClassshapeDB>
	Pkgelt_ClassdiagramsDBID?: NullInt64
	Pkgelt_Classdiagrams_reverse?: PkgeltDB

}
