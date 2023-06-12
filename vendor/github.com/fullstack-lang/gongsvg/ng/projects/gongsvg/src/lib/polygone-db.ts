// insertion point for imports
import { AnimateDB } from './animate-db'
import { LayerDB } from './layer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolygoneDB {

	static GONGSTRUCT_NAME = "Polygone"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Points: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	Layer_PolygonesDBID: NullInt64 = new NullInt64
	Layer_PolygonesDBID_Index: NullInt64  = new NullInt64 // store the index of the polygone instance in Layer.Polygones
	Layer_Polygones_reverse?: LayerDB 

}
