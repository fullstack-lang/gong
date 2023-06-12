import { Component, OnInit, Input, OnDestroy, DoCheck, OnChanges, SimpleChanges } from '@angular/core';
import { ElementRef, AfterViewInit } from '@angular/core';


import { RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { Subscription } from 'rxjs';
import { ShapeMouseEvent } from '../shape.mouse.event';
import { createPoint } from '../link/draw.segments';
import { MouseEventService } from '../mouse-event.service';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit, OnDestroy, DoCheck, OnChanges {

  @Input() Rect: gongsvg.RectDB = new gongsvg.RectDB()
  @Input() GONG__StackPath: string = ""

  // for use in the template
  RectAnchorType = gongsvg.RectAnchorType

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' | null = null;

  rectDragging: boolean = false;

  // RectAtMouseDown is the clone of the Rect when mouse down
  private RectAtMouseDown: gongsvg.RectDB | undefined

  // to compute wether it was a select / dragging event
  distanceMoved = 0
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private dragThreshold = 5;

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private rectService: gongsvg.RectService,
    private rectangleEventService: RectangleEventService,
    private mouseEventService: MouseEventService,
    private svgEventService: SvgEventService,
    private isEditableService: IsEditableService,
    private refreshService: RefreshService,
  ) {

    this.subscriptions.push(
      mouseEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (this.anchorDragging || this.rectDragging || this.Rect.IsSelected) {

            console.log("rect: mouseMouseDownEvent, ", this.Rect!.Name)

            this.distanceMoved = 0
            this.RectAtMouseDown = structuredClone(this.Rect)
            this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)
          }
        }
      )
    )

    this.subscriptions.push(
      mouseEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (!this.isEditableService.getIsEditable()) {
            return
          }

          if (this.anchorDragging) {
            if (this.activeAnchor === 'left') {
              this.Rect.X = this.RectAtMouseDown!.X + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
              this.Rect.Width = this.RectAtMouseDown!.Width - (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'right') {
              this.Rect.Width = this.RectAtMouseDown!.Width + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'top') {
              this.Rect.Y = this.RectAtMouseDown!.Y + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
              this.Rect.Height = this.RectAtMouseDown!.Height - (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            } else if (this.activeAnchor === 'bottom') {
              this.Rect.Height = this.RectAtMouseDown!.Height + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            }
            return // we don't want the move move to be interpreted by the rect
          }

          if (this.PointAtMouseDown && (this.rectDragging || this.Rect.IsSelected)) {
            const deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
            const deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
            this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

            if (this.Rect?.CanMoveHorizontaly) {
              this.Rect.X = this.RectAtMouseDown!.X + deltaX
            }
            if (this.Rect?.CanMoveVerticaly) {
              this.Rect.Y = this.RectAtMouseDown!.Y + deltaY
            }
          }
        }
      )
    )

    this.subscriptions.push(
      mouseEventService.mouseMouseUpEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (!this.isEditableService.getIsEditable()) {
            return
          }

          if (shapeMouseEvent.ShapeID != 0 && this.distanceMoved > this.dragThreshold) {
            if (this.anchorDragging || this.rectDragging || this.Rect.IsSelected) {
              this.Rect.IsSelected = false
              this.manageHandles()
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
                _ => {
                  this.refreshService.emitRefreshRequestEvent(0)
                }
              )
            }

          }
          if (this.distanceMoved <= this.dragThreshold) {
            if (this.Rect?.IsSelectable &&
              shapeMouseEvent.ShapeType == gongsvg.RectDB.GONGSTRUCT_NAME &&
              shapeMouseEvent.ShapeID == this.Rect.ID) {
              console.log("rect, mouseEventService.mouseMouseUpEvent$.subscribe, from the shape: ", this.Rect?.Name)
              this.Rect.IsSelected = !this.Rect.IsSelected
              this.manageHandles()
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
                _ => {
                  this.refreshService.emitRefreshRequestEvent(0)
                }
              )
            }

            // mouseup emited from the background let to unselect selected shapes
            if (this.Rect.IsSelectable && this.Rect.IsSelected && shapeMouseEvent.ShapeID == 0) {
              console.log("rect, mouseEventService.mouseMouseUpEvent$.subscribe: from the svg", this.Rect?.Name)
              this.Rect.IsSelected = false
              this.manageHandles()
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
                _ => {
                  this.refreshService.emitRefreshRequestEvent(0)
                }
              )
            }
          }

          this.rectDragging = false
          this.anchorDragging = false
          // console.log('Rect ', this.Rect.Name, 'Mouse down event occurred on rectangle ', rectangleID);
        })
    )

    this.subscriptions.push(
      svgEventService.multiShapeSelectEndEvent$.subscribe(
        (selectAreaConfig: SelectAreaConfig) => {

          if (this.Rect.IsSelected) {
            return
          }

          switch (selectAreaConfig.SweepDirection) {
            case SweepDirection.LEFT_TO_RIGHT:

              // rectangle has to be in boxed in the rect
              if (
                this.Rect.X > selectAreaConfig.TopLeft[0] &&
                this.Rect.X + this.Rect.Width < selectAreaConfig.BottomRigth[0] &&
                this.Rect.Y > selectAreaConfig.TopLeft[1] &&
                this.Rect.Y + this.Rect.Height < selectAreaConfig.BottomRigth[1]
              ) {
                this.Rect.IsSelected = true
                this.manageHandles()
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
                  _ => {
                    this.refreshService.emitRefreshRequestEvent(0)
                  }
                )
              }
              break
            case SweepDirection.RIGHT_TO_LEFT:

              // rectangle has to be partialy boxed in the rect
              if (
                this.Rect.X < selectAreaConfig.BottomRigth[0] &&
                this.Rect.X + this.Rect.Width > selectAreaConfig.TopLeft[0] &&
                this.Rect.Y < selectAreaConfig.BottomRigth[1] &&
                this.Rect.Y + this.Rect.Height > selectAreaConfig.TopLeft[1]
              ) {
                this.Rect.IsSelected = true
                this.manageHandles()
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
                  _ => {
                    this.refreshService.emitRefreshRequestEvent(0)
                  }
                )
              }
              break
          }
        })
    )
  }

  ngOnInit(): void {
    this.rectDragging = false
    this.anchorDragging = false
  }

  ngOnDestroy() {
    gongsvg.RectAnchorType.RECT_ANCHOR_TOP
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  ngDoCheck(): void {
    // console.log("DoCheck")
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes['Link']) {
      // console.log('Previous value: ', changes['Link'].previousValue);
      // console.log('Current value: ', changes['Link'].currentValue);
    }
  }

  rectMouseDown(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    } else {
      console.log("startRectDrag + shiftKey : ", this.Rect?.Name)

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.rectangleEventService.emitRectAltKeyMouseDownEvent(shapeMouseEvent)

    }
  }

  rectMouseMove(event: MouseEvent): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: this.Rect.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    // console.log("RectComponent DragRect : ", deltaX, deltaY, distanceMoved)

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', this.Rect.Name, event.clientX, event.clientY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(shapeMouseEvent)
      return
    }

    if (event.shiftKey) {
      this.svgEventService.emitMultiShapeSelectDrag(shapeMouseEvent)
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  rectMouseUp(event: MouseEvent): void {

    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: this.Rect!.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    if (!event.altKey && !event.shiftKey) {
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }

    if (event.altKey) {
      console.log("endRectDrag + shiftKey rect : ", this.Rect?.Name)

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(this.Rect.ID);
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent(shapeMouseEvent)
    }
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  anchorMouseMove(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  anchorMouseUp(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = null;
      this.Rect!.IsSelected = false
      this.manageHandles()
      this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe(
        _ => {
          this.refreshService.emitRefreshRequestEvent(0)
        }
      )
    }
  }

  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }

  // display or not handles if selected or not
  manageHandles() {

    if (this.Rect!.IsSelected && this.Rect!.CanHaveLeftHandle) {
      this.Rect!.HasLeftHandle = true
    } else {
      this.Rect!.HasLeftHandle = false
    }
    if (this.Rect!.IsSelected && this.Rect!.CanHaveRightHandle) {
      this.Rect!.HasRightHandle = true
    } else {
      this.Rect!.HasRightHandle = false
    }
    if (this.Rect!.IsSelected && this.Rect!.CanHaveTopHandle) {
      this.Rect!.HasTopHandle = true
    } else {
      this.Rect!.HasTopHandle = false
    }
    if (this.Rect!.IsSelected && this.Rect!.CanHaveBottomHandle) {
      this.Rect!.HasBottomHandle = true
    } else {
      this.Rect!.HasBottomHandle = false
    }
  }
}
