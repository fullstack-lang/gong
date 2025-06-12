// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LogoAPI {

	static GONGSTRUCT_NAME = "Logo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	SVG: string = ""

	// insertion point for other decls

	LogoPointersEncoding: LogoPointersEncoding = new LogoPointersEncoding
}

export class LogoPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
