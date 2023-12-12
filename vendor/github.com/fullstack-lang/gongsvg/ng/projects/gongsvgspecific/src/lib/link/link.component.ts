import { AfterViewInit, Component, ElementRef, Input, DoCheck, OnInit, SimpleChanges, ViewChild, AfterViewChecked, OnChanges } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { SegmentsParams, Segment, createPoint, drawSegments, Offset } from '../draw.segments';
import { Subscription } from 'rxjs';
import { ShapeMouseEvent } from '../shape.mouse.event';
import { MouseEventService } from '../mouse-event.service';
import { compareRectGeometries } from '../compare.rect.geometries'

import { SplitComponent } from 'angular-split'
import { AngularDragEndEventService } from '../angular-drag-end-event.service';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { drawLineFromRectToB } from '../draw.line.from.rect.to.point';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { swapSegment } from '../swap.segment';
import { computeLinkFromMouseEvent } from '../compute.link.from.mouse.event';

@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent implements OnInit, DoCheck, AfterViewChecked {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

  // the list of segments to draw the link
  // every change on this will provoke a redrawing
  segments: Segment[] | undefined

  // @ViewChild('textElement', { static: false }) textElement: ElementRef | undefined
  textWidth: number = 0
  textHeight: number = 0

  nbControlPoints = 0
  isFloatingOrthogonal = false

  // to compute wether it was a select / dragging event
  dragging = false
  draggedLink: gongsvg.LinkDB | undefined
  draggedSegmentNumber = 0
  draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType = gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START

  // dragged anchored text
  textDragging = false
  draggedTextIndex = 0

  // offset between the cursor at the start and the top left corner
  distanceMoved = 0

  private dragThreshold = 5;
  isSelected = false
  // LinkAtMouseDown is the clone of the Link when mouse down
  PointAtMouseDown: gongsvg.PointDB | undefined
  AnchoredTextAtMouseDown: gongsvg.LinkAnchoredTextDB | undefined

  //
  // for events management
  //
  private subscriptions: Subscription[] = []

  // for change detection, we need to store start and end rect
  previousStart: gongsvg.RectDB | undefined
  previousEnd: gongsvg.RectDB | undefined

  //
  // indicate wether the link is being updated
  // no drawing must happen then
  linkUpdating: boolean = false


  previousStartX = 0
  previousStartY = 0
  previousEndX = 0
  previousEndY = 0

  // to allow conformity to the drawSegment API
  map_Link_Segment: Map<gongsvg.LinkDB, Segment[]> = new (Map<gongsvg.LinkDB, Segment[]>)

  constructor(
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,
    private angularDragEndEventService: AngularDragEndEventService,
    private mouseEventService: MouseEventService,
    private elementRef: ElementRef,
    private isEditableService: IsEditableService,
    private refreshService: RefreshService,
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
  ) {

    this.subscriptions.push(mouseEventService.mouseMouseDownEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        if (shapeMouseEvent.ShapeID == 0) {
          return
        }

        let link = this.gongsvgFrontRepoService.frontRepo.Links.get(shapeMouseEvent.ShapeID)
        if (link == undefined) {
          return
        }

        if (shapeMouseEvent.ShapeType != gongsvg.LinkDB.GONGSTRUCT_NAME ||
          shapeMouseEvent.ShapeID != this.Link!.ID) {
          return
        }

        this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)

        if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
          if (link.TextAtArrowEnd && link.TextAtArrowEnd[this.draggedTextIndex]) {
            this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowEnd[this.draggedTextIndex])
          }
        }
        if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
          if (link.TextAtArrowStart && link.TextAtArrowStart[this.draggedTextIndex]) {
            this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowStart[this.draggedTextIndex])
          }
        }

      })
    )

    this.subscriptions.push(mouseEventService.mouseMouseMoveEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        if (!this.isEditableService.getIsEditable()) {
          return
        }

        if (!this.dragging && !this.textDragging) {
          return
        }

        if (this.dragging) {
          computeLinkFromMouseEvent(this, shapeMouseEvent)
        }
        if (this.textDragging) {
          let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
          let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y

          console.log("Text dragging, deltaX", deltaX, "deltaY", deltaY)

          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            let text = this.Link!.TextAtArrowEnd![this.draggedTextIndex]
            text.X_Offset = this.AnchoredTextAtMouseDown!.X_Offset + deltaX
            text.Y_Offset = this.AnchoredTextAtMouseDown!.Y_Offset + deltaY
          }
          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
            let text = this.Link!.TextAtArrowStart![this.draggedTextIndex]
            text.X_Offset = this.AnchoredTextAtMouseDown!.X_Offset + deltaX
            text.Y_Offset = this.AnchoredTextAtMouseDown!.Y_Offset + deltaY
          }

        }
      })
    )

    this.subscriptions.push(mouseEventService.mouseMouseUpEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        if (this.dragging && this.isEditableService.getIsEditable()) {
          document.body.style.cursor = ''

          let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
          let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
          this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

          if (this.distanceMoved < this.dragThreshold) {
            console.log("Link, link selected selected: ", this.Link?.Name)
          } else {
            this.linkUpdating = true
            this.linkService.updateLink(this.Link!, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
              link => {
                // this.Link = link
                this.linkUpdating = false
                this.refreshService.emitRefreshRequestEvent(0)
              }
            )
          }
        }

        if (this.textDragging && this.isEditableService.getIsEditable()) {
          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            let text = this.Link!.TextAtArrowEnd![this.draggedTextIndex]
            this.anchoredTextService.updateLinkAnchoredText(text, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()
          }
          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
            let text = this.Link!.TextAtArrowStart![this.draggedTextIndex]
            this.anchoredTextService.updateLinkAnchoredText(text, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()
          }
        }
        this.dragging = false
        this.textDragging = false
      })
    )
  }


  ngOnInit(): void {
    // console.log("LinkComponent init: ", this.Link?.Name)

    this.isFloatingOrthogonal = this.Link!.Type == gongsvg.LinkType.LINK_TYPE_FLOATING_ORTHOGONAL

    if (this.Link) {
      if (this.Link.ControlPoints) {
        this.nbControlPoints = this.Link.ControlPoints.length
      }
    }

    this.drawSegments(this.Link!)
    this.resetPreviousState()
    this.drawSegments(this.Link!)
  }

  // this is invoked in order to modify the link if the end rects are changed
  ngDoCheck(): void {

    let hasStartChanged = !compareRectGeometries(this.previousStart!, this.Link!.Start!)
    let hasEndChanged = !compareRectGeometries(this.previousEnd!, this.Link!.End!)
    if (hasStartChanged || hasEndChanged) {
      this.drawSegments(this.Link!)
      this.resetPreviousState()
    }
  }

  ngAfterViewChecked() {
    //  console.log('Change detection run on MySvgComponent');
  }

  resetPreviousState() {
    this.previousStart = structuredClone(this.Link!.Start!)
    this.previousEnd = structuredClone(this.Link!.End!)

    this.previousStartX = this.Link!.Start!.X
    this.previousStartY = this.Link!.Start!.Y
    this.previousEndX = this.Link!.End!.X
    this.previousEndY = this.Link!.End!.Y
  }

  public getPosition(
    startRect: gongsvg.RectDB | undefined,
    position: string | undefined,
    endRect?: gongsvg.RectDB | undefined
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

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.LinkDB): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      // this link shit to dragging state
      this.dragging = true
      this.draggedLink = link
      this.draggedSegmentNumber = segmentNumber

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  linkMouseMove(event: MouseEvent, link: gongsvg.LinkDB): void {

    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  linkMouseUp(event: MouseEvent, link: gongsvg.LinkDB, segmentNumber: number = 0): void {

    // console.log("Link : linkMouseUp", this.Link?.Name)
    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  drawSegments(link: gongsvg.LinkDB): boolean {

    if (link == undefined) {
      return false
    }


    if (this.linkUpdating) {
      return true
    }

    let segmentsParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartOrientation! as gongsvg.OrientationType,
      EndDirection: link.EndOrientation! as gongsvg.OrientationType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
      CornerRadius: link.CornerRadius,
    }

    this.segments = drawSegments(segmentsParams)

    return true
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
      return 'horizontal';
    } else if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
      return 'vertical';
    } else {
      return null; // You can return null or another value if the line is not strictly horizontal or vertical
    }
  }

  // Add this method to ArcComponent
  getArcPath(segment: Segment, nextSegment: Segment): string {

    const startDegree = 180
    const endDegree = 270
    const startRadians = (startDegree * Math.PI) / 180;
    const endRadians = (endDegree * Math.PI) / 180;
    const startX = segment.EndPoint.X
    const startY = segment.EndPoint.Y
    const endX = nextSegment.StartPoint.X
    const endY = nextSegment.StartPoint.Y
    const largeArcFlag = endDegree - startDegree <= 180 ? 0 : 1;

    // 1 is positive angle direction
    // 0 otherwise
    let sweepFlag = 0
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

      let segmentDirection = 0
      if (segment.EndPoint.X > segment.StartPoint.X) {
        segmentDirection = 1
      } else {
        segmentDirection = -1
      }

      let cornerDirection = 0
      if (segment.EndPoint.Y < nextSegment.StartPoint.Y) {
        cornerDirection = 1
      } else {
        cornerDirection = -1
      }

      if (segmentDirection * cornerDirection == 1) {
        sweepFlag = 1
      } else {
        sweepFlag = 0
      }

    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
      let segmentDirection = 0
      if (segment.EndPoint.Y > segment.StartPoint.Y) {
        segmentDirection = 1
      } else {
        segmentDirection = -1
      }

      let cornerDirection = 0
      if (segment.EndPoint.X < nextSegment.StartPoint.X) {
        cornerDirection = 1
      } else {
        cornerDirection = -1
      }

      if (segmentDirection * cornerDirection == 1) {
        sweepFlag = 0
      } else {
        sweepFlag = 1
      }
    }
    return `M ${startX} ${startY} A ${this.Link!.CornerRadius} ${this.Link!.CornerRadius} 0 ${largeArcFlag} ${sweepFlag} ${endX} ${endY}`;
  }

  getEndArrowPath(segment: Segment, arrowSize: number): string {
    const ratio = 0.707106781 / 2 // (1/sqrt(2)) / 2

    let firstStartX = segment.EndPoint.X
    let firstStartY = segment.EndPoint.Y

    let secondStartX = segment.EndPoint.X
    let secondStartY = segment.EndPoint.Y

    let firstTipX = segment.EndPoint.X
    let firstTipY = segment.EndPoint.Y
    let secondTipX = segment.EndPoint.X
    let secondTipY = segment.EndPoint.Y

    {
      let { x, y } = this.rotateToSegmentDirection(segment, - arrowSize, - arrowSize)

      firstTipX += x
      firstTipY += y
    }
    {
      let { x, y } = this.rotateToSegmentDirection(segment, this.Link!.StrokeWidth * ratio, this.Link!.StrokeWidth * ratio)
      firstStartX += x
      firstStartY += y
    }
    {
      let { x, y } = this.rotateToSegmentDirection(segment, - arrowSize, arrowSize)

      secondTipX += x
      secondTipY += y
    }
    {
      let { x, y } = this.rotateToSegmentDirection(segment, this.Link!.StrokeWidth * ratio, - this.Link!.StrokeWidth * ratio)

      secondStartX += x
      secondStartY += y
    }

    let path = `M ${firstStartX} ${firstStartY} L ${firstTipX} ${firstTipY} M ${secondStartX} ${secondStartY} L ${secondTipX} ${secondTipY}`

    return path
  }

  getStartArrowPath(segment: Segment, arrowSize: number): string {

    let inverseSegment = swapSegment(segment)

    let path = this.getEndArrowPath(inverseSegment, arrowSize)

    return path
  }

  rotateToSegmentDirection(segment: Segment, x: number, y: number): { x: number, y: number } {
    let x_res = 0
    let y_res = 0

    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
      if (segment.EndPoint.X > segment.StartPoint.X) { // 0'
        x_res = x
        y_res = y
      } else { // pi
        x_res = -x
        y_res = y
      }
    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
      if (segment.EndPoint.Y > segment.StartPoint.Y) { // pi/2
        x_res = y
        y_res = x
      } else { // 3*pi/2
        x_res = -y
        y_res = -x
      }
    }

    return { x: x_res, y: y_res }
  }

  adjustToSegmentDirection(segment: Segment, x: number, y: number): { x: number, y: number } {
    let x_res = 0
    let y_res = 0

    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
      if (segment.EndPoint.X > segment.StartPoint.X) { // 0'
        x_res = x
        y_res = y
      } else { // pi
        x_res = -x
        y_res = y
      }
    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
      if (segment.EndPoint.Y > segment.StartPoint.Y) { // pi/2
        x_res = y
        y_res = x
      } else { // 3*pi/2
        x_res = -y
        y_res = -x
      }
    }

    return { x: x_res, y: y_res }
  }

  textAnchoredMouseDown(
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      // this text shift to dragging state
      this.textDragging = true
      this.draggedTextIndex = anchoredTextIndex
      this.draggedSegmentPositionOnArrow = draggedSegmentPositionOnArrow as gongsvg.PositionOnArrowType

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  textAnchoredMouseMove(event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  textAnchoredMouseUp(event: MouseEvent): void {

    // console.log("Link : linkMouseUp", this.Link?.Name)
    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }
}
