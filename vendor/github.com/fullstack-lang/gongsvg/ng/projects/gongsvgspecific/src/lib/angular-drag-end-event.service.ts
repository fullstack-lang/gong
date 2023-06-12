import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AngularDragEndEventService {

  private angularSplitDragEndSource = new Subject<number>();
  angularSplitDragEndEvent$ = this.angularSplitDragEndSource.asObservable();
  emitMouseUpEvent(x: number) {
    this.angularSplitDragEndSource.next(x);
  }
}
