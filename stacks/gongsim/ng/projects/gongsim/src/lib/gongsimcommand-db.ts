// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GongsimCommandDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Command?: string
	CommandDate?: string
	SpeedCommandType?: string
	DateSpeedCommand?: string

	// insertion point for other declarations
}
