import { AfterViewInit, Component, ElementRef, Input, DoCheck, OnInit, SimpleChanges, ViewChild, AfterViewChecked, OnChanges } from '@angular/core';
import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { SegmentsParams, Segment, createPoint, drawSegments, Offset } from './draw.segments';
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

@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent implements OnInit, AfterViewInit, DoCheck, AfterViewChecked, OnChanges {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

  // @ViewChild('textElement', { static: false }) textElement: ElementRef | undefined
  textWidth: number = 0
  textHeight: number = 0

  nbControlPoints = 0
  isFloatingOrthogonal = false

  // to compute wether it was a select / dragging event
  dragging = false
  draggedSegment = 0
  draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType = gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START

  // dragged anchored text
  textDragging = false
  draggedTextIndex = 0

  // offset between the cursor at the start and the top left corner
  distanceMoved = 0

  private dragThreshold = 5;
  isSelected = false
  // LinkAtMouseDown is the clone of the Link when mouse down
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private LinkAtMouseDown: gongsvg.LinkDB | undefined
  private AnchoredTextAtMouseDown: gongsvg.LinkAnchoredTextDB | undefined

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

  constructor(
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,
    private angularDragEndEventService: AngularDragEndEventService,
    private mouseEventService: MouseEventService,
    private elementRef: ElementRef,
    private isEditableService: IsEditableService,
    private refreshService: RefreshService,
  ) {

    this.subscriptions.push(
      mouseEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (shapeMouseEvent.ShapeType != gongsvg.LinkDB.GONGSTRUCT_NAME ||
            shapeMouseEvent.ShapeID != this.Link!.ID) {
            return
          }

          this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)
          this.LinkAtMouseDown = structuredClone(this.Link!)

          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            if (this.Link!.TextAtArrowEnd && this.Link!.TextAtArrowEnd[this.draggedTextIndex]) {
              this.AnchoredTextAtMouseDown = structuredClone(this.Link!.TextAtArrowEnd[this.draggedTextIndex])
            }
          }
          if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
            if (this.Link!.TextAtArrowStart && this.Link!.TextAtArrowStart[this.draggedTextIndex]) {
              this.AnchoredTextAtMouseDown = structuredClone(this.Link!.TextAtArrowStart[this.draggedTextIndex])
            }
          }

        })
    )

    this.subscriptions.push(
      mouseEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (!this.isEditableService.getIsEditable()) {
            return
          }

          // offSetForNewMidlleSegment denotes the standard distance between
          // a rect and the middle segment that is created when going from a
          // two segment link to a three segment link
          const offSetForNewMidlleSegment = 100

          // in order to have middle segment always visible
          const offsetFromBorderForNewMidlleSegment = 50

          if (this.dragging) {

            if (this.draggedSegment >= this.segments!.length) {
              this.draggedSegment = this.segments!.length - 1
            }

            let segment = this.segments![this.draggedSegment]

            if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

              // set up the cursor style
              document.body.style.cursor = 'ns-resize'

              let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
              if (this.draggedSegment == 0 && deltaY != 0) {

                let newStartRatio = (shapeMouseEvent.Point.Y - this.previousStart!.Y) / this.previousStart!.Height
                this.Link!.StartRatio = newStartRatio

                if (newStartRatio < 0 || newStartRatio > 1) {

                  let nextStartRatio = (shapeMouseEvent.Point.X - this.Link!.Start!.X) / this.Link!.Start!.Width

                  // we swap first segment from horizontal to vertical only if the new ratio
                  // is within 0 and 1 
                  if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    // the cursor is at the left of the start rectangle
                    if (newStartRatio < 0) {

                      // we compute the CornerOffsetRatio in order to have the vertical middle segment
                      // at the left of the End Rect
                      let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                      // Math.max(this.Link!.Start!.Y - offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)

                      this.Link!.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - this.Link!.Start!.Y) / this.Link!.Start!.Height
                    }

                    // 
                    if (newStartRatio > 1) {
                      let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y

                      this.Link!.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - this.Link!.Start!.Y) / this.Link!.Start!.Height

                      console.log("this.Link!.CornerOffsetRatio", this.Link!.CornerOffsetRatio)
                    }
                    // switch dragged element to next segment
                    this.draggedSegment = 1

                    newStartRatio = nextStartRatio
                    this.Link!.StartOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    this.Link!.StartRatio = newStartRatio
                    this.drawSegments()
                  } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                  }
                }
              }

              // last segment
              if (this.draggedSegment == this.segments!.length - 1 && deltaY != 0) {

                let newRatio = (shapeMouseEvent.Point.Y - this.previousEnd!.Y) / this.previousEnd!.Height

                if (newRatio < 0 || newRatio > 1) {
                  let nextStartRatio = (shapeMouseEvent.Point.X - this.Link!.End!.X) / this.Link!.End!.Width

                  if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    if (newRatio < 0) {
                      // we compute the CornerOffsetRatio in order to have the vertical middle segment
                      // at the left of the End Rect
                      let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                      this.Link!.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - this.Link!.Start!.Y) / this.Link!.Start!.Height
                    }
                    if (newRatio > 1) {
                      let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                      this.Link!.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - this.Link!.Start!.Y) / this.Link!.Start!.Height
                    }
                    // switch dragged element to previous segment
                    this.draggedSegment = this.segments!.length - 1

                    newRatio = nextStartRatio
                    this.Link!.EndOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    this.Link!.EndRatio = newRatio
                    this.drawSegments()
                  } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newRatio < 0) { newRatio = 0 }
                    if (newRatio > 1) { newRatio = 1 }
                  }
                }
                this.Link!.EndRatio = newRatio
              }

              // middle segment
              if (this.segments!.length == 3 && this.draggedSegment == 1 && deltaY != 0) {

                let newCornerOffsetRatio =
                  (shapeMouseEvent.Point.Y - this.Link!.Start!.Y) / this.Link!.Start!.Height

                this.Link!.CornerOffsetRatio = newCornerOffsetRatio
              }
            }
            if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {

              // set up the cursor style
              document.body.style.cursor = 'ew-resize'

              let deltaX = (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)

              // first segment
              if (this.draggedSegment == 0 && deltaX != 0) {

                let newStartRatio = (shapeMouseEvent.Point.X - this.previousStart!.X) / this.previousStart!.Width

                this.Link!.StartRatio = newStartRatio

                // special case where we need to change orientation
                if (newStartRatio < 0 || newStartRatio > 1) {

                  // we change orientation (horizontal to vertical), therefore we compute the new start ratio
                  let newStartVerticalRatio = (shapeMouseEvent.Point.Y - this.Link!.Start!.Y) / this.Link!.Start!.Height

                  if (newStartVerticalRatio >= 0 && newStartVerticalRatio <= 1) {
                    if (newStartRatio < 0) {
                      // we compute the CornerOffsetRatio in order to have the vertical middle segment
                      // at the left of the End Rect
                      let targetX_ForVerticalMiddleSegment = shapeMouseEvent.Point.X

                      this.Link!.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - this.Link!.Start!.X) / this.Link!.Start!.Width
                    }
                    if (newStartRatio > 1) {
                      let targetX_ForVerticalMiddleSegment = shapeMouseEvent.Point.X
                      this.Link!.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - this.Link!.Start!.X) / this.Link!.Start!.Width
                    }

                    newStartRatio = newStartVerticalRatio
                    this.Link!.StartOrientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
                    this.Link!.StartRatio = newStartVerticalRatio
                  } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                  }
                }

                // in all case, we are finished here
                this.drawSegments()
                return
              }

              // last segment
              if (this.draggedSegment == this.segments!.length - 1 && deltaX != 0) {

                let newEndRatio = (shapeMouseEvent.Point.X - this.previousEnd!.X) / this.previousEnd!.Width

                if (newEndRatio < 0 || newEndRatio > 1) {
                  let newOrientationRatio = (shapeMouseEvent.Point.Y - this.Link!.End!.Y) / this.Link!.End!.Height

                  if (newOrientationRatio >= 0 && newOrientationRatio <= 1) {
                    if (newEndRatio < 0) {
                      // we compute the CornerOffsetRatio in order to have the vertical middle segment
                      // at the left of the End Rect
                      let targetX_ForVerticalMiddleSegment =
                        Math.max(this.Link!.End!.X - offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)
                      this.Link!.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - this.Link!.Start!.X) / this.Link!.Start!.Width
                    }
                    if (newEndRatio > 1) {
                      let targetX_ForVerticalMiddleSegment =
                        Math.max(this.Link!.End!.X + this.Link!.End!.Width + offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)
                      this.Link!.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - this.Link!.Start!.X) / this.Link!.Start!.Width
                    }

                    newEndRatio = newOrientationRatio
                    this.Link!.EndOrientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
                    this.Link!.EndRatio = newEndRatio
                    this.drawSegments()
                    // switch dragged element to previous segment
                    this.draggedSegment = this.segments!.length - 1
                  } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newEndRatio < 0) { newEndRatio = 0 }
                    if (newEndRatio > 1) { newEndRatio = 1 }
                  }
                }
                this.Link!.EndRatio = newEndRatio
              }

              // middle segment
              if (this.segments!.length == 3 && this.draggedSegment == 1 && deltaX != 0) {

                let newCornerOffsetRatio = (shapeMouseEvent.Point.X - this.previousStart!.X) / this.previousStart!.Width

                this.Link!.CornerOffsetRatio = newCornerOffsetRatio
              }
            }
            this.drawSegments()
          }
          if (this.textDragging) {
            let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
            let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y

            // console.log("Text dragging, deltaX", deltaX, "deltaY", deltaY)

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

    this.subscriptions.push(
      mouseEventService.mouseMouseUpEvent$.subscribe(
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
              this.linkService.updateLink(this.Link!, this.GONG__StackPath).subscribe(
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
              this.anchoredTextService.updateLinkAnchoredText(text, this.GONG__StackPath).subscribe()
            }
            if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
              let text = this.Link!.TextAtArrowStart![this.draggedTextIndex]
              this.anchoredTextService.updateLinkAnchoredText(text, this.GONG__StackPath).subscribe()
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

    this.drawSegments()
    this.resetPreviousState()
    this.drawSegments()
  }

  ngDoCheck(): void {

    let hasStartChanged = !compareRectGeometries(this.previousStart!, this.Link!.Start!)
    let hasEndChanged = !compareRectGeometries(this.previousEnd!, this.Link!.End!)
    if (hasStartChanged || hasEndChanged) {
      this.drawSegments()
      this.resetPreviousState()
    }
  }

  ngAfterViewChecked() {
    //  console.log('Change detection run on MySvgComponent');
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes['Link']) {
      // console.log('Previous value: ', changes['Link'].previousValue);
      // console.log('Current value: ', changes['Link'].currentValue);
    }
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

  linkMouseDown(
    event: MouseEvent,
    segmentNumber: number): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      // this link shit to dragging state
      this.dragging = true
      this.draggedSegment = segmentNumber

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  linkMouseMove(event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Link!.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  linkMouseUp(event: MouseEvent, segmentNumber: number = 0): void {

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

  segmentsParams: SegmentsParams | undefined
  segments: Segment[] | undefined

  drawSegments(): boolean {

    if (this.linkUpdating) {
      return true
    }

    let link = this.Link!

    this.segmentsParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartOrientation! as gongsvg.OrientationType,
      EndDirection: link.EndOrientation! as gongsvg.OrientationType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
      CornerRadius: link.CornerRadius,
    }

    this.segments = drawSegments(this.segmentsParams)

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

  getArrowPath(segment: Segment): string {
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
      let { x, y } = this.rotateToSegmentDirection(segment, - this.Link!.EndArrowSize, - this.Link!.EndArrowSize)

      firstTipX += x
      firstTipY += y
    }
    {
      let { x, y } = this.rotateToSegmentDirection(segment, this.Link!.StrokeWidth * ratio, this.Link!.StrokeWidth * ratio)
      firstStartX += x
      firstStartY += y
    }
    {
      let { x, y } = this.rotateToSegmentDirection(segment, - this.Link!.EndArrowSize, this.Link!.EndArrowSize)

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

  ngAfterViewInit() {

    // const bbox = this.textElement?.nativeElement.getBBox()

    // if (bbox != undefined) {
    //   this.textWidth = bbox.width;
    //   this.textHeight = bbox.height;
    // }
  }

  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }
}
