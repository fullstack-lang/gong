import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, HostListener, Input, OnDestroy, OnInit, QueryList, ViewChild, ViewChildren } from '@angular/core';

import * as svg from '../../../../svg/src/public-api'

import { CommonModule } from '@angular/common';

import { FormsModule } from '@angular/forms'; // <-- Import FormsModule
import { MatSliderModule } from '@angular/material/slider'; // <-- Import MatSliderModule
import { MatInputModule } from '@angular/material/input'; // <-- Might be needed for slider styling/labels
import { MatDividerModule } from '@angular/material/divider'

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
import { debounceTime, Observable, Subject, takeUntil, timer } from 'rxjs';
import { StateEnumType } from '../state.enum';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { getFunctionName } from '../get.function.name';
import { informBackEndOfEndOfLinkDrawing } from '../inform-backend-end-of-link-drawing';
import { selectRectsByArea } from '../select-rects-by-area';
import { LinkConf, computeLinkFromMouseEvent } from '../compute.link.from.mouse.event';
import { updateLinkFromCursor } from '../update.link.from.cursor';
import { TextWidthCalculatorComponent } from '../calc/calc.component';
import { auto_Y_offset } from '../auto-y-offset';
import { drawLineFromRectToB } from '../draw.line.from.rect.to.point';
import { LinkSegmentsPipe } from '../link-segments.pipe'

import { formatSVG, processSVG } from '../cleanandresizesvg'
import { LayoutService } from '../layout.service';

@Component({
  selector: 'lib-svg-specific',
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    FormsModule,
    MatSliderModule,
    MatInputModule,     // <-- Add if not present
    MatDividerModule,


    TextWidthCalculatorComponent,
    LinkSegmentsPipe,
  ],
  templateUrl: './svg-specific.component.html',
  styleUrl: './svg-specific.component.css'
})
export class SvgSpecificComponent implements OnInit, OnDestroy, AfterViewInit {


  @ViewChild('controlsContainer') controlsContainerRef!: ElementRef<HTMLDivElement>;

  private destroy$ = new Subject<void>();
  private resizeSubject = new Subject<void>();

  @ViewChild('svgContainer', { static: true })
  private svgContainer!: ElementRef<SVGSVGElement>

  @Input() Name: string = ""
  zoom: number = 1;
  shiftX: number = 0;
  shiftY: number = 0;

  @ViewChild('textWidthCalculator') textWidthCalculator: TextWidthCalculatorComponent | undefined
  map_text_textWidth: Map<string, number> = new Map<string, number>
  ngAfterViewInit() {
    // Now you can use textWidthCalculator
    this.changeDetectorRef.detectChanges() // this is necessary to have the width configuration working
    this.oneEm = this.textWidthCalculator!.measureTextHeight("A");
    this.changeDetectorRef.detectChanges() // this is necessary to have the width configuration working

    // Initial height calculation after view is ready
    // Use setTimeout to ensure rendering is fully complete, especially with complex elements
    setTimeout(() => this.updateHeight(), 0);
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
  svg = new svg.SVG

  //
  // USER INTERACTION MANAGEMENT
  //
  PointAtMouseDown = new svg.Point
  PointAtMouseMove = new svg.Point
  PointAtMouseUp = new svg.Point

  dragThreshold = 5

  //
  // RECT MANAGEMENT
  //
  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }
  // map of rects to link
  // it is used to update the connected link
  map_Rect_ConnectedLinks: Map<svg.Rect, Set<svg.Link>> = new (Map<svg.Rect, Set<svg.Link>>)

  //
  // LINK MANAGEMENT
  //
  // the components draws all svg elements directly.
  //
  // however, links of type LINK_TYPE_FLOATING_ORTHOGONAL are a set of line
  // in this case, each link is associated with a set of segment
  //
  LinkType = svg.LinkType
  map_Link_Segment: Map<svg.Link, Segment[]> = new (Map<svg.Link, Segment[]>)
  getSegments(link: svg.Link): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    return getOrientation(segment)
  }

  getArcPath(link: svg.Link, segment: Segment, nextSegment: Segment): string {
    return getArcPath(link, segment, nextSegment)
  }

  getEndArrowPath(link: svg.Link, segment: Segment, arrowSize: number): string {
    return getEndArrowPath(link, segment, arrowSize)
  }

  getStartArrowPath(link: svg.Link, segment: Segment, arrowSize: number): string {
    let inverseSegment = swapSegment(segment)
    let path = this.getEndArrowPath(link, inverseSegment, arrowSize)
    return path
  }

  resetAllLinksPreviousStartEndRects() {

    for (let link of this.gongsvgFrontRepo?.getFrontArray<svg.Link>(svg.Link.GONGSTRUCT_NAME)!) {
      // TODO : why do we need to store the previous thing
      // one remived the srture clone call
      this.map_Link_PreviousStart.set(link, link.Start!)
      this.map_Link_PreviousEnd.set(link, link.End!)
    }
  }

  // for change detection, we need to store start and end rect of all links
  map_Link_PreviousStart: Map<svg.Link, svg.Rect> = new (Map<svg.Link, svg.Rect>)
  map_Link_PreviousEnd: Map<svg.Link, svg.Rect> = new (Map<svg.Link, svg.Rect>)

  draggedLink: svg.Link | undefined
  draggedSegmentNumber = 0

  //
  // RECT ANCHOR MANAGEMENT
  //
  RectAnchorType = svg.RectAnchorType
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  draggedRect: svg.Rect | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'
  RectAtMouseDown: svg.Rect | undefined
  map_SelectedRectAtMouseDown: Map<svg.Rect, svg.Rect> = new (Map<svg.Rect, svg.Rect>)

  // a rect that is scaled might have anchored path to it
  // if the path scale proportionnaly with the shape, one
  // have to store a clone at mouse down to compute the scaled path
  RectAnchoredPathsAtMouseDown: Map<svg.RectAnchoredPath, svg.RectAnchoredPath> =
    new Map<svg.RectAnchoredPath, svg.RectAnchoredPath>

  // display or not handles if selected or not
  manageHandles(rect: svg.Rect) {
    manageHandles(rect)
  }

  //
  // LINK ANCHORED TEXT MANAGEMENT
  //

  // the link whose anchored text is dragged
  draggedLinkAnchoredTextLink: svg.Link | undefined
  draggedTextIndex = 0
  draggedText: svg.LinkAnchoredText | undefined

  // LinkAtMouseDown is the clone of the Link when mouse down
  AnchoredTextAtMouseDown: svg.LinkAnchoredText | undefined

  // is the link anchored text at the start or the end of the arrows
  PositionOnArrowType = svg.PositionOnArrowType
  draggedSegmentPositionOnArrow: svg.PositionOnArrowType = svg.PositionOnArrowType.POSITION_ON_ARROW_START

  //
  // RECT LINK LINK MANAGEMENT
  //
  // Elements for managing links between a rect and a link
  public getStartPosition(rectLinkLink: svg.RectLinkLink): Coordinate {
    return getStartPosition(rectLinkLink, this.map_Link_Segment)
  }

  public getEndPosition(rectLinkLink: svg.RectLinkLink): Coordinate {
    return getEndPosition(rectLinkLink, this.map_Link_Segment)
  }

  //
  // BACKEND MANAGEMENT
  //
  public gongsvgFrontRepo?: svg.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    public gongsvgFrontRepoService: svg.FrontRepoService,
    private gongsvgNbFromBackService: svg.CommitNbFromBackService,
    public svgService: svg.SVGService,
    public isEditableService: IsEditableService,

    public rectService: svg.RectService,
    private linkService: svg.LinkService,
    private anchoredTextService: svg.LinkAnchoredTextService,
    private rectAnchoredPathService: svg.RectAnchoredPathService,
    private svgTextService: svg.SvgTextService,
    private commandService: svg.CommandService,

    private changeDetectorRef: ChangeDetectorRef,

    private layoutService: LayoutService
  ) {

    // Debounce resize events to avoid excessive calculations
    this.resizeSubject.pipe(
      debounceTime(100), // Wait for 100ms of silence before recalculating
      takeUntil(this.destroy$)
    ).subscribe(() => this.updateHeight());
  }

  updateHeight(): void {
    if (this.controlsContainerRef?.nativeElement) {
      const height = this.controlsContainerRef.nativeElement.offsetHeight;
      // Send the calculated height to the shared service
      this.layoutService.updateControlsHeight(height);
    } else {
      console.warn('Controls container ref not ready in updateHeight');
    }
  }

  @ViewChildren('#background2') backgroundElement: QueryList<ElementRef> | undefined;

  @ViewChildren('svgRef') svgElements: QueryList<ElementRef> | undefined;


  ngOnInit(): void {

    // console.log("Material component->ngOnInit : Name, " + this.Name)

    // this component is a "push mode component"
    // because the template calls many functions whose state cannot be computed
    // by the change detector ("dirty" or "clean").
    // therefore, the extra complexity is needed otherwise the template is perpetually
    // computed
    this.changeDetectorRef.detach()

    this.gongsvgFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo
        //   "in promise to front repose servive pull", "gongsvgFrontRepo not good")

        if (this.gongsvgFrontRepo.getFrontArray(svg.SVG.GONGSTRUCT_NAME).length == 1) {
          this.svg = this.gongsvgFrontRepo.getFrontArray<svg.SVG>(svg.SVG.GONGSTRUCT_NAME)[0]

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
        for (let layer of this.gongsvgFrontRepo.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)) {
          for (let link of layer.Links) {
            let segments = drawSegmentsFromLink(link)
            this.map_Link_Segment.set(link, segments)

            console.assert(link.Start != undefined, "link without start rect")
            console.assert(link.End != undefined, "link without end rect")

            let connectedLinksViaStart = this.map_Rect_ConnectedLinks.get(link.Start!)
            if (connectedLinksViaStart == undefined) {
              connectedLinksViaStart = new Set<svg.Link>
              this.map_Rect_ConnectedLinks.set(link.Start!, connectedLinksViaStart)
            }
            connectedLinksViaStart.add(link)

            let connectedLinksViaEnd = this.map_Rect_ConnectedLinks.get(link.End!)
            if (connectedLinksViaEnd == undefined) {
              connectedLinksViaEnd = new Set<svg.Link>
              this.map_Rect_ConnectedLinks.set(link.End!, connectedLinksViaEnd)
            }
            connectedLinksViaEnd.add(link)
          }
        }

        this.resetAllLinksPreviousStartEndRects()

        // Reset all animations
        this.resetAnimationsProgrammatically()

        // Manually trigger change detection
        this.changeDetectorRef.detectChanges()

        if (this.svg.IsSVGFileGenerated) {
          this.generatesSVG(false)
        }

        console.assert(this.gongsvgFrontRepo?.getFrontArray(svg.SVG.GONGSTRUCT_NAME).length == 1,
          "in promise to front repose servive pull", "gongsvgFrontRepo not good")
      }
    )
  }

  private resetAnimationsProgrammatically() {
    const allAnimateElements = this.svgContainer.nativeElement.querySelectorAll('animate');

    allAnimateElements.forEach((animateEl: SVGAnimateElement) => {

      console.log("animate modif")

      const parentElement = animateEl.parentElement;
      const attributeName = animateEl.getAttribute('attributeName');
      const fromValue = animateEl.getAttribute('from');

      if (parentElement && attributeName && fromValue) {
        // Reset to initial value
        parentElement.setAttribute(attributeName, fromValue);

        console.log("reseted")
        animateEl.beginElement()

        // Force a reflow to ensure reset
        void parentElement.offsetWidth;
      }
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  // Listen for window resize
  @HostListener('window:resize')
  onResize(): void {
    this.resizeSubject.next();
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
    for (let layer of this.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
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
    console.assert(this.gongsvgFrontRepo?.getFrontArray(svg.SVG.GONGSTRUCT_NAME).length == 1,
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

      this.rectService.updateFront(this.draggedRect!, this.Name).subscribe(
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
      this.anchoredTextService.updateFront(this.draggedText!, this.Name).subscribe()
    }

    if (this.State == StateEnumType.LINK_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
      this.linkService.updateFront(this.draggedLink!, this.Name).subscribe(
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
        this.rectAnchoredPathService.updateFront(path, this.Name).subscribe()
      }

      this.rectService.updateFront(this.draggedRect!, this.Name).subscribe(
        () => {
        }
      )
    }

    if (this.State == StateEnumType.RECTS_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
      this.rectService.updateFront(this.draggedRect!, this.Name).subscribe(
        () => {
        }
      )
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

  Math = Math

  public selectRect(rect: svg.Rect) {
    console.log(getFunctionName(), "selecting rect", rect.Name);
    rect.IsSelected = true;
    this.manageHandles(rect);
    this.rectService.updateFront(rect, this.Name).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectRect(rect: svg.Rect) {
    console.log(getFunctionName(), "unselecting rect", rect.Name);
    rect.IsSelected = false;
    this.manageHandles(rect);
    this.rectService.updateFront(rect, this.Name).subscribe(
      _ => {
        this.changeDetectorRef.detectChanges()
      }
    );
  }

  private unselectAllRects() {
    for (let layer of this.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
      for (let rect of layer.Rects) {
        if (rect.IsSelected) {
          this.unselectRect(rect)
        }
      }
    }
  }

  onmousemove(event: MouseEvent, source?: string): void {
    this.PointAtMouseMove = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
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

      for (let layer of this.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
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
    this.PointAtMouseDown = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && event.shiftKey) {

      console.log(getFunctionName(), "state switch, before", this.State)
      this.State = StateEnumType.MULTI_RECTS_SELECTION
      console.log(getFunctionName(), "state switch, current", this.State)
    }

    this.computeShapeStates()
    this.changeDetectorRef.detectChanges()
  }

  backgroundDragOver(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnClick(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnDragEnd(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  backgroundOnMouseUp(event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  rectMouseDown(event: MouseEvent, rect: svg.Rect): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
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
      for (let layer of this.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
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

  rectMouseUp(event: MouseEvent, rect: svg.Rect): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    if (this.State == StateEnumType.LINK_DRAWING) {
      this.svg.EndRect = rect
    }
    this.processMouseUp(event)
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: svg.Rect): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
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
  anchorMouseUp(event: MouseEvent, rect: svg.Rect): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: svg.Link): void {
    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_DRAGGING
      console.log(getFunctionName(), "state at entry", this.State)
      console.assert(this.gongsvgFrontRepo?.getFrontArray(svg.SVG.GONGSTRUCT_NAME).length == 1,
        getFunctionName(), "gongsvgFrontRepo not good")

      console.assert(link.Start != undefined, getFunctionName(), "dragged link without start rect")

      // this link shit to dragging state
      this.draggedLink = link
      this.draggedSegmentNumber = segmentNumber
    }
  }
  linkMouseUp(event: MouseEvent, link: svg.Link, segmentNumber: number = 0): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  textAnchoredMouseDown(
    link: svg.Link,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      this.State = StateEnumType.LINK_ANCHORED_TEXT_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      this.draggedTextIndex = anchoredTextIndex
      this.draggedSegmentPositionOnArrow = draggedSegmentPositionOnArrow as svg.PositionOnArrowType
      if (this.draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        this.draggedText = link.TextAtArrowStart![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowStart[this.draggedTextIndex])

      }
      if (this.draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        this.draggedText = link.TextAtArrowEnd![this.draggedTextIndex]
        this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowEnd[this.draggedTextIndex])
      }
    }
  }

  textAnchoredMouseUp(link: svg.Link, event: MouseEvent): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    this.processMouseUp(event)
  }

  getBezierPath(link: svg.Link): string {

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
    link: svg.Link,
    segment: Segment,
    text: svg.LinkAnchoredText,
    draggedSegmentPositionOnArrow: string): number {
    // console.log(getFunctionName(), "text", text.Content)

    return auto_Y_offset(link, segment, text, draggedSegmentPositionOnArrow, this.oneEm)
  }

  public getPosition(
    startRect: svg.Rect | undefined,
    position: string | undefined,
    endRect?: svg.Rect | undefined
  ): Coordinate {

    let coordinate: Coordinate = [0, 0]

    if (startRect == undefined || position == undefined) {
      return coordinate
    }

    switch (position) {
      case svg.AnchorType.ANCHOR_BOTTOM:
        coordinate = [startRect.X + startRect.Width / 2, startRect.Y + startRect.Height]
        break;
      case svg.AnchorType.ANCHOR_TOP:
        coordinate = [startRect.X + startRect.Width / 2, startRect.Y]
        break;
      case svg.AnchorType.ANCHOR_LEFT:
        coordinate = [startRect.X, startRect.Y + startRect.Height / 2]
        break;
      case svg.AnchorType.ANCHOR_RIGHT:
        coordinate = [startRect.X + startRect.Width, startRect.Y + startRect.Height / 2]
        break;
      case svg.AnchorType.ANCHOR_CENTER:
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

  generatesSVG(download: boolean) {
    // Retrieve the native SVG element through the ViewChild/ElementRef
    const svgElement: SVGSVGElement = this.svgContainer.nativeElement;

    // Find the main content group element (the first <g> inside the <svg>)
    // This assumes your drawable content is within the first <g> tag directly under <svg>
    const contentGroup = svgElement.querySelector('g');
    if (!contentGroup) {
      console.error("Could not find the main content group <g> element.");
      return; // Exit if the group isn't found
    }

    // Calculate the bounding box of the content group
    // Note: getBBox might not be perfectly accurate if elements have transforms applied.
    // Consider iterating through elements if needed for complex cases.
    const bbox = contentGroup.getBBox();

    // Add some padding around the bounding box (optional, adjust as needed)
    const padding = 20;
    const viewBoxX = bbox.x - padding;
    const viewBoxY = bbox.y - padding;
    const viewBoxWidth = bbox.width + (padding * 2);
    const viewBoxHeight = bbox.height + (padding * 2);

    // --- Store original attributes ---
    const originalWidth = svgElement.getAttribute('width');
    const originalHeight = svgElement.getAttribute('height');
    const originalViewBox = svgElement.getAttribute('viewBox');

    // --- Set attributes for download ---
    // Set viewBox to encompass the calculated bounding box
    // Ensure width/height are positive, fallback if bbox is empty
    const finalViewBoxWidth = viewBoxWidth > 0 ? viewBoxWidth : 100; // Min width 100
    const finalViewBoxHeight = viewBoxHeight > 0 ? viewBoxHeight : 100; // Min height 100
    svgElement.setAttribute('viewBox', `${viewBoxX} ${viewBoxY} ${finalViewBoxWidth} ${finalViewBoxHeight}`);

    // Set width/height based on viewBox aspect ratio for clarity in downloaded file
    // These attributes often help standalone SVG viewers determine initial size
    svgElement.setAttribute('width', `${finalViewBoxWidth}`);
    svgElement.setAttribute('height', `${finalViewBoxHeight}`);

    // Create a serializer to convert the SVG DOM node to a string
    const serializer: XMLSerializer = new XMLSerializer();

    // Serialize the SVG element (now with the correct viewBox/dimensions)
    let svgData: string = serializer.serializeToString(svgElement);

    // --- Restore original attributes (important!) ---
    // Restore original width, height, and viewBox so the on-screen display isn't affected
    if (originalWidth !== null) svgElement.setAttribute('width', originalWidth); else svgElement.removeAttribute('width');
    if (originalHeight !== null) svgElement.setAttribute('height', originalHeight); else svgElement.removeAttribute('height');
    if (originalViewBox !== null) svgElement.setAttribute('viewBox', originalViewBox); else svgElement.removeAttribute('viewBox');

    // --- Continue with the rest of the processing ---

    // Remove any existing HTML comments in the serialized SVG (if '//g' was intended for comments)
    // This regex might need adjustment depending on actual comment format
    // let withoutComments: string = svgData.replace(/\/\/g/g, ''); // Example regex for "//g" comment
    let withoutComments: string = svgData; // Assuming no specific comment format "//g"

    // Remove Angular's auto-generated attributes
    let res: string = withoutComments
      .replace(/\s*_ngcontent-[^="]*=""/g, '')
      .replace(/\s+_nghost-[^="]*=""/g, '');

    // Perform additional custom processing if these functions exist
    let svg2: string = processSVG(res); // Ensure processSVG is defined and imported
    let svg3: string = formatSVG(svg2); // Ensure formatSVG is defined and imported

    // Inject Roboto Font Style
    const fontStyle: string = '<style>text { font-family: Roboto, Arial, sans-serif !important; }</style>';
    const svgTagEndIndex: number = svg3.indexOf('>');
    if (svgTagEndIndex > -1) {
      svg3 = svg3.slice(0, svgTagEndIndex + 1) + fontStyle + svg3.slice(svgTagEndIndex + 1);
    }

    // --- SvgText Update Logic (Keep if needed) ---
    let svgText: svg.SvgText | undefined;
    if (this.gongsvgFrontRepo?.array_SvgTexts) {
      // Assuming only one SvgText exists or you want the last one
      svgText = this.gongsvgFrontRepo.array_SvgTexts[this.gongsvgFrontRepo.array_SvgTexts.length - 1];
    }

    if (svgText !== undefined && this.svgTextService) { // Check if svgTextService is available
      svgText.Text = svg3;
      console.log("svgText updating svgText with", svg3);
      this.svgTextService.updateFront(svgText, this.Name).subscribe(
        () => {
          console.log("svgText updated with downloaded content");
        }
      );
    }
    // --- End SvgText Update Logic ---

    if (download == true) {
      // Create Blob
      const blob: Blob = new Blob([svg3], { type: 'image/svg+xml' });

      // Create download link
      const url: string = URL.createObjectURL(blob);
      const link: HTMLAnchorElement = document.createElement('a');
      link.href = url;
      link.download = (this.svg?.Name || 'download') + ".svg"; // Use optional chaining for safety
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      URL.revokeObjectURL(url);
    }
  }

  /**
 * Called when zoom, shiftX, or shiftY slider value changes.
 * Manually triggers change detection because the detector is detached.
 */
  onTransformChange(): void {
    this.changeDetectorRef.detectChanges();
  }

  generatesSVGInTheBack(arg0: boolean) {

    console.log("generatesSVGInTheBack begin")

    let generateSVG = new (svg.CommandAPI)

    generateSVG.CommandType = svg.CommandType.CommandTypeSVGInTheBack

    this.commandService.post(generateSVG, this.Name, this.gongsvgFrontRepo!).subscribe(

      generateCommand => {

        console.log("Command post completed")

        // destroy the command
        this.commandService.delete(generateCommand, this.Name).subscribe(
          () => {
            console.log("Command delete completed")
          }
        )
      }
    )

  }
}
