package angular

const NgFileAppConfig = `import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, withHashLocation } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { APP_BASE_HREF } from '@angular/common';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideNgtRenderer } from 'angular-three/dom';

import { routes } from './app.routes';
// import { wasmInterceptor } from './wasm.interceptor'; // Assuming this might not be available globally, but wait! The original template didn't have it.

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes, withHashLocation()), 
    provideAnimationsAsync(),
    provideHttpClient(),
    provideNgtRenderer(),
  ]
};
`
