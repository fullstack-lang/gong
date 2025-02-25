import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AnimateAPI } from './animate-api'
import { Animate, CopyAnimateAPIToAnimate } from './animate'
import { AnimateService } from './animate.service'

import { CircleAPI } from './circle-api'
import { Circle, CopyCircleAPIToCircle } from './circle'
import { CircleService } from './circle.service'

import { EllipseAPI } from './ellipse-api'
import { Ellipse, CopyEllipseAPIToEllipse } from './ellipse'
import { EllipseService } from './ellipse.service'

import { LayerAPI } from './layer-api'
import { Layer, CopyLayerAPIToLayer } from './layer'
import { LayerService } from './layer.service'

import { LineAPI } from './line-api'
import { Line, CopyLineAPIToLine } from './line'
import { LineService } from './line.service'

import { LinkAPI } from './link-api'
import { Link, CopyLinkAPIToLink } from './link'
import { LinkService } from './link.service'

import { LinkAnchoredTextAPI } from './linkanchoredtext-api'
import { LinkAnchoredText, CopyLinkAnchoredTextAPIToLinkAnchoredText } from './linkanchoredtext'
import { LinkAnchoredTextService } from './linkanchoredtext.service'

import { PathAPI } from './path-api'
import { Path, CopyPathAPIToPath } from './path'
import { PathService } from './path.service'

import { PointAPI } from './point-api'
import { Point, CopyPointAPIToPoint } from './point'
import { PointService } from './point.service'

import { PolygoneAPI } from './polygone-api'
import { Polygone, CopyPolygoneAPIToPolygone } from './polygone'
import { PolygoneService } from './polygone.service'

import { PolylineAPI } from './polyline-api'
import { Polyline, CopyPolylineAPIToPolyline } from './polyline'
import { PolylineService } from './polyline.service'

import { RectAPI } from './rect-api'
import { Rect, CopyRectAPIToRect } from './rect'
import { RectService } from './rect.service'

import { RectAnchoredPathAPI } from './rectanchoredpath-api'
import { RectAnchoredPath, CopyRectAnchoredPathAPIToRectAnchoredPath } from './rectanchoredpath'
import { RectAnchoredPathService } from './rectanchoredpath.service'

import { RectAnchoredRectAPI } from './rectanchoredrect-api'
import { RectAnchoredRect, CopyRectAnchoredRectAPIToRectAnchoredRect } from './rectanchoredrect'
import { RectAnchoredRectService } from './rectanchoredrect.service'

import { RectAnchoredTextAPI } from './rectanchoredtext-api'
import { RectAnchoredText, CopyRectAnchoredTextAPIToRectAnchoredText } from './rectanchoredtext'
import { RectAnchoredTextService } from './rectanchoredtext.service'

import { RectLinkLinkAPI } from './rectlinklink-api'
import { RectLinkLink, CopyRectLinkLinkAPIToRectLinkLink } from './rectlinklink'
import { RectLinkLinkService } from './rectlinklink.service'

import { SVGAPI } from './svg-api'
import { SVG, CopySVGAPIToSVG } from './svg'
import { SVGService } from './svg.service'

import { SvgTextAPI } from './svgtext-api'
import { SvgText, CopySvgTextAPIToSvgText } from './svgtext'
import { SvgTextService } from './svgtext.service'

import { TextAPI } from './text-api'
import { Text, CopyTextAPIToText } from './text'
import { TextService } from './text.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongsvg/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Animates = new Array<Animate>() // array of front instances
	map_ID_Animate = new Map<number, Animate>() // map of front instances

	array_Circles = new Array<Circle>() // array of front instances
	map_ID_Circle = new Map<number, Circle>() // map of front instances

	array_Ellipses = new Array<Ellipse>() // array of front instances
	map_ID_Ellipse = new Map<number, Ellipse>() // map of front instances

	array_Layers = new Array<Layer>() // array of front instances
	map_ID_Layer = new Map<number, Layer>() // map of front instances

	array_Lines = new Array<Line>() // array of front instances
	map_ID_Line = new Map<number, Line>() // map of front instances

	array_Links = new Array<Link>() // array of front instances
	map_ID_Link = new Map<number, Link>() // map of front instances

	array_LinkAnchoredTexts = new Array<LinkAnchoredText>() // array of front instances
	map_ID_LinkAnchoredText = new Map<number, LinkAnchoredText>() // map of front instances

	array_Paths = new Array<Path>() // array of front instances
	map_ID_Path = new Map<number, Path>() // map of front instances

	array_Points = new Array<Point>() // array of front instances
	map_ID_Point = new Map<number, Point>() // map of front instances

	array_Polygones = new Array<Polygone>() // array of front instances
	map_ID_Polygone = new Map<number, Polygone>() // map of front instances

	array_Polylines = new Array<Polyline>() // array of front instances
	map_ID_Polyline = new Map<number, Polyline>() // map of front instances

	array_Rects = new Array<Rect>() // array of front instances
	map_ID_Rect = new Map<number, Rect>() // map of front instances

	array_RectAnchoredPaths = new Array<RectAnchoredPath>() // array of front instances
	map_ID_RectAnchoredPath = new Map<number, RectAnchoredPath>() // map of front instances

	array_RectAnchoredRects = new Array<RectAnchoredRect>() // array of front instances
	map_ID_RectAnchoredRect = new Map<number, RectAnchoredRect>() // map of front instances

	array_RectAnchoredTexts = new Array<RectAnchoredText>() // array of front instances
	map_ID_RectAnchoredText = new Map<number, RectAnchoredText>() // map of front instances

	array_RectLinkLinks = new Array<RectLinkLink>() // array of front instances
	map_ID_RectLinkLink = new Map<number, RectLinkLink>() // map of front instances

	array_SVGs = new Array<SVG>() // array of front instances
	map_ID_SVG = new Map<number, SVG>() // map of front instances

	array_SvgTexts = new Array<SvgText>() // array of front instances
	map_ID_SvgText = new Map<number, SvgText>() // map of front instances

	array_Texts = new Array<Text>() // array of front instances
	map_ID_Text = new Map<number, Text>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Animate':
				return this.array_Animates as unknown as Array<Type>
			case 'Circle':
				return this.array_Circles as unknown as Array<Type>
			case 'Ellipse':
				return this.array_Ellipses as unknown as Array<Type>
			case 'Layer':
				return this.array_Layers as unknown as Array<Type>
			case 'Line':
				return this.array_Lines as unknown as Array<Type>
			case 'Link':
				return this.array_Links as unknown as Array<Type>
			case 'LinkAnchoredText':
				return this.array_LinkAnchoredTexts as unknown as Array<Type>
			case 'Path':
				return this.array_Paths as unknown as Array<Type>
			case 'Point':
				return this.array_Points as unknown as Array<Type>
			case 'Polygone':
				return this.array_Polygones as unknown as Array<Type>
			case 'Polyline':
				return this.array_Polylines as unknown as Array<Type>
			case 'Rect':
				return this.array_Rects as unknown as Array<Type>
			case 'RectAnchoredPath':
				return this.array_RectAnchoredPaths as unknown as Array<Type>
			case 'RectAnchoredRect':
				return this.array_RectAnchoredRects as unknown as Array<Type>
			case 'RectAnchoredText':
				return this.array_RectAnchoredTexts as unknown as Array<Type>
			case 'RectLinkLink':
				return this.array_RectLinkLinks as unknown as Array<Type>
			case 'SVG':
				return this.array_SVGs as unknown as Array<Type>
			case 'SvgText':
				return this.array_SvgTexts as unknown as Array<Type>
			case 'Text':
				return this.array_Texts as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Animate':
				return this.map_ID_Animate as unknown as Map<number, Type>
			case 'Circle':
				return this.map_ID_Circle as unknown as Map<number, Type>
			case 'Ellipse':
				return this.map_ID_Ellipse as unknown as Map<number, Type>
			case 'Layer':
				return this.map_ID_Layer as unknown as Map<number, Type>
			case 'Line':
				return this.map_ID_Line as unknown as Map<number, Type>
			case 'Link':
				return this.map_ID_Link as unknown as Map<number, Type>
			case 'LinkAnchoredText':
				return this.map_ID_LinkAnchoredText as unknown as Map<number, Type>
			case 'Path':
				return this.map_ID_Path as unknown as Map<number, Type>
			case 'Point':
				return this.map_ID_Point as unknown as Map<number, Type>
			case 'Polygone':
				return this.map_ID_Polygone as unknown as Map<number, Type>
			case 'Polyline':
				return this.map_ID_Polyline as unknown as Map<number, Type>
			case 'Rect':
				return this.map_ID_Rect as unknown as Map<number, Type>
			case 'RectAnchoredPath':
				return this.map_ID_RectAnchoredPath as unknown as Map<number, Type>
			case 'RectAnchoredRect':
				return this.map_ID_RectAnchoredRect as unknown as Map<number, Type>
			case 'RectAnchoredText':
				return this.map_ID_RectAnchoredText as unknown as Map<number, Type>
			case 'RectLinkLink':
				return this.map_ID_RectLinkLink as unknown as Map<number, Type>
			case 'SVG':
				return this.map_ID_SVG as unknown as Map<number, Type>
			case 'SvgText':
				return this.map_ID_SvgText as unknown as Map<number, Type>
			case 'Text':
				return this.map_ID_Text as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private animateService: AnimateService,
		private circleService: CircleService,
		private ellipseService: EllipseService,
		private layerService: LayerService,
		private lineService: LineService,
		private linkService: LinkService,
		private linkanchoredtextService: LinkAnchoredTextService,
		private pathService: PathService,
		private pointService: PointService,
		private polygoneService: PolygoneService,
		private polylineService: PolylineService,
		private rectService: RectService,
		private rectanchoredpathService: RectAnchoredPathService,
		private rectanchoredrectService: RectAnchoredRectService,
		private rectanchoredtextService: RectAnchoredTextService,
		private rectlinklinkService: RectLinkLinkService,
		private svgService: SVGService,
		private svgtextService: SvgTextService,
		private textService: TextService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AnimateAPI[]>,
		Observable<CircleAPI[]>,
		Observable<EllipseAPI[]>,
		Observable<LayerAPI[]>,
		Observable<LineAPI[]>,
		Observable<LinkAPI[]>,
		Observable<LinkAnchoredTextAPI[]>,
		Observable<PathAPI[]>,
		Observable<PointAPI[]>,
		Observable<PolygoneAPI[]>,
		Observable<PolylineAPI[]>,
		Observable<RectAPI[]>,
		Observable<RectAnchoredPathAPI[]>,
		Observable<RectAnchoredRectAPI[]>,
		Observable<RectAnchoredTextAPI[]>,
		Observable<RectLinkLinkAPI[]>,
		Observable<SVGAPI[]>,
		Observable<SvgTextAPI[]>,
		Observable<TextAPI[]>,
	];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.animateService.getAnimates(this.GONG__StackPath, this.frontRepo),
			this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
			this.ellipseService.getEllipses(this.GONG__StackPath, this.frontRepo),
			this.layerService.getLayers(this.GONG__StackPath, this.frontRepo),
			this.lineService.getLines(this.GONG__StackPath, this.frontRepo),
			this.linkService.getLinks(this.GONG__StackPath, this.frontRepo),
			this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath, this.frontRepo),
			this.pathService.getPaths(this.GONG__StackPath, this.frontRepo),
			this.pointService.getPoints(this.GONG__StackPath, this.frontRepo),
			this.polygoneService.getPolygones(this.GONG__StackPath, this.frontRepo),
			this.polylineService.getPolylines(this.GONG__StackPath, this.frontRepo),
			this.rectService.getRects(this.GONG__StackPath, this.frontRepo),
			this.rectanchoredpathService.getRectAnchoredPaths(this.GONG__StackPath, this.frontRepo),
			this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath, this.frontRepo),
			this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath, this.frontRepo),
			this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath, this.frontRepo),
			this.svgService.getSVGs(this.GONG__StackPath, this.frontRepo),
			this.svgtextService.getSvgTexts(this.GONG__StackPath, this.frontRepo),
			this.textService.getTexts(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						animates_,
						circles_,
						ellipses_,
						layers_,
						lines_,
						links_,
						linkanchoredtexts_,
						paths_,
						points_,
						polygones_,
						polylines_,
						rects_,
						rectanchoredpaths_,
						rectanchoredrects_,
						rectanchoredtexts_,
						rectlinklinks_,
						svgs_,
						svgtexts_,
						texts_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var animates: AnimateAPI[]
						animates = animates_ as AnimateAPI[]
						var circles: CircleAPI[]
						circles = circles_ as CircleAPI[]
						var ellipses: EllipseAPI[]
						ellipses = ellipses_ as EllipseAPI[]
						var layers: LayerAPI[]
						layers = layers_ as LayerAPI[]
						var lines: LineAPI[]
						lines = lines_ as LineAPI[]
						var links: LinkAPI[]
						links = links_ as LinkAPI[]
						var linkanchoredtexts: LinkAnchoredTextAPI[]
						linkanchoredtexts = linkanchoredtexts_ as LinkAnchoredTextAPI[]
						var paths: PathAPI[]
						paths = paths_ as PathAPI[]
						var points: PointAPI[]
						points = points_ as PointAPI[]
						var polygones: PolygoneAPI[]
						polygones = polygones_ as PolygoneAPI[]
						var polylines: PolylineAPI[]
						polylines = polylines_ as PolylineAPI[]
						var rects: RectAPI[]
						rects = rects_ as RectAPI[]
						var rectanchoredpaths: RectAnchoredPathAPI[]
						rectanchoredpaths = rectanchoredpaths_ as RectAnchoredPathAPI[]
						var rectanchoredrects: RectAnchoredRectAPI[]
						rectanchoredrects = rectanchoredrects_ as RectAnchoredRectAPI[]
						var rectanchoredtexts: RectAnchoredTextAPI[]
						rectanchoredtexts = rectanchoredtexts_ as RectAnchoredTextAPI[]
						var rectlinklinks: RectLinkLinkAPI[]
						rectlinklinks = rectlinklinks_ as RectLinkLinkAPI[]
						var svgs: SVGAPI[]
						svgs = svgs_ as SVGAPI[]
						var svgtexts: SvgTextAPI[]
						svgtexts = svgtexts_ as SvgTextAPI[]
						var texts: TextAPI[]
						texts = texts_ as TextAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Animates = []
						this.frontRepo.map_ID_Animate.clear()

						animates.forEach(
							animateAPI => {
								let animate = new Animate
								this.frontRepo.array_Animates.push(animate)
								this.frontRepo.map_ID_Animate.set(animateAPI.ID, animate)
							}
						)

						// init the arrays
						this.frontRepo.array_Circles = []
						this.frontRepo.map_ID_Circle.clear()

						circles.forEach(
							circleAPI => {
								let circle = new Circle
								this.frontRepo.array_Circles.push(circle)
								this.frontRepo.map_ID_Circle.set(circleAPI.ID, circle)
							}
						)

						// init the arrays
						this.frontRepo.array_Ellipses = []
						this.frontRepo.map_ID_Ellipse.clear()

						ellipses.forEach(
							ellipseAPI => {
								let ellipse = new Ellipse
								this.frontRepo.array_Ellipses.push(ellipse)
								this.frontRepo.map_ID_Ellipse.set(ellipseAPI.ID, ellipse)
							}
						)

						// init the arrays
						this.frontRepo.array_Layers = []
						this.frontRepo.map_ID_Layer.clear()

						layers.forEach(
							layerAPI => {
								let layer = new Layer
								this.frontRepo.array_Layers.push(layer)
								this.frontRepo.map_ID_Layer.set(layerAPI.ID, layer)
							}
						)

						// init the arrays
						this.frontRepo.array_Lines = []
						this.frontRepo.map_ID_Line.clear()

						lines.forEach(
							lineAPI => {
								let line = new Line
								this.frontRepo.array_Lines.push(line)
								this.frontRepo.map_ID_Line.set(lineAPI.ID, line)
							}
						)

						// init the arrays
						this.frontRepo.array_Links = []
						this.frontRepo.map_ID_Link.clear()

						links.forEach(
							linkAPI => {
								let link = new Link
								this.frontRepo.array_Links.push(link)
								this.frontRepo.map_ID_Link.set(linkAPI.ID, link)
							}
						)

						// init the arrays
						this.frontRepo.array_LinkAnchoredTexts = []
						this.frontRepo.map_ID_LinkAnchoredText.clear()

						linkanchoredtexts.forEach(
							linkanchoredtextAPI => {
								let linkanchoredtext = new LinkAnchoredText
								this.frontRepo.array_LinkAnchoredTexts.push(linkanchoredtext)
								this.frontRepo.map_ID_LinkAnchoredText.set(linkanchoredtextAPI.ID, linkanchoredtext)
							}
						)

						// init the arrays
						this.frontRepo.array_Paths = []
						this.frontRepo.map_ID_Path.clear()

						paths.forEach(
							pathAPI => {
								let path = new Path
								this.frontRepo.array_Paths.push(path)
								this.frontRepo.map_ID_Path.set(pathAPI.ID, path)
							}
						)

						// init the arrays
						this.frontRepo.array_Points = []
						this.frontRepo.map_ID_Point.clear()

						points.forEach(
							pointAPI => {
								let point = new Point
								this.frontRepo.array_Points.push(point)
								this.frontRepo.map_ID_Point.set(pointAPI.ID, point)
							}
						)

						// init the arrays
						this.frontRepo.array_Polygones = []
						this.frontRepo.map_ID_Polygone.clear()

						polygones.forEach(
							polygoneAPI => {
								let polygone = new Polygone
								this.frontRepo.array_Polygones.push(polygone)
								this.frontRepo.map_ID_Polygone.set(polygoneAPI.ID, polygone)
							}
						)

						// init the arrays
						this.frontRepo.array_Polylines = []
						this.frontRepo.map_ID_Polyline.clear()

						polylines.forEach(
							polylineAPI => {
								let polyline = new Polyline
								this.frontRepo.array_Polylines.push(polyline)
								this.frontRepo.map_ID_Polyline.set(polylineAPI.ID, polyline)
							}
						)

						// init the arrays
						this.frontRepo.array_Rects = []
						this.frontRepo.map_ID_Rect.clear()

						rects.forEach(
							rectAPI => {
								let rect = new Rect
								this.frontRepo.array_Rects.push(rect)
								this.frontRepo.map_ID_Rect.set(rectAPI.ID, rect)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredPaths = []
						this.frontRepo.map_ID_RectAnchoredPath.clear()

						rectanchoredpaths.forEach(
							rectanchoredpathAPI => {
								let rectanchoredpath = new RectAnchoredPath
								this.frontRepo.array_RectAnchoredPaths.push(rectanchoredpath)
								this.frontRepo.map_ID_RectAnchoredPath.set(rectanchoredpathAPI.ID, rectanchoredpath)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredRects = []
						this.frontRepo.map_ID_RectAnchoredRect.clear()

						rectanchoredrects.forEach(
							rectanchoredrectAPI => {
								let rectanchoredrect = new RectAnchoredRect
								this.frontRepo.array_RectAnchoredRects.push(rectanchoredrect)
								this.frontRepo.map_ID_RectAnchoredRect.set(rectanchoredrectAPI.ID, rectanchoredrect)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredTexts = []
						this.frontRepo.map_ID_RectAnchoredText.clear()

						rectanchoredtexts.forEach(
							rectanchoredtextAPI => {
								let rectanchoredtext = new RectAnchoredText
								this.frontRepo.array_RectAnchoredTexts.push(rectanchoredtext)
								this.frontRepo.map_ID_RectAnchoredText.set(rectanchoredtextAPI.ID, rectanchoredtext)
							}
						)

						// init the arrays
						this.frontRepo.array_RectLinkLinks = []
						this.frontRepo.map_ID_RectLinkLink.clear()

						rectlinklinks.forEach(
							rectlinklinkAPI => {
								let rectlinklink = new RectLinkLink
								this.frontRepo.array_RectLinkLinks.push(rectlinklink)
								this.frontRepo.map_ID_RectLinkLink.set(rectlinklinkAPI.ID, rectlinklink)
							}
						)

						// init the arrays
						this.frontRepo.array_SVGs = []
						this.frontRepo.map_ID_SVG.clear()

						svgs.forEach(
							svgAPI => {
								let svg = new SVG
								this.frontRepo.array_SVGs.push(svg)
								this.frontRepo.map_ID_SVG.set(svgAPI.ID, svg)
							}
						)

						// init the arrays
						this.frontRepo.array_SvgTexts = []
						this.frontRepo.map_ID_SvgText.clear()

						svgtexts.forEach(
							svgtextAPI => {
								let svgtext = new SvgText
								this.frontRepo.array_SvgTexts.push(svgtext)
								this.frontRepo.map_ID_SvgText.set(svgtextAPI.ID, svgtext)
							}
						)

						// init the arrays
						this.frontRepo.array_Texts = []
						this.frontRepo.map_ID_Text.clear()

						texts.forEach(
							textAPI => {
								let text = new Text
								this.frontRepo.array_Texts.push(text)
								this.frontRepo.map_ID_Text.set(textAPI.ID, text)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						animates.forEach(
							animateAPI => {
								let animate = this.frontRepo.map_ID_Animate.get(animateAPI.ID)
								CopyAnimateAPIToAnimate(animateAPI, animate!, this.frontRepo)
							}
						)

						// fill up front objects
						circles.forEach(
							circleAPI => {
								let circle = this.frontRepo.map_ID_Circle.get(circleAPI.ID)
								CopyCircleAPIToCircle(circleAPI, circle!, this.frontRepo)
							}
						)

						// fill up front objects
						ellipses.forEach(
							ellipseAPI => {
								let ellipse = this.frontRepo.map_ID_Ellipse.get(ellipseAPI.ID)
								CopyEllipseAPIToEllipse(ellipseAPI, ellipse!, this.frontRepo)
							}
						)

						// fill up front objects
						layers.forEach(
							layerAPI => {
								let layer = this.frontRepo.map_ID_Layer.get(layerAPI.ID)
								CopyLayerAPIToLayer(layerAPI, layer!, this.frontRepo)
							}
						)

						// fill up front objects
						lines.forEach(
							lineAPI => {
								let line = this.frontRepo.map_ID_Line.get(lineAPI.ID)
								CopyLineAPIToLine(lineAPI, line!, this.frontRepo)
							}
						)

						// fill up front objects
						links.forEach(
							linkAPI => {
								let link = this.frontRepo.map_ID_Link.get(linkAPI.ID)
								CopyLinkAPIToLink(linkAPI, link!, this.frontRepo)
							}
						)

						// fill up front objects
						linkanchoredtexts.forEach(
							linkanchoredtextAPI => {
								let linkanchoredtext = this.frontRepo.map_ID_LinkAnchoredText.get(linkanchoredtextAPI.ID)
								CopyLinkAnchoredTextAPIToLinkAnchoredText(linkanchoredtextAPI, linkanchoredtext!, this.frontRepo)
							}
						)

						// fill up front objects
						paths.forEach(
							pathAPI => {
								let path = this.frontRepo.map_ID_Path.get(pathAPI.ID)
								CopyPathAPIToPath(pathAPI, path!, this.frontRepo)
							}
						)

						// fill up front objects
						points.forEach(
							pointAPI => {
								let point = this.frontRepo.map_ID_Point.get(pointAPI.ID)
								CopyPointAPIToPoint(pointAPI, point!, this.frontRepo)
							}
						)

						// fill up front objects
						polygones.forEach(
							polygoneAPI => {
								let polygone = this.frontRepo.map_ID_Polygone.get(polygoneAPI.ID)
								CopyPolygoneAPIToPolygone(polygoneAPI, polygone!, this.frontRepo)
							}
						)

						// fill up front objects
						polylines.forEach(
							polylineAPI => {
								let polyline = this.frontRepo.map_ID_Polyline.get(polylineAPI.ID)
								CopyPolylineAPIToPolyline(polylineAPI, polyline!, this.frontRepo)
							}
						)

						// fill up front objects
						rects.forEach(
							rectAPI => {
								let rect = this.frontRepo.map_ID_Rect.get(rectAPI.ID)
								CopyRectAPIToRect(rectAPI, rect!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredpaths.forEach(
							rectanchoredpathAPI => {
								let rectanchoredpath = this.frontRepo.map_ID_RectAnchoredPath.get(rectanchoredpathAPI.ID)
								CopyRectAnchoredPathAPIToRectAnchoredPath(rectanchoredpathAPI, rectanchoredpath!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredrects.forEach(
							rectanchoredrectAPI => {
								let rectanchoredrect = this.frontRepo.map_ID_RectAnchoredRect.get(rectanchoredrectAPI.ID)
								CopyRectAnchoredRectAPIToRectAnchoredRect(rectanchoredrectAPI, rectanchoredrect!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredtexts.forEach(
							rectanchoredtextAPI => {
								let rectanchoredtext = this.frontRepo.map_ID_RectAnchoredText.get(rectanchoredtextAPI.ID)
								CopyRectAnchoredTextAPIToRectAnchoredText(rectanchoredtextAPI, rectanchoredtext!, this.frontRepo)
							}
						)

						// fill up front objects
						rectlinklinks.forEach(
							rectlinklinkAPI => {
								let rectlinklink = this.frontRepo.map_ID_RectLinkLink.get(rectlinklinkAPI.ID)
								CopyRectLinkLinkAPIToRectLinkLink(rectlinklinkAPI, rectlinklink!, this.frontRepo)
							}
						)

						// fill up front objects
						svgs.forEach(
							svgAPI => {
								let svg = this.frontRepo.map_ID_SVG.get(svgAPI.ID)
								CopySVGAPIToSVG(svgAPI, svg!, this.frontRepo)
							}
						)

						// fill up front objects
						svgtexts.forEach(
							svgtextAPI => {
								let svgtext = this.frontRepo.map_ID_SvgText.get(svgtextAPI.ID)
								CopySvgTextAPIToSvgText(svgtextAPI, svgtext!, this.frontRepo)
							}
						)

						// fill up front objects
						texts.forEach(
							textAPI => {
								let text = this.frontRepo.map_ID_Text.get(textAPI.ID)
								CopyTextAPIToText(textAPI, text!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongsvg/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Animates = []
				frontRepo.map_ID_Animate.clear()

				backRepoData.AnimateAPIs.forEach(
					animateAPI => {
						let animate = new Animate
						frontRepo.array_Animates.push(animate)
						frontRepo.map_ID_Animate.set(animateAPI.ID, animate)
					}
				)

				// init the arrays
				frontRepo.array_Circles = []
				frontRepo.map_ID_Circle.clear()

				backRepoData.CircleAPIs.forEach(
					circleAPI => {
						let circle = new Circle
						frontRepo.array_Circles.push(circle)
						frontRepo.map_ID_Circle.set(circleAPI.ID, circle)
					}
				)

				// init the arrays
				frontRepo.array_Ellipses = []
				frontRepo.map_ID_Ellipse.clear()

				backRepoData.EllipseAPIs.forEach(
					ellipseAPI => {
						let ellipse = new Ellipse
						frontRepo.array_Ellipses.push(ellipse)
						frontRepo.map_ID_Ellipse.set(ellipseAPI.ID, ellipse)
					}
				)

				// init the arrays
				frontRepo.array_Layers = []
				frontRepo.map_ID_Layer.clear()

				backRepoData.LayerAPIs.forEach(
					layerAPI => {
						let layer = new Layer
						frontRepo.array_Layers.push(layer)
						frontRepo.map_ID_Layer.set(layerAPI.ID, layer)
					}
				)

				// init the arrays
				frontRepo.array_Lines = []
				frontRepo.map_ID_Line.clear()

				backRepoData.LineAPIs.forEach(
					lineAPI => {
						let line = new Line
						frontRepo.array_Lines.push(line)
						frontRepo.map_ID_Line.set(lineAPI.ID, line)
					}
				)

				// init the arrays
				frontRepo.array_Links = []
				frontRepo.map_ID_Link.clear()

				backRepoData.LinkAPIs.forEach(
					linkAPI => {
						let link = new Link
						frontRepo.array_Links.push(link)
						frontRepo.map_ID_Link.set(linkAPI.ID, link)
					}
				)

				// init the arrays
				frontRepo.array_LinkAnchoredTexts = []
				frontRepo.map_ID_LinkAnchoredText.clear()

				backRepoData.LinkAnchoredTextAPIs.forEach(
					linkanchoredtextAPI => {
						let linkanchoredtext = new LinkAnchoredText
						frontRepo.array_LinkAnchoredTexts.push(linkanchoredtext)
						frontRepo.map_ID_LinkAnchoredText.set(linkanchoredtextAPI.ID, linkanchoredtext)
					}
				)

				// init the arrays
				frontRepo.array_Paths = []
				frontRepo.map_ID_Path.clear()

				backRepoData.PathAPIs.forEach(
					pathAPI => {
						let path = new Path
						frontRepo.array_Paths.push(path)
						frontRepo.map_ID_Path.set(pathAPI.ID, path)
					}
				)

				// init the arrays
				frontRepo.array_Points = []
				frontRepo.map_ID_Point.clear()

				backRepoData.PointAPIs.forEach(
					pointAPI => {
						let point = new Point
						frontRepo.array_Points.push(point)
						frontRepo.map_ID_Point.set(pointAPI.ID, point)
					}
				)

				// init the arrays
				frontRepo.array_Polygones = []
				frontRepo.map_ID_Polygone.clear()

				backRepoData.PolygoneAPIs.forEach(
					polygoneAPI => {
						let polygone = new Polygone
						frontRepo.array_Polygones.push(polygone)
						frontRepo.map_ID_Polygone.set(polygoneAPI.ID, polygone)
					}
				)

				// init the arrays
				frontRepo.array_Polylines = []
				frontRepo.map_ID_Polyline.clear()

				backRepoData.PolylineAPIs.forEach(
					polylineAPI => {
						let polyline = new Polyline
						frontRepo.array_Polylines.push(polyline)
						frontRepo.map_ID_Polyline.set(polylineAPI.ID, polyline)
					}
				)

				// init the arrays
				frontRepo.array_Rects = []
				frontRepo.map_ID_Rect.clear()

				backRepoData.RectAPIs.forEach(
					rectAPI => {
						let rect = new Rect
						frontRepo.array_Rects.push(rect)
						frontRepo.map_ID_Rect.set(rectAPI.ID, rect)
					}
				)

				// init the arrays
				frontRepo.array_RectAnchoredPaths = []
				frontRepo.map_ID_RectAnchoredPath.clear()

				backRepoData.RectAnchoredPathAPIs.forEach(
					rectanchoredpathAPI => {
						let rectanchoredpath = new RectAnchoredPath
						frontRepo.array_RectAnchoredPaths.push(rectanchoredpath)
						frontRepo.map_ID_RectAnchoredPath.set(rectanchoredpathAPI.ID, rectanchoredpath)
					}
				)

				// init the arrays
				frontRepo.array_RectAnchoredRects = []
				frontRepo.map_ID_RectAnchoredRect.clear()

				backRepoData.RectAnchoredRectAPIs.forEach(
					rectanchoredrectAPI => {
						let rectanchoredrect = new RectAnchoredRect
						frontRepo.array_RectAnchoredRects.push(rectanchoredrect)
						frontRepo.map_ID_RectAnchoredRect.set(rectanchoredrectAPI.ID, rectanchoredrect)
					}
				)

				// init the arrays
				frontRepo.array_RectAnchoredTexts = []
				frontRepo.map_ID_RectAnchoredText.clear()

				backRepoData.RectAnchoredTextAPIs.forEach(
					rectanchoredtextAPI => {
						let rectanchoredtext = new RectAnchoredText
						frontRepo.array_RectAnchoredTexts.push(rectanchoredtext)
						frontRepo.map_ID_RectAnchoredText.set(rectanchoredtextAPI.ID, rectanchoredtext)
					}
				)

				// init the arrays
				frontRepo.array_RectLinkLinks = []
				frontRepo.map_ID_RectLinkLink.clear()

				backRepoData.RectLinkLinkAPIs.forEach(
					rectlinklinkAPI => {
						let rectlinklink = new RectLinkLink
						frontRepo.array_RectLinkLinks.push(rectlinklink)
						frontRepo.map_ID_RectLinkLink.set(rectlinklinkAPI.ID, rectlinklink)
					}
				)

				// init the arrays
				frontRepo.array_SVGs = []
				frontRepo.map_ID_SVG.clear()

				backRepoData.SVGAPIs.forEach(
					svgAPI => {
						let svg = new SVG
						frontRepo.array_SVGs.push(svg)
						frontRepo.map_ID_SVG.set(svgAPI.ID, svg)
					}
				)

				// init the arrays
				frontRepo.array_SvgTexts = []
				frontRepo.map_ID_SvgText.clear()

				backRepoData.SvgTextAPIs.forEach(
					svgtextAPI => {
						let svgtext = new SvgText
						frontRepo.array_SvgTexts.push(svgtext)
						frontRepo.map_ID_SvgText.set(svgtextAPI.ID, svgtext)
					}
				)

				// init the arrays
				frontRepo.array_Texts = []
				frontRepo.map_ID_Text.clear()

				backRepoData.TextAPIs.forEach(
					textAPI => {
						let text = new Text
						frontRepo.array_Texts.push(text)
						frontRepo.map_ID_Text.set(textAPI.ID, text)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AnimateAPIs.forEach(
					animateAPI => {
						let animate = frontRepo.map_ID_Animate.get(animateAPI.ID)
						CopyAnimateAPIToAnimate(animateAPI, animate!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CircleAPIs.forEach(
					circleAPI => {
						let circle = frontRepo.map_ID_Circle.get(circleAPI.ID)
						CopyCircleAPIToCircle(circleAPI, circle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.EllipseAPIs.forEach(
					ellipseAPI => {
						let ellipse = frontRepo.map_ID_Ellipse.get(ellipseAPI.ID)
						CopyEllipseAPIToEllipse(ellipseAPI, ellipse!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LayerAPIs.forEach(
					layerAPI => {
						let layer = frontRepo.map_ID_Layer.get(layerAPI.ID)
						CopyLayerAPIToLayer(layerAPI, layer!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LineAPIs.forEach(
					lineAPI => {
						let line = frontRepo.map_ID_Line.get(lineAPI.ID)
						CopyLineAPIToLine(lineAPI, line!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LinkAPIs.forEach(
					linkAPI => {
						let link = frontRepo.map_ID_Link.get(linkAPI.ID)
						CopyLinkAPIToLink(linkAPI, link!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LinkAnchoredTextAPIs.forEach(
					linkanchoredtextAPI => {
						let linkanchoredtext = frontRepo.map_ID_LinkAnchoredText.get(linkanchoredtextAPI.ID)
						CopyLinkAnchoredTextAPIToLinkAnchoredText(linkanchoredtextAPI, linkanchoredtext!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PathAPIs.forEach(
					pathAPI => {
						let path = frontRepo.map_ID_Path.get(pathAPI.ID)
						CopyPathAPIToPath(pathAPI, path!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PointAPIs.forEach(
					pointAPI => {
						let point = frontRepo.map_ID_Point.get(pointAPI.ID)
						CopyPointAPIToPoint(pointAPI, point!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PolygoneAPIs.forEach(
					polygoneAPI => {
						let polygone = frontRepo.map_ID_Polygone.get(polygoneAPI.ID)
						CopyPolygoneAPIToPolygone(polygoneAPI, polygone!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PolylineAPIs.forEach(
					polylineAPI => {
						let polyline = frontRepo.map_ID_Polyline.get(polylineAPI.ID)
						CopyPolylineAPIToPolyline(polylineAPI, polyline!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RectAPIs.forEach(
					rectAPI => {
						let rect = frontRepo.map_ID_Rect.get(rectAPI.ID)
						CopyRectAPIToRect(rectAPI, rect!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RectAnchoredPathAPIs.forEach(
					rectanchoredpathAPI => {
						let rectanchoredpath = frontRepo.map_ID_RectAnchoredPath.get(rectanchoredpathAPI.ID)
						CopyRectAnchoredPathAPIToRectAnchoredPath(rectanchoredpathAPI, rectanchoredpath!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RectAnchoredRectAPIs.forEach(
					rectanchoredrectAPI => {
						let rectanchoredrect = frontRepo.map_ID_RectAnchoredRect.get(rectanchoredrectAPI.ID)
						CopyRectAnchoredRectAPIToRectAnchoredRect(rectanchoredrectAPI, rectanchoredrect!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RectAnchoredTextAPIs.forEach(
					rectanchoredtextAPI => {
						let rectanchoredtext = frontRepo.map_ID_RectAnchoredText.get(rectanchoredtextAPI.ID)
						CopyRectAnchoredTextAPIToRectAnchoredText(rectanchoredtextAPI, rectanchoredtext!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RectLinkLinkAPIs.forEach(
					rectlinklinkAPI => {
						let rectlinklink = frontRepo.map_ID_RectLinkLink.get(rectlinklinkAPI.ID)
						CopyRectLinkLinkAPIToRectLinkLink(rectlinklinkAPI, rectlinklink!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SVGAPIs.forEach(
					svgAPI => {
						let svg = frontRepo.map_ID_SVG.get(svgAPI.ID)
						CopySVGAPIToSVG(svgAPI, svg!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SvgTextAPIs.forEach(
					svgtextAPI => {
						let svgtext = frontRepo.map_ID_SvgText.get(svgtextAPI.ID)
						CopySvgTextAPIToSvgText(svgtextAPI, svgtext!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TextAPIs.forEach(
					textAPI => {
						let text = frontRepo.map_ID_Text.get(textAPI.ID)
						CopyTextAPIToText(textAPI, text!, frontRepo)
					}
				)



				observer.next(frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getAnimateUniqueID(id: number): number {
	return 31 * id
}
export function getCircleUniqueID(id: number): number {
	return 37 * id
}
export function getEllipseUniqueID(id: number): number {
	return 41 * id
}
export function getLayerUniqueID(id: number): number {
	return 43 * id
}
export function getLineUniqueID(id: number): number {
	return 47 * id
}
export function getLinkUniqueID(id: number): number {
	return 53 * id
}
export function getLinkAnchoredTextUniqueID(id: number): number {
	return 59 * id
}
export function getPathUniqueID(id: number): number {
	return 61 * id
}
export function getPointUniqueID(id: number): number {
	return 67 * id
}
export function getPolygoneUniqueID(id: number): number {
	return 71 * id
}
export function getPolylineUniqueID(id: number): number {
	return 73 * id
}
export function getRectUniqueID(id: number): number {
	return 79 * id
}
export function getRectAnchoredPathUniqueID(id: number): number {
	return 83 * id
}
export function getRectAnchoredRectUniqueID(id: number): number {
	return 89 * id
}
export function getRectAnchoredTextUniqueID(id: number): number {
	return 97 * id
}
export function getRectLinkLinkUniqueID(id: number): number {
	return 101 * id
}
export function getSVGUniqueID(id: number): number {
	return 103 * id
}
export function getSvgTextUniqueID(id: number): number {
	return 107 * id
}
export function getTextUniqueID(id: number): number {
	return 109 * id
}
