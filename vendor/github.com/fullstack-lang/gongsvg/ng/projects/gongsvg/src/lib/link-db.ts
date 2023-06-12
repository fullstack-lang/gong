// insertion point for imports
import { RectDB } from './rect-db'
import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { PointDB } from './point-db'
import { LayerDB } from './layer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: string = ""
	StartAnchorType: string = ""
	EndAnchorType: string = ""
	StartOrientation: string = ""
	StartRatio: number = 0
	EndOrientation: string = ""
	EndRatio: number = 0
	CornerOffsetRatio: number = 0
	CornerRadius: number = 0
	HasEndArrow: boolean = false
	EndArrowSize: number = 0
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

	End?: RectDB
	EndID: NullInt64 = new NullInt64 // if pointer is null, End.ID = 0

	TextAtArrowEnd?: Array<LinkAnchoredTextDB>
	TextAtArrowStart?: Array<LinkAnchoredTextDB>
	ControlPoints?: Array<PointDB>
	Layer_LinksDBID: NullInt64 = new NullInt64
	Layer_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the link instance in Layer.Links
	Layer_Links_reverse?: LayerDB 

}
