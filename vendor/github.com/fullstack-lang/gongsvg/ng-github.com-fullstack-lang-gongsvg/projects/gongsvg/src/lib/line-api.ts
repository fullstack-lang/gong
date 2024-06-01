// insertion point for imports
import { AnimateAPI } from './animate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LineAPI {

	static GONGSTRUCT_NAME = "Line"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X1: number = 0
	Y1: number = 0
	X2: number = 0
	Y2: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	MouseClickX: number = 0
	MouseClickY: number = 0

	// insertion point for other decls

	LinePointersEncoding: LinePointersEncoding = new LinePointersEncoding
}

export class LinePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
