// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class StatusAPI {

	static GONGSTRUCT_NAME = "Status"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CurrentCommand: string = ""
	CompletionDate: string = ""
	CurrentSpeedCommand: string = ""
	SpeedCommandCompletionDate: string = ""

	// insertion point for other decls

	StatusPointersEncoding: StatusPointersEncoding = new StatusPointersEncoding
}

export class StatusPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
