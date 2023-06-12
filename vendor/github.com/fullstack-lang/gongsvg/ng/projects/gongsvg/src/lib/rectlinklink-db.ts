// insertion point for imports
import { RectDB } from './rect-db'
import { LinkDB } from './link-db'
import { LayerDB } from './layer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectLinkLinkDB {

	static GONGSTRUCT_NAME = "RectLinkLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	TargetAnchorPosition: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Start?: RectDB
	StartID: NullInt64 = new NullInt64 // if pointer is null, Start.ID = 0

	End?: LinkDB
	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

	Layer_RectLinkLinksDBID: NullInt64 = new NullInt64
	Layer_RectLinkLinksDBID_Index: NullInt64  = new NullInt64 // store the index of the rectlinklink instance in Layer.RectLinkLinks
	Layer_RectLinkLinks_reverse?: LayerDB 

}
