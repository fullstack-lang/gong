import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { AnimateDB } from './animate-db'
import { AnimateService } from './animate.service'

import { CircleDB } from './circle-db'
import { CircleService } from './circle.service'

import { EllipseDB } from './ellipse-db'
import { EllipseService } from './ellipse.service'

import { LayerDB } from './layer-db'
import { LayerService } from './layer.service'

import { LineDB } from './line-db'
import { LineService } from './line.service'

import { LinkDB } from './link-db'
import { LinkService } from './link.service'

import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { LinkAnchoredTextService } from './linkanchoredtext.service'

import { PathDB } from './path-db'
import { PathService } from './path.service'

import { PointDB } from './point-db'
import { PointService } from './point.service'

import { PolygoneDB } from './polygone-db'
import { PolygoneService } from './polygone.service'

import { PolylineDB } from './polyline-db'
import { PolylineService } from './polyline.service'

import { RectDB } from './rect-db'
import { RectService } from './rect.service'

import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { RectAnchoredRectService } from './rectanchoredrect.service'

import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { RectAnchoredTextService } from './rectanchoredtext.service'

import { RectLinkLinkDB } from './rectlinklink-db'
import { RectLinkLinkService } from './rectlinklink.service'

import { SVGDB } from './svg-db'
import { SVGService } from './svg.service'

import { TextDB } from './text-db'
import { TextService } from './text.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Animates_array = new Array<AnimateDB>(); // array of repo instances
  Animates = new Map<number, AnimateDB>(); // map of repo instances
  Animates_batch = new Map<number, AnimateDB>(); // same but only in last GET (for finding repo instances to delete)
  Circles_array = new Array<CircleDB>(); // array of repo instances
  Circles = new Map<number, CircleDB>(); // map of repo instances
  Circles_batch = new Map<number, CircleDB>(); // same but only in last GET (for finding repo instances to delete)
  Ellipses_array = new Array<EllipseDB>(); // array of repo instances
  Ellipses = new Map<number, EllipseDB>(); // map of repo instances
  Ellipses_batch = new Map<number, EllipseDB>(); // same but only in last GET (for finding repo instances to delete)
  Layers_array = new Array<LayerDB>(); // array of repo instances
  Layers = new Map<number, LayerDB>(); // map of repo instances
  Layers_batch = new Map<number, LayerDB>(); // same but only in last GET (for finding repo instances to delete)
  Lines_array = new Array<LineDB>(); // array of repo instances
  Lines = new Map<number, LineDB>(); // map of repo instances
  Lines_batch = new Map<number, LineDB>(); // same but only in last GET (for finding repo instances to delete)
  Links_array = new Array<LinkDB>(); // array of repo instances
  Links = new Map<number, LinkDB>(); // map of repo instances
  Links_batch = new Map<number, LinkDB>(); // same but only in last GET (for finding repo instances to delete)
  LinkAnchoredTexts_array = new Array<LinkAnchoredTextDB>(); // array of repo instances
  LinkAnchoredTexts = new Map<number, LinkAnchoredTextDB>(); // map of repo instances
  LinkAnchoredTexts_batch = new Map<number, LinkAnchoredTextDB>(); // same but only in last GET (for finding repo instances to delete)
  Paths_array = new Array<PathDB>(); // array of repo instances
  Paths = new Map<number, PathDB>(); // map of repo instances
  Paths_batch = new Map<number, PathDB>(); // same but only in last GET (for finding repo instances to delete)
  Points_array = new Array<PointDB>(); // array of repo instances
  Points = new Map<number, PointDB>(); // map of repo instances
  Points_batch = new Map<number, PointDB>(); // same but only in last GET (for finding repo instances to delete)
  Polygones_array = new Array<PolygoneDB>(); // array of repo instances
  Polygones = new Map<number, PolygoneDB>(); // map of repo instances
  Polygones_batch = new Map<number, PolygoneDB>(); // same but only in last GET (for finding repo instances to delete)
  Polylines_array = new Array<PolylineDB>(); // array of repo instances
  Polylines = new Map<number, PolylineDB>(); // map of repo instances
  Polylines_batch = new Map<number, PolylineDB>(); // same but only in last GET (for finding repo instances to delete)
  Rects_array = new Array<RectDB>(); // array of repo instances
  Rects = new Map<number, RectDB>(); // map of repo instances
  Rects_batch = new Map<number, RectDB>(); // same but only in last GET (for finding repo instances to delete)
  RectAnchoredRects_array = new Array<RectAnchoredRectDB>(); // array of repo instances
  RectAnchoredRects = new Map<number, RectAnchoredRectDB>(); // map of repo instances
  RectAnchoredRects_batch = new Map<number, RectAnchoredRectDB>(); // same but only in last GET (for finding repo instances to delete)
  RectAnchoredTexts_array = new Array<RectAnchoredTextDB>(); // array of repo instances
  RectAnchoredTexts = new Map<number, RectAnchoredTextDB>(); // map of repo instances
  RectAnchoredTexts_batch = new Map<number, RectAnchoredTextDB>(); // same but only in last GET (for finding repo instances to delete)
  RectLinkLinks_array = new Array<RectLinkLinkDB>(); // array of repo instances
  RectLinkLinks = new Map<number, RectLinkLinkDB>(); // map of repo instances
  RectLinkLinks_batch = new Map<number, RectLinkLinkDB>(); // same but only in last GET (for finding repo instances to delete)
  SVGs_array = new Array<SVGDB>(); // array of repo instances
  SVGs = new Map<number, SVGDB>(); // map of repo instances
  SVGs_batch = new Map<number, SVGDB>(); // same but only in last GET (for finding repo instances to delete)
  Texts_array = new Array<TextDB>(); // array of repo instances
  Texts = new Map<number, TextDB>(); // map of repo instances
  Texts_batch = new Map<number, TextDB>(); // same but only in last GET (for finding repo instances to delete)
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
  SourceStruct: string = ""  // The "Aclass"
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
  observableFrontRepo: [ // insertion point sub template 
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
    Observable<RectAnchoredRectDB[]>,
    Observable<RectAnchoredTextDB[]>,
    Observable<RectLinkLinkDB[]>,
    Observable<SVGDB[]>,
    Observable<TextDB[]>,
  ] = [ // insertion point sub template
      this.animateService.getAnimates(this.GONG__StackPath),
      this.circleService.getCircles(this.GONG__StackPath),
      this.ellipseService.getEllipses(this.GONG__StackPath),
      this.layerService.getLayers(this.GONG__StackPath),
      this.lineService.getLines(this.GONG__StackPath),
      this.linkService.getLinks(this.GONG__StackPath),
      this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath),
      this.pathService.getPaths(this.GONG__StackPath),
      this.pointService.getPoints(this.GONG__StackPath),
      this.polygoneService.getPolygones(this.GONG__StackPath),
      this.polylineService.getPolylines(this.GONG__StackPath),
      this.rectService.getRects(this.GONG__StackPath),
      this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath),
      this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath),
      this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath),
      this.svgService.getSVGs(this.GONG__StackPath),
      this.textService.getTexts(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.animateService.getAnimates(this.GONG__StackPath),
      this.circleService.getCircles(this.GONG__StackPath),
      this.ellipseService.getEllipses(this.GONG__StackPath),
      this.layerService.getLayers(this.GONG__StackPath),
      this.lineService.getLines(this.GONG__StackPath),
      this.linkService.getLinks(this.GONG__StackPath),
      this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath),
      this.pathService.getPaths(this.GONG__StackPath),
      this.pointService.getPoints(this.GONG__StackPath),
      this.polygoneService.getPolygones(this.GONG__StackPath),
      this.polylineService.getPolylines(this.GONG__StackPath),
      this.rectService.getRects(this.GONG__StackPath),
      this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath),
      this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath),
      this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath),
      this.svgService.getSVGs(this.GONG__StackPath),
      this.textService.getTexts(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
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
            rectanchoredrects_,
            rectanchoredtexts_,
            rectlinklinks_,
            svgs_,
            texts_,
          ]) => {
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
            // init the array
            this.frontRepo.Animates_array = animates

            // clear the map that counts Animate in the GET
            this.frontRepo.Animates_batch.clear()

            animates.forEach(
              animate => {
                this.frontRepo.Animates.set(animate.ID, animate)
                this.frontRepo.Animates_batch.set(animate.ID, animate)
              }
            )

            // clear animates that are absent from the batch
            this.frontRepo.Animates.forEach(
              animate => {
                if (this.frontRepo.Animates_batch.get(animate.ID) == undefined) {
                  this.frontRepo.Animates.delete(animate.ID)
                }
              }
            )

            // sort Animates_array array
            this.frontRepo.Animates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Circles_array = circles

            // clear the map that counts Circle in the GET
            this.frontRepo.Circles_batch.clear()

            circles.forEach(
              circle => {
                this.frontRepo.Circles.set(circle.ID, circle)
                this.frontRepo.Circles_batch.set(circle.ID, circle)
              }
            )

            // clear circles that are absent from the batch
            this.frontRepo.Circles.forEach(
              circle => {
                if (this.frontRepo.Circles_batch.get(circle.ID) == undefined) {
                  this.frontRepo.Circles.delete(circle.ID)
                }
              }
            )

            // sort Circles_array array
            this.frontRepo.Circles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            this.frontRepo.Ellipses_batch.clear()

            ellipses.forEach(
              ellipse => {
                this.frontRepo.Ellipses.set(ellipse.ID, ellipse)
                this.frontRepo.Ellipses_batch.set(ellipse.ID, ellipse)
              }
            )

            // clear ellipses that are absent from the batch
            this.frontRepo.Ellipses.forEach(
              ellipse => {
                if (this.frontRepo.Ellipses_batch.get(ellipse.ID) == undefined) {
                  this.frontRepo.Ellipses.delete(ellipse.ID)
                }
              }
            )

            // sort Ellipses_array array
            this.frontRepo.Ellipses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Layers_array = layers

            // clear the map that counts Layer in the GET
            this.frontRepo.Layers_batch.clear()

            layers.forEach(
              layer => {
                this.frontRepo.Layers.set(layer.ID, layer)
                this.frontRepo.Layers_batch.set(layer.ID, layer)
              }
            )

            // clear layers that are absent from the batch
            this.frontRepo.Layers.forEach(
              layer => {
                if (this.frontRepo.Layers_batch.get(layer.ID) == undefined) {
                  this.frontRepo.Layers.delete(layer.ID)
                }
              }
            )

            // sort Layers_array array
            this.frontRepo.Layers_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Lines_array = lines

            // clear the map that counts Line in the GET
            this.frontRepo.Lines_batch.clear()

            lines.forEach(
              line => {
                this.frontRepo.Lines.set(line.ID, line)
                this.frontRepo.Lines_batch.set(line.ID, line)
              }
            )

            // clear lines that are absent from the batch
            this.frontRepo.Lines.forEach(
              line => {
                if (this.frontRepo.Lines_batch.get(line.ID) == undefined) {
                  this.frontRepo.Lines.delete(line.ID)
                }
              }
            )

            // sort Lines_array array
            this.frontRepo.Lines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            links.forEach(
              link => {
                this.frontRepo.Links.set(link.ID, link)
                this.frontRepo.Links_batch.set(link.ID, link)
              }
            )

            // clear links that are absent from the batch
            this.frontRepo.Links.forEach(
              link => {
                if (this.frontRepo.Links_batch.get(link.ID) == undefined) {
                  this.frontRepo.Links.delete(link.ID)
                }
              }
            )

            // sort Links_array array
            this.frontRepo.Links_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.LinkAnchoredTexts_array = linkanchoredtexts

            // clear the map that counts LinkAnchoredText in the GET
            this.frontRepo.LinkAnchoredTexts_batch.clear()

            linkanchoredtexts.forEach(
              linkanchoredtext => {
                this.frontRepo.LinkAnchoredTexts.set(linkanchoredtext.ID, linkanchoredtext)
                this.frontRepo.LinkAnchoredTexts_batch.set(linkanchoredtext.ID, linkanchoredtext)
              }
            )

            // clear linkanchoredtexts that are absent from the batch
            this.frontRepo.LinkAnchoredTexts.forEach(
              linkanchoredtext => {
                if (this.frontRepo.LinkAnchoredTexts_batch.get(linkanchoredtext.ID) == undefined) {
                  this.frontRepo.LinkAnchoredTexts.delete(linkanchoredtext.ID)
                }
              }
            )

            // sort LinkAnchoredTexts_array array
            this.frontRepo.LinkAnchoredTexts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Paths_array = paths

            // clear the map that counts Path in the GET
            this.frontRepo.Paths_batch.clear()

            paths.forEach(
              path => {
                this.frontRepo.Paths.set(path.ID, path)
                this.frontRepo.Paths_batch.set(path.ID, path)
              }
            )

            // clear paths that are absent from the batch
            this.frontRepo.Paths.forEach(
              path => {
                if (this.frontRepo.Paths_batch.get(path.ID) == undefined) {
                  this.frontRepo.Paths.delete(path.ID)
                }
              }
            )

            // sort Paths_array array
            this.frontRepo.Paths_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Points_array = points

            // clear the map that counts Point in the GET
            this.frontRepo.Points_batch.clear()

            points.forEach(
              point => {
                this.frontRepo.Points.set(point.ID, point)
                this.frontRepo.Points_batch.set(point.ID, point)
              }
            )

            // clear points that are absent from the batch
            this.frontRepo.Points.forEach(
              point => {
                if (this.frontRepo.Points_batch.get(point.ID) == undefined) {
                  this.frontRepo.Points.delete(point.ID)
                }
              }
            )

            // sort Points_array array
            this.frontRepo.Points_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            this.frontRepo.Polygones_batch.clear()

            polygones.forEach(
              polygone => {
                this.frontRepo.Polygones.set(polygone.ID, polygone)
                this.frontRepo.Polygones_batch.set(polygone.ID, polygone)
              }
            )

            // clear polygones that are absent from the batch
            this.frontRepo.Polygones.forEach(
              polygone => {
                if (this.frontRepo.Polygones_batch.get(polygone.ID) == undefined) {
                  this.frontRepo.Polygones.delete(polygone.ID)
                }
              }
            )

            // sort Polygones_array array
            this.frontRepo.Polygones_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            this.frontRepo.Polylines_batch.clear()

            polylines.forEach(
              polyline => {
                this.frontRepo.Polylines.set(polyline.ID, polyline)
                this.frontRepo.Polylines_batch.set(polyline.ID, polyline)
              }
            )

            // clear polylines that are absent from the batch
            this.frontRepo.Polylines.forEach(
              polyline => {
                if (this.frontRepo.Polylines_batch.get(polyline.ID) == undefined) {
                  this.frontRepo.Polylines.delete(polyline.ID)
                }
              }
            )

            // sort Polylines_array array
            this.frontRepo.Polylines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Rects_array = rects

            // clear the map that counts Rect in the GET
            this.frontRepo.Rects_batch.clear()

            rects.forEach(
              rect => {
                this.frontRepo.Rects.set(rect.ID, rect)
                this.frontRepo.Rects_batch.set(rect.ID, rect)
              }
            )

            // clear rects that are absent from the batch
            this.frontRepo.Rects.forEach(
              rect => {
                if (this.frontRepo.Rects_batch.get(rect.ID) == undefined) {
                  this.frontRepo.Rects.delete(rect.ID)
                }
              }
            )

            // sort Rects_array array
            this.frontRepo.Rects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectAnchoredRects_array = rectanchoredrects

            // clear the map that counts RectAnchoredRect in the GET
            this.frontRepo.RectAnchoredRects_batch.clear()

            rectanchoredrects.forEach(
              rectanchoredrect => {
                this.frontRepo.RectAnchoredRects.set(rectanchoredrect.ID, rectanchoredrect)
                this.frontRepo.RectAnchoredRects_batch.set(rectanchoredrect.ID, rectanchoredrect)
              }
            )

            // clear rectanchoredrects that are absent from the batch
            this.frontRepo.RectAnchoredRects.forEach(
              rectanchoredrect => {
                if (this.frontRepo.RectAnchoredRects_batch.get(rectanchoredrect.ID) == undefined) {
                  this.frontRepo.RectAnchoredRects.delete(rectanchoredrect.ID)
                }
              }
            )

            // sort RectAnchoredRects_array array
            this.frontRepo.RectAnchoredRects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectAnchoredTexts_array = rectanchoredtexts

            // clear the map that counts RectAnchoredText in the GET
            this.frontRepo.RectAnchoredTexts_batch.clear()

            rectanchoredtexts.forEach(
              rectanchoredtext => {
                this.frontRepo.RectAnchoredTexts.set(rectanchoredtext.ID, rectanchoredtext)
                this.frontRepo.RectAnchoredTexts_batch.set(rectanchoredtext.ID, rectanchoredtext)
              }
            )

            // clear rectanchoredtexts that are absent from the batch
            this.frontRepo.RectAnchoredTexts.forEach(
              rectanchoredtext => {
                if (this.frontRepo.RectAnchoredTexts_batch.get(rectanchoredtext.ID) == undefined) {
                  this.frontRepo.RectAnchoredTexts.delete(rectanchoredtext.ID)
                }
              }
            )

            // sort RectAnchoredTexts_array array
            this.frontRepo.RectAnchoredTexts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectLinkLinks_array = rectlinklinks

            // clear the map that counts RectLinkLink in the GET
            this.frontRepo.RectLinkLinks_batch.clear()

            rectlinklinks.forEach(
              rectlinklink => {
                this.frontRepo.RectLinkLinks.set(rectlinklink.ID, rectlinklink)
                this.frontRepo.RectLinkLinks_batch.set(rectlinklink.ID, rectlinklink)
              }
            )

            // clear rectlinklinks that are absent from the batch
            this.frontRepo.RectLinkLinks.forEach(
              rectlinklink => {
                if (this.frontRepo.RectLinkLinks_batch.get(rectlinklink.ID) == undefined) {
                  this.frontRepo.RectLinkLinks.delete(rectlinklink.ID)
                }
              }
            )

            // sort RectLinkLinks_array array
            this.frontRepo.RectLinkLinks_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            this.frontRepo.SVGs_batch.clear()

            svgs.forEach(
              svg => {
                this.frontRepo.SVGs.set(svg.ID, svg)
                this.frontRepo.SVGs_batch.set(svg.ID, svg)
              }
            )

            // clear svgs that are absent from the batch
            this.frontRepo.SVGs.forEach(
              svg => {
                if (this.frontRepo.SVGs_batch.get(svg.ID) == undefined) {
                  this.frontRepo.SVGs.delete(svg.ID)
                }
              }
            )

            // sort SVGs_array array
            this.frontRepo.SVGs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            texts.forEach(
              text => {
                this.frontRepo.Texts.set(text.ID, text)
                this.frontRepo.Texts_batch.set(text.ID, text)
              }
            )

            // clear texts that are absent from the batch
            this.frontRepo.Texts.forEach(
              text => {
                if (this.frontRepo.Texts_batch.get(text.ID) == undefined) {
                  this.frontRepo.Texts.delete(text.ID)
                }
              }
            )

            // sort Texts_array array
            this.frontRepo.Texts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            animates.forEach(
              animate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Circle.Animations redeeming
                {
                  let _circle = this.frontRepo.Circles.get(animate.Circle_AnimationsDBID.Int64)
                  if (_circle) {
                    if (_circle.Animations == undefined) {
                      _circle.Animations = new Array<AnimateDB>()
                    }
                    _circle.Animations.push(animate)
                    if (animate.Circle_Animations_reverse == undefined) {
                      animate.Circle_Animations_reverse = _circle
                    }
                  }
                }
                // insertion point for slice of pointer field Ellipse.Animates redeeming
                {
                  let _ellipse = this.frontRepo.Ellipses.get(animate.Ellipse_AnimatesDBID.Int64)
                  if (_ellipse) {
                    if (_ellipse.Animates == undefined) {
                      _ellipse.Animates = new Array<AnimateDB>()
                    }
                    _ellipse.Animates.push(animate)
                    if (animate.Ellipse_Animates_reverse == undefined) {
                      animate.Ellipse_Animates_reverse = _ellipse
                    }
                  }
                }
                // insertion point for slice of pointer field Line.Animates redeeming
                {
                  let _line = this.frontRepo.Lines.get(animate.Line_AnimatesDBID.Int64)
                  if (_line) {
                    if (_line.Animates == undefined) {
                      _line.Animates = new Array<AnimateDB>()
                    }
                    _line.Animates.push(animate)
                    if (animate.Line_Animates_reverse == undefined) {
                      animate.Line_Animates_reverse = _line
                    }
                  }
                }
                // insertion point for slice of pointer field LinkAnchoredText.Animates redeeming
                {
                  let _linkanchoredtext = this.frontRepo.LinkAnchoredTexts.get(animate.LinkAnchoredText_AnimatesDBID.Int64)
                  if (_linkanchoredtext) {
                    if (_linkanchoredtext.Animates == undefined) {
                      _linkanchoredtext.Animates = new Array<AnimateDB>()
                    }
                    _linkanchoredtext.Animates.push(animate)
                    if (animate.LinkAnchoredText_Animates_reverse == undefined) {
                      animate.LinkAnchoredText_Animates_reverse = _linkanchoredtext
                    }
                  }
                }
                // insertion point for slice of pointer field Path.Animates redeeming
                {
                  let _path = this.frontRepo.Paths.get(animate.Path_AnimatesDBID.Int64)
                  if (_path) {
                    if (_path.Animates == undefined) {
                      _path.Animates = new Array<AnimateDB>()
                    }
                    _path.Animates.push(animate)
                    if (animate.Path_Animates_reverse == undefined) {
                      animate.Path_Animates_reverse = _path
                    }
                  }
                }
                // insertion point for slice of pointer field Polygone.Animates redeeming
                {
                  let _polygone = this.frontRepo.Polygones.get(animate.Polygone_AnimatesDBID.Int64)
                  if (_polygone) {
                    if (_polygone.Animates == undefined) {
                      _polygone.Animates = new Array<AnimateDB>()
                    }
                    _polygone.Animates.push(animate)
                    if (animate.Polygone_Animates_reverse == undefined) {
                      animate.Polygone_Animates_reverse = _polygone
                    }
                  }
                }
                // insertion point for slice of pointer field Polyline.Animates redeeming
                {
                  let _polyline = this.frontRepo.Polylines.get(animate.Polyline_AnimatesDBID.Int64)
                  if (_polyline) {
                    if (_polyline.Animates == undefined) {
                      _polyline.Animates = new Array<AnimateDB>()
                    }
                    _polyline.Animates.push(animate)
                    if (animate.Polyline_Animates_reverse == undefined) {
                      animate.Polyline_Animates_reverse = _polyline
                    }
                  }
                }
                // insertion point for slice of pointer field Rect.Animations redeeming
                {
                  let _rect = this.frontRepo.Rects.get(animate.Rect_AnimationsDBID.Int64)
                  if (_rect) {
                    if (_rect.Animations == undefined) {
                      _rect.Animations = new Array<AnimateDB>()
                    }
                    _rect.Animations.push(animate)
                    if (animate.Rect_Animations_reverse == undefined) {
                      animate.Rect_Animations_reverse = _rect
                    }
                  }
                }
                // insertion point for slice of pointer field RectAnchoredText.Animates redeeming
                {
                  let _rectanchoredtext = this.frontRepo.RectAnchoredTexts.get(animate.RectAnchoredText_AnimatesDBID.Int64)
                  if (_rectanchoredtext) {
                    if (_rectanchoredtext.Animates == undefined) {
                      _rectanchoredtext.Animates = new Array<AnimateDB>()
                    }
                    _rectanchoredtext.Animates.push(animate)
                    if (animate.RectAnchoredText_Animates_reverse == undefined) {
                      animate.RectAnchoredText_Animates_reverse = _rectanchoredtext
                    }
                  }
                }
                // insertion point for slice of pointer field Text.Animates redeeming
                {
                  let _text = this.frontRepo.Texts.get(animate.Text_AnimatesDBID.Int64)
                  if (_text) {
                    if (_text.Animates == undefined) {
                      _text.Animates = new Array<AnimateDB>()
                    }
                    _text.Animates.push(animate)
                    if (animate.Text_Animates_reverse == undefined) {
                      animate.Text_Animates_reverse = _text
                    }
                  }
                }
              }
            )
            circles.forEach(
              circle => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Circles redeeming
                {
                  let _layer = this.frontRepo.Layers.get(circle.Layer_CirclesDBID.Int64)
                  if (_layer) {
                    if (_layer.Circles == undefined) {
                      _layer.Circles = new Array<CircleDB>()
                    }
                    _layer.Circles.push(circle)
                    if (circle.Layer_Circles_reverse == undefined) {
                      circle.Layer_Circles_reverse = _layer
                    }
                  }
                }
              }
            )
            ellipses.forEach(
              ellipse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Ellipses redeeming
                {
                  let _layer = this.frontRepo.Layers.get(ellipse.Layer_EllipsesDBID.Int64)
                  if (_layer) {
                    if (_layer.Ellipses == undefined) {
                      _layer.Ellipses = new Array<EllipseDB>()
                    }
                    _layer.Ellipses.push(ellipse)
                    if (ellipse.Layer_Ellipses_reverse == undefined) {
                      ellipse.Layer_Ellipses_reverse = _layer
                    }
                  }
                }
              }
            )
            layers.forEach(
              layer => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Layers redeeming
                {
                  let _svg = this.frontRepo.SVGs.get(layer.SVG_LayersDBID.Int64)
                  if (_svg) {
                    if (_svg.Layers == undefined) {
                      _svg.Layers = new Array<LayerDB>()
                    }
                    _svg.Layers.push(layer)
                    if (layer.SVG_Layers_reverse == undefined) {
                      layer.SVG_Layers_reverse = _svg
                    }
                  }
                }
              }
            )
            lines.forEach(
              line => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Lines redeeming
                {
                  let _layer = this.frontRepo.Layers.get(line.Layer_LinesDBID.Int64)
                  if (_layer) {
                    if (_layer.Lines == undefined) {
                      _layer.Lines = new Array<LineDB>()
                    }
                    _layer.Lines.push(line)
                    if (line.Layer_Lines_reverse == undefined) {
                      line.Layer_Lines_reverse = _layer
                    }
                  }
                }
              }
            )
            links.forEach(
              link => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.StartID.Int64)
                  if (_rect) {
                    link.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.EndID.Int64)
                  if (_rect) {
                    link.End = _rect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Links redeeming
                {
                  let _layer = this.frontRepo.Layers.get(link.Layer_LinksDBID.Int64)
                  if (_layer) {
                    if (_layer.Links == undefined) {
                      _layer.Links = new Array<LinkDB>()
                    }
                    _layer.Links.push(link)
                    if (link.Layer_Links_reverse == undefined) {
                      link.Layer_Links_reverse = _layer
                    }
                  }
                }
              }
            )
            linkanchoredtexts.forEach(
              linkanchoredtext => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Link.TextAtArrowEnd redeeming
                {
                  let _link = this.frontRepo.Links.get(linkanchoredtext.Link_TextAtArrowEndDBID.Int64)
                  if (_link) {
                    if (_link.TextAtArrowEnd == undefined) {
                      _link.TextAtArrowEnd = new Array<LinkAnchoredTextDB>()
                    }
                    _link.TextAtArrowEnd.push(linkanchoredtext)
                    if (linkanchoredtext.Link_TextAtArrowEnd_reverse == undefined) {
                      linkanchoredtext.Link_TextAtArrowEnd_reverse = _link
                    }
                  }
                }
                // insertion point for slice of pointer field Link.TextAtArrowStart redeeming
                {
                  let _link = this.frontRepo.Links.get(linkanchoredtext.Link_TextAtArrowStartDBID.Int64)
                  if (_link) {
                    if (_link.TextAtArrowStart == undefined) {
                      _link.TextAtArrowStart = new Array<LinkAnchoredTextDB>()
                    }
                    _link.TextAtArrowStart.push(linkanchoredtext)
                    if (linkanchoredtext.Link_TextAtArrowStart_reverse == undefined) {
                      linkanchoredtext.Link_TextAtArrowStart_reverse = _link
                    }
                  }
                }
              }
            )
            paths.forEach(
              path => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Paths redeeming
                {
                  let _layer = this.frontRepo.Layers.get(path.Layer_PathsDBID.Int64)
                  if (_layer) {
                    if (_layer.Paths == undefined) {
                      _layer.Paths = new Array<PathDB>()
                    }
                    _layer.Paths.push(path)
                    if (path.Layer_Paths_reverse == undefined) {
                      path.Layer_Paths_reverse = _layer
                    }
                  }
                }
              }
            )
            points.forEach(
              point => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Link.ControlPoints redeeming
                {
                  let _link = this.frontRepo.Links.get(point.Link_ControlPointsDBID.Int64)
                  if (_link) {
                    if (_link.ControlPoints == undefined) {
                      _link.ControlPoints = new Array<PointDB>()
                    }
                    _link.ControlPoints.push(point)
                    if (point.Link_ControlPoints_reverse == undefined) {
                      point.Link_ControlPoints_reverse = _link
                    }
                  }
                }
              }
            )
            polygones.forEach(
              polygone => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Polygones redeeming
                {
                  let _layer = this.frontRepo.Layers.get(polygone.Layer_PolygonesDBID.Int64)
                  if (_layer) {
                    if (_layer.Polygones == undefined) {
                      _layer.Polygones = new Array<PolygoneDB>()
                    }
                    _layer.Polygones.push(polygone)
                    if (polygone.Layer_Polygones_reverse == undefined) {
                      polygone.Layer_Polygones_reverse = _layer
                    }
                  }
                }
              }
            )
            polylines.forEach(
              polyline => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Polylines redeeming
                {
                  let _layer = this.frontRepo.Layers.get(polyline.Layer_PolylinesDBID.Int64)
                  if (_layer) {
                    if (_layer.Polylines == undefined) {
                      _layer.Polylines = new Array<PolylineDB>()
                    }
                    _layer.Polylines.push(polyline)
                    if (polyline.Layer_Polylines_reverse == undefined) {
                      polyline.Layer_Polylines_reverse = _layer
                    }
                  }
                }
              }
            )
            rects.forEach(
              rect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Rects redeeming
                {
                  let _layer = this.frontRepo.Layers.get(rect.Layer_RectsDBID.Int64)
                  if (_layer) {
                    if (_layer.Rects == undefined) {
                      _layer.Rects = new Array<RectDB>()
                    }
                    _layer.Rects.push(rect)
                    if (rect.Layer_Rects_reverse == undefined) {
                      rect.Layer_Rects_reverse = _layer
                    }
                  }
                }
              }
            )
            rectanchoredrects.forEach(
              rectanchoredrect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Rect.RectAnchoredRects redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectanchoredrect.Rect_RectAnchoredRectsDBID.Int64)
                  if (_rect) {
                    if (_rect.RectAnchoredRects == undefined) {
                      _rect.RectAnchoredRects = new Array<RectAnchoredRectDB>()
                    }
                    _rect.RectAnchoredRects.push(rectanchoredrect)
                    if (rectanchoredrect.Rect_RectAnchoredRects_reverse == undefined) {
                      rectanchoredrect.Rect_RectAnchoredRects_reverse = _rect
                    }
                  }
                }
              }
            )
            rectanchoredtexts.forEach(
              rectanchoredtext => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Rect.RectAnchoredTexts redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectanchoredtext.Rect_RectAnchoredTextsDBID.Int64)
                  if (_rect) {
                    if (_rect.RectAnchoredTexts == undefined) {
                      _rect.RectAnchoredTexts = new Array<RectAnchoredTextDB>()
                    }
                    _rect.RectAnchoredTexts.push(rectanchoredtext)
                    if (rectanchoredtext.Rect_RectAnchoredTexts_reverse == undefined) {
                      rectanchoredtext.Rect_RectAnchoredTexts_reverse = _rect
                    }
                  }
                }
              }
            )
            rectlinklinks.forEach(
              rectlinklink => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectlinklink.StartID.Int64)
                  if (_rect) {
                    rectlinklink.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _link = this.frontRepo.Links.get(rectlinklink.EndID.Int64)
                  if (_link) {
                    rectlinklink.End = _link
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.RectLinkLinks redeeming
                {
                  let _layer = this.frontRepo.Layers.get(rectlinklink.Layer_RectLinkLinksDBID.Int64)
                  if (_layer) {
                    if (_layer.RectLinkLinks == undefined) {
                      _layer.RectLinkLinks = new Array<RectLinkLinkDB>()
                    }
                    _layer.RectLinkLinks.push(rectlinklink)
                    if (rectlinklink.Layer_RectLinkLinks_reverse == undefined) {
                      rectlinklink.Layer_RectLinkLinks_reverse = _layer
                    }
                  }
                }
              }
            )
            svgs.forEach(
              svg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field StartRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.StartRectID.Int64)
                  if (_rect) {
                    svg.StartRect = _rect
                  }
                }
                // insertion point for pointer field EndRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.EndRectID.Int64)
                  if (_rect) {
                    svg.EndRect = _rect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            texts.forEach(
              text => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Texts redeeming
                {
                  let _layer = this.frontRepo.Layers.get(text.Layer_TextsDBID.Int64)
                  if (_layer) {
                    if (_layer.Texts == undefined) {
                      _layer.Texts = new Array<TextDB>()
                    }
                    _layer.Texts.push(text)
                    if (text.Layer_Texts_reverse == undefined) {
                      text.Layer_Texts_reverse = _layer
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // AnimatePull performs a GET on Animate of the stack and redeem association pointers 
  AnimatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.animateService.getAnimates(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            animates,
          ]) => {
            // init the array
            this.frontRepo.Animates_array = animates

            // clear the map that counts Animate in the GET
            this.frontRepo.Animates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            animates.forEach(
              animate => {
                this.frontRepo.Animates.set(animate.ID, animate)
                this.frontRepo.Animates_batch.set(animate.ID, animate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Circle.Animations redeeming
                {
                  let _circle = this.frontRepo.Circles.get(animate.Circle_AnimationsDBID.Int64)
                  if (_circle) {
                    if (_circle.Animations == undefined) {
                      _circle.Animations = new Array<AnimateDB>()
                    }
                    _circle.Animations.push(animate)
                    if (animate.Circle_Animations_reverse == undefined) {
                      animate.Circle_Animations_reverse = _circle
                    }
                  }
                }
                // insertion point for slice of pointer field Ellipse.Animates redeeming
                {
                  let _ellipse = this.frontRepo.Ellipses.get(animate.Ellipse_AnimatesDBID.Int64)
                  if (_ellipse) {
                    if (_ellipse.Animates == undefined) {
                      _ellipse.Animates = new Array<AnimateDB>()
                    }
                    _ellipse.Animates.push(animate)
                    if (animate.Ellipse_Animates_reverse == undefined) {
                      animate.Ellipse_Animates_reverse = _ellipse
                    }
                  }
                }
                // insertion point for slice of pointer field Line.Animates redeeming
                {
                  let _line = this.frontRepo.Lines.get(animate.Line_AnimatesDBID.Int64)
                  if (_line) {
                    if (_line.Animates == undefined) {
                      _line.Animates = new Array<AnimateDB>()
                    }
                    _line.Animates.push(animate)
                    if (animate.Line_Animates_reverse == undefined) {
                      animate.Line_Animates_reverse = _line
                    }
                  }
                }
                // insertion point for slice of pointer field LinkAnchoredText.Animates redeeming
                {
                  let _linkanchoredtext = this.frontRepo.LinkAnchoredTexts.get(animate.LinkAnchoredText_AnimatesDBID.Int64)
                  if (_linkanchoredtext) {
                    if (_linkanchoredtext.Animates == undefined) {
                      _linkanchoredtext.Animates = new Array<AnimateDB>()
                    }
                    _linkanchoredtext.Animates.push(animate)
                    if (animate.LinkAnchoredText_Animates_reverse == undefined) {
                      animate.LinkAnchoredText_Animates_reverse = _linkanchoredtext
                    }
                  }
                }
                // insertion point for slice of pointer field Path.Animates redeeming
                {
                  let _path = this.frontRepo.Paths.get(animate.Path_AnimatesDBID.Int64)
                  if (_path) {
                    if (_path.Animates == undefined) {
                      _path.Animates = new Array<AnimateDB>()
                    }
                    _path.Animates.push(animate)
                    if (animate.Path_Animates_reverse == undefined) {
                      animate.Path_Animates_reverse = _path
                    }
                  }
                }
                // insertion point for slice of pointer field Polygone.Animates redeeming
                {
                  let _polygone = this.frontRepo.Polygones.get(animate.Polygone_AnimatesDBID.Int64)
                  if (_polygone) {
                    if (_polygone.Animates == undefined) {
                      _polygone.Animates = new Array<AnimateDB>()
                    }
                    _polygone.Animates.push(animate)
                    if (animate.Polygone_Animates_reverse == undefined) {
                      animate.Polygone_Animates_reverse = _polygone
                    }
                  }
                }
                // insertion point for slice of pointer field Polyline.Animates redeeming
                {
                  let _polyline = this.frontRepo.Polylines.get(animate.Polyline_AnimatesDBID.Int64)
                  if (_polyline) {
                    if (_polyline.Animates == undefined) {
                      _polyline.Animates = new Array<AnimateDB>()
                    }
                    _polyline.Animates.push(animate)
                    if (animate.Polyline_Animates_reverse == undefined) {
                      animate.Polyline_Animates_reverse = _polyline
                    }
                  }
                }
                // insertion point for slice of pointer field Rect.Animations redeeming
                {
                  let _rect = this.frontRepo.Rects.get(animate.Rect_AnimationsDBID.Int64)
                  if (_rect) {
                    if (_rect.Animations == undefined) {
                      _rect.Animations = new Array<AnimateDB>()
                    }
                    _rect.Animations.push(animate)
                    if (animate.Rect_Animations_reverse == undefined) {
                      animate.Rect_Animations_reverse = _rect
                    }
                  }
                }
                // insertion point for slice of pointer field RectAnchoredText.Animates redeeming
                {
                  let _rectanchoredtext = this.frontRepo.RectAnchoredTexts.get(animate.RectAnchoredText_AnimatesDBID.Int64)
                  if (_rectanchoredtext) {
                    if (_rectanchoredtext.Animates == undefined) {
                      _rectanchoredtext.Animates = new Array<AnimateDB>()
                    }
                    _rectanchoredtext.Animates.push(animate)
                    if (animate.RectAnchoredText_Animates_reverse == undefined) {
                      animate.RectAnchoredText_Animates_reverse = _rectanchoredtext
                    }
                  }
                }
                // insertion point for slice of pointer field Text.Animates redeeming
                {
                  let _text = this.frontRepo.Texts.get(animate.Text_AnimatesDBID.Int64)
                  if (_text) {
                    if (_text.Animates == undefined) {
                      _text.Animates = new Array<AnimateDB>()
                    }
                    _text.Animates.push(animate)
                    if (animate.Text_Animates_reverse == undefined) {
                      animate.Text_Animates_reverse = _text
                    }
                  }
                }
              }
            )

            // clear animates that are absent from the GET
            this.frontRepo.Animates.forEach(
              animate => {
                if (this.frontRepo.Animates_batch.get(animate.ID) == undefined) {
                  this.frontRepo.Animates.delete(animate.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // CirclePull performs a GET on Circle of the stack and redeem association pointers 
  CirclePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.circleService.getCircles(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            circles,
          ]) => {
            // init the array
            this.frontRepo.Circles_array = circles

            // clear the map that counts Circle in the GET
            this.frontRepo.Circles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            circles.forEach(
              circle => {
                this.frontRepo.Circles.set(circle.ID, circle)
                this.frontRepo.Circles_batch.set(circle.ID, circle)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Circles redeeming
                {
                  let _layer = this.frontRepo.Layers.get(circle.Layer_CirclesDBID.Int64)
                  if (_layer) {
                    if (_layer.Circles == undefined) {
                      _layer.Circles = new Array<CircleDB>()
                    }
                    _layer.Circles.push(circle)
                    if (circle.Layer_Circles_reverse == undefined) {
                      circle.Layer_Circles_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear circles that are absent from the GET
            this.frontRepo.Circles.forEach(
              circle => {
                if (this.frontRepo.Circles_batch.get(circle.ID) == undefined) {
                  this.frontRepo.Circles.delete(circle.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // EllipsePull performs a GET on Ellipse of the stack and redeem association pointers 
  EllipsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ellipseService.getEllipses(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            ellipses,
          ]) => {
            // init the array
            this.frontRepo.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            this.frontRepo.Ellipses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ellipses.forEach(
              ellipse => {
                this.frontRepo.Ellipses.set(ellipse.ID, ellipse)
                this.frontRepo.Ellipses_batch.set(ellipse.ID, ellipse)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Ellipses redeeming
                {
                  let _layer = this.frontRepo.Layers.get(ellipse.Layer_EllipsesDBID.Int64)
                  if (_layer) {
                    if (_layer.Ellipses == undefined) {
                      _layer.Ellipses = new Array<EllipseDB>()
                    }
                    _layer.Ellipses.push(ellipse)
                    if (ellipse.Layer_Ellipses_reverse == undefined) {
                      ellipse.Layer_Ellipses_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear ellipses that are absent from the GET
            this.frontRepo.Ellipses.forEach(
              ellipse => {
                if (this.frontRepo.Ellipses_batch.get(ellipse.ID) == undefined) {
                  this.frontRepo.Ellipses.delete(ellipse.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LayerPull performs a GET on Layer of the stack and redeem association pointers 
  LayerPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.layerService.getLayers(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            layers,
          ]) => {
            // init the array
            this.frontRepo.Layers_array = layers

            // clear the map that counts Layer in the GET
            this.frontRepo.Layers_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            layers.forEach(
              layer => {
                this.frontRepo.Layers.set(layer.ID, layer)
                this.frontRepo.Layers_batch.set(layer.ID, layer)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Layers redeeming
                {
                  let _svg = this.frontRepo.SVGs.get(layer.SVG_LayersDBID.Int64)
                  if (_svg) {
                    if (_svg.Layers == undefined) {
                      _svg.Layers = new Array<LayerDB>()
                    }
                    _svg.Layers.push(layer)
                    if (layer.SVG_Layers_reverse == undefined) {
                      layer.SVG_Layers_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear layers that are absent from the GET
            this.frontRepo.Layers.forEach(
              layer => {
                if (this.frontRepo.Layers_batch.get(layer.ID) == undefined) {
                  this.frontRepo.Layers.delete(layer.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinePull performs a GET on Line of the stack and redeem association pointers 
  LinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.lineService.getLines(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            lines,
          ]) => {
            // init the array
            this.frontRepo.Lines_array = lines

            // clear the map that counts Line in the GET
            this.frontRepo.Lines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            lines.forEach(
              line => {
                this.frontRepo.Lines.set(line.ID, line)
                this.frontRepo.Lines_batch.set(line.ID, line)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Lines redeeming
                {
                  let _layer = this.frontRepo.Layers.get(line.Layer_LinesDBID.Int64)
                  if (_layer) {
                    if (_layer.Lines == undefined) {
                      _layer.Lines = new Array<LineDB>()
                    }
                    _layer.Lines.push(line)
                    if (line.Layer_Lines_reverse == undefined) {
                      line.Layer_Lines_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear lines that are absent from the GET
            this.frontRepo.Lines.forEach(
              line => {
                if (this.frontRepo.Lines_batch.get(line.ID) == undefined) {
                  this.frontRepo.Lines.delete(line.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinkPull performs a GET on Link of the stack and redeem association pointers 
  LinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkService.getLinks(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            links,
          ]) => {
            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            links.forEach(
              link => {
                this.frontRepo.Links.set(link.ID, link)
                this.frontRepo.Links_batch.set(link.ID, link)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.StartID.Int64)
                  if (_rect) {
                    link.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.EndID.Int64)
                  if (_rect) {
                    link.End = _rect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Links redeeming
                {
                  let _layer = this.frontRepo.Layers.get(link.Layer_LinksDBID.Int64)
                  if (_layer) {
                    if (_layer.Links == undefined) {
                      _layer.Links = new Array<LinkDB>()
                    }
                    _layer.Links.push(link)
                    if (link.Layer_Links_reverse == undefined) {
                      link.Layer_Links_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear links that are absent from the GET
            this.frontRepo.Links.forEach(
              link => {
                if (this.frontRepo.Links_batch.get(link.ID) == undefined) {
                  this.frontRepo.Links.delete(link.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinkAnchoredTextPull performs a GET on LinkAnchoredText of the stack and redeem association pointers 
  LinkAnchoredTextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            linkanchoredtexts,
          ]) => {
            // init the array
            this.frontRepo.LinkAnchoredTexts_array = linkanchoredtexts

            // clear the map that counts LinkAnchoredText in the GET
            this.frontRepo.LinkAnchoredTexts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            linkanchoredtexts.forEach(
              linkanchoredtext => {
                this.frontRepo.LinkAnchoredTexts.set(linkanchoredtext.ID, linkanchoredtext)
                this.frontRepo.LinkAnchoredTexts_batch.set(linkanchoredtext.ID, linkanchoredtext)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Link.TextAtArrowEnd redeeming
                {
                  let _link = this.frontRepo.Links.get(linkanchoredtext.Link_TextAtArrowEndDBID.Int64)
                  if (_link) {
                    if (_link.TextAtArrowEnd == undefined) {
                      _link.TextAtArrowEnd = new Array<LinkAnchoredTextDB>()
                    }
                    _link.TextAtArrowEnd.push(linkanchoredtext)
                    if (linkanchoredtext.Link_TextAtArrowEnd_reverse == undefined) {
                      linkanchoredtext.Link_TextAtArrowEnd_reverse = _link
                    }
                  }
                }
                // insertion point for slice of pointer field Link.TextAtArrowStart redeeming
                {
                  let _link = this.frontRepo.Links.get(linkanchoredtext.Link_TextAtArrowStartDBID.Int64)
                  if (_link) {
                    if (_link.TextAtArrowStart == undefined) {
                      _link.TextAtArrowStart = new Array<LinkAnchoredTextDB>()
                    }
                    _link.TextAtArrowStart.push(linkanchoredtext)
                    if (linkanchoredtext.Link_TextAtArrowStart_reverse == undefined) {
                      linkanchoredtext.Link_TextAtArrowStart_reverse = _link
                    }
                  }
                }
              }
            )

            // clear linkanchoredtexts that are absent from the GET
            this.frontRepo.LinkAnchoredTexts.forEach(
              linkanchoredtext => {
                if (this.frontRepo.LinkAnchoredTexts_batch.get(linkanchoredtext.ID) == undefined) {
                  this.frontRepo.LinkAnchoredTexts.delete(linkanchoredtext.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PathPull performs a GET on Path of the stack and redeem association pointers 
  PathPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pathService.getPaths(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            paths,
          ]) => {
            // init the array
            this.frontRepo.Paths_array = paths

            // clear the map that counts Path in the GET
            this.frontRepo.Paths_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paths.forEach(
              path => {
                this.frontRepo.Paths.set(path.ID, path)
                this.frontRepo.Paths_batch.set(path.ID, path)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Paths redeeming
                {
                  let _layer = this.frontRepo.Layers.get(path.Layer_PathsDBID.Int64)
                  if (_layer) {
                    if (_layer.Paths == undefined) {
                      _layer.Paths = new Array<PathDB>()
                    }
                    _layer.Paths.push(path)
                    if (path.Layer_Paths_reverse == undefined) {
                      path.Layer_Paths_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear paths that are absent from the GET
            this.frontRepo.Paths.forEach(
              path => {
                if (this.frontRepo.Paths_batch.get(path.ID) == undefined) {
                  this.frontRepo.Paths.delete(path.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PointPull performs a GET on Point of the stack and redeem association pointers 
  PointPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pointService.getPoints(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            points,
          ]) => {
            // init the array
            this.frontRepo.Points_array = points

            // clear the map that counts Point in the GET
            this.frontRepo.Points_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            points.forEach(
              point => {
                this.frontRepo.Points.set(point.ID, point)
                this.frontRepo.Points_batch.set(point.ID, point)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Link.ControlPoints redeeming
                {
                  let _link = this.frontRepo.Links.get(point.Link_ControlPointsDBID.Int64)
                  if (_link) {
                    if (_link.ControlPoints == undefined) {
                      _link.ControlPoints = new Array<PointDB>()
                    }
                    _link.ControlPoints.push(point)
                    if (point.Link_ControlPoints_reverse == undefined) {
                      point.Link_ControlPoints_reverse = _link
                    }
                  }
                }
              }
            )

            // clear points that are absent from the GET
            this.frontRepo.Points.forEach(
              point => {
                if (this.frontRepo.Points_batch.get(point.ID) == undefined) {
                  this.frontRepo.Points.delete(point.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PolygonePull performs a GET on Polygone of the stack and redeem association pointers 
  PolygonePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polygoneService.getPolygones(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            polygones,
          ]) => {
            // init the array
            this.frontRepo.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            this.frontRepo.Polygones_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polygones.forEach(
              polygone => {
                this.frontRepo.Polygones.set(polygone.ID, polygone)
                this.frontRepo.Polygones_batch.set(polygone.ID, polygone)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Polygones redeeming
                {
                  let _layer = this.frontRepo.Layers.get(polygone.Layer_PolygonesDBID.Int64)
                  if (_layer) {
                    if (_layer.Polygones == undefined) {
                      _layer.Polygones = new Array<PolygoneDB>()
                    }
                    _layer.Polygones.push(polygone)
                    if (polygone.Layer_Polygones_reverse == undefined) {
                      polygone.Layer_Polygones_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear polygones that are absent from the GET
            this.frontRepo.Polygones.forEach(
              polygone => {
                if (this.frontRepo.Polygones_batch.get(polygone.ID) == undefined) {
                  this.frontRepo.Polygones.delete(polygone.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PolylinePull performs a GET on Polyline of the stack and redeem association pointers 
  PolylinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polylineService.getPolylines(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            polylines,
          ]) => {
            // init the array
            this.frontRepo.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            this.frontRepo.Polylines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polylines.forEach(
              polyline => {
                this.frontRepo.Polylines.set(polyline.ID, polyline)
                this.frontRepo.Polylines_batch.set(polyline.ID, polyline)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Polylines redeeming
                {
                  let _layer = this.frontRepo.Layers.get(polyline.Layer_PolylinesDBID.Int64)
                  if (_layer) {
                    if (_layer.Polylines == undefined) {
                      _layer.Polylines = new Array<PolylineDB>()
                    }
                    _layer.Polylines.push(polyline)
                    if (polyline.Layer_Polylines_reverse == undefined) {
                      polyline.Layer_Polylines_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear polylines that are absent from the GET
            this.frontRepo.Polylines.forEach(
              polyline => {
                if (this.frontRepo.Polylines_batch.get(polyline.ID) == undefined) {
                  this.frontRepo.Polylines.delete(polyline.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectPull performs a GET on Rect of the stack and redeem association pointers 
  RectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectService.getRects(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            rects,
          ]) => {
            // init the array
            this.frontRepo.Rects_array = rects

            // clear the map that counts Rect in the GET
            this.frontRepo.Rects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rects.forEach(
              rect => {
                this.frontRepo.Rects.set(rect.ID, rect)
                this.frontRepo.Rects_batch.set(rect.ID, rect)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Rects redeeming
                {
                  let _layer = this.frontRepo.Layers.get(rect.Layer_RectsDBID.Int64)
                  if (_layer) {
                    if (_layer.Rects == undefined) {
                      _layer.Rects = new Array<RectDB>()
                    }
                    _layer.Rects.push(rect)
                    if (rect.Layer_Rects_reverse == undefined) {
                      rect.Layer_Rects_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear rects that are absent from the GET
            this.frontRepo.Rects.forEach(
              rect => {
                if (this.frontRepo.Rects_batch.get(rect.ID) == undefined) {
                  this.frontRepo.Rects.delete(rect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectAnchoredRectPull performs a GET on RectAnchoredRect of the stack and redeem association pointers 
  RectAnchoredRectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            rectanchoredrects,
          ]) => {
            // init the array
            this.frontRepo.RectAnchoredRects_array = rectanchoredrects

            // clear the map that counts RectAnchoredRect in the GET
            this.frontRepo.RectAnchoredRects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectanchoredrects.forEach(
              rectanchoredrect => {
                this.frontRepo.RectAnchoredRects.set(rectanchoredrect.ID, rectanchoredrect)
                this.frontRepo.RectAnchoredRects_batch.set(rectanchoredrect.ID, rectanchoredrect)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Rect.RectAnchoredRects redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectanchoredrect.Rect_RectAnchoredRectsDBID.Int64)
                  if (_rect) {
                    if (_rect.RectAnchoredRects == undefined) {
                      _rect.RectAnchoredRects = new Array<RectAnchoredRectDB>()
                    }
                    _rect.RectAnchoredRects.push(rectanchoredrect)
                    if (rectanchoredrect.Rect_RectAnchoredRects_reverse == undefined) {
                      rectanchoredrect.Rect_RectAnchoredRects_reverse = _rect
                    }
                  }
                }
              }
            )

            // clear rectanchoredrects that are absent from the GET
            this.frontRepo.RectAnchoredRects.forEach(
              rectanchoredrect => {
                if (this.frontRepo.RectAnchoredRects_batch.get(rectanchoredrect.ID) == undefined) {
                  this.frontRepo.RectAnchoredRects.delete(rectanchoredrect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectAnchoredTextPull performs a GET on RectAnchoredText of the stack and redeem association pointers 
  RectAnchoredTextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            rectanchoredtexts,
          ]) => {
            // init the array
            this.frontRepo.RectAnchoredTexts_array = rectanchoredtexts

            // clear the map that counts RectAnchoredText in the GET
            this.frontRepo.RectAnchoredTexts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectanchoredtexts.forEach(
              rectanchoredtext => {
                this.frontRepo.RectAnchoredTexts.set(rectanchoredtext.ID, rectanchoredtext)
                this.frontRepo.RectAnchoredTexts_batch.set(rectanchoredtext.ID, rectanchoredtext)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Rect.RectAnchoredTexts redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectanchoredtext.Rect_RectAnchoredTextsDBID.Int64)
                  if (_rect) {
                    if (_rect.RectAnchoredTexts == undefined) {
                      _rect.RectAnchoredTexts = new Array<RectAnchoredTextDB>()
                    }
                    _rect.RectAnchoredTexts.push(rectanchoredtext)
                    if (rectanchoredtext.Rect_RectAnchoredTexts_reverse == undefined) {
                      rectanchoredtext.Rect_RectAnchoredTexts_reverse = _rect
                    }
                  }
                }
              }
            )

            // clear rectanchoredtexts that are absent from the GET
            this.frontRepo.RectAnchoredTexts.forEach(
              rectanchoredtext => {
                if (this.frontRepo.RectAnchoredTexts_batch.get(rectanchoredtext.ID) == undefined) {
                  this.frontRepo.RectAnchoredTexts.delete(rectanchoredtext.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectLinkLinkPull performs a GET on RectLinkLink of the stack and redeem association pointers 
  RectLinkLinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            rectlinklinks,
          ]) => {
            // init the array
            this.frontRepo.RectLinkLinks_array = rectlinklinks

            // clear the map that counts RectLinkLink in the GET
            this.frontRepo.RectLinkLinks_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectlinklinks.forEach(
              rectlinklink => {
                this.frontRepo.RectLinkLinks.set(rectlinklink.ID, rectlinklink)
                this.frontRepo.RectLinkLinks_batch.set(rectlinklink.ID, rectlinklink)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectlinklink.StartID.Int64)
                  if (_rect) {
                    rectlinklink.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _link = this.frontRepo.Links.get(rectlinklink.EndID.Int64)
                  if (_link) {
                    rectlinklink.End = _link
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.RectLinkLinks redeeming
                {
                  let _layer = this.frontRepo.Layers.get(rectlinklink.Layer_RectLinkLinksDBID.Int64)
                  if (_layer) {
                    if (_layer.RectLinkLinks == undefined) {
                      _layer.RectLinkLinks = new Array<RectLinkLinkDB>()
                    }
                    _layer.RectLinkLinks.push(rectlinklink)
                    if (rectlinklink.Layer_RectLinkLinks_reverse == undefined) {
                      rectlinklink.Layer_RectLinkLinks_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear rectlinklinks that are absent from the GET
            this.frontRepo.RectLinkLinks.forEach(
              rectlinklink => {
                if (this.frontRepo.RectLinkLinks_batch.get(rectlinklink.ID) == undefined) {
                  this.frontRepo.RectLinkLinks.delete(rectlinklink.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // SVGPull performs a GET on SVG of the stack and redeem association pointers 
  SVGPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.svgService.getSVGs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            svgs,
          ]) => {
            // init the array
            this.frontRepo.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            this.frontRepo.SVGs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            svgs.forEach(
              svg => {
                this.frontRepo.SVGs.set(svg.ID, svg)
                this.frontRepo.SVGs_batch.set(svg.ID, svg)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field StartRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.StartRectID.Int64)
                  if (_rect) {
                    svg.StartRect = _rect
                  }
                }
                // insertion point for pointer field EndRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.EndRectID.Int64)
                  if (_rect) {
                    svg.EndRect = _rect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear svgs that are absent from the GET
            this.frontRepo.SVGs.forEach(
              svg => {
                if (this.frontRepo.SVGs_batch.get(svg.ID) == undefined) {
                  this.frontRepo.SVGs.delete(svg.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TextPull performs a GET on Text of the stack and redeem association pointers 
  TextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.textService.getTexts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            texts,
          ]) => {
            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            texts.forEach(
              text => {
                this.frontRepo.Texts.set(text.ID, text)
                this.frontRepo.Texts_batch.set(text.ID, text)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Layer.Texts redeeming
                {
                  let _layer = this.frontRepo.Layers.get(text.Layer_TextsDBID.Int64)
                  if (_layer) {
                    if (_layer.Texts == undefined) {
                      _layer.Texts = new Array<TextDB>()
                    }
                    _layer.Texts.push(text)
                    if (text.Layer_Texts_reverse == undefined) {
                      text.Layer_Texts_reverse = _layer
                    }
                  }
                }
              }
            )

            // clear texts that are absent from the GET
            this.frontRepo.Texts.forEach(
              text => {
                if (this.frontRepo.Texts_batch.get(text.ID) == undefined) {
                  this.frontRepo.Texts.delete(text.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

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
export function getRectAnchoredRectUniqueID(id: number): number {
  return 83 * id
}
export function getRectAnchoredTextUniqueID(id: number): number {
  return 89 * id
}
export function getRectLinkLinkUniqueID(id: number): number {
  return 97 * id
}
export function getSVGUniqueID(id: number): number {
  return 101 * id
}
export function getTextUniqueID(id: number): number {
  return 103 * id
}
