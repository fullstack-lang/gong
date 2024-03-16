// insertion point for imports
import { RectAPI } from './rect-api'
import { TextAPI } from './text-api'
import { CircleAPI } from './circle-api'
import { LineAPI } from './line-api'
import { EllipseAPI } from './ellipse-api'
import { PolylineAPI } from './polyline-api'
import { PolygoneAPI } from './polygone-api'
import { PathAPI } from './path-api'
import { LinkAPI } from './link-api'
import { RectLinkLinkAPI } from './rectlinklink-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerAPI {

	static GONGSTRUCT_NAME = "Layer"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Display: boolean = false
	Name: string = ""

	// insertion point for other decls

	LayerPointersEncoding: LayerPointersEncoding = new LayerPointersEncoding
}

export class LayerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Rects: number[] = []
	Texts: number[] = []
	Circles: number[] = []
	Lines: number[] = []
	Ellipses: number[] = []
	Polylines: number[] = []
	Polygones: number[] = []
	Paths: number[] = []
	Links: number[] = []
	RectLinkLinks: number[] = []
}
