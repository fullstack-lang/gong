package angular

const WebSocketServiceTemplate = `import { Injectable, Inject } from '@angular/core';
import { Observable } from 'rxjs';
import { DOCUMENT } from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class WebSocketService {

  constructor(@Inject(DOCUMENT) private document: Document) { }

  public connect(stackPath: string): Observable<any> {
    return new Observable<any>(subscriber => {

      // Determine if we are running in the offline HTML file
      const isOfflineMode = this.document.location.protocol === 'file:';

      const attemptConnection = (retries: number) => {
        // 1. WASM OFFLINE MODE (Check if Go is ready)
        if ((window as any).openWasmSocket) {
          const onMessageFromGo = (message: string) => {
            subscriber.next(JSON.parse(message));
          };
          (window as any).openWasmSocket(stackPath, onMessageFromGo);
          return; // Successfully connected to WASM
        } 
        
        // 2. STILL WAITING FOR WASM
        if (isOfflineMode && retries > 0) {
          // If we are offline, WASM might just be decoding. Wait 100ms and check again.
          setTimeout(() => attemptConnection(retries - 1), 100);
          return;
        }

        // 3. STANDARD SERVER MODE (Fallback)
        if (!isOfflineMode) {
          let protocol = this.document.location.protocol === 'https:' ? 'wss://' : 'ws://';
          let port = this.document.location.port ? ':' + this.document.location.port : '';
          let host = this.document.location.hostname;
          
          const wsUrl = ` + "`" + `${protocol}${host}${port}/api/github.com/fullstack-lang/gong/test/test4/go/v1/ws/stage?Name=${stackPath}` + "`" + `;
          
          const ws = new WebSocket(wsUrl);
          ws.onmessage = (event) => subscriber.next(JSON.parse(event.data));
          ws.onerror = (error) => subscriber.error(error);
          ws.onclose = () => subscriber.complete();
        } else {
           subscriber.error("Offline mode detected, but WASM backend failed to load.");
        }
      };

      // Start the connection loop, allowing up to 5 seconds (50 * 100ms) for WASM to boot
      attemptConnection(50); 
    });
  }
}`
