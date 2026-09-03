import { Component, OnInit, OnDestroy, NgZone, Input, ChangeDetectorRef } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';

import * as Tone from 'tone';
import * as tonelocal from '../../../../tone/src/public-api';

import { Subject } from 'rxjs';
import { takeUntil, catchError } from 'rxjs/operators';
import { SALAMANDER_SAMPLES } from './salamander-samples';

@Component({
  selector: 'lib-tone-specific',
  imports: [
    MatButtonModule,
    MatDividerModule,
    MatIconModule
  ],
  templateUrl: './tone-specific.component.html',
  styleUrl: './tone-specific.component.css'
})
export class ToneSpecificComponent {

  @Input() Name: string = ""

  private synth: Tone.PolySynth | undefined;
  private sampler: Tone.Sampler | undefined;
  private currentLoop: Tone.Loop | undefined;
  private destroy$ = new Subject<void>();


  frontRepo?: tonelocal.FrontRepo;
  isLoading = false;
  isPlaying = false;

  constructor(
    private frontRepoService: tonelocal.FrontRepoService,
    private playerService: tonelocal.PlayerService,
    private ngZone: NgZone,
    private snackBar: MatSnackBar,
    private cdr: ChangeDetectorRef,
  ) { }

  ngOnInit(): void {
    this.initializeSynth();
    this.connectToWebSocket();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
    this.stopPlayback();
  }

  private initializeSynth(): void {
    try {
      this.synth = new Tone.PolySynth().toDestination();
    } catch (error) {
      this.handleAudioInitError(error);
    }
  }

  private connectToWebSocket(): void {
    this.frontRepoService.connectToWebSocket(this.Name)
      .pipe(
        takeUntil(this.destroy$),
        catchError(error => {
          this.handleWebSocketError(error);
          throw error;
        })
      )
      .subscribe(gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo;
        this.cdr.markForCheck();
      });
  }

  stopPlayback(): void {
    this.ngZone.runOutsideAngular(() => {
      try {
        // Stop and dispose of the current loop
        if (this.currentLoop) {
          this.currentLoop.stop();
          this.currentLoop.dispose();
          this.currentLoop = undefined;
        }

        // Stop the transport and dispose of the sampler
        Tone.getTransport().stop();
        Tone.getTransport().cancel(); // Cancel all scheduled events
        if (this.sampler) {
          this.sampler.dispose();
          this.sampler = undefined;
        }

        this.isPlaying = false;
      } catch (error) {
        console.error('Error stopping playback:', error);
      }
    });

    const players = this.frontRepo?.getFrontArray<tonelocal.Player>(tonelocal.Player.GONGSTRUCT_NAME);
    if (players && players.length === 1) {
      const player = players[0];
      player.Status = tonelocal.Status.PAUSED;
      this.playerService.updateFront(player, this.Name).subscribe(
        () => {
          console.log("gongtone: status set to PAUSED");
        }
      );
    }
  }

  async play(): Promise<void> {
    if (!this.frontRepo) {
      this.showError('No data available for playback');
      return;
    }

    try {
      await Tone.start();
    } catch (e) {
      console.warn('Tone.start() error:', e);
    }

    // Stop any existing playback before starting new one
    this.stopPlayback();

    this.isLoading = true;
    this.ngZone.runOutsideAngular(() => {
      try {
        const notes = this.frontRepo!.getFrontArray<tonelocal.Note>(tonelocal.Note.GONGSTRUCT_NAME);
        const duration = this.calculateTotalDuration(notes);
        this.initializeSampler(duration, notes);
      } catch (error) {
        this.handlePlaybackError(error);
      }
    });

    const players = this.frontRepo.getFrontArray<tonelocal.Player>(tonelocal.Player.GONGSTRUCT_NAME);
    if (players.length === 1) {
      const player = players[0];
      player.Status = tonelocal.Status.PLAYING;
      this.playerService.updateFront(player, this.Name).subscribe(
        () => {
          console.log("gongtone: status set to PLAYING");
        }
      );
    }
  }

  private initializeSampler(duration: number, notes: tonelocal.Note[]): void {
    this.sampler = new Tone.Sampler({
      urls: SALAMANDER_SAMPLES,
      release: 1,
      onload: () => {
        console.log('Sampler loaded successfully');
        this.startPlayback(duration, notes);
      },
      onerror: (error) => {
        console.error('Sampler load error:', error);
        this.handleSamplerLoadError(error);
      }
    }).toDestination();
  }

  private calculateTotalDuration(notes: tonelocal.Note[]): number {
    return notes.reduce((maxDuration, note) =>
      Math.max(maxDuration, note.Start + note.Duration), 0);
  }

  private startPlayback(duration: number, notes: tonelocal.Note[]): void {
    this.ngZone.runOutsideAngular(() => {
      try {
        this.isLoading = false;
        this.isPlaying = true;

        Tone.getTransport().stop();
        Tone.getTransport().position = 0;

        // Create and store the new loop
        this.currentLoop = new Tone.Loop((time) => {
          notes.forEach(note => {
            const frequencies = note.Frequencies.map(freq => freq.Name);
            const noteTime = Math.max(Tone.now(), time + note.Start);
            this.sampler?.triggerAttackRelease(frequencies, note.Duration, noteTime);
          });
        }, duration).start(0);

        Tone.getTransport().start();
      } catch (error) {
        this.handlePlaybackError(error);
      }
    });
  }

  private handleWebSocketError(error: any): void {
    console.error('WebSocket connection error:', error);
    this.showError('Failed to connect to WebSocket');
  }

  private handleAudioInitError(error: any): void {
    console.error('Audio initialization error:', error);
    this.showError('Failed to initialize audio');
  }

  private handleSamplerLoadError(error: any): void {
    console.error('Sampler load error:', error);
    this.showError('Failed to load audio samples');
    this.isLoading = false;
  }

  private handlePlaybackError(error: any): void {
    console.error('Playback error:', error);
    this.showError('Playback failed');
    this.isLoading = false;
    this.isPlaying = false;
  }

  private showError(message: string): void {
    this.ngZone.run(() => {
      this.snackBar.open(message, 'Close', {
        duration: 3000,
        horizontalPosition: 'center',
        verticalPosition: 'top'
      });
    });
  }
}
