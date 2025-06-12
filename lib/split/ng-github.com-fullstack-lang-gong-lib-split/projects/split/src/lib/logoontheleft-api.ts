// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LogoOnTheLeftAPI {

	static GONGSTRUCT_NAME = "LogoOnTheLeft"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Height: number = 0
	SVG: string = ""

	// insertion point for other decls

	LogoOnTheLeftPointersEncoding: LogoOnTheLeftPointersEncoding = new LogoOnTheLeftPointersEncoding
}

export class LogoOnTheLeftPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
