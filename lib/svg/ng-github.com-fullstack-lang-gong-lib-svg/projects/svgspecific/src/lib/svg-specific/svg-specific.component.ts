import { AfterViewInit, ChangeDetectorRef, ChangeDetectionStrategy, Component, ElementRef, HostListener, Input, OnDestroy, OnInit, QueryList, ViewChild, ViewChildren } from '@angular/core';

import * as svg from '../../../../svg/src/public-api'

import { CommonModule } from '@angular/common';

import { FormsModule } from '@angular/forms'; // <-- Import FormsModule
import { MatSliderModule } from '@angular/material/slider'; // <-- Import MatSliderModule
import { MatInputModule } from '@angular/material/input'; // <-- Might be needed for slider styling/labels
import { MatDividerModule } from '@angular/material/divider'

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule, TooltipPosition } from '@angular/material/tooltip';
import { provideAnimations } from '@angular/platform-browser/animations';

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
import { getPosition } from '../get-position';
import { controlPointToPoint, pointToControlPoint } from '../control-points';

@Component({
  selector: 'lib-svg-specific',
  changeDetection: ChangeDetectionStrategy.OnPush,
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    FormsModule,
    MatSliderModule,
    MatInputModule,     // <-- Add if not present
    MatDividerModule,
    MatTooltipModule,

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
    // console.log("this.oneEm", this.oneEm)
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
  svgHeight = 8000

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
  map_PeerRectAtMouseDown: Map<svg.Rect, svg.Rect> = new (Map<svg.Rect, svg.Rect>)

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
  // HOVERING MANAGEMENT
  //
  // Map to store hover states for Conditions by their ID
  public conditionHoverStates: Map<number, boolean> = new Map<number, boolean>()


  // for control point dragging
  controlPointDragging: boolean = false
  activeControlPointLink: svg.Link | undefined
  activeControlPointIndex: number = 0
  ControlPointAtMouseDown: svg.ControlPoint | undefined

  //
  // BACKEND MANAGEMENT
  //
  public gongsvgFrontRepo?: svg.FrontRepo
  public fileToDownload?: svg.FileToDownload
  private downloadedFileIds = new Set<number>();

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
    private controlPointService: svg.ControlPointService,
    private anchoredTextService: svg.LinkAnchoredTextService,
    private rectAnchoredPathService: svg.RectAnchoredPathService,
    private svgTextService: svg.SvgTextService,
    public fileToDownloadService: svg.FileToDownloadService,

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
    // this.changeDetectorRef.detach()

    this.gongsvgFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo
        this.svg = new svg.SVG()
        //   "in promise to front repose servive pull", "gongsvgFrontRepo not good")

        // Initialize conditionHoverStates after repo is loaded
        this.initializeConditionHoverStates()

        this.fileToDownload = undefined
        for (let file_ of this.gongsvgFrontRepo.getFrontArray<svg.FileToDownload>(svg.FileToDownload.GONGSTRUCT_NAME)) {
          this.fileToDownload = file_;
        }

        if (this.fileToDownload && !this.downloadedFileIds.has(this.fileToDownload.ID)) {
          this.downloadedFileIds.add(this.fileToDownload.ID);

          // Decode the base64 string to binary data
          const binaryString = window.atob(this.fileToDownload.Base64EncodedContent);
          const len = binaryString.length;
          const bytes = new Uint8Array(len);
          for (let i = 0; i < len; i++) {
            bytes[i] = binaryString.charCodeAt(i);
          }

          // Create Blob from the binary array instead of the raw string
          const blob = new Blob([bytes], { type: 'application/octet-stream' });
          const url = URL.createObjectURL(blob);
          const link = document.createElement('a');
          link.href = url;
          link.download = this.fileToDownload.Name;
          link.click();
          URL.revokeObjectURL(url);
        }

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
          this.changeDetectorRef.detectChanges()
          return
        }

        if (this.svg.Layers == undefined) {
          this.changeDetectorRef.detectChanges()
          return
        }

        if (this.svg.OverrideWidth == true) {
          this.svgWidth = this.svg.OverriddenWidth
        } else {
          this.svgWidth = 3000
        }
        if (this.svg.OverrideHeight == true) {
          this.svgHeight = this.svg.OverriddenHeight
        } else {
          this.svgHeight = 4000
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

        if (this.svg.IsSVGFrontEndFileGenerated) {
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

  @HostListener('window:mousemove', ['$event'])
  onWindowMouseMove(event: MouseEvent): void {
    if (this.State !== StateEnumType.WAITING_FOR_USER_INPUT && this.State !== StateEnumType.NOT_EDITABLE) {
      let pt = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
      
      if (pt.X > this.svgWidth) {
        this.svgWidth = pt.X + 100
      }
      if (pt.Y > this.svgHeight) {
        this.svgHeight = pt.Y + 100
      }

      this.onmousemove(event, 'window')
    }
  }

  @HostListener('window:mouseup', ['$event'])
  onWindowMouseUp(event: MouseEvent): void {
    if (this.State !== StateEnumType.WAITING_FOR_USER_INPUT && this.State !== StateEnumType.NOT_EDITABLE) {
      this.processMouseUp(event)
    }
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

      this.draggedRect!.MouseX = this.PointAtMouseUp.X
      this.draggedRect!.MouseY = this.PointAtMouseUp.Y
      if (event.shiftKey) {
        this.draggedRect!.MouseEventKey = svg.MouseEventKey.MouseEventKeyShift
      }
      if (event.altKey) {
        this.draggedRect!.MouseEventKey = svg.MouseEventKey.MouseEventKeyAlt
      }
      if (event.metaKey) {
        this.draggedRect!.MouseEventKey = svg.MouseEventKey.MouseEventKeyMeta
      }
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
      let point = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
      this.draggedLink!.MouseX = point.X
      this.draggedLink!.MouseY = point.Y
      if (event.shiftKey) {
        this.draggedLink!.MouseEventKey = svg.MouseEventKey.MouseEventKeyShift
      }
      if (event.altKey) {
        this.draggedLink!.MouseEventKey = svg.MouseEventKey.MouseEventKeyAlt
      }
      if (event.metaKey) {
        this.draggedLink!.MouseEventKey = svg.MouseEventKey.MouseEventKeyMeta
      }
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

      // Create a set to ensure we don't update the same rect twice 
      // (in case draggedRect is also IsSelected)
      let rectsToUpdate = new Set<svg.Rect>()

      if (this.draggedRect) {
        rectsToUpdate.add(this.draggedRect)
      }
      for (let peer of this.map_PeerRectAtMouseDown.keys()) {
        rectsToUpdate.add(peer)
      }

      for (let rect of this.map_SelectedRectAtMouseDown.keys()) {
        rectsToUpdate.add(rect)
      }

      for (let rect of rectsToUpdate) {
        this.rectService.updateFront(rect, this.Name).subscribe(
          () => {
            // console.log("RECTS_DRAGGING, updated", rect.Name)
          }
        )
      }
    }

    if (this.State == StateEnumType.CONTROL_POINT_DRAGGING) {
      this.State = StateEnumType.WAITING_FOR_USER_INPUT
      console.log(getFunctionName(), "state at exit", this.State)
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

  private constrainRect(rect: svg.Rect, originalX?: number, originalY?: number) {
    if (rect.AnchoredTo) {
      let anchor = rect.AnchoredTo;
      let cx = rect.X + rect.Width / 2;
      let cy = rect.Y + rect.Height / 2;

      let minX = anchor.X;
      let maxX = anchor.X + anchor.Width;
      let minY = anchor.Y;
      let maxY = anchor.Y + anchor.Height;

      let clampedCX = Math.max(minX, Math.min(cx, maxX));
      let clampedCY = Math.max(minY, Math.min(cy, maxY));

      let distLeft = Math.abs(clampedCX - minX);
      let distRight = Math.abs(clampedCX - maxX);
      let distTop = Math.abs(clampedCY - minY);
      let distBottom = Math.abs(clampedCY - maxY);

      let minDist = Math.min(distLeft, distRight, distTop, distBottom);

      if (minDist === distLeft) {
        clampedCX = minX;
      } else if (minDist === distRight) {
        clampedCX = maxX;
      } else if (minDist === distTop) {
        clampedCY = minY;
      } else if (minDist === distBottom) {
        clampedCY = maxY;
      }

      rect.X = clampedCX - rect.Width / 2;
      rect.Y = clampedCY - rect.Height / 2;
    }

    if (rect.EnclosingRect) {
      // Priority 2: if X too big, resize X
      // (This is performed first to guarantee that the shifting below can succeed)
      if (rect.Width > rect.EnclosingRect.Width) {
        rect.Width = rect.EnclosingRect.Width;
      }
      if (rect.Height > rect.EnclosingRect.Height) {
        rect.Height = rect.EnclosingRect.Height;
      }

      // Priority 1: move X inside Y
      if (rect.X < rect.EnclosingRect.X) {
        rect.X = rect.EnclosingRect.X;
      } else if (rect.X + rect.Width > rect.EnclosingRect.X + rect.EnclosingRect.Width) {
        rect.X = rect.EnclosingRect.X + rect.EnclosingRect.Width - rect.Width;
      }

      if (rect.Y < rect.EnclosingRect.Y) {
        rect.Y = rect.EnclosingRect.Y;
      } else if (rect.Y + rect.Height > rect.EnclosingRect.Y + rect.EnclosingRect.Height) {
        rect.Y = rect.EnclosingRect.Y + rect.EnclosingRect.Height - rect.Height;
      }
    }

    // Obstacle Avoidance (Sliding Collision System)
    if (originalX !== undefined && originalY !== undefined) {
      const obstacles: any[] = (rect as any).Obstacles || [];
      for (const obstacle of obstacles) {
        if (!obstacle || obstacle.X === undefined || obstacle.Width === undefined) {
          continue;
        }

        let clampedX = rect.X;
        let clampedY = rect.Y;

        // X-axis constraint
        const overlapY = clampedY < obstacle.Y + obstacle.Height && clampedY + rect.Height > obstacle.Y;
        if (overlapY) {
          const overlapX = clampedX < obstacle.X + obstacle.Width && clampedX + rect.Width > obstacle.X;
          if (overlapX) {
            const startedLeft = originalX + rect.Width <= obstacle.X;
            const startedRight = originalX >= obstacle.X + obstacle.Width;

            if (startedLeft) {
              clampedX = Math.min(clampedX, obstacle.X - rect.Width);
            } else if (startedRight) {
              clampedX = Math.max(clampedX, obstacle.X + obstacle.Width);
            }
          }
        }
        rect.X = clampedX;

        // Y-axis constraint (re-evaluate with new X)
        const overlapX2 = clampedX < obstacle.X + obstacle.Width && clampedX + rect.Width > obstacle.X;
        if (overlapX2) {
          const overlapY2 = clampedY < obstacle.Y + obstacle.Height && clampedY + rect.Height > obstacle.Y;
          if (overlapY2) {
            const startedAbove = originalY + rect.Height <= obstacle.Y;
            const startedBelow = originalY >= obstacle.Y + obstacle.Height;

            if (startedAbove) {
              clampedY = Math.min(clampedY, obstacle.Y - rect.Height);
            } else if (startedBelow) {
              clampedY = Math.max(clampedY, obstacle.Y + obstacle.Height);
            }
          }
        }
        rect.Y = clampedY;
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

      // Constrain movement to horizontal or vertical if Shift key is pressed
      if (event.shiftKey) {
        if (Math.abs(deltaX) > Math.abs(deltaY)) {
          deltaY = 0
        } else {
          deltaX = 0
        }
      }

      // 1. Gather all rects that are moving together
      let movingRects = new Set<svg.Rect>();
      if (this.draggedRect) movingRects.add(this.draggedRect);
      for (let peer of this.map_PeerRectAtMouseDown.keys()) movingRects.add(peer);
      for (let rect_ of this.map_SelectedRectAtMouseDown.keys()) movingRects.add(rect_);

      const getOriginal = (r: svg.Rect) => {
        if (r === this.draggedRect) return this.RectAtMouseDown!;
        if (this.map_PeerRectAtMouseDown.has(r)) return this.map_PeerRectAtMouseDown.get(r)!;
        if (this.map_SelectedRectAtMouseDown.has(r)) return this.map_SelectedRectAtMouseDown.get(r)!;
        return r; 
      };

      // 2. Calculate the maximum allowed movement for the entire group
      let allowedDeltaX = deltaX;
      let allowedDeltaY = deltaY;

      for (let r of movingRects) {
        let orig = getOriginal(r);

        // Create a shallow copy to safely test constrainRect without applying it yet
        let testRect = Object.assign({}, r);
        if (r.CanMoveHorizontaly) testRect.X = orig.X + deltaX;
        if (r.CanMoveVerticaly) testRect.Y = orig.Y + deltaY;

        this.constrainRect(testRect as svg.Rect, orig.X, orig.Y);

        let dx = testRect.X - orig.X;
        let dy = testRect.Y - orig.Y;

        // Keep the most restrictive delta (closest to 0) to ensure the group stays rigid
        // For AnchoredTo, the snap MUST take precedence, as it can introduce orthogonal movement
        if (r.AnchoredTo) {
          allowedDeltaX = dx;
          allowedDeltaY = dy;
        } else {
          if (Math.abs(dx) < Math.abs(allowedDeltaX)) allowedDeltaX = dx;
          if (Math.abs(dy) < Math.abs(allowedDeltaY)) allowedDeltaY = dy;
        }
      }

      // 3. Apply the unified safe movement to all rects and update links
      for (let r of movingRects) {
        let orig = getOriginal(r);
        
        if (r.CanMoveHorizontaly) r.X = orig.X + allowedDeltaX;
        if (r.CanMoveVerticaly) r.Y = orig.Y + allowedDeltaY;

        // Recompute segments of links connected to the moved rect
        let set = this.map_Rect_ConnectedLinks.get(r);
        if (set != undefined) {
          for (let link of set) {
            let segments = drawSegmentsFromLink(link);
            this.map_Link_Segment.set(link, segments);
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

      this.constrainRect(this.draggedRect!)

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

    if (this.State == StateEnumType.CONTROL_POINT_DRAGGING) {
      if (!this.activeControlPointLink || !this.ControlPointAtMouseDown) {
        return
      }

      // Get the actual point object from the link
      let draggedPoint = this.activeControlPointLink.ControlPoints[this.activeControlPointIndex]

      // Update its coordinates based on the drag delta
      let point = new (svg.Point)
      point = controlPointToPoint(this.ControlPointAtMouseDown)

      point.X = point.X + deltaX
      point.Y = point.Y + deltaY

      let draggedPointTmp = pointToControlPoint(point, this.activeControlPointLink)
      draggedPoint.X_Relative = draggedPointTmp.X_Relative
      draggedPoint.Y_Relative = draggedPointTmp.Y_Relative
      draggedPoint.ClosestRect = draggedPointTmp.ClosestRect

      // Recompute the link's segments so the line redraws live
      let segments = drawSegmentsFromLink(this.activeControlPointLink!)
      this.map_Link_Segment.set(this.activeControlPointLink!, segments)
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
      this.map_PeerRectAtMouseDown.clear()
      for (let layer of this.gongsvgFrontRepo?.getFrontArray<svg.Layer>(svg.Layer.GONGSTRUCT_NAME)!) {
        for (let rect of layer.Rects) {
          if (rect.IsSelected) {
            this.map_SelectedRectAtMouseDown.set(rect, structuredClone(rect))
            if (rect.Peers) {
              for (let peer of rect.Peers) {
                this.map_PeerRectAtMouseDown.set(peer, structuredClone(peer))
              }
            }
          }
        }
      }
      if (this.draggedRect.Peers) {
        for (let peer of this.draggedRect.Peers) {
          this.map_PeerRectAtMouseDown.set(peer, structuredClone(peer))
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
    endRect: svg.Rect | undefined,
    offset: number
  ): Coordinate {
    return getPosition(startRect, position, endRect, offset)
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

    console.log("generatesSVGInTheBack", "begin")
    this.svg.IsSVGBackEndFileGenerated = true

    this.svgService.updateFront(this.svg, this.Name).subscribe(
      svg => {
        console.log("generatesSVGInTheBack", "request for back end generation sent")
        this.svg.IsSVGBackEndFileGenerated = false
        this.svgService.updateFront(this.svg, this.Name).subscribe(
          svg => {
            console.log("generatesSVGInTheBack", "svg set back to normal")
          }
        )
      }
    )
  }

  onRectHover(rect: svg.Rect, isHovered: boolean) {
    if (rect.ChangeColorWhenHovered == true) {
      if (isHovered) {
        console.log("entrer is hovered, color", rect.ColorWhenHovered)
        rect.OriginalColor = rect.Color; // Store original
        rect.Color = rect.ColorWhenHovered; // Hover color
        rect.OriginalFillOpacity = rect.FillOpacity
        rect.FillOpacity = rect.FillOpacityWhenHovered
        this.changeDetectorRef.detectChanges();
      } else {
        console.log("exit is hovered")
        rect.Color = rect.OriginalColor; // Restore original
        rect.FillOpacity = rect.OriginalFillOpacity
        this.changeDetectorRef.detectChanges();
      }
    }

    if (rect.HoveringTrigger && rect.HoveringTrigger.length > 0) {
      rect.HoveringTrigger.forEach(condition => {
        if (this.conditionHoverStates.has(condition.ID)) {
          this.conditionHoverStates.set(condition.ID, isHovered);
          console.log(`Condition ID ${condition.ID} hover state set to: ${isHovered}`);
        } else {
          console.warn(`Condition ID ${condition.ID} not found in hover states map during hover event.`);
        }
      });
    }
  }


  getContextForAnchoredRect(anchoredRect: svg.RectAnchoredRect, parentRect: svg.Rect) {
    let anchorX = 0;
    let anchorY = 0;

    // Calculate width based on whether it should follow the parent
    const width = anchoredRect.WidthFollowRect
      ? parentRect.Width
      : anchoredRect.Width;

    // The switch logic is now neatly contained in the component method
    switch (anchoredRect.RectAnchorType) {
      case svg.RectAnchorType.RECT_TOP:
        anchorX = parentRect.X + parentRect.Width / 2 + anchoredRect.X_Offset;
        anchorY = parentRect.Y + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_LEFT:
        anchorX = parentRect.X + anchoredRect.X_Offset;
        anchorY = parentRect.Y + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_RIGHT:
        anchorX = parentRect.X + parentRect.Width + anchoredRect.X_Offset;
        anchorY = parentRect.Y + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM:
        anchorX = parentRect.X + parentRect.Width / 2 + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT:
        anchorX = parentRect.X + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT_LEFT:
        anchorX = parentRect.X - parentRect.Height + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_BOTTOM_LEFT:
        anchorX = parentRect.X + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + parentRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_RIGHT:
        anchorX = parentRect.X + parentRect.Width + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_INSIDE_RIGHT:
        anchorX = parentRect.X + parentRect.Width - anchoredRect.Width + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height - anchoredRect.Height + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_LEFT:
        anchorX = parentRect.X + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_RIGHT:
        anchorX = parentRect.X + parentRect.Width + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + anchoredRect.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_CENTER:
        anchorX = parentRect.X + parentRect.Width / 2 + anchoredRect.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + anchoredRect.Y_Offset;
        break;
    }

    // Return the complete context object that the template needs
    return {
      rect: anchoredRect,
      anchorX: anchorX,
      anchorY: anchorY,
      width: width,
    };
  }

  getContextForAnchoredPath(path: svg.RectAnchoredPath, parentRect: svg.Rect) {
    let anchorX = 0;
    let anchorY = 0;

    // The switch logic is moved here from the template - now matches rect method exactly
    switch (path.RectAnchorType) {
      case svg.RectAnchorType.RECT_TOP:
        anchorX = parentRect.X + parentRect.Width / 2 + path.X_Offset;
        anchorY = parentRect.Y + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_LEFT:
        anchorX = parentRect.X + path.X_Offset;
        anchorY = parentRect.Y + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_RIGHT:
        anchorX = parentRect.X + parentRect.Width + path.X_Offset;
        anchorY = parentRect.Y + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM:
        anchorX = parentRect.X + parentRect.Width / 2 + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT:
        anchorX = parentRect.X + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT_LEFT:
        anchorX = parentRect.X - parentRect.Height + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_BOTTOM_LEFT:
        anchorX = parentRect.X + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_RIGHT:
        anchorX = parentRect.X + parentRect.Width + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_INSIDE_RIGHT:
        anchorX = parentRect.X + parentRect.Width + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_LEFT:
        anchorX = parentRect.X + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_RIGHT:
        anchorX = parentRect.X + parentRect.Width + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + path.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_CENTER:
        anchorX = parentRect.X + parentRect.Width / 2 + path.X_Offset;
        anchorY = parentRect.Y + parentRect.Height / 2 + path.Y_Offset;
        break;
    }

    // Return the context object for the template
    return {
      path: path,
      anchorX: anchorX,
      anchorY: anchorY,
    };
  }

  getContextForAnchoredText(text: svg.RectAnchoredText, rect: svg.Rect) {
    let anchorX = 0;
    let anchorY = 0;

    // The same switch logic is now consistent with rect method
    switch (text.RectAnchorType) {
      case svg.RectAnchorType.RECT_TOP:
        anchorX = rect.X + rect.Width / 2 + text.X_Offset;
        anchorY = rect.Y + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_LEFT:
        anchorX = rect.X + text.X_Offset;
        anchorY = rect.Y + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_TOP_RIGHT:
        anchorX = rect.X + rect.Width + text.X_Offset;
        anchorY = rect.Y + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM:
        anchorX = rect.X + rect.Width / 2 + text.X_Offset;
        anchorY = rect.Y + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT:
        anchorX = rect.X + text.X_Offset;
        anchorY = rect.Y + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_LEFT_LEFT:
        anchorX = rect.X - rect.Height + text.X_Offset;
        anchorY = rect.Y + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_BOTTOM_LEFT:
        anchorX = rect.X + text.X_Offset;
        anchorY = rect.Y + rect.Height + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_RIGHT:
        anchorX = rect.X + rect.Width + text.X_Offset;
        anchorY = rect.Y + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_BOTTOM_INSIDE_RIGHT:
        anchorX = rect.X + rect.Width + text.X_Offset;
        anchorY = rect.Y + rect.Height + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_LEFT:
        anchorX = rect.X + text.X_Offset;
        anchorY = rect.Y + rect.Height / 2 + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_RIGHT:
        anchorX = rect.X + rect.Width + text.X_Offset;
        anchorY = rect.Y + rect.Height / 2 + text.Y_Offset;
        break;
      case svg.RectAnchorType.RECT_CENTER:
        anchorX = rect.X + rect.Width / 2 + text.X_Offset;
        anchorY = rect.Y + rect.Height / 2 + text.Y_Offset;
        break;
    }

    // Return the context object the template needs
    return {
      text: text,
      anchorX: anchorX,
      anchorY: anchorY,
    };
  }

  getTooltipPosition(rect: svg.Rect): TooltipPosition {
    // Check if ToolTipPosition is a valid, non-empty string
    if (rect.ToolTipPosition) {
      // If it exists, cast it to the TooltipPosition type
      return rect.ToolTipPosition as TooltipPosition;
    } else {
      // Otherwise, provide a default position
      return 'above'; // You can change this default to 'below', 'left', 'right', etc.
    }
  }

  private initializeConditionHoverStates(): void {
    this.conditionHoverStates.clear(); // Clear previous states if any
    if (this.gongsvgFrontRepo) {
      const conditions = this.gongsvgFrontRepo.getFrontArray<svg.Condition>(svg.Condition.GONGSTRUCT_NAME);
      conditions.forEach(condition => {
        this.conditionHoverStates.set(condition.ID, false); // Initialize all to false
      });
    }
  }

  // --- Helper function to check display conditions ---
  shouldDisplayRect(rect: svg.Rect): boolean {
    if (!rect.DisplayConditions || rect.DisplayConditions.length === 0) {
      return true; // No conditions, always display
    }

    // Check if ALL display conditions are currently true in the state map
    for (const condition of rect.DisplayConditions) {
      const conditionState = this.conditionHoverStates.get(condition.ID);
      // console.log(`Checking display for Rect ${rect.Name}: Condition ${condition.ID} state is ${conditionState}`);
      if (conditionState === undefined || conditionState === false) {
        // console.log(`Rect ${rect.Name} should NOT display because Condition ${condition.ID} is ${conditionState}`);
        return false; // If any condition is missing or false, do not display
      }
    }
    // console.log(`Rect ${rect.Name} should display`);
    return true; // All conditions are true
  }

  isConditionHovered(conditionID: number): boolean {
    return this.conditionHoverStates.get(conditionID) ?? false;
  }

  public getPolylinePoints(segments: Segment[]): string {
    if (!segments || segments.length === 0) {
      return "";
    }
    // Build the string: "x1,y1 x2,y2 x3,y3 ..."
    // The first point
    let points = `${segments[0].StartPoint.X},${segments[0].StartPoint.Y}`;
    // All subsequent end points
    for (const segment of segments) {
      points += ` ${segment.EndPoint.X},${segment.EndPoint.Y}`;
    }
    return points;
  }

  controlPointMouseDown(event: MouseEvent, link: svg.Link, pointIndex: number): void {
    this.PointAtMouseDown = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)

    if (this.State == StateEnumType.WAITING_FOR_USER_INPUT && !event.altKey && !event.shiftKey) {
      // Set the state to CONTROL_POINT_DRAGGING
      // (This assumes you've added CONTROL_POINT_DRAGGING to your StateEnumType)
      this.State = StateEnumType.CONTROL_POINT_DRAGGING
      console.log(getFunctionName(), "state at exit", this.State)

      // Set tracking properties
      this.controlPointDragging = true
      this.activeControlPointLink = link
      this.activeControlPointIndex = pointIndex

      // Store a clone of the point's original position
      this.ControlPointAtMouseDown = structuredClone(link.ControlPoints[pointIndex])

      // Stop the event from bubbling up to the background rect
      event.stopPropagation()
    }
  }

  controlPointMouseUp(event: MouseEvent, controlPoint: svg.ControlPoint): void {
    this.PointAtMouseUp = mouseCoordInComponentRef(event, this.zoom, this.shiftX, this.shiftY)
    console.log(getFunctionName(), "state at entry", this.State)

    // let draggedPoint = this.activeControlPointLink!.ControlPoints[this.activeControlPointIndex]

    this.controlPointService.updateFront(controlPoint, this.Name).subscribe(
      () => {
      }
    )
    this.processMouseUp(event)
  }

  pointToControlPoint(point: svg.Point, link: svg.Link): svg.ControlPoint {
    return pointToControlPoint(point, link)
  }

  controlPointToPoint(controlPoint: svg.ControlPoint): svg.Point {
    return controlPointToPoint(controlPoint)
  }

  manualDetect() {
    this.changeDetectorRef.detectChanges()
  }

  getContextForAnchoredPngImage(rectAnchoredPngImage: any, rect: any): any {
    // Basic anchoring logic. 
    // You should calculate anchorX and anchorY based on rectAnchoredPngImage's specific 
    // AnchorType, X_Offset, Y_Offset, etc. mirroring how you position paths or text.
    let anchorX = rect.X; 
    let anchorY = rect.Y;
    
    if (rectAnchoredPngImage.X_Offset) {
      anchorX += rectAnchoredPngImage.X_Offset;
    }
    
    if (rectAnchoredPngImage.Y_Offset) {
      anchorY += rectAnchoredPngImage.Y_Offset;
    }

    return {
      pngImage: rectAnchoredPngImage,
      anchorX: anchorX,
      anchorY: anchorY,
      width: rectAnchoredPngImage.Width,
      height: rectAnchoredPngImage.Height
    };
  }

  // --- Helper methods for TextAtCorner positioning ---

  getCornerTextX(segment: Segment, segments: Segment[], text: svg.LinkAnchoredText): number {
    if (!text.AutomaticLayout) {
        return segment.EndPoint.X + text.X_Offset;
    }
    if (segments.length < 2) return segment.EndPoint.X;
    
    let segment0 = segments[0];
    let isSeg0Horizontal = this.getOrientation(segment0) === 'horizontal';
    
    if (isSeg0Horizontal) {
        // Segment 1 is vertical: Center the text exactly on the corner's X coordinate
        return segment0.EndPoint.X;
    } else {
        // Segment 1 is horizontal: Place text opposite to the horizontal direction
        let segment1 = segments[1];
        let dirX = segment1.EndPoint.X - segment1.StartPoint.X;
        let signX = dirX >= 0 ? 1 : -1;
        let paddingX = 12; // Base padding distance from the corner
        
        return segment0.EndPoint.X - signX * paddingX;
    }
}

  getCornerTextY(segment: Segment, segments: Segment[], text: svg.LinkAnchoredText): number {
    if (!text.AutomaticLayout) {
        return segment.EndPoint.Y + text.Y_Offset;
    }
    if (segments.length < 2) return segment.EndPoint.Y;
    
    let segment0 = segments[0];
    let segment1 = segments[1];
    let isSeg0Vertical = this.getOrientation(segment0) === 'vertical';
    
    // Determine the direction of the vertical segment
    let dirY = isSeg0Vertical ? segment0.EndPoint.Y - segment0.StartPoint.Y : segment1.EndPoint.Y - segment1.StartPoint.Y;
    let signY = dirY >= 0 ? 1 : -1;
    let paddingY = 8;
    
    // If the first segment is vertical, we push in the SAME direction as dirY to get away from the segment
    // If the second segment is vertical, we push in the OPPOSITE direction of dirY to get away from it
    let yOffset = isSeg0Vertical ? (signY * paddingY) : (-signY * paddingY);
    
    if (yOffset > 0) {
        // Text is BELOW the corner; shift down by roughly 1em so the top of the text doesn't overlap the line
        yOffset += this.oneEm * 0.8;
    } else {
        // Text is ABOVE the corner; shift up based on multiline count so the bottom line rests above the corner
        let lines = this.splitTextIntoLines(text.Content);
        if (lines.length > 1) {
            yOffset -= (lines.length - 1) * this.oneEm;
        }
    }

    return segment0.EndPoint.Y + yOffset;
  }

  getCornerTextAnchor(segment: Segment, segments: Segment[], text: svg.LinkAnchoredText): string {
    if (!text.AutomaticLayout) {
        return 'start'; 
    }
    if (segments.length < 2) return 'start';
    
    let segment0 = segments[0];
    let isSeg0Horizontal = this.getOrientation(segment0) === 'horizontal';
    
    if (isSeg0Horizontal) {
        // Segment 1 is vertical: Text should be centered relative to the corner
        return 'middle';
    } else {
        // Segment 1 is horizontal: Flow text away from the corner
        let segment1 = segments[1];
        let dirX = segment1.EndPoint.X - segment1.StartPoint.X;
        return dirX >= 0 ? 'end' : 'start';
    }
 }
}
