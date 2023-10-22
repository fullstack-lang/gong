// insertion point for imports
import { AnimateDB } from './animate-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CircleDB {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	Radius: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animations: Array<AnimateDB> = []

	CirclePointersEncoding: CirclePointersEncoding = new CirclePointersEncoding
}

export class CirclePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animations: number[] = []
}
