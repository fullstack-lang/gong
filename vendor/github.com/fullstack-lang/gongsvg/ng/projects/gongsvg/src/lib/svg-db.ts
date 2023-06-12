// insertion point for imports
import { LayerDB } from './layer-db'
import { RectDB } from './rect-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGDB {

	static GONGSTRUCT_NAME = "SVG"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DrawingState: string = ""
	IsEditable: boolean = false

	// insertion point for other declarations
	Layers?: Array<LayerDB>
	StartRect?: RectDB
	StartRectID: NullInt64 = new NullInt64 // if pointer is null, StartRect.ID = 0

	EndRect?: RectDB
	EndRectID: NullInt64 = new NullInt64 // if pointer is null, EndRect.ID = 0

}
