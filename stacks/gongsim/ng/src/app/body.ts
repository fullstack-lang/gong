import { MatHeaderCell } from '@angular/material/table';

// 
const CanvasSunX = 500
const CanvasSunY = 500

const pxPerMeters = 0.8 / (1000 * 1000 * 1000.0) // 100 pixels = 100 MKm

const bodiesRadiusExageration = 40
const satelliteOrbitExageration = 40

const speedOfLight = 300 * 1000 * 1000 // 300 000 km/s

export class Body {

    offsetX = 0
    offsetY = 0

    constructor(private ctx: CanvasRenderingContext2D) { }

    OrbitSemiMajoraxis: number // in meters
    OrbitalPeriod: number // in seconds
    Radius: number // in meters

    OrbitCenter: Body

    draw() {
        var radiusInPx = Math.max(this.Radius * pxPerMeters * bodiesRadiusExageration, 1)
        this.ctx.beginPath();
        this.ctx.arc(CanvasSunX + this.offsetX, CanvasSunY + this.offsetY, radiusInPx, 0, Math.PI * 2, true)
        this.ctx.stroke();
    }

    // update position seen from earth (take speed of light into account)
    public updatePosition(secondsSinceStart: number, earth: Body) {

        var distanceToEarth = Math.sqrt(
            Math.pow(this.offsetX - earth.offsetX, 2) +
            Math.pow(this.offsetY - earth.offsetY, 2)
        ) / pxPerMeters

        var timeShiftSeconds = distanceToEarth / speedOfLight

        if (this.OrbitCenter != undefined) {

            var angleOfRotation = (secondsSinceStart - timeShiftSeconds) / this.OrbitalPeriod * Math.PI * 2

            this.offsetX = this.OrbitSemiMajoraxis * pxPerMeters * Math.cos(angleOfRotation)
            this.offsetY = -this.OrbitSemiMajoraxis * pxPerMeters * Math.sin(angleOfRotation)

            if (this.OrbitCenter.offsetX != 0.0) {
                this.offsetX *= satelliteOrbitExageration
                this.offsetY *= satelliteOrbitExageration
            }
            this.offsetX += this.OrbitCenter.offsetX
            this.offsetY += this.OrbitCenter.offsetY
        }
    }
}
