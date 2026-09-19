import { Injectable } from '@angular/core';
import { OverlayContainer } from '@angular/cdk/overlay';

@Injectable({ providedIn: 'root' })
export class CustomOverlayContainer extends OverlayContainer {
  protected override _createContainer(): void {
    super._createContainer();
  }
}
