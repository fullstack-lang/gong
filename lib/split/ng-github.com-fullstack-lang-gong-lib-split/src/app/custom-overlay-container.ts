import { Injectable } from '@angular/core';
import { OverlayContainer } from '@angular/cdk/overlay';

@Injectable({ providedIn: 'root' })
export class CustomOverlayContainer extends OverlayContainer {
  protected override _createContainer(): void {
    console.log("CustomOverlayContainer: _createContainer called");
    console.log("CustomOverlayContainer: document.body before append:", document.body);
    const container = document.createElement('div');
    container.classList.add('cdk-overlay-container');
    document.body.appendChild(container);
    this._containerElement = container;
    console.log("CustomOverlayContainer: appended", container, "to body");
    console.log("CustomOverlayContainer: stack trace:", new Error().stack);
  }
}
