import { Injectable, Inject } from '@angular/core';
import { Observable } from 'rxjs';
import { DOCUMENT } from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class WebSocketService {

  constructor(@Inject(DOCUMENT) private document: Document) { }

  public connect(stackPath: string): Observable<any> {
    return new Observable<any>(subscriber => {

      // 1. WASM OFFLINE MODE INTERCEPT
      if ((window as any).openWasmSocket) {
        const onMessageFromGo = (message: string) => {
          subscriber.next(JSON.parse(message));
        };
        // Call the Go function exported in main_wasm.go
        (window as any).openWasmSocket(stackPath, onMessageFromGo);
        return; 
      }

      // 2. STANDARD NETWORK MODE
      let protocol = this.document.location.protocol === 'https:' ? 'wss://' : 'ws://';
      let port = this.document.location.port ? ':' + this.document.location.port : '';
      let host = this.document.location.hostname;
      
      const wsUrl = `${protocol}${host}${port}/api/github.com/fullstack-lang/gong/test/test4/go/v1/ws/stage?Name=${stackPath}`;
      
      const ws = new WebSocket(wsUrl);
      ws.onmessage = (event) => subscriber.next(JSON.parse(event.data));
      ws.onerror = (error) => subscriber.error(error);
      ws.onclose = () => subscriber.complete();

      return () => ws.close();
    });
  }
}