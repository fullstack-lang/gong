// insertion point for imports
import { AnimateAPI } from './animate-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class EllipseAPI {

	static GONGSTRUCT_NAME = "Ellipse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	RX: number = 0
	RY: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other decls

	EllipsePointersEncoding: EllipsePointersEncoding = new EllipsePointersEncoding
}

export class EllipsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
}
