// insertion point for imports
import { MachineDB } from './machine-db'
import { WasherDB } from './washer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class SimulationDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	LastCommitNb?: number

	// insertion point for other declarations
	Machine?: MachineDB
	MachineID?: NullInt64
	MachineName?: string

	Washer?: WasherDB
	WasherID?: NullInt64
	WasherName?: string

}
