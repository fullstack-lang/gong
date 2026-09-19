import { Component, ElementRef, Input, ChangeDetectionStrategy, HostBinding } from '@angular/core';

@Component({
  selector: 'gong-split-area, as-split-area',
  standalone: true,
  template: `<ng-content></ng-content>`,
  styles: [`
    :host {
      display: block;
      position: relative;
      overflow: auto;
      box-sizing: border-box;
      min-width: 0;
      min-height: 0;
    }
  `],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class GongSplitAreaComponent {
  @Input() size: number | string | null = null;
  @Input() minSize: number = 0;
  @Input() maxSize: number | null = null;
  @Input() order: number | null = null;
  @Input() visible: boolean = true;

  public appliedSize: number | string | null = null;

  constructor(public elementRef: ElementRef<HTMLElement>) {}

  get isWildcard(): boolean {
    return (
      this.size === '*' ||
      this.size === '$any(*)' ||
      this.size === "$any('*')" ||
      this.size === null ||
      this.size === undefined ||
      this.size === ''
    );
  }

  get parsedNumericSize(): number | null {
    if (this.isWildcard) return null;
    const num = typeof this.size === 'number' ? this.size : parseFloat(String(this.size));
    return isNaN(num) ? null : num;
  }

  @HostBinding('style.order')
  get hostOrder(): number | null {
    return this.order;
  }

  @HostBinding('style.display')
  get hostDisplay(): string {
    return this.visible ? 'block' : 'none';
  }

  public applyStyle(
    direction: 'horizontal' | 'vertical',
    unit: 'percent' | 'pixel',
    sizeValue: number | string | null
  ): void {
    this.appliedSize = sizeValue;
    const el = this.elementRef.nativeElement;

    if (sizeValue === '*' || sizeValue === null || this.isWildcard) {
      el.style.flex = '1 1 0px';
      if (direction === 'horizontal') {
        el.style.width = 'auto';
        el.style.maxWidth = 'none';
        el.style.minWidth = this.minSize ? `${this.minSize}px` : '0px';
        el.style.height = '100%';
        el.style.maxHeight = 'none';
        el.style.minHeight = '0px';
      } else {
        el.style.height = 'auto';
        el.style.maxHeight = 'none';
        el.style.minHeight = this.minSize ? `${this.minSize}px` : '0px';
        el.style.width = '100%';
        el.style.maxWidth = 'none';
        el.style.minWidth = '0px';
      }
    } else if (unit === 'pixel') {
      const px = typeof sizeValue === 'number' ? sizeValue : parseFloat(String(sizeValue));
      el.style.flex = `0 0 ${px}px`;
      if (direction === 'horizontal') {
        el.style.width = `${px}px`;
        el.style.maxWidth = `${px}px`;
        el.style.minWidth = this.minSize ? `${this.minSize}px` : '0px';
        el.style.height = '100%';
      } else {
        el.style.height = `${px}px`;
        el.style.maxHeight = `${px}px`;
        el.style.minHeight = this.minSize ? `${this.minSize}px` : '0px';
        el.style.width = '100%';
      }
    } else {
      const pct = typeof sizeValue === 'number' ? sizeValue : parseFloat(String(sizeValue));
      el.style.flex = `0 0 ${pct}%`;
      if (direction === 'horizontal') {
        el.style.width = `${pct}%`;
        el.style.maxWidth = `${pct}%`;
        el.style.minWidth = this.minSize ? `${this.minSize}%` : '0px';
        el.style.height = '100%';
      } else {
        el.style.height = `${pct}%`;
        el.style.maxHeight = `${pct}%`;
        el.style.minHeight = this.minSize ? `${this.minSize}%` : '0px';
        el.style.width = '100%';
      }
    }
  }
}
