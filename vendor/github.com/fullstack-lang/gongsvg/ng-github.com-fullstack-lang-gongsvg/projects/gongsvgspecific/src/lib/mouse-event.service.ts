import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { ShapeMouseEvent } from './shape.mouse.event';

@Injectable({
  providedIn: 'root'
})
export class MouseEventService {

  constructor() { }

  private mouseMouseDownEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseDownEvent$ = this.mouseMouseDownEventSource.asObservable();
  emitMouseDownEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse Down event', ShapeMouseEvent.shapeID)
    this.mouseMouseDownEventSource.next(ShapeMouseEvent);
  }

  private mouseMouseMoveEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseMoveEvent$ = this.mouseMouseMoveEventSource.asObservable();
  emitMouseMoveEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse move event', ShapeMouseEvent.shapeID)
    this.mouseMouseMoveEventSource.next(ShapeMouseEvent);
  }

  private mouseMouseUpEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseUpEvent$ = this.mouseMouseUpEventSource.asObservable();
  emitMouseUpEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse Up event', ShapeMouseEvent.shapeID)
    this.mouseMouseUpEventSource.next(ShapeMouseEvent);
  }
}
