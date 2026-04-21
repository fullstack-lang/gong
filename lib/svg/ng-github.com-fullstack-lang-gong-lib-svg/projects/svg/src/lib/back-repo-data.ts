// generated code - do not edit

//insertion point for imports
import { AnimateAPI } from './animate-api'

import { CircleAPI } from './circle-api'

import { ConditionAPI } from './condition-api'

import { ControlPointAPI } from './controlpoint-api'

import { EllipseAPI } from './ellipse-api'

import { LayerAPI } from './layer-api'

import { LineAPI } from './line-api'

import { LinkAPI } from './link-api'

import { LinkAnchoredTextAPI } from './linkanchoredtext-api'

import { PathAPI } from './path-api'

import { PointAPI } from './point-api'

import { PolygoneAPI } from './polygone-api'

import { PolylineAPI } from './polyline-api'

import { RectAPI } from './rect-api'

import { RectAnchoredJpgImageAPI } from './rectanchoredjpgimage-api'

import { RectAnchoredPathAPI } from './rectanchoredpath-api'

import { RectAnchoredPngImageAPI } from './rectanchoredpngimage-api'

import { RectAnchoredRectAPI } from './rectanchoredrect-api'

import { RectAnchoredSvgImageAPI } from './rectanchoredsvgimage-api'

import { RectAnchoredTextAPI } from './rectanchoredtext-api'

import { RectLinkLinkAPI } from './rectlinklink-api'

import { SVGAPI } from './svg-api'

import { SvgTextAPI } from './svgtext-api'

import { TextAPI } from './text-api'


export class BackRepoData {
	// insertion point for declarations
	AnimateAPIs = new Array<AnimateAPI>()

	CircleAPIs = new Array<CircleAPI>()

	ConditionAPIs = new Array<ConditionAPI>()

	ControlPointAPIs = new Array<ControlPointAPI>()

	EllipseAPIs = new Array<EllipseAPI>()

	LayerAPIs = new Array<LayerAPI>()

	LineAPIs = new Array<LineAPI>()

	LinkAPIs = new Array<LinkAPI>()

	LinkAnchoredTextAPIs = new Array<LinkAnchoredTextAPI>()

	PathAPIs = new Array<PathAPI>()

	PointAPIs = new Array<PointAPI>()

	PolygoneAPIs = new Array<PolygoneAPI>()

	PolylineAPIs = new Array<PolylineAPI>()

	RectAPIs = new Array<RectAPI>()

	RectAnchoredJpgImageAPIs = new Array<RectAnchoredJpgImageAPI>()

	RectAnchoredPathAPIs = new Array<RectAnchoredPathAPI>()

	RectAnchoredPngImageAPIs = new Array<RectAnchoredPngImageAPI>()

	RectAnchoredRectAPIs = new Array<RectAnchoredRectAPI>()

	RectAnchoredSvgImageAPIs = new Array<RectAnchoredSvgImageAPI>()

	RectAnchoredTextAPIs = new Array<RectAnchoredTextAPI>()

	RectLinkLinkAPIs = new Array<RectLinkLinkAPI>()

	SVGAPIs = new Array<SVGAPI>()

	SvgTextAPIs = new Array<SvgTextAPI>()

	TextAPIs = new Array<TextAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AnimateAPIs = data?.AnimateAPIs || [];

		this.CircleAPIs = data?.CircleAPIs || [];

		this.ConditionAPIs = data?.ConditionAPIs || [];

		this.ControlPointAPIs = data?.ControlPointAPIs || [];

		this.EllipseAPIs = data?.EllipseAPIs || [];

		this.LayerAPIs = data?.LayerAPIs || [];

		this.LineAPIs = data?.LineAPIs || [];

		this.LinkAPIs = data?.LinkAPIs || [];

		this.LinkAnchoredTextAPIs = data?.LinkAnchoredTextAPIs || [];

		this.PathAPIs = data?.PathAPIs || [];

		this.PointAPIs = data?.PointAPIs || [];

		this.PolygoneAPIs = data?.PolygoneAPIs || [];

		this.PolylineAPIs = data?.PolylineAPIs || [];

		this.RectAPIs = data?.RectAPIs || [];

		this.RectAnchoredJpgImageAPIs = data?.RectAnchoredJpgImageAPIs || [];

		this.RectAnchoredPathAPIs = data?.RectAnchoredPathAPIs || [];

		this.RectAnchoredPngImageAPIs = data?.RectAnchoredPngImageAPIs || [];

		this.RectAnchoredRectAPIs = data?.RectAnchoredRectAPIs || [];

		this.RectAnchoredSvgImageAPIs = data?.RectAnchoredSvgImageAPIs || [];

		this.RectAnchoredTextAPIs = data?.RectAnchoredTextAPIs || [];

		this.RectLinkLinkAPIs = data?.RectLinkLinkAPIs || [];

		this.SVGAPIs = data?.SVGAPIs || [];

		this.SvgTextAPIs = data?.SvgTextAPIs || [];

		this.TextAPIs = data?.TextAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}