
import { AfterViewInit, Component, DoCheck, ElementRef, Input, OnChanges, OnDestroy, OnInit, SimpleChanges, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { ShapeMouseEvent } from '../shape.mouse.event';
import { MouseEventService } from '../mouse-event.service';
import { AngularDragEndEventService } from '../angular-drag-end-event.service';
import { mouseCoordInComponentRef } from '../mouse.coord.in.component.ref';
import { IsEditableService } from '../is-editable.service';
import { RefreshService } from '../refresh.service';
import { SizeTrackerService } from '../size-tracker.service';

import { SegmentsParams, Segment, createPoint, drawSegments, Offset } from '../draw.segments';
import { swapSegment } from '../swap.segment';
import { manageHandles } from '../manage.handles';
import { getArcPath } from '../get.arc.path';
import { getOrientation } from '../get.orientation';
import { getEndArrowPath } from '../get.end.arrow.path';
import { adjustToSegmentDirection } from '../adjust.to.segment.direction';
import { LinkConf, computeLinkFromMouseEvent } from '../compute.link.from.mouse.event';
import { drawLineFromRectToB } from '../draw.line.from.rect.to.point';
import { getAnchorPoint } from '../get.anchor.point';


@Component({
  selector: 'lib-gongsvg',
  templateUrl: './gongsvg.component.html',
  styleUrls: ['./gongsvg.component.css']
})
export class GongsvgComponent implements OnInit, OnDestroy {

  @Input() GONG__StackPath: string = ""
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

  // for use in the template
  RectAnchorType = gongsvg.RectAnchorType
  LinkType = gongsvg.LinkType

  // the components draws all elements directly but the links when they are 
  // LINK_TYPE_FLOATING_ORTHOGONAL, in this case, each link is associated with a 
  // set of segments
  map_Link_Segment: Map<gongsvg.LinkDB, Segment[]> = new (Map<gongsvg.LinkDB, Segment[]>)
  getSegments(link: gongsvg.LinkDB): Segment[] {
    return this.map_Link_Segment.get(link)!
  }

  // for change detection, we need to store start and end rect of all links
  map_Link_PreviousStart: Map<gongsvg.LinkDB, gongsvg.RectDB> = new (Map<gongsvg.LinkDB, gongsvg.RectDB>)
  map_Link_PreviousEnd: Map<gongsvg.LinkDB, gongsvg.RectDB> = new (Map<gongsvg.LinkDB, gongsvg.RectDB>)

  // temporary, will be computed dynamicaly
  svgWidth = 3000
  svgHeight = 4000
  private onResize = (): void => {
    this.updateSize();
  }

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  rectDragging: boolean = false
  draggedRect: gongsvg.RectDB | undefined
  anchorDragging: boolean = false
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' = 'left'

  // RectAtMouseDown is the clone of the Rect when mouse down
  private RectAtMouseDown: gongsvg.RectDB | undefined
  map_SelectedRectAtMouseDown: Map<gongsvg.RectDB, gongsvg.RectDB> = new (Map<gongsvg.RectDB, gongsvg.RectDB>)

  // to compute wether it was a select / dragging event
  distanceMoved = 0
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private dragThreshold = 5;

  private updateSize() {
    this.svgWidth = window.innerWidth;
    this.svgHeight = window.innerHeight;
    console.log(`New size: Width = ${this.width}, Height = ${this.height}`);

    this.refresh()
  }

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
    private sizeTracker: SizeTrackerService,

    private rectService: gongsvg.RectService,
    private linkService: gongsvg.LinkService,
    private anchoredTextService: gongsvg.LinkAnchoredTextService,

    private refreshService: RefreshService,
  ) {

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseDownEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {
        console.log('SvgComponent, Alt Mouse down event occurred on rectangle ', shapeMouseEvent.ShapeID)
        this.linkStartRectangleID = shapeMouseEvent.ShapeID

        // refactorable of Rect name
        let rect = this.gongsvgFrontRepo?.getMap<gongsvg.RectDB>(gongsvg.RectDB.GONGSTRUCT_NAME).get(shapeMouseEvent.ShapeID)

        if (rect == undefined) {
          return
        }

        this.linkDrawing = true
        this.startX = shapeMouseEvent.Point.X
        this.startY = shapeMouseEvent.Point.Y
      }))

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseDragEvent$.subscribe((shapeMouseEvent: ShapeMouseEvent) => {

      this.endX = shapeMouseEvent.Point.X
      this.endY = shapeMouseEvent.Point.Y
      // console.log('SvgComponent, Received Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
    }))

    this.subscriptions.push(rectangleEventService.mouseRectAltKeyMouseUpEvent$.subscribe((rectangleID: number) => {
      console.log('SvgComponent, Mouse up event occurred on rectangle ', rectangleID);
      this.linkDrawing = false

      this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
    }))

    this.subscriptions.push(svgEventService.multiShapeSelectDragEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        let actualX = shapeMouseEvent.Point.X
        let actualY = shapeMouseEvent.Point.Y

        this.rectX = Math.min(this.startX, actualX);
        this.rectY = Math.min(this.startY, actualY);
        this.width = Math.abs(actualX - this.startX);
        this.height = Math.abs(actualY - this.startY);
      }))

    this.subscriptions.push(svgEventService.mouseShiftKeyMouseUpEvent$.subscribe(
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
      }))

    this.subscriptions.push(svgEventService.multiShapeSelectEndEvent$.subscribe(
      (selectAreaConfig: SelectAreaConfig) => {

        for (let layer of this.gongsvgFrontRepo!.Layers_array) {
          for (let rect of layer.Rects) {
            switch (selectAreaConfig.SweepDirection) {
              case SweepDirection.LEFT_TO_RIGHT:
                if (
                  rect.X > selectAreaConfig.TopLeft[0] &&
                  rect.X + rect.Width < selectAreaConfig.BottomRigth[0] &&
                  rect.Y > selectAreaConfig.TopLeft[1] &&
                  rect.Y + rect.Height < selectAreaConfig.BottomRigth[1]
                ) {
                  rect.IsSelected = true
                  this.manageHandles(rect)
                  this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                    _ => {
                      this.refreshService.emitRefreshRequestEvent(0)
                    }
                  )
                }
                break;
              case SweepDirection.RIGHT_TO_LEFT:
                // rectangle has to be partialy boxed in the rect
                if (
                  rect.X < selectAreaConfig.BottomRigth[0] &&
                  rect.X + rect.Width > selectAreaConfig.TopLeft[0] &&
                  rect.Y < selectAreaConfig.BottomRigth[1] &&
                  rect.Y + rect.Height > selectAreaConfig.TopLeft[1]
                ) {
                  rect.IsSelected = true
                  this.manageHandles(rect)
                  this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                    _ => {
                      this.refreshService.emitRefreshRequestEvent(0)
                    }
                  )
                }
                break;
            }


            this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
              _ => {
                this.refreshService.emitRefreshRequestEvent(0)
              }
            )
          }
        }

      })
    )

    this.subscriptions.push(refreshRequestService.refreshRequest$.subscribe(
      _ => {
        this.refresh()
      }))

    this.subscriptions.push(mouseEventService.mouseMouseDownEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        switch (shapeMouseEvent.ShapeType) {
          case gongsvg.RectDB.GONGSTRUCT_NAME:
            let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
            if (rect == undefined) {
              return
            }

            if (this.anchorDragging || this.rectDragging || rect.IsSelected) {
              console.log("rect: mouseMouseDownEvent, ", rect!.Name)

              this.distanceMoved = 0
              this.RectAtMouseDown = structuredClone(rect)
              this.map_SelectedRectAtMouseDown = new (Map<gongsvg.RectDB, gongsvg.RectDB>)
              for (let layer of this.gongsvgFrontRepo!.Layers_array) {
                for (let rect of layer.Rects) {
                  if (rect.IsSelected) {
                    this.map_SelectedRectAtMouseDown.set(rect, structuredClone(rect))
                  }
                }
              }

              this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)
            }
            break;
          case gongsvg.LinkDB.GONGSTRUCT_NAME:
            let link = this.gongsvgFrontRepoService.frontRepo.Links.get(shapeMouseEvent.ShapeID)
            if (link == undefined) {
              return
            }

            if (shapeMouseEvent.ShapeType != gongsvg.LinkDB.GONGSTRUCT_NAME ||
              shapeMouseEvent.ShapeID != link.ID) {
              return
            }

            this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)

            if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
              if (link.TextAtArrowEnd && link.TextAtArrowEnd[this.draggedTextIndex]) {
                this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowEnd[this.draggedTextIndex])
              }
            }
            if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
              if (link.TextAtArrowStart && link.TextAtArrowStart[this.draggedTextIndex]) {
                this.AnchoredTextAtMouseDown = structuredClone(link.TextAtArrowStart[this.draggedTextIndex])
              }
            }
            break;
        }

      }
    )
    )

    this.subscriptions.push(mouseEventService.mouseMouseMoveEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {
        if (!this.isEditableService.getIsEditable()) {
          return
        }

        if (shapeMouseEvent.ShapeType == gongsvg.RectDB.GONGSTRUCT_NAME) {
          let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
          if (rect == undefined) {
            return
          }

          if (this.anchorDragging) {
            if (this.activeAnchor === 'left') {
              rect.X = this.RectAtMouseDown!.X + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
              rect.Width = this.RectAtMouseDown!.Width - (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'right') {
              rect.Width = this.RectAtMouseDown!.Width + (shapeMouseEvent.Point.X - this.PointAtMouseDown!.X)
            } else if (this.activeAnchor === 'top') {
              rect.Y = this.RectAtMouseDown!.Y + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
              rect.Height = this.RectAtMouseDown!.Height - (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            } else if (this.activeAnchor === 'bottom') {
              rect.Height = this.RectAtMouseDown!.Height + (shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y)
            }
            return // we don't want the move move to be interpreted by the rect
          }

          // console.log("material svg mouse move event", this.PointAtMouseDown != undefined, (this.rectDragging || rect.IsSelected))
          if (this.PointAtMouseDown && (this.rectDragging || rect.IsSelected)) {
            const deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
            const deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
            this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

            if (this.draggedRect == undefined) {
              return
            }
            if (this.draggedRect.CanMoveHorizontaly) {
              this.draggedRect.X = this.RectAtMouseDown!.X + deltaX
            }
            if (this.draggedRect.CanMoveVerticaly) {
              this.draggedRect.Y = this.RectAtMouseDown!.Y + deltaY
            }

            for (let layer of this.gongsvgFrontRepo!.Layers_array) {
              for (let rect_ of layer.Rects) {
                if (rect_.IsSelected) {
                  let rectAtMouseDown_ = this.map_SelectedRectAtMouseDown.get(rect_)
                  if (rect_.CanMoveHorizontaly) {
                    rect_.X = rectAtMouseDown_!.X + deltaX
                  }
                  if (rect_.CanMoveVerticaly) {
                    rect_.Y = rectAtMouseDown_!.Y + deltaY
                  }
                }
              }
            }

            for (let layer of this.gongsvgFrontRepo!.Layers_array) {
              for (let link of layer.Links) {

                // some other rect might have move this test
                // if is a bug
                if (link.End != rect && link.Start != rect) {
                  // return
                }

                let segmentsParams = {
                  StartRect: link.Start!,
                  EndRect: link.End!,
                  StartDirection: link.StartOrientation! as gongsvg.OrientationType,
                  EndDirection: link.EndOrientation! as gongsvg.OrientationType,
                  StartRatio: link.StartRatio,
                  EndRatio: link.EndRatio,
                  CornerOffsetRatio: link.CornerOffsetRatio,
                  CornerRadius: link.CornerRadius,
                }

                let segments = drawSegments(segmentsParams)
                this.map_Link_Segment.set(link, segments)
              }
            }
          }
        }


        if (this.linkDragging && this.draggedLink) {
          let linkConf: LinkConf = {
            drawSegments: this.updateLinkSegments,
            dragging: this.linkDragging,
            draggedLink: this.draggedLink,
            draggedSegmentNumber: this.draggedSegmentNumber,
            draggedSegmentPositionOnArrow: this.draggedSegmentPositionOnArrow,
            segments: this.map_Link_Segment.get(this.draggedLink)!,
            PointAtMouseDown: this.PointAtMouseDown,
            previousStart: this.map_Link_PreviousStart.get(this.draggedLink)!,
            previousEnd: this.map_Link_PreviousEnd.get(this.draggedLink)!,
            linkUpdating: this.linkUpdating,
            map_Link_Segment: this.map_Link_Segment
          }

          computeLinkFromMouseEvent(linkConf, shapeMouseEvent)
        }
        if (this.textDragging) {
          let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
          let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y

          console.log("gongsvg Text dragging, deltaX", deltaX, "deltaY", deltaY)

          if (this.draggedText == undefined) {
            console.log("Problem : this.draggedText should not be undefined")
            return
          }
          this.draggedText.X_Offset = this.AnchoredTextAtMouseDown!.X_Offset + deltaX
          this.draggedText.Y_Offset = this.AnchoredTextAtMouseDown!.Y_Offset + deltaY
        }
      }
    )
    )

    this.subscriptions.push(mouseEventService.mouseMouseUpEvent$.subscribe(
      (shapeMouseEvent: ShapeMouseEvent) => {

        switch (shapeMouseEvent.ShapeType) {
          case gongsvg.RectDB.GONGSTRUCT_NAME:
            let rect = this.gongsvgFrontRepo!.Rects.get(shapeMouseEvent.ShapeID)
            if (rect == undefined) {
              return
            }

            if (!this.isEditableService.getIsEditable()) {
              //
              // if the rect has been clicked on, inform the back end
              //
              this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                _ => {
                  console.log("Rect ", rect?.Name, "clicked")
                }
              )
              return
            }

            if (shapeMouseEvent.ShapeID != 0 && this.distanceMoved > this.dragThreshold) {
              if (this.anchorDragging || this.rectDragging || rect.IsSelected) {
                rect.IsSelected = false
                this.manageHandles(rect)
                this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                  _ => {
                    this.refreshService.emitRefreshRequestEvent(0)
                  }
                )

                // for all other rects
                for (let [rect_, value] of this.map_SelectedRectAtMouseDown) {
                  rect_.IsSelected = false
                  this.manageHandles(rect_)
                  this.rectService.updateRect(rect_, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                    _ => {
                      this.refreshService.emitRefreshRequestEvent(0)
                    }
                  )
                }

              }

            }
            if (this.distanceMoved <= this.dragThreshold) {
              if (rect.IsSelectable &&
                shapeMouseEvent.ShapeType == gongsvg.RectDB.GONGSTRUCT_NAME &&
                shapeMouseEvent.ShapeID == rect.ID) {
                // console.log("rect (distance < threshold), mouseEventService.mouseMouseUpEvent$.subscribe, from the shape: ", rect?.Name)
                rect.IsSelected = !rect.IsSelected
                this.manageHandles(rect)
                this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                  _ => {
                    this.refreshService.emitRefreshRequestEvent(0)

                    // one reset te point at mouse down. Otherwise, the rect will move with the mouse
                    this.PointAtMouseDown = undefined
                  }
                )
              }

              // mouseup emited from the background let to unselect selected shapes
              if (rect.IsSelectable && rect.IsSelected && shapeMouseEvent.ShapeID == 0) {
                console.log("rect, mouseEventService.mouseMouseUpEvent$.subscribe: from the svg", rect.Name)
                rect.IsSelected = false
                this.manageHandles(rect)
                this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                  _ => {
                    this.refreshService.emitRefreshRequestEvent(0)
                  }
                )
              }
            }

            this.rectDragging = false
            this.draggedRect = undefined
            this.anchorDragging = false
            break;
          case gongsvg.LinkDB.GONGSTRUCT_NAME:
            if (this.linkDragging && this.isEditableService.getIsEditable()) {
              document.body.style.cursor = ''

              let deltaX = shapeMouseEvent.Point.X - this.PointAtMouseDown!.X
              let deltaY = shapeMouseEvent.Point.Y - this.PointAtMouseDown!.Y
              this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

              if (this.distanceMoved < this.dragThreshold) {
                console.log("Link, link selected: ", this.draggedLink!.Name, this.draggedLink!.EndRatio)
              } else {
                console.log("Link, link moved: ", this.draggedLink!.Name, this.draggedLink!.EndRatio)

                this.linkUpdating = true
                this.linkService.updateLink(this.draggedLink!,
                  this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                    link => {
                      // this.Link = link
                      this.linkUpdating = false
                      this.refreshService.emitRefreshRequestEvent(0)
                    }
                  )
              }
            }

            if (this.textDragging && this.isEditableService.getIsEditable()) {
              if (this.draggedText == undefined) {
                console.log("Problem : this.draggedText should not be undefined")
                return
              }
              this.anchoredTextService.updateLinkAnchoredText(
                this.draggedText, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()
            }
            this.linkDragging = false
            this.textDragging = false
            break;
          default:
            if (this.distanceMoved <= this.dragThreshold) {
              console.log('Material svg: Mouse down event occurred with no shape, distance', this.distanceMoved);

              // one have to deselect all
              for (let layer of this.gongsvgFrontRepo!.Layers_array) {
                for (let rect_ of layer.Rects) {
                  if (rect_.IsSelected) {
                    rect_.IsSelected = false
                    this.manageHandles(rect_)
                    this.rectService.updateRect(rect_, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
                      _ => {
                        this.refreshService.emitRefreshRequestEvent(0)
                      }
                    )
                  }
                }
              }
            }
        }
        // console.log('Rect ', this.Rect.Name, 'Mouse down event occurred on rectangle ', rectangleID);
      })
    )
  }

  ngOnInit(): void {

    this.updateSize();
    window.addEventListener('resize', this.onResize);

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

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


  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        if (this.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
          this.svg = this.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0]

          // set the isEditable
          this.isEditableService.setIsEditable(this.svg!.IsEditable)
        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        // compute segments for links
        this.map_Link_Segment.clear()

        for (let layer of this.gongsvgFrontRepo.Layers_array) {
          for (let link of layer.Links) {
            let segmentsParams = {
              StartRect: link.Start!,
              EndRect: link.End!,
              StartDirection: link.StartOrientation! as gongsvg.OrientationType,
              EndDirection: link.EndOrientation! as gongsvg.OrientationType,
              StartRatio: link.StartRatio,
              EndRatio: link.EndRatio,
              CornerOffsetRatio: link.CornerOffsetRatio,
              CornerRadius: link.CornerRadius,
            }

            let segments = drawSegments(segmentsParams)
            this.map_Link_Segment.set(link, segments)
          }
        }

        this.resetAllLinksPreviousStartEndRects()

      }

    )
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
    window.removeEventListener('resize', this.onResize);
  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    let svgArray = this.gongsvgFrontRepo?.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)
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

    let rectMap = this.gongsvgFrontRepo!.getMap<gongsvg.RectDB>(gongsvg.RectDB.GONGSTRUCT_NAME)

    this.svg.StartRect = rectMap.get(startRectangleID)
    this.svg.EndRect = rectMap.get(endRectangleID)

    this.svgService.updateSVG(this.svg, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
      () => {

        this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
          gongsvgsFrontRepo => {
            this.gongsvgFrontRepo = gongsvgsFrontRepo

            if (this.gongsvgFrontRepo.getArray(gongsvg.SVGDB.GONGSTRUCT_NAME).length == 1) {
              this.svg = this.gongsvgFrontRepo.getArray<gongsvg.SVGDB>(gongsvg.SVGDB.GONGSTRUCT_NAME)[0]

              // back to normal state
              this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE
              this.svgService.updateSVG(this.svg, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe()

              // set the isEditable
              this.isEditableService.setIsEditable(this.svg!.IsEditable)
            } else {
              return
            }
          }
        )
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
      // console.log("svg background, mouse move", shapeMouseEvent.Point.X, shapeMouseEvent.Point.Y)
    }
  }

  onmouseup(event: MouseEvent): void {
    this.selectionRectDrawing = false

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: 0,
      ShapeType: "",
      Point: mouseCoordInComponentRef(event),
    }

    // strangely, when mouse up with shift key, this is not allways triggering
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


  rectMouseDown(event: MouseEvent, rect: gongsvg.RectDB): void {
    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true
      this.draggedRect = rect

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    } else {
      console.log('RectComponent, Alt Mouse down event occurred on rectangle ', rect.Name, rect.ID, event.clientX, event.clientY);

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.rectangleEventService.emitRectAltKeyMouseDownEvent(shapeMouseEvent)

    }
  }

  rectMouseMove(event: MouseEvent, rect: gongsvg.RectDB): void {

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: rect.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    // console.log("RectComponent DragRect : ", deltaX, deltaY, distanceMoved)

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      // console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', rect.Name, event.clientX, event.clientY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(shapeMouseEvent)
      return
    }

    if (event.shiftKey) {
      this.svgEventService.emitMultiShapeSelectDrag(shapeMouseEvent)
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  rectMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {

    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    let shapeMouseEvent: ShapeMouseEvent = {
      ShapeID: rect.ID,
      ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
      Point: mouseCoordInComponentRef(event),
    }

    if (!event.altKey && !event.shiftKey) {
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }

    if (event.altKey) {
      console.log('RectComponent, Alt Mouse up event occurred on rectangle ', rect.Name, rect.ID, event.clientX, event.clientY);

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(rect.ID);
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent(shapeMouseEvent)
    }
  }

  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  anchorMouseMove(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom', rect: gongsvg.RectDB): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: rect.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  anchorMouseUp(event: MouseEvent, rect: gongsvg.RectDB): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = 'bottom'
      rect.IsSelected = false
      this.manageHandles(rect)
      this.rectService.updateRect(rect, this.GONG__StackPath, this.gongsvgFrontRepoService.frontRepo).subscribe(
        _ => {
          this.refreshService.emitRefreshRequestEvent(0)
        }
      )
    }
  }

  splitTextIntoLines(text: string): string[] {
    return text.split('\n')
  }

  // display or not handles if selected or not
  manageHandles(rect: gongsvg.RectDB) {
    manageHandles(rect)
  }

  //
  // for links management
  //

  // to compute wether it was a select / dragging event
  linkDragging = false
  draggedLink: gongsvg.LinkDB | undefined
  draggedSegmentNumber = 0
  draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType = gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START

  // indicate wether the link is being updated
  // no drawing must happen then
  linkUpdating: boolean = false

  // dragged anchored text
  textDragging = false
  draggedTextIndex = 0
  draggedText: gongsvg.LinkAnchoredTextDB | undefined

  // LinkAtMouseDown is the clone of the Link when mouse down
  AnchoredTextAtMouseDown: gongsvg.LinkAnchoredTextDB | undefined

  linkMouseDown(event: MouseEvent, segmentNumber: number, link: gongsvg.LinkDB): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      // this link shit to dragging state
      this.linkDragging = true
      this.draggedLink = link
      this.draggedSegmentNumber = segmentNumber

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  linkMouseMove(event: MouseEvent, link: gongsvg.LinkDB): void {

    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  linkMouseUp(event: MouseEvent, link: gongsvg.LinkDB, segmentNumber: number = 0): void {

    // console.log("Link : linkMouseUp", this.Link?.Name)
    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {
    return getOrientation(segment)
  }

  // Add this method to ArcComponent
  getArcPath(link: gongsvg.LinkDB, segment: Segment, nextSegment: Segment): string {
    return getArcPath(link, segment, nextSegment)
  }

  getEndArrowPath(link: gongsvg.LinkDB, segment: Segment, arrowSize: number): string {
    return getEndArrowPath(link, segment, arrowSize)
  }

  getStartArrowPath(link: gongsvg.LinkDB, segment: Segment, arrowSize: number): string {

    let inverseSegment = swapSegment(segment)
    let path = this.getEndArrowPath(link, inverseSegment, arrowSize)
    return path
  }

  adjustToSegmentDirection(segment: Segment, x: number, y: number): { x: number, y: number } {
    return adjustToSegmentDirection(segment, x, y)
  }

  textAnchoredMouseDown(
    link: gongsvg.LinkDB,
    event: MouseEvent,
    anchoredTextIndex: number,
    draggedSegmentPositionOnArrow: string): void {

    if (!event.altKey && !event.shiftKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      // this text shift to dragging state
      this.textDragging = true
      this.draggedTextIndex = anchoredTextIndex
      this.draggedSegmentPositionOnArrow = draggedSegmentPositionOnArrow as gongsvg.PositionOnArrowType
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        this.draggedText = link.TextAtArrowStart![this.draggedTextIndex]
      }
      if (this.draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        this.draggedText = link.TextAtArrowEnd![this.draggedTextIndex]
      }
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseDownEvent(shapeMouseEvent)
    }
  }

  textAnchoredMouseMove(link: gongsvg.LinkDB, event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  textAnchoredMouseUp(link: gongsvg.LinkDB, event: MouseEvent): void {

    // console.log("Link : linkMouseUp", this.Link?.Name)
    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: link.ID,
        ShapeType: gongsvg.LinkDB.GONGSTRUCT_NAME,
        Point: mouseCoordInComponentRef(event),
      }
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  updateLinkSegments(link: gongsvg.LinkDB, linkUpdating: boolean, map_Link_Segment: Map<gongsvg.LinkDB, Segment[]>): boolean {

    if (link == undefined) {
      return false
    }

    if (linkUpdating) {
      return true
    }

    let segmentsParams = {
      StartRect: link.Start!,
      EndRect: link.End!,
      StartDirection: link.StartOrientation! as gongsvg.OrientationType,
      EndDirection: link.EndOrientation! as gongsvg.OrientationType,
      StartRatio: link.StartRatio,
      EndRatio: link.EndRatio,
      CornerOffsetRatio: link.CornerOffsetRatio,
      CornerRadius: link.CornerRadius,
    }

    let segments = drawSegments(segmentsParams)
    if (segments == undefined) {
      return false
    }
    if (map_Link_Segment == undefined) {
      console.log("strange")
    }

    this.map_Link_Segment.set(link, segments)

    return true
  }

  resetAllLinksPreviousStartEndRects() {
    for (let link of this.gongsvgFrontRepo!.Links_array) {
      this.map_Link_PreviousStart.set(link, structuredClone(link.Start!))
      this.map_Link_PreviousEnd.set(link, structuredClone(link.End!))
    }
  }

  public getStartPosition(rectLinkLink: gongsvg.RectLinkLinkDB): Coordinate {

    let coordinate: Coordinate = [0, 0]
    if (rectLinkLink.End == undefined || rectLinkLink.Start == undefined) {
      return coordinate
    }

    let sourceSegments = this.map_Link_Segment.get(rectLinkLink.End)
    if (sourceSegments == undefined) {
      return coordinate
    }

    let target = getAnchorPoint(sourceSegments, rectLinkLink.TargetAnchorPosition)
    if (target == undefined) {
      return coordinate
    }
    let source = drawLineFromRectToB(rectLinkLink.Start, target)

    coordinate[0] = source.X
    coordinate[1] = source.Y

    return coordinate
  }

  public getEndPosition(rectLinkLink: gongsvg.RectLinkLinkDB): Coordinate {

    let coordinate: Coordinate = [0, 0]
    if (rectLinkLink.End == undefined || rectLinkLink.Start == undefined) {
      return coordinate
    }

    let sourceSegments = this.map_Link_Segment.get(rectLinkLink.End)
    if (sourceSegments == undefined) {
      return coordinate
    }

    let target = getAnchorPoint(sourceSegments, rectLinkLink.TargetAnchorPosition)
    if (target == undefined) {
      return coordinate
    }
    let source = drawLineFromRectToB(rectLinkLink.Start, target)

    coordinate[0] = target.X
    coordinate[1] = target.Y

    return coordinate
  }
}