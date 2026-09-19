import {
  Component,
  ElementRef,
  Input,
  Output,
  EventEmitter,
  ContentChildren,
  QueryList,
  AfterContentInit,
  OnDestroy,
  OnChanges,
  SimpleChanges,
  Renderer2,
  ChangeDetectionStrategy,
  HostBinding,
  inject
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { Subscription } from 'rxjs';
import { GongSplitAreaComponent } from './gong-split-area.component';

export { GongSplitAreaComponent };

export interface SplitGutterClickEvent {
  gutterNum: number;
}

export interface SplitDragEvent {
  gutterNum: number;
  sizes: (number | string | null)[];
}

@Component({
  selector: 'gong-split, as-split',
  standalone: true,
  imports: [CommonModule],
  template: `<ng-content></ng-content>`,
  styleUrl: './gong-split.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class GongSplitComponent implements AfterContentInit, OnDestroy, OnChanges {
  private elementRef = inject<ElementRef<HTMLElement>>(ElementRef);
  private renderer = inject(Renderer2);

  @Input() direction: 'horizontal' | 'vertical' | string = 'horizontal';
  @Input() unit: 'percent' | 'pixel' | string = 'percent';
  @Input() gutterSize: number | null = 6;
  @Input() disabled: boolean = false;

  @Output() dragStart = new EventEmitter<SplitDragEvent>();
  @Output() dragEnd = new EventEmitter<SplitDragEvent>();
  @Output() gutterClick = new EventEmitter<SplitGutterClickEvent>();
  @Output() transitionEnd = new EventEmitter<void>();

  @ContentChildren(GongSplitAreaComponent, { descendants: false })
  areas!: QueryList<GongSplitAreaComponent>;

  private areasSubscription?: Subscription;
  private gutterElements: HTMLElement[] = [];
  private unlistenGutterFns: (() => void)[] = [];

  private isDragging = false;
  private activeGutterIndex = -1;
  private dragStartPos = 0;
  private dragContainerSize = 0;
  private dragInitialPxSizes: number[] = [];
  private dragInitialPctSizes: number[] = [];
  private dragAvailablePx = 0;

  get normalizedDirection(): 'horizontal' | 'vertical' {
    if (this.direction === 'vertical' || this.direction === 'Vertical') {
      return 'vertical';
    }
    return 'horizontal';
  }

  get normalizedUnit(): 'percent' | 'pixel' {
    if (this.unit === 'pixel' || this.unit === 'Pixel') {
      return 'pixel';
    }
    return 'percent';
  }

  get effectiveGutterSize(): number {
    return this.gutterSize !== null && this.gutterSize !== undefined && this.gutterSize >= 0
      ? this.gutterSize
      : 6;
  }

  @HostBinding('class.gong-split-horizontal')
  get isHorizontalClass(): boolean {
    return this.normalizedDirection === 'horizontal';
  }

  @HostBinding('class.gong-split-vertical')
  get isVerticalClass(): boolean {
    return this.normalizedDirection === 'vertical';
  }

  @HostBinding('class.gong-split-dragging')
  get isDraggingClass(): boolean {
    return this.isDragging;
  }

  ngAfterContentInit(): void {
    this.refreshAreasAndGutters();

    this.areasSubscription = this.areas.changes.subscribe(() => {
      this.refreshAreasAndGutters();
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['direction'] || changes['unit'] || changes['gutterSize']) {
      if (this.areas) {
        this.refreshAreasAndGutters();
      }
    }
  }

  ngOnDestroy(): void {
    if (this.areasSubscription) {
      this.areasSubscription.unsubscribe();
    }
    this.cleanupGutters();
  }

  private refreshAreasAndGutters(): void {
    this.cleanupGutters();

    if (!this.areas || this.areas.length === 0) {
      return;
    }

    const areaList = this.areas.toArray();
    this.applyInitialAreaSizes(areaList);

    if (areaList.length <= 1) {
      return;
    }

    const host = this.elementRef.nativeElement;
    const isHoriz = this.normalizedDirection === 'horizontal';
    const gSize = this.effectiveGutterSize;

    for (let i = 0; i < areaList.length - 1; i++) {
      const gutter = this.renderer.createElement('div') as HTMLElement;
      this.renderer.addClass(gutter, 'gong-split-gutter');
      this.renderer.addClass(gutter, 'as-split-gutter');

      if (isHoriz) {
        this.renderer.setStyle(gutter, 'width', `${gSize}px`);
        this.renderer.setStyle(gutter, 'height', '100%');
      } else {
        this.renderer.setStyle(gutter, 'height', `${gSize}px`);
        this.renderer.setStyle(gutter, 'width', '100%');
      }

      const icon = this.renderer.createElement('div') as HTMLElement;
      this.renderer.addClass(icon, 'gong-split-gutter-icon');
      this.renderer.appendChild(gutter, icon);

      const nextAreaEl = areaList[i + 1].elementRef.nativeElement;
      this.renderer.insertBefore(host, gutter, nextAreaEl);

      const unlistenPointerDown = this.renderer.listen(
        gutter,
        'pointerdown',
        (event: PointerEvent) => this.onGutterPointerDown(event, i)
      );

      const unlistenClick = this.renderer.listen(
        gutter,
        'click',
        (event: MouseEvent) => {
          this.gutterClick.emit({ gutterNum: i + 1 });
        }
      );

      this.gutterElements.push(gutter);
      this.unlistenGutterFns.push(unlistenPointerDown, unlistenClick);
    }
  }

  private cleanupGutters(): void {
    for (const fn of this.unlistenGutterFns) {
      try {
        fn();
      } catch (_) {}
    }
    this.unlistenGutterFns = [];

    for (const gutter of this.gutterElements) {
      if (gutter.parentNode) {
        gutter.parentNode.removeChild(gutter);
      }
    }
    this.gutterElements = [];
  }

  private applyInitialAreaSizes(areaList: GongSplitAreaComponent[]): void {
    const dir = this.normalizedDirection;
    const u = this.normalizedUnit;

    const visibleAreas = areaList.filter(a => a.visible);
    if (visibleAreas.length === 0) return;

    if (u === 'percent') {
      let totalAssignedPercent = 0;
      let wildcardCount = 0;

      for (const area of visibleAreas) {
        const num = area.parsedNumericSize;
        if (num !== null) {
          totalAssignedPercent += num;
        } else {
          wildcardCount++;
        }
      }

      if (wildcardCount === 0 && totalAssignedPercent > 0) {
        for (const area of visibleAreas) {
          const num = area.parsedNumericSize ?? (100 / visibleAreas.length);
          const normalizedPct = (num / totalAssignedPercent) * 100;
          area.applyStyle(dir, u, normalizedPct);
        }
      } else if (wildcardCount === 0 && totalAssignedPercent === 0) {
        const equalPct = 100 / visibleAreas.length;
        for (const area of visibleAreas) {
          area.applyStyle(dir, u, equalPct);
        }
      } else {
        const remainingPct = Math.max(0, 100 - totalAssignedPercent);
        const perWildcardPct = wildcardCount > 0 ? remainingPct / wildcardCount : 0;
        for (const area of visibleAreas) {
          if (area.parsedNumericSize !== null) {
            area.applyStyle(dir, u, area.parsedNumericSize);
          } else {
            area.applyStyle(dir, u, perWildcardPct > 0 ? perWildcardPct : '*');
          }
        }
      }
    } else {
      // Pixel mode
      for (const area of visibleAreas) {
        if (area.parsedNumericSize !== null) {
          area.applyStyle(dir, u, area.parsedNumericSize);
        } else {
          area.applyStyle(dir, u, '*');
        }
      }
    }
  }

  private onGutterPointerDown(event: PointerEvent, gutterIndex: number): void {
    if (this.disabled || event.button !== 0) {
      return;
    }

    event.preventDefault();
    event.stopPropagation();

    const gutterEl = event.currentTarget as HTMLElement;
    if (gutterEl && typeof gutterEl.setPointerCapture === 'function') {
      try {
        gutterEl.setPointerCapture(event.pointerId);
      } catch (_) {}
    }

    this.isDragging = true;
    this.activeGutterIndex = gutterIndex;

    const host = this.elementRef.nativeElement;
    const isHoriz = this.normalizedDirection === 'horizontal';

    this.dragStartPos = isHoriz ? event.clientX : event.clientY;
    this.dragContainerSize = isHoriz ? host.clientWidth : host.clientHeight;

    const areaList = this.areas.toArray().filter(a => a.visible);
    const numGutters = this.gutterElements.length;
    const totalGutterPx = numGutters * this.effectiveGutterSize;
    this.dragAvailablePx = Math.max(1, this.dragContainerSize - totalGutterPx);

    this.dragInitialPxSizes = areaList.map(a => {
      const el = a.elementRef.nativeElement;
      return isHoriz ? el.offsetWidth : el.offsetHeight;
    });

    const sumInitialPx = this.dragInitialPxSizes.reduce((acc, px) => acc + px, 0);
    this.dragInitialPctSizes = this.dragInitialPxSizes.map(px =>
      sumInitialPx > 0 ? (px / sumInitialPx) * 100 : 100 / areaList.length
    );

    this.renderer.addClass(gutterEl, 'gong-split-gutter-active');

    this.dragStart.emit({
      gutterNum: gutterIndex + 1,
      sizes: areaList.map(a => a.appliedSize)
    });

    const unlistenMove = this.renderer.listen('window', 'pointermove', (e: PointerEvent) => {
      this.onPointerMove(e);
    });

    const onPointerUp = (e: PointerEvent) => {
      unlistenMove();
      unlistenUp();
      unlistenCancel();
      this.onPointerUp(e, gutterEl);
    };

    const unlistenUp = this.renderer.listen('window', 'pointerup', onPointerUp);
    const unlistenCancel = this.renderer.listen('window', 'pointercancel', onPointerUp);
  }

  private onPointerMove(event: PointerEvent): void {
    if (!this.isDragging || this.activeGutterIndex < 0) {
      return;
    }

    const areaList = this.areas.toArray().filter(a => a.visible);
    const i = this.activeGutterIndex;
    if (i < 0 || i >= areaList.length - 1) {
      return;
    }

    const isHoriz = this.normalizedDirection === 'horizontal';
    const currentPos = isHoriz ? event.clientX : event.clientY;
    const deltaPx = currentPos - this.dragStartPos;

    const dir = this.normalizedDirection;
    const u = this.normalizedUnit;
    const areaA = areaList[i];
    const areaB = areaList[i + 1];

    if (u === 'pixel') {
      const startPxA = this.dragInitialPxSizes[i];
      const startPxB = this.dragInitialPxSizes[i + 1];
      const minA = areaA.minSize || 20;
      const minB = areaB.minSize || 20;

      if (areaA.isWildcard && !areaB.isWildcard) {
        const newPxB = Math.max(minB, startPxB - deltaPx);
        areaB.applyStyle(dir, u, newPxB);
      } else if (!areaA.isWildcard && areaB.isWildcard) {
        const newPxA = Math.max(minA, startPxA + deltaPx);
        areaA.applyStyle(dir, u, newPxA);
      } else {
        const combined = startPxA + startPxB;
        let newPxA = Math.max(minA, Math.min(combined - minB, startPxA + deltaPx));
        let newPxB = combined - newPxA;
        areaA.applyStyle(dir, u, newPxA);
        areaB.applyStyle(dir, u, newPxB);
      }
    } else {
      // Percentage
      const startPctA = this.dragInitialPctSizes[i];
      const startPctB = this.dragInitialPctSizes[i + 1];
      const combinedPct = startPctA + startPctB;

      const deltaPct = (deltaPx / this.dragAvailablePx) * 100;
      const minPctA = areaA.minSize ? (areaA.minSize / this.dragAvailablePx) * 100 : 2;
      const minPctB = areaB.minSize ? (areaB.minSize / this.dragAvailablePx) * 100 : 2;

      let newPctA = Math.max(minPctA, Math.min(combinedPct - minPctB, startPctA + deltaPct));
      let newPctB = combinedPct - newPctA;

      areaA.applyStyle(dir, u, newPctA);
      areaB.applyStyle(dir, u, newPctB);
    }
  }

  private onPointerUp(event: PointerEvent, gutterEl: HTMLElement): void {
    if (!this.isDragging) {
      return;
    }

    if (gutterEl && typeof gutterEl.releasePointerCapture === 'function') {
      try {
        gutterEl.releasePointerCapture(event.pointerId);
      } catch (_) {}
    }

    this.renderer.removeClass(gutterEl, 'gong-split-gutter-active');
    this.isDragging = false;

    const areaList = this.areas.toArray().filter(a => a.visible);
    this.dragEnd.emit({
      gutterNum: this.activeGutterIndex + 1,
      sizes: areaList.map(a => a.appliedSize)
    });

    this.activeGutterIndex = -1;
  }
}
