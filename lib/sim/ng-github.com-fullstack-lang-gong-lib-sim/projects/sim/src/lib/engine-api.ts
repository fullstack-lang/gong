// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class EngineAPI {

	static GONGSTRUCT_NAME = "Engine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EndTime: string = ""
	CurrentTime: string = ""
	DisplayFormat: string = ""
	SecondsSinceStart: number = 0
	Fired: number = 0
	ControlMode: string = ""
	State: string = ""
	Speed: number = 0

	// insertion point for other decls

	EnginePointersEncoding: EnginePointersEncoding = new EnginePointersEncoding
}

export class EnginePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
