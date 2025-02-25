// insertion point for imports
import { LayerAPI } from './layer-api'
import { RectAPI } from './rect-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGAPI {

	static GONGSTRUCT_NAME = "SVG"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DrawingState: string = ""
	IsEditable: boolean = false
	IsSVGFileGenerated: boolean = false

	// insertion point for other decls

	SVGPointersEncoding: SVGPointersEncoding = new SVGPointersEncoding
}

export class SVGPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Layers: number[] = []
	StartRectID: NullInt64 = new NullInt64 // if pointer is null, StartRect.ID = 0

	EndRectID: NullInt64 = new NullInt64 // if pointer is null, EndRect.ID = 0

}
