import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { ShapeMouseEvent } from './shape.mouse.event';

export type Coordinate = [number, number]

// Define the interface for the event object
interface RectMouseEvent {
  rectangleID: number;
  MousePosRelativeSVG: [number, number]
}

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  //
  // mouse ALT events
  //

  private mouseRectAltKeyMouseDownEventSource = new Subject<ShapeMouseEvent>();
  mouseRectAltKeyMouseDownEvent$ = this.mouseRectAltKeyMouseDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(ShapeMouseEvent: ShapeMouseEvent) {
    this.mouseRectAltKeyMouseDownEventSource.next(ShapeMouseEvent);
  }

  private mouseRectAltKeyMouseDragEventSource = new Subject<ShapeMouseEvent>()
  mouseRectAltKeyMouseDragEvent$ = this.mouseRectAltKeyMouseDragEventSource.asObservable()
  emitRectAltKeyMouseDragEvent(shapeMouseEvent: ShapeMouseEvent) {
    this.mouseRectAltKeyMouseDragEventSource.next(shapeMouseEvent)
  }

  private mouseRectAltKeyMouseUpEventSource = new Subject<number>();
  mouseRectAltKeyMouseUpEvent$ = this.mouseRectAltKeyMouseUpEventSource.asObservable();
  emitMouseRectAltKeyMouseUpEvent(rectangleID: number) {
    this.mouseRectAltKeyMouseUpEventSource.next(rectangleID);
  }

}