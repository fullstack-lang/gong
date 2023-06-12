import { AfterViewInit, Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { ShapeMouseEvent } from '../shape.mouse.event';
import { createPoint } from '../link/draw.segments';
import { MouseEventService } from '../mouse-event.service';
import { AngularDragEndEventService } from '../angular-drag-end-event.service';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';


@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

  public gongsvgFrontRepo?: gongsvg.FrontRepo


  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  svg = new gongsvg.SVGDB
  linkStartRectangleID: number = 0

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];


  // if true, the end user is shiftKey + mouse down from one rectangle
  // to another
  linkDrawing: boolean = false
  startX = 0;
  startY = 0;
  endX = 0;
  endY = 0;

  selectionRectDrawing: boolean = false
  rectX = 100;
  rectY = 100;
  width = 300;
  height = 40;

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private mouseEventService: MouseEventService,
    private isEditableService: IsEditableService,
    private refreshRequestService: RefreshService,
  ) {

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {
          // console.log('SvgComponent, Mouse down event occurred on rectangle ', rectangleID, " at ", coordinate)
          this.linkStartRectangleID = shapeMouseEvent.ShapeID

          let rect = this.gongsvgFrontRepo?.Rects.get(shapeMouseEvent.ShapeID)

          if (rect == undefined) {
            return
          }

          this.linkDrawing = true
          this.startX = shapeMouseEvent.Point.X
          this.startY = shapeMouseEvent.Point.Y
        })
    );

    this.subscriptions.push(

      rectangleEventService.mouseRectAltKeyMouseDragEvent$.subscribe((shapeMouseEvent: ShapeMouseEvent) => {

        this.endX = shapeMouseEvent.Point.X
        this.endY = shapeMouseEvent.Point.Y
        // console.log('SvgComponent, Received Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
      })
    )

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseUpEvent$.subscribe((rectangleID: number) => {
        // console.log('SvgComponent, Mouse up event occurred on rectangle ', rectangleID);
        this.linkDrawing = false

        this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
      })
    )

    this.subscriptions.push(
      svgEventService.multiShapeSelectDragEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          let actualX = shapeMouseEvent.Point.X
          let actualY = shapeMouseEvent.Point.Y

          this.rectX = Math.min(this.startX, actualX);
          this.rectY = Math.min(this.startY, actualY);
          this.width = Math.abs(actualX - this.startX);
          this.height = Math.abs(actualY - this.startY);
        })
    );

    this.subscriptions.push(
      svgEventService.mouseShiftKeyMouseUpEvent$.subscribe(
        (shapeMouseEvent) => {

          this.selectionRectDrawing = false
          this.endX = shapeMouseEvent.Point.X
          this.endY = shapeMouseEvent.Point.Y

          let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

          if (this.endX > this.startX) {
            selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
            selectAreaConfig.TopLeft = [this.startX, this.startY]
            selectAreaConfig.BottomRigth = [this.endX, this.endY]
          } else {
            selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
            selectAreaConfig.TopLeft = [this.endX, this.endY]
            selectAreaConfig.BottomRigth = [this.startX, this.startY]
          }

          this.svgEventService.emitMultiShapeSelectEnd(selectAreaConfig)
        }
      )
    )

    this.subscriptions.push(
      refreshRequestService.refreshRequest$.subscribe(
        _ => {
          this.refresh()
        }
      )
    )
  }

  ngOnInit(): void {

    // console.log("Svg component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.refresh()
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack
        }
      }
    )

    // this.gongsvgPushFromFrontNbService.getPushNbFromFront(500, this.GONG__StackPath).subscribe(
    //   lastPushFromFrontNb => {
    //     if (this.lastPushFromFrontNb < lastPushFromFrontNb) {

    //       // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromFront + " new: " + commiNbFromBagetCommitNbFromFront)
    //       this.refresh()
    //       this.lastPushFromFrontNb = lastPushFromFrontNb
    //     }
    //   }
    // )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        if (this.gongsvgFrontRepo.SVGs_array.length == 1) {
          this.svg = this.gongsvgFrontRepo.SVGs_array[0]

          // set the isEditable
          this.isEditableService.setIsEditable(this.svg!.IsEditable)
        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        this.svg.Layers.sort((t1, t2) => {
          let t1_revPointerID_Index = t1.SVG_LayersDBID_Index
          let t2_revPointerID_Index = t2.SVG_LayersDBID_Index

          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }

    )
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    let svgArray = this.gongsvgFrontRepo?.SVGs_array
    // update the svg
    if (svgArray?.length == 1) {
      this.svg = svgArray[0]
    } else {
      return
    }

    if (this.svg.Layers == undefined) {
      return
    }

    if (this.svg.DrawingState != gongsvg.DrawingState.NOT_DRAWING_LINE) {
      // console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.NOT_DRAWING_LINE)
    }

    this.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINE

    this.svg.StartRectID.Valid = true
    this.svg.StartRectID.Int64 = startRectangleID

    this.svg.EndRectID.Valid = true
    this.svg.EndRectID.Int64 = endRectangleID

    this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe(
      () => {
        // back to normal state
        this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE
        this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe()
      }
    )
  }

  mousedown(event: MouseEvent): void {
    if (event.shiftKey) {

      this.selectionRectDrawing = true

      let point = mouseCoordInComponentRef(event)

      this.startX = point.X
      this.startY = point.Y
    }
  }

  mousemove(event: MouseEvent): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: 0,
      ShapeType: "",
      Point: mouseCoordInComponentRef(event),
    }

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      // console.log('SvgComponent, ALT Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(shapeMouseEvent)
      return
    }
    if (event.shiftKey) {

      this.svgEventService.emitMultiShapeSelectDrag(shapeMouseEvent)
      // console.log('SvgComponent, SHIFT Mouse drag event occurred', this.selectionRectDrawing, this.rectX, this.rectY, this.width, this.height);
    }

    if (!event.shiftKey && !event.altKey) {
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
      // console.log("svg background, mouse move", x, y)
    }
  }


  onmouseup(event: MouseEvent): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: 0,
      ShapeType: "",
      Point: mouseCoordInComponentRef(event),
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent(shapeMouseEvent)
    }

    if (!event.shiftKey && !event.altKey) {
      // in case of dragging something, when the mouse moves fast, it can reach the SVG background
      // in this case, one forward the mouse event on the event bus
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  exportSVG() {
    const serializer = new XMLSerializer();
    const svgString = serializer.serializeToString(this.drawingArea!.nativeElement);
    return svgString;
  }

  downloadSVG() {
    const svgString = this.exportSVG();
    const blob = new Blob([svgString], { type: 'image/svg+xml' });
    const url = window.URL.createObjectURL(blob);

    // Create a link element
    const downloadLink = document.createElement('a');
    downloadLink.href = url;
    downloadLink.download = 'image.svg';

    // Attach the link to the document
    document.body.appendChild(downloadLink);
    downloadLink.click();

    // Clean up after to avoid memory leaks
    document.body.removeChild(downloadLink);
  }

}
