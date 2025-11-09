// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SliderAPI {

	static GONGSTRUCT_NAME = "Slider"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	SliderPointersEncoding: SliderPointersEncoding = new SliderPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class SliderPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
