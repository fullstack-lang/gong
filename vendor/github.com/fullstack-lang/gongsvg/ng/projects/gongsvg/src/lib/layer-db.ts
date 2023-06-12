// insertion point for imports
import { RectDB } from './rect-db'
import { TextDB } from './text-db'
import { CircleDB } from './circle-db'
import { LineDB } from './line-db'
import { EllipseDB } from './ellipse-db'
import { PolylineDB } from './polyline-db'
import { PolygoneDB } from './polygone-db'
import { PathDB } from './path-db'
import { LinkDB } from './link-db'
import { RectLinkLinkDB } from './rectlinklink-db'
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerDB {

	static GONGSTRUCT_NAME = "Layer"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Display: boolean = false
	Name: string = ""

	// insertion point for other declarations
	Rects?: Array<RectDB>
	Texts?: Array<TextDB>
	Circles?: Array<CircleDB>
	Lines?: Array<LineDB>
	Ellipses?: Array<EllipseDB>
	Polylines?: Array<PolylineDB>
	Polygones?: Array<PolygoneDB>
	Paths?: Array<PathDB>
	Links?: Array<LinkDB>
	RectLinkLinks?: Array<RectLinkLinkDB>
	SVG_LayersDBID: NullInt64 = new NullInt64
	SVG_LayersDBID_Index: NullInt64  = new NullInt64 // store the index of the layer instance in SVG.Layers
	SVG_Layers_reverse?: SVGDB 

}
