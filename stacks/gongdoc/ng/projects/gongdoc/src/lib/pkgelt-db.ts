// insertion point for imports
import { ClassdiagramDB } from './classdiagram-db'
import { UmlscDB } from './umlsc-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class PkgeltDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Path?: string

	// insertion point for other declarations
	Classdiagrams?: Array<ClassdiagramDB>
	Umlscs?: Array<UmlscDB>
}
