// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class MachineDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	TechName?: string
	Name?: string
	DrumLoad?: number
	RemainingTime?: number
	RemainingTimeMinutes?: number
	Cleanedlaundry?: string
	State?: string

	// insertion point for other declarations
}
