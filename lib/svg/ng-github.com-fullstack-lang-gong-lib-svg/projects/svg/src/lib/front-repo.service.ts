// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AnimateAPI } from './animate-api'
import { Animate, CopyAnimateAPIToAnimate } from './animate'

import { CircleAPI } from './circle-api'
import { Circle, CopyCircleAPIToCircle } from './circle'

import { ConditionAPI } from './condition-api'
import { Condition, CopyConditionAPIToCondition } from './condition'

import { ControlPointAPI } from './controlpoint-api'
import { ControlPoint, CopyControlPointAPIToControlPoint } from './controlpoint'

import { EllipseAPI } from './ellipse-api'
import { Ellipse, CopyEllipseAPIToEllipse } from './ellipse'

import { FileToDownloadAPI } from './filetodownload-api'
import { FileToDownload, CopyFileToDownloadAPIToFileToDownload } from './filetodownload'

import { LayerAPI } from './layer-api'
import { Layer, CopyLayerAPIToLayer } from './layer'

import { LineAPI } from './line-api'
import { Line, CopyLineAPIToLine } from './line'

import { LinkAPI } from './link-api'
import { Link, CopyLinkAPIToLink } from './link'

import { LinkAnchoredPathAPI } from './linkanchoredpath-api'
import { LinkAnchoredPath, CopyLinkAnchoredPathAPIToLinkAnchoredPath } from './linkanchoredpath'

import { LinkAnchoredTextAPI } from './linkanchoredtext-api'
import { LinkAnchoredText, CopyLinkAnchoredTextAPIToLinkAnchoredText } from './linkanchoredtext'

import { PathAPI } from './path-api'
import { Path, CopyPathAPIToPath } from './path'

import { PointAPI } from './point-api'
import { Point, CopyPointAPIToPoint } from './point'

import { PolygoneAPI } from './polygone-api'
import { Polygone, CopyPolygoneAPIToPolygone } from './polygone'

import { PolylineAPI } from './polyline-api'
import { Polyline, CopyPolylineAPIToPolyline } from './polyline'

import { RectAPI } from './rect-api'
import { Rect, CopyRectAPIToRect } from './rect'

import { RectAnchoredPathAPI } from './rectanchoredpath-api'
import { RectAnchoredPath, CopyRectAnchoredPathAPIToRectAnchoredPath } from './rectanchoredpath'

import { RectAnchoredPngImageAPI } from './rectanchoredpngimage-api'
import { RectAnchoredPngImage, CopyRectAnchoredPngImageAPIToRectAnchoredPngImage } from './rectanchoredpngimage'

import { RectAnchoredRectAPI } from './rectanchoredrect-api'
import { RectAnchoredRect, CopyRectAnchoredRectAPIToRectAnchoredRect } from './rectanchoredrect'

import { RectAnchoredTextAPI } from './rectanchoredtext-api'
import { RectAnchoredText, CopyRectAnchoredTextAPIToRectAnchoredText } from './rectanchoredtext'

import { RectLinkLinkAPI } from './rectlinklink-api'
import { RectLinkLink, CopyRectLinkLinkAPIToRectLinkLink } from './rectlinklink'

import { SVGAPI } from './svg-api'
import { SVG, CopySVGAPIToSVG } from './svg'

import { SvgTextAPI } from './svgtext-api'
import { SvgText, CopySvgTextAPIToSvgText } from './svgtext'

import { TextAPI } from './text-api'
import { Text, CopyTextAPIToText } from './text'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/svg/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Animates = new Array<Animate>() // array of front instances
	map_ID_Animate = new Map<number, Animate>() // map of front instances

	array_Circles = new Array<Circle>() // array of front instances
	map_ID_Circle = new Map<number, Circle>() // map of front instances

	array_Conditions = new Array<Condition>() // array of front instances
	map_ID_Condition = new Map<number, Condition>() // map of front instances

	array_ControlPoints = new Array<ControlPoint>() // array of front instances
	map_ID_ControlPoint = new Map<number, ControlPoint>() // map of front instances

	array_Ellipses = new Array<Ellipse>() // array of front instances
	map_ID_Ellipse = new Map<number, Ellipse>() // map of front instances

	array_FileToDownloads = new Array<FileToDownload>() // array of front instances
	map_ID_FileToDownload = new Map<number, FileToDownload>() // map of front instances

	array_Layers = new Array<Layer>() // array of front instances
	map_ID_Layer = new Map<number, Layer>() // map of front instances

	array_Lines = new Array<Line>() // array of front instances
	map_ID_Line = new Map<number, Line>() // map of front instances

	array_Links = new Array<Link>() // array of front instances
	map_ID_Link = new Map<number, Link>() // map of front instances

	array_LinkAnchoredPaths = new Array<LinkAnchoredPath>() // array of front instances
	map_ID_LinkAnchoredPath = new Map<number, LinkAnchoredPath>() // map of front instances

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

	array_RectAnchoredPngImages = new Array<RectAnchoredPngImage>() // array of front instances
	map_ID_RectAnchoredPngImage = new Map<number, RectAnchoredPngImage>() // map of front instances

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
			case 'Condition':
				return this.array_Conditions as unknown as Array<Type>
			case 'ControlPoint':
				return this.array_ControlPoints as unknown as Array<Type>
			case 'Ellipse':
				return this.array_Ellipses as unknown as Array<Type>
			case 'FileToDownload':
				return this.array_FileToDownloads as unknown as Array<Type>
			case 'Layer':
				return this.array_Layers as unknown as Array<Type>
			case 'Line':
				return this.array_Lines as unknown as Array<Type>
			case 'Link':
				return this.array_Links as unknown as Array<Type>
			case 'LinkAnchoredPath':
				return this.array_LinkAnchoredPaths as unknown as Array<Type>
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
			case 'RectAnchoredPngImage':
				return this.array_RectAnchoredPngImages as unknown as Array<Type>
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
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Animate':
				return this.map_ID_Animate as unknown as Map<number, Type>
			case 'Circle':
				return this.map_ID_Circle as unknown as Map<number, Type>
			case 'Condition':
				return this.map_ID_Condition as unknown as Map<number, Type>
			case 'ControlPoint':
				return this.map_ID_ControlPoint as unknown as Map<number, Type>
			case 'Ellipse':
				return this.map_ID_Ellipse as unknown as Map<number, Type>
			case 'FileToDownload':
				return this.map_ID_FileToDownload as unknown as Map<number, Type>
			case 'Layer':
				return this.map_ID_Layer as unknown as Map<number, Type>
			case 'Line':
				return this.map_ID_Line as unknown as Map<number, Type>
			case 'Link':
				return this.map_ID_Link as unknown as Map<number, Type>
			case 'LinkAnchoredPath':
				return this.map_ID_LinkAnchoredPath as unknown as Map<number, Type>
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
			case 'RectAnchoredPngImage':
				return this.map_ID_RectAnchoredPngImage as unknown as Map<number, Type>
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
				throw new Error("Type not recognized")
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

	Name: string = ""
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

	Name: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient,
		private ngZone: NgZone,
	) { }

	//
	// pull returns the FrontRepo Observable
	//
	pull(Name: string = ""): Observable<FrontRepo> {
		this.Name = Name
		return of(this.frontRepo)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		// console.log("github.com/fullstack-lang/gong/lib/svg/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/svg/go; connectToWebSocket: returning existing connection")
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/svg/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/svg/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/svg/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
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
				frontRepo.array_Conditions = []
				frontRepo.map_ID_Condition.clear()

				backRepoData.ConditionAPIs.forEach(
					conditionAPI => {
						let condition = new Condition
						frontRepo.array_Conditions.push(condition)
						frontRepo.map_ID_Condition.set(conditionAPI.ID, condition)
					}
				)

				// init the arrays
				frontRepo.array_ControlPoints = []
				frontRepo.map_ID_ControlPoint.clear()

				backRepoData.ControlPointAPIs.forEach(
					controlpointAPI => {
						let controlpoint = new ControlPoint
						frontRepo.array_ControlPoints.push(controlpoint)
						frontRepo.map_ID_ControlPoint.set(controlpointAPI.ID, controlpoint)
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
				frontRepo.array_FileToDownloads = []
				frontRepo.map_ID_FileToDownload.clear()

				backRepoData.FileToDownloadAPIs.forEach(
					filetodownloadAPI => {
						let filetodownload = new FileToDownload
						frontRepo.array_FileToDownloads.push(filetodownload)
						frontRepo.map_ID_FileToDownload.set(filetodownloadAPI.ID, filetodownload)
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
				frontRepo.array_LinkAnchoredPaths = []
				frontRepo.map_ID_LinkAnchoredPath.clear()

				backRepoData.LinkAnchoredPathAPIs.forEach(
					linkanchoredpathAPI => {
						let linkanchoredpath = new LinkAnchoredPath
						frontRepo.array_LinkAnchoredPaths.push(linkanchoredpath)
						frontRepo.map_ID_LinkAnchoredPath.set(linkanchoredpathAPI.ID, linkanchoredpath)
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
				frontRepo.array_RectAnchoredPngImages = []
				frontRepo.map_ID_RectAnchoredPngImage.clear()

				backRepoData.RectAnchoredPngImageAPIs.forEach(
					rectanchoredpngimageAPI => {
						let rectanchoredpngimage = new RectAnchoredPngImage
						frontRepo.array_RectAnchoredPngImages.push(rectanchoredpngimage)
						frontRepo.map_ID_RectAnchoredPngImage.set(rectanchoredpngimageAPI.ID, rectanchoredpngimage)
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
				backRepoData.ConditionAPIs.forEach(
					conditionAPI => {
						let condition = frontRepo.map_ID_Condition.get(conditionAPI.ID)
						CopyConditionAPIToCondition(conditionAPI, condition!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ControlPointAPIs.forEach(
					controlpointAPI => {
						let controlpoint = frontRepo.map_ID_ControlPoint.get(controlpointAPI.ID)
						CopyControlPointAPIToControlPoint(controlpointAPI, controlpoint!, frontRepo)
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
				backRepoData.FileToDownloadAPIs.forEach(
					filetodownloadAPI => {
						let filetodownload = frontRepo.map_ID_FileToDownload.get(filetodownloadAPI.ID)
						CopyFileToDownloadAPIToFileToDownload(filetodownloadAPI, filetodownload!, frontRepo)
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
				backRepoData.LinkAnchoredPathAPIs.forEach(
					linkanchoredpathAPI => {
						let linkanchoredpath = frontRepo.map_ID_LinkAnchoredPath.get(linkanchoredpathAPI.ID)
						CopyLinkAnchoredPathAPIToLinkAnchoredPath(linkanchoredpathAPI, linkanchoredpath!, frontRepo)
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
				backRepoData.RectAnchoredPngImageAPIs.forEach(
					rectanchoredpngimageAPI => {
						let rectanchoredpngimage = frontRepo.map_ID_RectAnchoredPngImage.get(rectanchoredpngimageAPI.ID)
						CopyRectAnchoredPngImageAPIToRectAnchoredPngImage(rectanchoredpngimageAPI, rectanchoredpngimage!, frontRepo)
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


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// Offline mode handling: Listen to the global event
			if (isOfflineMode) {
				console.log("github.com/fullstack-lang/gong/lib/svg/go; Offline mode detected. Skipping WebSocket connection.")

				window.addEventListener('message', (event) => {
					if (event.data && event.data.type === 'STAGE_UPDATE') {
						console.log("github.com/fullstack-lang/gong/lib/svg/go; Received STAGE_UPDATE message.")
						processData(JSON.stringify(event.data.data))
					}
				})

				return () => {
					console.log("github.com/fullstack-lang/gong/lib/svg/go; Cleaning up offline message listener.")
				}
			}

			// Fallback: If not offline, create normal WebSocket
			const attemptConnection = () => {
				// Offline check inside attemptConnection: if window.openWasmSocket is available, use it!
				if (typeof window !== 'undefined' && (window as any).openWasmSocket) {
					(window as any).openWasmSocket('github.com/fullstack-lang/gong/lib/svg/go', Name, (data: any) => {
						processData(data)
					})
					return
				}

				if (isOfflineMode && retryCount > 0) {
					console.log("github.com/fullstack-lang/gong/lib/svg/go; Waiting for wasm socket provider...")
					setTimeout(() => attemptConnection(), 100)
					return
				}

				if (isOfflineMode) {
					console.error("github.com/fullstack-lang/gong/lib/svg/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.")
					return
				}

				socket = new WebSocket(url)

				socket.onopen = () => {
					// console.log("github.com/fullstack-lang/gong/lib/svg/go; WebSocket connection opened successfully:", url)
				}

				socket.onmessage = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/svg/go; WebSocket message received:", event.data)
					processData(event.data)
				}

				socket.onerror = (error) => {
					console.error("github.com/fullstack-lang/gong/lib/svg/go WebSocket: onerror", error)
					observer.error(error)
				}

				socket.onclose = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/svg/go; WebSocket connection closed:", event)
					observer.complete()
				}
			}

			let retryCount = 10
			attemptConnection()

			return () => {
				// console.log("github.com/fullstack-lang/gong/lib/svg/go; Cleaning up WebSocket connection")
				if (socket) {
					socket.close()
				}
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getAnimateUniqueID(id: number): number {
	return 31 * id
}
export function getCircleUniqueID(id: number): number {
	return 37 * id
}
export function getConditionUniqueID(id: number): number {
	return 41 * id
}
export function getControlPointUniqueID(id: number): number {
	return 43 * id
}
export function getEllipseUniqueID(id: number): number {
	return 47 * id
}
export function getFileToDownloadUniqueID(id: number): number {
	return 53 * id
}
export function getLayerUniqueID(id: number): number {
	return 59 * id
}
export function getLineUniqueID(id: number): number {
	return 61 * id
}
export function getLinkUniqueID(id: number): number {
	return 67 * id
}
export function getLinkAnchoredPathUniqueID(id: number): number {
	return 71 * id
}
export function getLinkAnchoredTextUniqueID(id: number): number {
	return 73 * id
}
export function getPathUniqueID(id: number): number {
	return 79 * id
}
export function getPointUniqueID(id: number): number {
	return 83 * id
}
export function getPolygoneUniqueID(id: number): number {
	return 89 * id
}
export function getPolylineUniqueID(id: number): number {
	return 97 * id
}
export function getRectUniqueID(id: number): number {
	return 101 * id
}
export function getRectAnchoredPathUniqueID(id: number): number {
	return 103 * id
}
export function getRectAnchoredPngImageUniqueID(id: number): number {
	return 107 * id
}
export function getRectAnchoredRectUniqueID(id: number): number {
	return 109 * id
}
export function getRectAnchoredTextUniqueID(id: number): number {
	return 113 * id
}
export function getRectLinkLinkUniqueID(id: number): number {
	return 127 * id
}
export function getSVGUniqueID(id: number): number {
	return 131 * id
}
export function getSvgTextUniqueID(id: number): number {
	return 137 * id
}
export function getTextUniqueID(id: number): number {
	return 139 * id
}
