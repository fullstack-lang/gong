import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AnimateDB } from './animate-db'
import { Animate, CopyAnimateDBToAnimate } from './animate'
import { AnimateService } from './animate.service'

import { CircleDB } from './circle-db'
import { Circle, CopyCircleDBToCircle } from './circle'
import { CircleService } from './circle.service'

import { EllipseDB } from './ellipse-db'
import { Ellipse, CopyEllipseDBToEllipse } from './ellipse'
import { EllipseService } from './ellipse.service'

import { LayerDB } from './layer-db'
import { Layer, CopyLayerDBToLayer } from './layer'
import { LayerService } from './layer.service'

import { LineDB } from './line-db'
import { Line, CopyLineDBToLine } from './line'
import { LineService } from './line.service'

import { LinkDB } from './link-db'
import { Link, CopyLinkDBToLink } from './link'
import { LinkService } from './link.service'

import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { LinkAnchoredText, CopyLinkAnchoredTextDBToLinkAnchoredText } from './linkanchoredtext'
import { LinkAnchoredTextService } from './linkanchoredtext.service'

import { PathDB } from './path-db'
import { Path, CopyPathDBToPath } from './path'
import { PathService } from './path.service'

import { PointDB } from './point-db'
import { Point, CopyPointDBToPoint } from './point'
import { PointService } from './point.service'

import { PolygoneDB } from './polygone-db'
import { Polygone, CopyPolygoneDBToPolygone } from './polygone'
import { PolygoneService } from './polygone.service'

import { PolylineDB } from './polyline-db'
import { Polyline, CopyPolylineDBToPolyline } from './polyline'
import { PolylineService } from './polyline.service'

import { RectDB } from './rect-db'
import { Rect, CopyRectDBToRect } from './rect'
import { RectService } from './rect.service'

import { RectAnchoredPathDB } from './rectanchoredpath-db'
import { RectAnchoredPath, CopyRectAnchoredPathDBToRectAnchoredPath } from './rectanchoredpath'
import { RectAnchoredPathService } from './rectanchoredpath.service'

import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { RectAnchoredRect, CopyRectAnchoredRectDBToRectAnchoredRect } from './rectanchoredrect'
import { RectAnchoredRectService } from './rectanchoredrect.service'

import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { RectAnchoredText, CopyRectAnchoredTextDBToRectAnchoredText } from './rectanchoredtext'
import { RectAnchoredTextService } from './rectanchoredtext.service'

import { RectLinkLinkDB } from './rectlinklink-db'
import { RectLinkLink, CopyRectLinkLinkDBToRectLinkLink } from './rectlinklink'
import { RectLinkLinkService } from './rectlinklink.service'

import { SVGDB } from './svg-db'
import { SVG, CopySVGDBToSVG } from './svg'
import { SVGService } from './svg.service'

import { TextDB } from './text-db'
import { Text, CopyTextDBToText } from './text'
import { TextService } from './text.service'

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

	array_Texts = new Array<Text>() // array of front instances
	map_ID_Text = new Map<number, Text>() // map of front instances


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
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AnimateDB[]>,
		Observable<CircleDB[]>,
		Observable<EllipseDB[]>,
		Observable<LayerDB[]>,
		Observable<LineDB[]>,
		Observable<LinkDB[]>,
		Observable<LinkAnchoredTextDB[]>,
		Observable<PathDB[]>,
		Observable<PointDB[]>,
		Observable<PolygoneDB[]>,
		Observable<PolylineDB[]>,
		Observable<RectDB[]>,
		Observable<RectAnchoredPathDB[]>,
		Observable<RectAnchoredRectDB[]>,
		Observable<RectAnchoredTextDB[]>,
		Observable<RectLinkLinkDB[]>,
		Observable<SVGDB[]>,
		Observable<TextDB[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
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
			this.textService.getTexts(this.GONG__StackPath, this.frontRepo),
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
						texts_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var animates: AnimateDB[]
						animates = animates_ as AnimateDB[]
						var circles: CircleDB[]
						circles = circles_ as CircleDB[]
						var ellipses: EllipseDB[]
						ellipses = ellipses_ as EllipseDB[]
						var layers: LayerDB[]
						layers = layers_ as LayerDB[]
						var lines: LineDB[]
						lines = lines_ as LineDB[]
						var links: LinkDB[]
						links = links_ as LinkDB[]
						var linkanchoredtexts: LinkAnchoredTextDB[]
						linkanchoredtexts = linkanchoredtexts_ as LinkAnchoredTextDB[]
						var paths: PathDB[]
						paths = paths_ as PathDB[]
						var points: PointDB[]
						points = points_ as PointDB[]
						var polygones: PolygoneDB[]
						polygones = polygones_ as PolygoneDB[]
						var polylines: PolylineDB[]
						polylines = polylines_ as PolylineDB[]
						var rects: RectDB[]
						rects = rects_ as RectDB[]
						var rectanchoredpaths: RectAnchoredPathDB[]
						rectanchoredpaths = rectanchoredpaths_ as RectAnchoredPathDB[]
						var rectanchoredrects: RectAnchoredRectDB[]
						rectanchoredrects = rectanchoredrects_ as RectAnchoredRectDB[]
						var rectanchoredtexts: RectAnchoredTextDB[]
						rectanchoredtexts = rectanchoredtexts_ as RectAnchoredTextDB[]
						var rectlinklinks: RectLinkLinkDB[]
						rectlinklinks = rectlinklinks_ as RectLinkLinkDB[]
						var svgs: SVGDB[]
						svgs = svgs_ as SVGDB[]
						var texts: TextDB[]
						texts = texts_ as TextDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Animates = []
						this.frontRepo.map_ID_Animate.clear()

						animates.forEach(
							animateDB => {
								let animate = new Animate
								this.frontRepo.array_Animates.push(animate)
								this.frontRepo.map_ID_Animate.set(animateDB.ID, animate)
							}
						)

						// init the arrays
						this.frontRepo.array_Circles = []
						this.frontRepo.map_ID_Circle.clear()

						circles.forEach(
							circleDB => {
								let circle = new Circle
								this.frontRepo.array_Circles.push(circle)
								this.frontRepo.map_ID_Circle.set(circleDB.ID, circle)
							}
						)

						// init the arrays
						this.frontRepo.array_Ellipses = []
						this.frontRepo.map_ID_Ellipse.clear()

						ellipses.forEach(
							ellipseDB => {
								let ellipse = new Ellipse
								this.frontRepo.array_Ellipses.push(ellipse)
								this.frontRepo.map_ID_Ellipse.set(ellipseDB.ID, ellipse)
							}
						)

						// init the arrays
						this.frontRepo.array_Layers = []
						this.frontRepo.map_ID_Layer.clear()

						layers.forEach(
							layerDB => {
								let layer = new Layer
								this.frontRepo.array_Layers.push(layer)
								this.frontRepo.map_ID_Layer.set(layerDB.ID, layer)
							}
						)

						// init the arrays
						this.frontRepo.array_Lines = []
						this.frontRepo.map_ID_Line.clear()

						lines.forEach(
							lineDB => {
								let line = new Line
								this.frontRepo.array_Lines.push(line)
								this.frontRepo.map_ID_Line.set(lineDB.ID, line)
							}
						)

						// init the arrays
						this.frontRepo.array_Links = []
						this.frontRepo.map_ID_Link.clear()

						links.forEach(
							linkDB => {
								let link = new Link
								this.frontRepo.array_Links.push(link)
								this.frontRepo.map_ID_Link.set(linkDB.ID, link)
							}
						)

						// init the arrays
						this.frontRepo.array_LinkAnchoredTexts = []
						this.frontRepo.map_ID_LinkAnchoredText.clear()

						linkanchoredtexts.forEach(
							linkanchoredtextDB => {
								let linkanchoredtext = new LinkAnchoredText
								this.frontRepo.array_LinkAnchoredTexts.push(linkanchoredtext)
								this.frontRepo.map_ID_LinkAnchoredText.set(linkanchoredtextDB.ID, linkanchoredtext)
							}
						)

						// init the arrays
						this.frontRepo.array_Paths = []
						this.frontRepo.map_ID_Path.clear()

						paths.forEach(
							pathDB => {
								let path = new Path
								this.frontRepo.array_Paths.push(path)
								this.frontRepo.map_ID_Path.set(pathDB.ID, path)
							}
						)

						// init the arrays
						this.frontRepo.array_Points = []
						this.frontRepo.map_ID_Point.clear()

						points.forEach(
							pointDB => {
								let point = new Point
								this.frontRepo.array_Points.push(point)
								this.frontRepo.map_ID_Point.set(pointDB.ID, point)
							}
						)

						// init the arrays
						this.frontRepo.array_Polygones = []
						this.frontRepo.map_ID_Polygone.clear()

						polygones.forEach(
							polygoneDB => {
								let polygone = new Polygone
								this.frontRepo.array_Polygones.push(polygone)
								this.frontRepo.map_ID_Polygone.set(polygoneDB.ID, polygone)
							}
						)

						// init the arrays
						this.frontRepo.array_Polylines = []
						this.frontRepo.map_ID_Polyline.clear()

						polylines.forEach(
							polylineDB => {
								let polyline = new Polyline
								this.frontRepo.array_Polylines.push(polyline)
								this.frontRepo.map_ID_Polyline.set(polylineDB.ID, polyline)
							}
						)

						// init the arrays
						this.frontRepo.array_Rects = []
						this.frontRepo.map_ID_Rect.clear()

						rects.forEach(
							rectDB => {
								let rect = new Rect
								this.frontRepo.array_Rects.push(rect)
								this.frontRepo.map_ID_Rect.set(rectDB.ID, rect)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredPaths = []
						this.frontRepo.map_ID_RectAnchoredPath.clear()

						rectanchoredpaths.forEach(
							rectanchoredpathDB => {
								let rectanchoredpath = new RectAnchoredPath
								this.frontRepo.array_RectAnchoredPaths.push(rectanchoredpath)
								this.frontRepo.map_ID_RectAnchoredPath.set(rectanchoredpathDB.ID, rectanchoredpath)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredRects = []
						this.frontRepo.map_ID_RectAnchoredRect.clear()

						rectanchoredrects.forEach(
							rectanchoredrectDB => {
								let rectanchoredrect = new RectAnchoredRect
								this.frontRepo.array_RectAnchoredRects.push(rectanchoredrect)
								this.frontRepo.map_ID_RectAnchoredRect.set(rectanchoredrectDB.ID, rectanchoredrect)
							}
						)

						// init the arrays
						this.frontRepo.array_RectAnchoredTexts = []
						this.frontRepo.map_ID_RectAnchoredText.clear()

						rectanchoredtexts.forEach(
							rectanchoredtextDB => {
								let rectanchoredtext = new RectAnchoredText
								this.frontRepo.array_RectAnchoredTexts.push(rectanchoredtext)
								this.frontRepo.map_ID_RectAnchoredText.set(rectanchoredtextDB.ID, rectanchoredtext)
							}
						)

						// init the arrays
						this.frontRepo.array_RectLinkLinks = []
						this.frontRepo.map_ID_RectLinkLink.clear()

						rectlinklinks.forEach(
							rectlinklinkDB => {
								let rectlinklink = new RectLinkLink
								this.frontRepo.array_RectLinkLinks.push(rectlinklink)
								this.frontRepo.map_ID_RectLinkLink.set(rectlinklinkDB.ID, rectlinklink)
							}
						)

						// init the arrays
						this.frontRepo.array_SVGs = []
						this.frontRepo.map_ID_SVG.clear()

						svgs.forEach(
							svgDB => {
								let svg = new SVG
								this.frontRepo.array_SVGs.push(svg)
								this.frontRepo.map_ID_SVG.set(svgDB.ID, svg)
							}
						)

						// init the arrays
						this.frontRepo.array_Texts = []
						this.frontRepo.map_ID_Text.clear()

						texts.forEach(
							textDB => {
								let text = new Text
								this.frontRepo.array_Texts.push(text)
								this.frontRepo.map_ID_Text.set(textDB.ID, text)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						animates.forEach(
							animateDB => {
								let animate = this.frontRepo.map_ID_Animate.get(animateDB.ID)
								CopyAnimateDBToAnimate(animateDB, animate!, this.frontRepo)
							}
						)

						// fill up front objects
						circles.forEach(
							circleDB => {
								let circle = this.frontRepo.map_ID_Circle.get(circleDB.ID)
								CopyCircleDBToCircle(circleDB, circle!, this.frontRepo)
							}
						)

						// fill up front objects
						ellipses.forEach(
							ellipseDB => {
								let ellipse = this.frontRepo.map_ID_Ellipse.get(ellipseDB.ID)
								CopyEllipseDBToEllipse(ellipseDB, ellipse!, this.frontRepo)
							}
						)

						// fill up front objects
						layers.forEach(
							layerDB => {
								let layer = this.frontRepo.map_ID_Layer.get(layerDB.ID)
								CopyLayerDBToLayer(layerDB, layer!, this.frontRepo)
							}
						)

						// fill up front objects
						lines.forEach(
							lineDB => {
								let line = this.frontRepo.map_ID_Line.get(lineDB.ID)
								CopyLineDBToLine(lineDB, line!, this.frontRepo)
							}
						)

						// fill up front objects
						links.forEach(
							linkDB => {
								let link = this.frontRepo.map_ID_Link.get(linkDB.ID)
								CopyLinkDBToLink(linkDB, link!, this.frontRepo)
							}
						)

						// fill up front objects
						linkanchoredtexts.forEach(
							linkanchoredtextDB => {
								let linkanchoredtext = this.frontRepo.map_ID_LinkAnchoredText.get(linkanchoredtextDB.ID)
								CopyLinkAnchoredTextDBToLinkAnchoredText(linkanchoredtextDB, linkanchoredtext!, this.frontRepo)
							}
						)

						// fill up front objects
						paths.forEach(
							pathDB => {
								let path = this.frontRepo.map_ID_Path.get(pathDB.ID)
								CopyPathDBToPath(pathDB, path!, this.frontRepo)
							}
						)

						// fill up front objects
						points.forEach(
							pointDB => {
								let point = this.frontRepo.map_ID_Point.get(pointDB.ID)
								CopyPointDBToPoint(pointDB, point!, this.frontRepo)
							}
						)

						// fill up front objects
						polygones.forEach(
							polygoneDB => {
								let polygone = this.frontRepo.map_ID_Polygone.get(polygoneDB.ID)
								CopyPolygoneDBToPolygone(polygoneDB, polygone!, this.frontRepo)
							}
						)

						// fill up front objects
						polylines.forEach(
							polylineDB => {
								let polyline = this.frontRepo.map_ID_Polyline.get(polylineDB.ID)
								CopyPolylineDBToPolyline(polylineDB, polyline!, this.frontRepo)
							}
						)

						// fill up front objects
						rects.forEach(
							rectDB => {
								let rect = this.frontRepo.map_ID_Rect.get(rectDB.ID)
								CopyRectDBToRect(rectDB, rect!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredpaths.forEach(
							rectanchoredpathDB => {
								let rectanchoredpath = this.frontRepo.map_ID_RectAnchoredPath.get(rectanchoredpathDB.ID)
								CopyRectAnchoredPathDBToRectAnchoredPath(rectanchoredpathDB, rectanchoredpath!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredrects.forEach(
							rectanchoredrectDB => {
								let rectanchoredrect = this.frontRepo.map_ID_RectAnchoredRect.get(rectanchoredrectDB.ID)
								CopyRectAnchoredRectDBToRectAnchoredRect(rectanchoredrectDB, rectanchoredrect!, this.frontRepo)
							}
						)

						// fill up front objects
						rectanchoredtexts.forEach(
							rectanchoredtextDB => {
								let rectanchoredtext = this.frontRepo.map_ID_RectAnchoredText.get(rectanchoredtextDB.ID)
								CopyRectAnchoredTextDBToRectAnchoredText(rectanchoredtextDB, rectanchoredtext!, this.frontRepo)
							}
						)

						// fill up front objects
						rectlinklinks.forEach(
							rectlinklinkDB => {
								let rectlinklink = this.frontRepo.map_ID_RectLinkLink.get(rectlinklinkDB.ID)
								CopyRectLinkLinkDBToRectLinkLink(rectlinklinkDB, rectlinklink!, this.frontRepo)
							}
						)

						// fill up front objects
						svgs.forEach(
							svgDB => {
								let svg = this.frontRepo.map_ID_SVG.get(svgDB.ID)
								CopySVGDBToSVG(svgDB, svg!, this.frontRepo)
							}
						)

						// fill up front objects
						texts.forEach(
							textDB => {
								let text = this.frontRepo.map_ID_Text.get(textDB.ID)
								CopyTextDBToText(textDB, text!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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
export function getTextUniqueID(id: number): number {
	return 107 * id
}
