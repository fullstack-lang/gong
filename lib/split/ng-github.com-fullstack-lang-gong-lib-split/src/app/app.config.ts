import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideNgtRenderer } from 'angular-three/dom';
import { provideRouter, withHashLocation } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { APP_BASE_HREF } from '@angular/common'; // 1. Import this

import { routes } from './app.routes';
import { wasmInterceptor } from './wasm.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes, withHashLocation()), 
    
    // 2. Set the base href dynamically to the exact current file path
    { provide: APP_BASE_HREF, useValue: window.location.pathname }, 
    
    provideHttpClient(withInterceptors([wasmInterceptor])),
    provideNgtRenderer()
  ]
};