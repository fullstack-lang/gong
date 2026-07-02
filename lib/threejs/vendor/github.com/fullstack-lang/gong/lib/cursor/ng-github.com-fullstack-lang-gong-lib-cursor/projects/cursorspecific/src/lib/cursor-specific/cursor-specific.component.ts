import { Component, Input, OnDestroy, OnInit } from '@angular/core';

import * as cursor from '../../../../cursor/src/public-api'

import { LayoutService } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/layout.service'; // Adjust path if needed
import { Subject, takeUntil } from 'rxjs';


@Component({
  selector: 'lib-cursor-specific',
  imports: [
    
  ],
  templateUrl: './cursor-specific.component.html',
  styleUrl: './cursor-specific.component.css'
})
export class CursorSpecificComponent implements OnInit, OnDestroy {

  @Input() Name: string = ""

  x = 0;
  xe = 500
  y = 0;
  ye = 0;
  private animationFrameId: number | null = null;  // Store animation frame ID

  public frontRepo?: cursor.FrontRepo;
  public cursor: cursor.Cursor | undefined

  controlsHeight: number = 0;
  private destroy$ = new Subject<void>();

  constructor(
    private frontRepoService: cursor.FrontRepoService,

    private layoutService: LayoutService
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    // Subscribe to the height changes from the service
    this.layoutService.controlsHeight$
      .pipe(takeUntil(this.destroy$)) // Unsubscribe automatically on component destruction
      .subscribe((height: number) => {
        if (this.controlsHeight !== height) { // Optional: Check if value actually changed
            this.controlsHeight = height;
            // console.log('SvgComponent: Controls height received', height);
            // Use the height here to adjust layout, transforms, viewbox, etc.
            // Example: setting margin-top (as shown in the template)
        }
      });


    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (gongtablesFrontRepo) => {
        this.frontRepo = gongtablesFrontRepo;

        let cursors = this.frontRepo.getFrontArray<cursor.Cursor>(cursor.Cursor.GONGSTRUCT_NAME);

        console.assert(cursors.length == 1);
        this.cursor = cursors[0];
        this.x = this.cursor.StartX
        this.xe = this.cursor.EndX

        this.y = this.cursor.Y1  + this.controlsHeight
        this.ye = this.cursor.Y2  + this.controlsHeight

        if (this.cursor.IsPlaying == true) {
          this.startEmittingPosition();
        } else {
          this.stopEmittingPosition();  // Stop the animation if IsPlaying is false
        }
      },
      error: (err) => {
        // this will capture any errors thrown by the Observable
        console.error("WebSocket error from Observable:", err);
      },
      complete: () => {
        console.log("WebSocket connection complete.");
      },
    });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  public startEmittingPosition(): void {
    if (this.cursor == undefined) {
      return;
    }

    const duration = 1000 * this.cursor.DurationSeconds;
    const progressAbciss = this.cursor.EndX - this.cursor.StartX;
    let startTime: number | null = null;
    let previousTimestamp: number | null = null;

    const animate = (timestamp: number) => {
      if (!startTime) {
        startTime = timestamp;
        previousTimestamp = timestamp;
      }

      // Calculate elapsed time since last frame
      const deltaTime = timestamp - previousTimestamp!;
      previousTimestamp = timestamp;

      // Subtract 50ms from the elapsed time calculation
      const elapsed = ((timestamp - startTime - 300) % duration);
      const progress = elapsed / duration;
      this.x = this.cursor!.StartX + progress * progressAbciss;

      this.animationFrameId = requestAnimationFrame(animate);
    }

    this.animationFrameId = requestAnimationFrame(animate);
  }

  /**
   * Stop emitting position by canceling the animation frame.
   */
  public stopEmittingPosition(): void {
    if (this.animationFrameId !== null) {
      cancelAnimationFrame(this.animationFrameId);
      this.animationFrameId = null;
    }
  }
}