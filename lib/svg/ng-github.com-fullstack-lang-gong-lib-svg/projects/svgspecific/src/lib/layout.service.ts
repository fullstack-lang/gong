import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root' // Makes it a singleton available app-wide
})
export class LayoutService {

  // BehaviorSubject holds the latest height value and emits it to new subscribers.
  // Initialize with 0 or null.
  private controlsHeightSubject = new BehaviorSubject<number>(0);

  // Expose the height as an Observable for other components to subscribe to.
  public readonly controlsHeight$: Observable<number> = this.controlsHeightSubject.asObservable();

  constructor() { }

  /**
   * Called by the component that owns the div to update the height.
   * @param height - The new height of the controls div.
   */
  updateControlsHeight(height: number): void {
    // Only emit if the height has actually changed to avoid unnecessary updates
    if (height !== this.controlsHeightSubject.getValue()) {
      this.controlsHeightSubject.next(height);
      console.log('LayoutService: Controls height updated to', height);
    }
  }

  /**
   * Optionally, provide a method to get the current value synchronously,
   * though subscribing to the Observable is generally preferred for reactivity.
   */
  getCurrentControlsHeight(): number {
    return this.controlsHeightSubject.getValue();
  }
}