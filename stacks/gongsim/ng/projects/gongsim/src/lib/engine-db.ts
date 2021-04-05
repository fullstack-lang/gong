// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class EngineDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	EndTime?: string
	CurrentTime?: string
	SecondsSinceStart?: number
	Fired?: number
	ControlMode?: string
	State?: string
	Speed?: number

	// insertion point for other declarations
}
