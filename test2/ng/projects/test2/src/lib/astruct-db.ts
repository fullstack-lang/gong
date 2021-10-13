// insertion point for imports
import { BstructDB } from './bstruct-db'
import { AstructBstructUseDB } from './astructbstructuse-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class AstructDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Date: Date = new Date
	Booleanfield: boolean = false
	Floatfield: number = 0
	Intfield: number = 0
	Anotherbooleanfield: boolean = false
	Duration1: number = 0

	// insertion point for other declarations
	Duration1_string?: string
	Associationtob?: BstructDB
	AssociationtobID?: NullInt64

	Anarrayofb?: Array<BstructDB>
	AnarrayofbUse?: Array<AstructBstructUseDB>
}
