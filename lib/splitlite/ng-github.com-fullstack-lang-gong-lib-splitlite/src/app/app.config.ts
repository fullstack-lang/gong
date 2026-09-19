import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, withHashLocation } from '@angular/router';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { APP_BASE_HREF } from '@angular/common';
import { OverlayContainer } from '@angular/cdk/overlay';

import { routes } from './app.routes';
import { wasmInterceptor } from './wasm.interceptor';
import { CustomOverlayContainer } from './custom-overlay-container';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes, withHashLocation()),
    provideAnimationsAsync(),
    { provide: APP_BASE_HREF, useValue: window.location.pathname },
    provideHttpClient(withInterceptors([wasmInterceptor])),
    { provide: OverlayContainer, useClass: CustomOverlayContainer }
  ]
};
