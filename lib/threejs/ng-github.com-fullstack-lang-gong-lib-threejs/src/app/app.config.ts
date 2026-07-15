import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideHttpClient } from '@angular/common/http';
import { provideNgtRenderer } from 'angular-three/dom';
import { OverlayContainer } from '@angular/cdk/overlay';
import { CustomOverlayContainer } from './custom-overlay-container';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideAnimationsAsync(),
    provideHttpClient(),
    provideNgtRenderer(),
    { provide: OverlayContainer, useClass: CustomOverlayContainer }
  ]
};
