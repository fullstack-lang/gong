import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, Input, OnDestroy, OnInit, QueryList, ViewChild, ViewChildren } from '@angular/core';

import * as gongsvg from '../../../../gongsvg/src/public-api'

import { CommonModule } from '@angular/common';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { manageHandles } from '../manage.handles';
import { Segment, createPoint, drawSegments, drawSegmentsFromLink } from '../draw.segments';
import { getOrientation } from '../get.orientation';
import { getArcPath } from '../get.arc.path';
import { getEndArrowPath } from '../get.end.arrow.path';
import { swapSegment } from '../swap.segment';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { getStartPosition } from '../get.start.position';
import { getEndPosition } from '../get.end.position';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { Observable, timer } from 'rxjs';
import { StateEnumType } from './state.enum';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { getFunctionName } from './get.function.name';
import { informBackEndOfEndOfLinkDrawing } from './inform-backend-end-of-link-drawing';
import { selectRectsByArea } from './select-rects-by-area';
import { LinkConf, computeLinkFromMouseEvent } from '../compute.link.from.mouse.event';
import { updateLinkFromCursor } from '../update.link.from.cursor';
import { TextWidthCalculatorComponent } from '../text-width-calculator/text-width-calculator.component';
import { auto_X_offset } from './auto-x-offset';
import { auto_Y_offset } from './auto-y-offset';
import { drawLineFromRectToB } from '../draw.line.from.rect.to.point';
import { LinkSegmentsPipe } from '../link-segments.pipe'


@Component({
  selector: 'lib-gongsvg-diagramming',
  templateUrl: './gongsvg-diagramming.html',
  styleUrls: ['./gongsvg-diagramming.css'],
  imports: [
    CommonModule,

    MatIconModule,
    MatButtonModule,

    TextWidthCalculatorComponent,

    LinkSegmentsPipe,
  ],
  standalone: true,
})
export class GongsvgDiagrammingComponent implements OnInit, OnDestroy, AfterViewInit {

  @Input() GONG__StackPath: string = ""

  @ViewChild('textWidthCalculator') textWidthCalculator: TextWidthCalculatorComponent | undefined
  map_text_textWidth: Map<string, number> = new Map<string, number>
  ngAfterViewInit() {
    // Now you can use textWidthCalculator
    this.changeDetectorRef.detectChanges() // this is necessary to have the width configuration working
    this.oneEm = this.textWidthCalculator!.measureTextHeight("A");
    this.changeDetectorRef.detectChanges() // this is necessary to have the width configuration working
  }

  //
  // state of the component
  //
  StateEnumType = StateEnumType
  State: StateEnumType = StateEnumType.NOT_EDITABLE


  // temporary, will be computed dynamicaly
  svgWidth = 3000
  svgHeight = 4000

  // svg is the singloton that is displayed. A svg
  // is the root of the directed acyclic graph containing
  // all svg elements
  svg = new gongsvg.SVG

  //
  // USER INTERACTION MANAGEMENT
  //
  PointAtMouseDown = new gongsvg.Point
  PointAtMouseMove = new gongsvg.Point
  PointAtMouseUp = new gongsvg.Point

  dragThreshold = 5

  //
  // RECT MANAGEMENT
  //
  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }
  // map of rects to link
  // it is used to update the connected link
  map_Rect_ConnectedLinks: Map<gongsvg.Rect, Set<gongsvg.Link>> = new (Map<gongsvg.Rect, Set<gongsvg.Link>>)

  //
  // LINK MANAGEMENT
  //
  // the components draws all svg elements directly.
  //
  // however, links of type LINK_TYPE_FLOATING_ORTHOGONAL are a set of line
  // in this case, each link is associated with a set of segment
  //
  LinkType = gongsvg.LinkType
  map_Link_Segment: Map<gongsvg.Link, Segment[]> = new (Map<gongsvg.Link, Segment[]>)
  getSegments(link: gongsvg.Link): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    return getOrientation(segment)
  }

  getArcPath(link: gongsvg.Link, segment: Segment, nextSegment: Segment): string {
    return getArcPath(link, segment, nextSegment)
  }

  getEndArrowPath(link: gongsvg.Link, segment: Segment, arrowSize: number): string {
    return getEndArrowPath(link, segment, arrowSize)
  }

  getStartArrowPath(link: gongsvg.Link, segment: Segment, arrowSize: number): string {
    let inverseSegment = swapSegment(segment)
    let path = this.getEndArrowPath(link, inverseSegment, arrowSize)
    return path
  }

  resetAllLinksPreviousStartEndRects() {
    console.log(getFunctionName())
    for (let link of this.gongsvgFrontRepo?.getFrontArray<gongsvg.Link>(gongsvg.Link.GONGSTRUCT_NAME)!) {
      // TODO : why do we need to store the previous thing
      // one remived the srture clone call
      this.map_Link_PreviousStart.set(link, link.Start!)
      this.map_Link_PreviousEnd.set(link, link.End!)
    }
  }

  // for change detection, we need to store start and end rect of all links
  map_Link_PreviousStart: Map<gongsvg.Link, gongsvg.Rect> = new (Map<gongsvg.Link, gongsvg.Rect>)
  map_Link_PreviousEnd: Map<gongsvg.Link, gongsvg.Rect> = new (Map<gongsvg.Link, gongsvg.Rect>)

  draggedLink: gongsvg.Link | undefined
  draggedSegmentNumber = 0

  //
  // RECT ANCHOR MANAGEMENT
  //
  RectAnchorType = gongsvg.RectAnchorType
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  draggedRect: gongsvg.Rect | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'
  RectAtMouseDown: gongsvg.Rect | undefined
  map_SelectedRectAtMouseDown: Map<gongsvg.Rect, gongsvg.Rect> = new (Map<gongsvg.Rect, gongsvg.Rect>)

  // a rect that is scaled might have anchored path to it
  // if the path scale proportionnaly with the shape, one
  // have to store a clone at mouse down to compute the scaled path
  RectAnchoredPathsAtMouseDown: Map<gongsvg.RectAnchoredPath, gongsvg.RectAnchoredPath> =
    new Map<gongsvg.RectAnchoredPath, gongsvg.RectAnchoredPath>

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.Rect) {
    manageHandles(rect)
  }

  //
  // LINK ANCHORED TEXT MANAGEMENT
  //

  // the link whose anchored text is dragged
  draggedLinkAnchoredTextLink: gongsvg.Link | undefined
  draggedTextIndex = 0
  draggedText: gongsvg.LinkAnchoredText | undefined

  // LinkAtMouseDown is the clone of the Link when mouse down
  AnchoredTextAtMouseDown: gongsvg.LinkAnchoredText | undefined

  // is the link anchored text at the start or the end of the arrows
  PositionOnArrowType = gongsvg.PositionOnArrowType
  draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType = gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START

  //
  // RECT LINK LINK MANAGEMENT
  //
  // Elements for managing links between a rect and a link
  public getStartPosition(rectLinkLink: gongsvg.RectLinkLink): Coordinate {
    return getStartPosition(rectLinkLink, this.map_Link_Segment)
  }

  public getEndPosition(rectLinkLink: gongsvg.RectLinkLink): Coordinate {
    return getEndPosition(rectLinkLink, this.map_Link_Segment)
  }

  //
  // BACKEND MANAGEMENT
  //
  public gongsvgFrontRepo?: gongsvg.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    public gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    public svgService: gongsvg.SVGService,
    public isEditableService: IsEditableService,

    public rectService: gongsvg.RectService,
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,
    private rectAnchoredPathService: gongsvg.RectAnchoredPathService,

    private changeDetectorRef: ChangeDetectorRef,
  ) { }

  @ViewChildren('#background2') backgroundElement: QueryList<ElementRef> | undefined;

  @ViewChildren('svgRef') svgElements: QueryList<ElementRef> | undefined;


  ngOnInit(): void {

    // console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // this component is a "push mode component"
    // because the template calls many functions whose state cannot be computed
    // by the change detector ("dirty" or "clean").
    // therefore, the extra complexity is needed otherwise the template is perpetually
    // computed
    this.changeDetectorRef.detach()

    this.gongsvgFrontRepoService.connectToWebSocket(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo
        //   "in promise to front repose servive pull", "gongsvgFrontRepo not good")

        if (this.gongsvgFrontRepo.getFrontArray(gongsvg.SVG.GONGSTRUCT_NAME).length == 1) {
          this.svg = this.gongsvgFrontRepo.getFrontArray<gongsvg.SVG>(gongsvg.SVG.GONGSTRUCT_NAME)[0]

          // set the isEditable
          this.isEditableService.setIsEditable(this.svg!.IsEditable)

          // console.log(getFunctionName(), "state switch, before", this.State)
          if (this.isEditableService.getIsEditable()) {
            this.State = StateEnumType.WAITING_FOR_USER_INPUT
          } else {
            this.State = StateEnumType.NOT_EDITABLE
          }
          // console.log(getFunctionName(), "state switch, current", this.State)

        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        // compute segments for links
        this.map_Link_Segment.clear()
        this.map_Rect_ConnectedLinks.clear()

        // compute map of rect to links
        for (let layer of this.gongsvgFrontRepo.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)) {
          for (let link of layer.Links) {
            let segments = drawSegmentsFromLink(link)
            this.map_Link_Segment.set(link, segments)

            console.assert(link.Start != undefined, "link without start rect")
            console.assert(link.End != undefined, "link without end rect")

            let connectedLinksViaStart = this.map_Rect_ConnectedLinks.get(link.Start!)
            if (connectedLinksViaStart == undefined) {
              connectedLinksViaStart = new Set<gongsvg.Link>
              this.map_Rect_ConnectedLinks.set(link.Start!, connectedLinksViaStart)
            }
            connectedLinksViaStart.add(link)

            let connectedLinksViaEnd = this.map_Rect_ConnectedLinks.get(link.End!)
            if (connectedLinksViaEnd == undefined) {
              connectedLinksViaEnd = new Set<gongsvg.Link>
              this.map_Rect_ConnectedLinks.set(link.End!, connectedLinksViaEnd)
            }
            connectedLinksViaEnd.add(link)
          }
        }

        this.resetAllLinksPreviousStartEndRects()

        // Manually trigger change detection
        this.changeDetectorRef.detectChanges()

        console.assert(this.gongsvgFrontRepo?.getFrontArray(gongsvg.SVG.GONGSTRUCT_NAME).length == 1,
          "in promise to front repose servive pull", "gongsvgFrontRepo not good")
      }
    )
  }

  ngOnDestroy() {

  }

  //
  // USER INTERACTION MNGT
  //

  //
  // computeShapeStates align shapes state on component state
  //
  // Shapes have states. For instance, Rect can be selected or not.
  // The State of shape must be conformed with the component state.
  computeShapeStates() {
    for (let layer of this.gongsvgFrontRepo?.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)!) {
      for (let rect of layer.Rects) {
        let unselectRect = false
        switch (this.State) {
          case StateEnumType.NOT_EDITABLE:
            unselectRect = true
            break;
          case StateEnumType.WAITING_FOR_USER_INPUT:
            break;
          case StateEnumType.MULTI_RECTS_SELECTION:
            break;
          case StateEnumType.LINK_DRAWING:
            unselectRect = true
            break;
          case StateEnumType.RECTS_DRAGGING:
            break;
          case StateEnumType.RECT_ANCHOR_DRAGGING:
            break;
          case StateEnumType.LINK_ANCHORED_TEXT_DRAGGING:
            unselectRect = true
            break;
          case StateEnumType.LINK_DRAGGING:
            unselectRect = true
            break;
        }
        if (unselectRect && rect.IsSelected) {
          this.unselectRect(rect)
        }
      }
    }
  }

  //
  // processGenericMouseUp performs all mouse up stuff
  processMouseUp(event: MouseEvent) {
    console.log(getFunctionName(), "state at entry", this.State)
    console.assert(this.gongsvgFrontRepo?.getFrontArray(gongsvg.SVG.GONGSTRUCT_NAME).length == 1,
      getFunctionName(), "gongsvgFrontRepo not good")

    // when the mouse has not moved more than a threshold
    // all rects are unselected
    let distanceMoved = this.Math.sqrt(
      (this.PointAtMouseDown.X - this.PointAtMouseUp.X) *
      (this.PointAtMouseDown.X - this.PointAtMouseUp.X) +
      (this.PointAtMouseDown.Y - this.PointAtMouseUp.Y) *
      (this.PointAtMouseDown.Y - this.PointAtMouseUp.Y))

    // the user clicks in the void to unselect all rect
    if (distanceMoved < this.dragThreshold && this.State == StateEnumType.WAITING_FOR_USER_INPUT) {
      console.log(getFunctionName(), "distanceMoved below threshold in state", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT

      this.unselectAllRects();
    }

    if (distanceMoved < this.dragThreshold && this.State == StateEnumType.NOT_EDITABLE) {
      console.log(getFunctionName(), "distanceMoved below threshold in state", this.State)

      this.rectService.updateFront(this.draggedRect!, this.GONG__StackPath).subscribe(
        _ => {
        }
      )
    }

    // the use clicks on a rect for selecting it if it is not selected or
    // unselecting it if it is already selected
    // if shift is pressed, all selected rect stay selected, otherwise, they are unselected
    if (distanceMoved < this.dragThreshold && this.State == StateEnumType.RECTS_DRAGGING) {
      console.log(getFunctionName(), "distanceMoved below threshold in state", this.State)

      // if the shift key is not pressed, unselect all other rects
      if (!event.shiftKey) {
        this.unselectAllRects()
      }

      console.assert(this.draggedRect != undefined, "no dragged rect")
      if (!this.draggedRect?.IsSelected) {
        this.selectRect(this.draggedRect!)
      } else {
        this.unselectRect(this.draggedRect!)
      }
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

    }

    if (this.State == StateEnumType.MULTI_RECTS_SELECTION) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

      selectRectsByArea(this)
    }

    if (this.State == StateEnumType.LINK_DRAWING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)

      informBackEndOfEndOfLinkDrawing(this)
    }

    if (this.State == StateEnumType.LINK_ANCHORED_TEXT_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)

      console.assert(this.draggedText != undefined, "no dragged text")
      this.anchoredTextService.updateFront(this.draggedText!, this.GONG__StackPath).subscribe()
    }

    if (this.State == StateEnumType.LINK_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
      this.linkService.updateFront(this.draggedLink!, this.GONG__StackPath).subscribe(
        () => {
        }
      )
      document.body.style.cursor = ''
    }

    if (this.State == StateEnumType.RECT_ANCHOR_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)

      // the path have to be updated first
      for (let path of this.draggedRect!.RectAnchoredPaths) {
        this.rectAnchoredPathService.updateFront(path, this.GONG__StackPath).subscribe()
      }

      this.rectService.updateFront(this.draggedRect!, this.GONG__StackPath).subscribe(
        () => {
        }
      )
    }

    if (this.State == StateEnumType.RECTS_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
      this.rectService.updateFront(this.draggedRect!, this.GONG__StackPath).subscribe(
        () => {
        }
      )
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

  Math = Math

  public selectRect(rect: gongsvg.Rect) {
    console.log(getFunctionName(), "selecting rect", rect.Name);
    rect.IsSelected = true;
    this.manageHandles(rect);
    this.rectService.updateFront(rect, this.GONG__StackPath).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectRect(rect: gongsvg.Rect) {
    console.log(getFunctionName(), "unselecting rect", rect.Name);
    rect.IsSelected = false;
    this.manageHandles(rect);
    this.rectService.updateFront(rect, this.GONG__StackPath).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectAllRects() {
    for (let layer of this.gongsvgFrontRepo?.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)!) {
      for (let rect of layer.Rects) {
        if (rect.IsSelected) {
          this.unselectRect(rect)
        }
      }
    }
  }

  onmousemove(event: MouseEvent, source?: string): void {
    this.PointAtMouseMove = mouseCoordInComponentRef(event)
    let deltaX = this.PointAtMouseMove.X - this.PointAtMouseDown!.X
    let deltaY = this.PointAtMouseMove.Y - this.PointAtMouseDown!.Y
    // console.log(getFunctionName(), this.PointAtMouseMove)

    // case when the user releases the shift key
    if (this.State == StateEnumType.MULTI_RECTS_SELECTION && !event.shiftKey) {
      console.log(getFunctionName(), "end user release shift key")
      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    if (this.State == StateEnumType.LINK_ANCHORED_TEXT_DRAGGING) {

      // console.log("gongsvg Text dragging, deltaX", deltaX, "deltaY", deltaY)

      if (this.draggedText == undefined) {
        console.log("Problem : this.draggedText should not be undefined")
        return
      }
      this.draggedText.X_Offset = this.AnchoredTextAtMouseDown!.X_Offset + deltaX
      this.draggedText.Y_Offset = this.AnchoredTextAtMouseDown!.Y_Offset + deltaY
    }

    if (this.State == StateEnumType.LINK_DRAGGING) {
      document.body.style.cursor = ''

      console.assert(this.draggedLink!.Start != undefined, "dragged link without start rect")
      console.assert(this.draggedLink!.End != undefined, "dragged link without end rect")

      updateLinkFromCursor(
        this.draggedLink!,
        this.draggedSegmentNumber,
        this.map_Link_Segment.get(this.draggedLink!)!,
        this.map_Link_PreviousStart.get(this.draggedLink!)!,
        this.map_Link_PreviousEnd.get(this.draggedLink!)!,
        this.PointAtMouseDown,
        this.PointAtMouseMove,
      )

      console.assert(this.draggedLink!.Start != undefined, "dragged link without start rect")
      console.assert(this.draggedLink!.End != undefined, "dragged link without end rect")

      let segments = drawSegmentsFromLink(this.draggedLink!)

      // case when one go from 2 segments to 3 segments
      if (this.map_Link_Segment.get(this.draggedLink!)!.length == 2 &&
        segments.length == 3) {
        this.draggedSegmentNumber = 1
      }

      this.map_Link_Segment.set(this.draggedLink!, segments)
    }

    if (this.State == StateEnumType.RECTS_DRAGGING) {
      console.assert(this.draggedRect != undefined, "no dragged rect")
      if (this.draggedRect!.CanMoveHorizontaly) {
        this.draggedRect!.X = this.RectAtMouseDown!.X + deltaX
      }
      if (this.draggedRect!.CanMoveVerticaly) {
        this.draggedRect!.Y = this.RectAtMouseDown!.Y + deltaY
      }

      // recompute segments of links connected to the resized rect
      let set = this.map_Rect_ConnectedLinks.get(this.draggedRect!)
      if (set != undefined) {
        for (let link of set) {
          let segments = drawSegmentsFromLink(link)
          this.map_Link_Segment.set(link, segments)
        }
      }

      for (let layer of this.gongsvgFrontRepo?.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)!) {
        for (let rect_ of layer.Rects) {
          if (rect_.IsSelected) {
            let rectAtMouseDown_ = this.map_SelectedRectAtMouseDown.get(rect_)
            if (rect_.CanMoveHorizontaly) {
              rect_.X = rectAtMouseDown_!.X + deltaX
            }
            if (rect_.CanMoveVerticaly) {
              rect_.Y = rectAtMouseDown_!.Y + deltaY
            }
            // recompute segments of links connected to the resized rect
            let set = this.map_Rect_ConnectedLinks.get(rect_)
            if (set != undefined) {
              for (let link of set) {
                let segments = drawSegmentsFromLink(link)
                this.map_Link_Segment.set(link, segments)
              }
            }
          }
        }
      }
    }

    if (this.State == StateEnumType.RECT_ANCHOR_DRAGGING) {

      let scaleProportionally = this.draggedRect?.IsScalingProportionally &&
        this.RectAtMouseDown!.Width > 0 &&
        this.RectAtMouseDown!.Height > 0

      let scale = 0

      if (this.activeAnchor === 'left') {
        this.draggedRect!.X = this.RectAtMouseDown!.X + deltaX
        this.draggedRect!.Width = this.RectAtMouseDown!.Width - deltaX
        if (scaleProportionally) {
          scale = this.draggedRect!.Width / this.RectAtMouseDown!.Width
          this.draggedRect!.Height = this.RectAtMouseDown!.Height * scale
        }
      } else if (this.activeAnchor === 'right') {
        this.draggedRect!.Width = this.RectAtMouseDown!.Width + deltaX
        if (scaleProportionally) {
          scale = this.draggedRect!.Width / this.RectAtMouseDown!.Width
          this.draggedRect!.Height = this.RectAtMouseDown!.Height * scale
        }
      } else if (this.activeAnchor === 'top') {
        this.draggedRect!.Y = this.RectAtMouseDown!.Y + deltaY
        this.draggedRect!.Height = this.RectAtMouseDown!.Height - deltaY

        if (scaleProportionally) {
          scale = this.draggedRect!.Height / this.RectAtMouseDown!.Height
          this.draggedRect!.Width = this.RectAtMouseDown!.Width * scale
        }
      } else if (this.activeAnchor === 'bottom') {
        this.draggedRect!.Height = this.RectAtMouseDown!.Height + deltaY
        if (scaleProportionally) {
          scale = this.draggedRect!.Height / this.RectAtMouseDown!.Height
          this.draggedRect!.Width = this.RectAtMouseDown!.Width * scale
        }
      }

      for (let path of this.draggedRect!.RectAnchoredPaths) {
        if (scaleProportionally && path.ScalePropotionnally) {
          // get clone at mouse down
          let clonedPath = this.RectAnchoredPathsAtMouseDown.get(path)
          if (clonedPath) {
            path.AppliedScaling = clonedPath.AppliedScaling * scale
          }
        }
      }

      // recompute segments of links connected to the resized rect
      let set = this.map_Rect_ConnectedLinks.get(this.draggedRect!)
      if (set != undefined) {
        for (let link of set) {
          let segments = drawSegmentsFromLink(link)
          this.map_Link_Segment.set(link, segments)
        }
      }
    }
    this.changeDetectorRef.detectChanges()
  }

  backgroundOnMouseDown(event: MouseEvent): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && event.shiftKey) {

      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.MULTI_RECTS_SELECTION
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

  backgroundDragOver(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnClick(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnDragEnd(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnMouseUp(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  rectMouseDown(event: MouseEvent, rect: gongsvg.Rect): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey) {
      this.State = StateEnumType.RECTS_DRAGGING
      this.RectAtMouseDown = structuredClone(rect)
      this.draggedRect = rect
      this.RectAnchoredPathsAtMouseDown.clear()
      for (let rectAnchoredPath of this.draggedRect!.RectAnchoredPaths) {
        if (rectAnchoredPath.ScalePropotionnally) {
          this.RectAnchoredPathsAtMouseDown.set(rectAnchoredPath, structuredClone(rectAnchoredPath))
        }
      }
      this.map_SelectedRectAtMouseDown.clear()
      for (let layer of this.gongsvgFrontRepo?.getFrontArray<gongsvg.Layer>(gongsvg.Layer.GONGSTRUCT_NAME)!) {
        for (let rect of layer.Rects) {
          if (rect.IsSelected) {
            this.map_SelectedRectAtMouseDown.set(rect, structuredClone(rect))
          }
        }
      }
    }
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && event.altKey) {
      this.State = StateEnumType.LINK_DRAWING
      this.svg.StartRect = rect
    }

    if (this.State == StateEnumType.NOT_EDITABLE) {
      this.draggedRect = rect
    }


    console.log(getFunctionName(), "state at exit", this.State)
  }

  rectMouseUp(event: MouseEvent, rect: gongsvg.Rect): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.LINK_DRAWING) {
      this.svg.EndRect = rect
    }
    this.processMouseUp(event)
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.Rect): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.RECT_ANCHOR_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      this.activeAnchor = anchor
      this.draggedRect = rect
      this.RectAtMouseDown = structuredClone(rect)
      this.RectAnchoredPathsAtMouseDown.clear
      for (let rectAnchoredPath of this.draggedRect!.RectAnchoredPaths) {
        if (rectAnchoredPath.ScalePropotionnally) {
          this.RectAnchoredPathsAtMouseDown.set(rectAnchoredPath, structuredClone(rectAnchoredPath))
        }
      }
    }
  }
  anchorMouseUp(event: MouseEvent, rect: gongsvg.Rect): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.Link): void {
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_DRAGGING
      console.log(getFunctionName(), "state at entry", this.State)
      console.assert(this.gongsvgFrontRepo?.getFrontArray(gongsvg.SVG.GONGSTRUCT_NAME).length == 1,
        getFunctionName(), "gongsvgFrontRepo not good")

      console.assert(link.Start != undefined, getFunctionName(), "dragged link without start rect")

      // this link shit to dragging state
      this.draggedLink = link
      this.draggedSegmentNumber = segmentNumber
    }
  }
  linkMouseUp(event: MouseEvent, link: gongsvg.Link, segmentNumber: number = 0): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  textAnchoredMouseDown(
    link: gongsvg.Link,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_ANCHORED_TEXT_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      this.draggedTextIndex = anchoredTextIndex
      this.draggedSegmentPositionOnArrow = draggedSegmentPositionOnArrow as gongsvg.PositionOnArrowType
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        this.draggedText = link.TextAtArrowStart![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowStart[this.draggedTextIndex])

      }
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        this.draggedText = link.TextAtArrowEnd![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowEnd[this.draggedTextIndex])
      }
    }
  }

  textAnchoredMouseUp(link: gongsvg.Link, event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  getBezierPath(link: gongsvg.Link): string {

    let segments = this.map_Link_Segment.get(link)

    if (segments == undefined) {
      return ""
    }

    let startPoint = segments[0].StartPoint
    let startPointVector = segments[0].EndPoint

    let endPoint = segments[segments.length - 1].EndPoint
    let endPointVectorPoint = segments[segments.length - 1].StartPoint
    let path = "M " +
      startPoint.X + " " + startPoint.Y +
      " c " +
      (startPointVector.X - startPoint.X) + ", " +
      (startPointVector.Y - startPoint.Y) + " " +

      (endPointVectorPoint.X - startPoint.X) + ", " +
      (endPointVectorPoint.Y - startPoint.Y) + " " +

      (endPoint.X - startPoint.X) + " " +
      (endPoint.Y - startPoint.Y)

    return path
  }

  // "1em" defaults to the size of the default font size applied by the browser or the user agent, 
  // which is typically 16 pixels.
  oneEm = 0

  auto_Y_offset(
    link: gongsvg.Link,
    segment: Segment,
    text: gongsvg.LinkAnchoredText,
    draggedSegmentPositionOnArrow: string): number {
    // console.log(getFunctionName(), "text", text.Content)

    return auto_Y_offset(link, segment, text, draggedSegmentPositionOnArrow, this.oneEm)
  }

  auto_X_offset(
    link: gongsvg.Link,
    segment: Segment,
    text: gongsvg.LinkAnchoredText,
    line: string,
    draggedSegmentPositionOnArrow: string): number {

    return auto_X_offset(link, segment, text, line, draggedSegmentPositionOnArrow,
      this.map_text_textWidth, this.oneEm, this.textWidthCalculator!)
  }

  public getPosition(
    startRect: gongsvg.Rect | undefined,
    position: string | undefined,
    endRect?: gongsvg.Rect | undefined
  ): Coordinate {

    let coordinate: Coordinate = [0, 0]

    if (startRect == undefined || position == undefined) {
      return coordinate
    }

    switch (position) {
      case gongsvg.AnchorType.ANCHOR_BOTTOM:
        coordinate = [startRect.X + startRect.Width / 2, startRect.Y + startRect.Height]
        break;
      case gongsvg.AnchorType.ANCHOR_TOP:
        coordinate = [startRect.X + startRect.Width / 2, startRect.Y]
        break;
      case gongsvg.AnchorType.ANCHOR_LEFT:
        coordinate = [startRect.X, startRect.Y + startRect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_RIGHT:
        coordinate = [startRect.X + startRect.Width, startRect.Y + startRect.Height / 2]
        break;
      case gongsvg.AnchorType.ANCHOR_CENTER:
        if (endRect == undefined) {
          coordinate = [startRect.X + startRect.Width / 2, startRect.Y + startRect.Height / 2]
        } else {
          let endRectCenter = createPoint(endRect.X + endRect.Width / 2, endRect.Y + endRect.Height / 2)
          let borderPoint = drawLineFromRectToB(startRect, endRectCenter)
          coordinate = [borderPoint.X, borderPoint.Y]
        }
        break;
    }

    return coordinate
  }
}
