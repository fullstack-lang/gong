// insertion point for imports
import { AnimateDB } from './animate-db'
import { LayerDB } from './layer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PathDB {

	static GONGSTRUCT_NAME = "Path"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	Layer_PathsDBID: NullInt64 = new NullInt64
	Layer_PathsDBID_Index: NullInt64  = new NullInt64 // store the index of the path instance in Layer.Paths
	Layer_Paths_reverse?: LayerDB 

}
