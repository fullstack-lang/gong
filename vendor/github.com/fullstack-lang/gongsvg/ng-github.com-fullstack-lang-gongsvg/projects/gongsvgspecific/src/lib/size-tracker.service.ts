import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SizeTrackerService {
  private width: number = 0
  private height: number = 0

  constructor() {
    this.updateWindowSize();
    window.addEventListener('resize', () => this.updateWindowSize());
  }

  private updateWindowSize() {
    this.width = window.innerWidth;
    this.height = window.innerHeight;
  }

  getWidth(): number {
    return this.width;
  }

  getHeight(): number {
    return this.height;
  }
}
