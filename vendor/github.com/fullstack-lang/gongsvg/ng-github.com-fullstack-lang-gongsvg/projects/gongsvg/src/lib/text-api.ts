// insertion point for imports
import { AnimateAPI } from './animate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TextAPI {

	static GONGSTRUCT_NAME = "Text"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Content: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	TextPointersEncoding: TextPointersEncoding = new TextPointersEncoding
}

export class TextPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
