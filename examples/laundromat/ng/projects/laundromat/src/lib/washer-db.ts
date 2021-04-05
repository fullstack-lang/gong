// insertion point for imports
import { MachineDB } from './machine-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class WasherDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	TechName?: string
	Name?: string
	DirtyLaundryWeight?: number
	State?: string
	CleanedLaundryWeight?: number

	// insertion point for other declarations
	Machine?: MachineDB
	MachineID?: NullInt64
	MachineName?: string

}
