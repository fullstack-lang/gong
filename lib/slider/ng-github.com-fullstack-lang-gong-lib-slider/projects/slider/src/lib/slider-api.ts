// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliderAPI {

	static GONGSTRUCT_NAME = "Slider"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsFloat64: boolean = false
	IsInt: boolean = false
	MinInt: number = 0
	MaxInt: number = 0
	StepInt: number = 0
	ValueInt: number = 0
	MinFloat64: number = 0
	MaxFloat64: number = 0
	StepFloat64: number = 0
	ValueFloat64: number = 0

	// insertion point for other decls

	SliderPointersEncoding: SliderPointersEncoding = new SliderPointersEncoding
}

export class SliderPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
