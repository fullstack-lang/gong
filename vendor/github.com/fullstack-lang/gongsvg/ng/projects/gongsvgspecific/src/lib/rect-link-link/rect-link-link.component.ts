import { Component, DoCheck, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { Coordinate } from '../rectangle-event.service';
import { compareRectGeometries } from '../compare.rect.geometries';
import { compareLinkGeometries } from '../compare.link.geometries';
import { Segment, SegmentsParams, drawSegments } from '../link/draw.segments';
import { getAnchorPoint } from '../get.anchor.point';
import { drawLineFromRectToB } from '../draw.line.from.rect.to.point';

@Component({
  selector: 'lib-rect-link-link',
  templateUrl: './rect-link-link.component.svg',
  styleUrls: ['./rect-link-link.component.css']
})
export class RectLinkLinkComponent implements OnInit, DoCheck {

  @Input() RectLinkLink: gongsvg.RectLinkLinkDB = new gongsvg.RectLinkLinkDB()
  @Input() GONG__StackPath: string = ""

  previousStart: gongsvg.RectDB | undefined
  previousEnd_StartRect: gongsvg.RectDB | undefined
  previousEnd_EndRect: gongsvg.RectDB | undefined
  previousLink: gongsvg.LinkDB | undefined
  isTargetX_a_NaN: boolean = false

  segments: Segment[] = []
  target: gongsvg.PointDB | null = null
  source: gongsvg.PointDB | null = null

  ngOnInit(): void {
    this.resetPreviousState()
  }

  public getStartPosition(): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = this.source!.X
    coordinate[1] = this.source!.Y

    return coordinate
  }

  public getEndPosition(): Coordinate {

    let coordinate: Coordinate = [0, 0]

    coordinate[0] = this.target!.X
    coordinate[1] = this.target!.Y

    if (Number.isNaN(coordinate[0])) {
      console.log("end is NaN")
    }

    return coordinate
  }

  ngDoCheck(): void {

    let hasStartChanged = !compareRectGeometries(this.previousStart!, this.RectLinkLink!.Start!)
    let hasEnd_StartChanged = !compareRectGeometries(this.previousEnd_StartRect!, this.RectLinkLink!.End!.Start!)
    let hasEnd_EndChanged = !compareRectGeometries(this.previousEnd_EndRect!, this.RectLinkLink!.End!.End!)
    let hasLinkChanged = !compareLinkGeometries(this.previousLink!, this.RectLinkLink!.End!)

    if (hasStartChanged || hasEnd_StartChanged || hasEnd_EndChanged || hasLinkChanged) {
      this.resetPreviousState()
    }
  }

  resetPreviousState() {
    this.previousStart = structuredClone(this.RectLinkLink!.Start!)
    this.previousEnd_StartRect = structuredClone(this.RectLinkLink!.End!.Start!)
    this.previousEnd_EndRect = structuredClone(this.RectLinkLink!.End!.End!)
    this.previousLink = structuredClone(this.RectLinkLink!.End!)

    let link = this.RectLinkLink!.End!

    let segmentsParams: SegmentsParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartOrientation! as gongsvg.OrientationType,
      EndDirection: link.EndOrientation! as gongsvg.OrientationType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
      CornerRadius: link.CornerRadius,
    }

    this.segments = drawSegments(segmentsParams)
    this.target = getAnchorPoint(this.segments, this.RectLinkLink!.TargetAnchorPosition)

    // taking into account the case where the target link has been
    // update but the front has not refreshed the link.Start and link.End associations
    if (this.target == null) {
      this.isTargetX_a_NaN = true
      return
    } else {
      this.isTargetX_a_NaN = false
    }

    this.source = drawLineFromRectToB(this.RectLinkLink!.Start!, this.target!)
  }
}
