import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RefreshService {

  private refreshRequestSource = new Subject<any>();
  refreshRequest$ = this.refreshRequestSource.asObservable();
  emitRefreshRequestEvent(any: any) {
    this.refreshRequestSource.next(any);
  }
}
