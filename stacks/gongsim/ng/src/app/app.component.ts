import { Component, ViewChild, ElementRef, OnInit, NgZone } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongsim from 'gongsim'

import { Body } from './body';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  @ViewChild('canvas', { static: true }) canvas: ElementRef<HTMLCanvasElement>;
  ctx: CanvasRenderingContext2D;
  requestId;
  interval;
  bodies: Body[] = [];

  SecondsSinceStart = 0

  Jupiter: Body
  Io: Body
  Earth: Body

  ioWasOnTheLeftOfJupiterAtLastCalculation = true
  timeEndOfEclipse = 0
  durationBetweenEclipse = 0
  minDurationBetweenEclipse = 3600 * 24 * 10
  maxDurationBetweenEclipse = 0
  minDurationBetweenEclipseSeconds = 0
  maxDurationBetweenEclipseSeconds = 0

  deltaFromEclipseToEclipseSeconds = 0

  earthToIoNorm = 1.0
  earthToJupiterNorm = 1.0
  sin: number;

  constructor(
    private engineService: gongsim.EngineService,
    private ngZone: NgZone
  ) { }

  ngOnInit() {

    this.ctx = this.canvas.nativeElement.getContext('2d');
    this.ctx.fillStyle = 'red';

    var Sun: Body = new Body(this.ctx)
    Sun.Radius = 696.342 * 1000 * 1000
    this.bodies = this.bodies.concat(Sun)


    this.Jupiter = new Body(this.ctx)
    this.Jupiter.OrbitSemiMajoraxis = 778.57 * 1000 * 1000 * 1000 // 778 millions of km in meters 
    this.Jupiter.OrbitalPeriod = 4432.59 * 24 * 3600 // 4432 days or 11 years in seconds
    this.Jupiter.Radius = 69.911 * 1000 * 1000 // in meters
    this.Jupiter.OrbitCenter = Sun
    this.bodies = this.bodies.concat(this.Jupiter)

    // Io is the jovian moon watched by romer
    this.Io = new Body(this.ctx)
    this.Io.OrbitalPeriod = 42.45930686 * 3600 // 42 hours in seconds
    this.Io.OrbitSemiMajoraxis = 421.7 * 1000 * 1000 // 421700 km in meters
    this.Io.Radius = 1821.6 * 1000.0 // 0.2 earth
    this.Io.OrbitCenter = this.Jupiter
    this.bodies = this.bodies.concat(this.Io)

    this.Earth = new Body(this.ctx)
    this.Earth.OrbitSemiMajoraxis = 149.598 * 1000 * 1000 * 1000 // 150 millions of km in meters 
    this.Earth.OrbitalPeriod = 365.256 * 24 * 3600 // 365 days in seconds
    this.Earth.Radius = 6 * 1000 * 1000 // in meters
    this.Earth.OrbitCenter = Sun
    this.bodies = this.bodies.concat(this.Earth)


    this.ngZone.runOutsideAngular(() => this.tick());
    setInterval(() => {
      this.tick();
    }, 200);

  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (updateDisplay: boolean): void => {

    if (updateDisplay) {

      // fetch the time
      this.engineService.getEngines().subscribe(
        engines => {
          // get the engine
          if (engines.length == 1) {
            // update position between two ticks
            for (this.SecondsSinceStart; this.SecondsSinceStart < engines[0].SecondsSinceStart; this.SecondsSinceStart++) {
              this.bodies.forEach((body: Body) => {
                body.updatePosition(this.SecondsSinceStart, this.Earth);
              });

              var ioIsOnTheLeftOfJupiter = this.isIoOnLeftOfJupiter(false)
              if (this.ioWasOnTheLeftOfJupiterAtLastCalculation != ioIsOnTheLeftOfJupiter) {
                if (ioIsOnTheLeftOfJupiter) {
                  this.isIoOnLeftOfJupiter(true)
                  this.deltaFromEclipseToEclipseSeconds =
                    (this.SecondsSinceStart - this.timeEndOfEclipse) - this.durationBetweenEclipse
                  this.durationBetweenEclipse = this.SecondsSinceStart - this.timeEndOfEclipse

                  this.timeEndOfEclipse = this.SecondsSinceStart

                  if (this.durationBetweenEclipse > this.maxDurationBetweenEclipse) {
                    this.maxDurationBetweenEclipse = this.durationBetweenEclipse
                    this.maxDurationBetweenEclipseSeconds = this.timeEndOfEclipse % this.Earth.OrbitalPeriod
                  }
                  if (this.durationBetweenEclipse < this.minDurationBetweenEclipse) {
                    if (this.durationBetweenEclipse > 3600 * 24) {
                      this.minDurationBetweenEclipse = this.durationBetweenEclipse
                      this.minDurationBetweenEclipseSeconds = this.timeEndOfEclipse % this.Earth.OrbitalPeriod
                    }
                  }
                }
                this.ioWasOnTheLeftOfJupiterAtLastCalculation = ioIsOnTheLeftOfJupiter
              }
            }
            this.SecondsSinceStart = engines[0].SecondsSinceStart
          }
        }
      )
    }
  }

  isIoOnLeftOfJupiter(display: boolean): boolean {

    let earthToIo: number[] = [
      this.Io.offsetX - this.Earth.offsetX,
      this.Io.offsetY - this.Earth.offsetY];

    let earthToJupiter: number[] = [
      this.Jupiter.offsetX - this.Earth.offsetX,
      this.Jupiter.offsetY - this.Earth.offsetY
    ];

    this.earthToIoNorm = Math.sqrt(earthToIo[0] * earthToIo[0] + earthToIo[1] * earthToIo[1])
    this.earthToJupiterNorm = Math.sqrt(earthToJupiter[0] * earthToJupiter[0] + earthToJupiter[1] * earthToJupiter[1])
    this.sin = (earthToIo[1] * earthToJupiter[0] - earthToIo[0] * earthToJupiter[1]) /
      (this.earthToIoNorm * this.earthToJupiterNorm)

    if (display == true) {
      console.log(
        this.SecondsSinceStart + " " +
        (this.SecondsSinceStart - this.timeEndOfEclipse) + " " +
        this.durationBetweenEclipse + " " +
        this.sin.toPrecision(10))
    }

    return this.sin < 0
  }

  displayDuration(duration: number): string {

    var days = Math.trunc(duration / (24 * 3600))
    var hours = Math.trunc((duration - days * (24 * 3600)) / 3600)
    var minutes = Math.trunc((duration - days * (24 * 3600) - hours * 3600) / 60)
    var seconds = Math.trunc(duration - days * (24 * 3600) - hours * 3600 - minutes * 60)

    return days + " day " +
      hours + " hour " +
      minutes + " minutes " +
      seconds + " seconds "
  }

  tick() {
    this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);

    this.bodies.forEach((body: Body) => {
      body.draw();
    });
    this.ctx.fillStyle = 'orange';
    this.ctx.font = '18px Arial';

    var left = "Io is on the left of Jupiter"
    var rigth = "Io is on the rigth of Jupiter"

    var text = "days since beginning " + (this.SecondsSinceStart / (3600 * 24)).toFixed(1)
    this.ctx.fillText(text, 10, 50);


    if (this.ioWasOnTheLeftOfJupiterAtLastCalculation) {
      text = left
    } else {
      text = rigth
    }
    this.ctx.fillText(text, 10, 75);
    this.ctx.fillText("Duration between Eclipse :", 10, 100);
    this.ctx.fillText("Cur:\t" + this.displayDuration(this.durationBetweenEclipse) +
      " delta " + this.deltaFromEclipseToEclipseSeconds, 10, 125);

    this.ctx.fillText("Max:\t" + this.displayDuration(this.maxDurationBetweenEclipse), 10, 150);
    this.ctx.fillText("Max:\t" + this.displayDuration(this.maxDurationBetweenEclipseSeconds), 10, 175);
    this.ctx.fillText("Min:\t" + this.displayDuration(this.minDurationBetweenEclipse), 10, 200);
    this.ctx.fillText("Min:\t" + this.displayDuration(this.minDurationBetweenEclipseSeconds), 10, 225);

    this.requestId = requestAnimationFrame(() => this.tick);

  }

  ngOnDestroy() {
    clearInterval(this.interval);
    cancelAnimationFrame(this.requestId);
  }
}
