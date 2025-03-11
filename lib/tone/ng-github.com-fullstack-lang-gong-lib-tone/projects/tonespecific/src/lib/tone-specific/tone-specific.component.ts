import { Component, OnInit, OnDestroy, NgZone } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';

import * as Tone from 'tone';
import * as tonelocal from '../../../../tone/src/public-api';

import { Subject } from 'rxjs';
import { takeUntil, catchError } from 'rxjs/operators';

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
  private synth: Tone.PolySynth | undefined;
  private sampler: Tone.Sampler | undefined;
  private currentLoop: Tone.Loop | undefined;
  private destroy$ = new Subject<void>();

  readonly StacksNames = tonelocal.StacksNames;

  frontRepo?: tonelocal.FrontRepo;
  isLoading = false;
  isPlaying = false;

  constructor(
    private frontRepoService: tonelocal.FrontRepoService,
    private playerService: tonelocal.PlayerService,
    private ngZone: NgZone,
    private snackBar: MatSnackBar
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
    this.frontRepoService.connectToWebSocket(this.StacksNames.Tone)
      .pipe(
        takeUntil(this.destroy$),
        catchError(error => {
          this.handleWebSocketError(error);
          throw error;
        })
      )
      .subscribe(gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo;
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
      this.playerService.updateFront(player, this.StacksNames.Tone).subscribe(
        () => {
          console.log("gongtone: status set to PAUSED");
        }
      );
    }
  }

  play(): void {
    if (!this.frontRepo) {
      this.showError('No data available for playback');
      return;
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
      this.playerService.updateFront(player, this.StacksNames.Tone).subscribe(
        () => {
          console.log("gongtone: status set to PLAYING");
        }
      );
    }
  }

  private initializeSampler(duration: number, notes: tonelocal.Note[]): void {

    // Use absolute URL with origin
    const audioBaseUrl = `${window.location.origin}/assets/audio/salamander/`;

    // Prefer OGG files, fall back to MP3
    const urls: { [key: string]: string } = {
      C3: `${audioBaseUrl}C3.mp3`,
      'D#3': `${audioBaseUrl}Ds3.mp3`,
      'F#3': `${audioBaseUrl}Fs3.mp3`,
      A3: `${audioBaseUrl}A3.mp3`,
      C4: `${audioBaseUrl}C4.mp3`,
      'D#4': `${audioBaseUrl}Ds4.mp3`,
      'F#4': `${audioBaseUrl}Fs4.mp3`,
      A4: `${audioBaseUrl}A4.mp3`,
      C5: `${audioBaseUrl}C5.mp3`,
      'D#5': `${audioBaseUrl}Ds5.mp3`
    };

    this.sampler = new Tone.Sampler({
      urls: urls,
      release: 1,
      onload: () => {
        console.log('Sampler loaded successfully');
        this.startPlayback(duration, notes);
      },
      onerror: (error) => {
        console.error('Sampler load error:', error);

        // Detailed error logging
        if (error instanceof Event) {
          const target = error.target as HTMLAudioElement;
          console.error('Failed Audio Source:', {
            src: target.src,
            error: target.error,
            errorCode: target.error?.code
          });
        }

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

        // Create and store the new loop
        this.currentLoop = new Tone.Loop((time) => {
          notes.forEach(note => {
            const frequencies = note.Frequencies.map(freq => freq.Name);
            this.sampler?.triggerAttackRelease(frequencies, note.Duration, time + note.Start);
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
