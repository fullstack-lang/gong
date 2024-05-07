import * as gongsvg from '../../../gongsvg/src/public-api'
import { Segment } from './draw.segments'
import { ShapeMouseEvent } from './shape.mouse.event'

export interface LinkConf {
    drawSegments(link: gongsvg.Link, linkUpdating: boolean, map_Link_Segment: Map<gongsvg.Link, Segment[]>): boolean

    dragging: boolean
    draggedLink: gongsvg.Link | undefined
    draggedSegmentNumber: number
    draggedSegmentPositionOnArrow: gongsvg.PositionOnArrowType
    segments: Segment[] | undefined
    PointAtMouseDown: gongsvg.Point | undefined

    // for change detection, we need to store start and end rect
    previousStart: gongsvg.Rect | undefined
    previousEnd: gongsvg.Rect | undefined
    map_Link_Segment: Map<gongsvg.Link, Segment[]>
    linkUpdating: boolean,
}



// computeLinkFromMouseEvent takes a mouse event
// 1. update the link conf
// 2. call drawSegment
export function computeLinkFromMouseEvent(linkConf: LinkConf, shapeMouseEvent: ShapeMouseEvent) {

    // offSetForNewMidlleSegment denotes the standard distance between
    // a rect and the middle segment that is created when going from a
    // two segment link to a three segment link
    const offSetForNewMidlleSegment = 100

    // in order to have middle segment always visible
    const offsetFromBorderForNewMidlleSegment = 50

    if (linkConf.segments == undefined) {
        return
    }

    if (linkConf.draggedSegmentNumber >= linkConf.segments!.length) {
        linkConf.draggedSegmentNumber = linkConf.segments!.length - 1
    }

    let segment = linkConf.segments![linkConf.draggedSegmentNumber]
    let link = linkConf.draggedLink
    if (link == undefined) {
        return
    }

    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        // set up the cursor style
        document.body.style.cursor = 'ns-resize'

        let deltaY = shapeMouseEvent.Point.Y - linkConf.PointAtMouseDown!.Y
        if (linkConf.draggedSegmentNumber == 0 && deltaY != 0) {

            let newStartRatio = (shapeMouseEvent.Point.Y - linkConf.previousStart!.Y) / linkConf.previousStart!.Height
            link.StartRatio = newStartRatio

            if (newStartRatio < 0 || newStartRatio > 1) {

                let nextStartRatio = (shapeMouseEvent.Point.X - link.Start!.X) / link.Start!.Width

                // we swap first segment from horizontal to vertical only if the new ratio
                // is within 0 and 1 
                if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    // the cursor is at the left of the start rectangle
                    if (newStartRatio < 0) {

                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                        // Math.max(link.Start!.Y - offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)

                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }

                    // 
                    if (newStartRatio > 1) {
                        let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y

                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height

                        console.log("link.CornerOffsetRatio", link.CornerOffsetRatio)
                    }
                    // switch dragged element to next segment
                    linkConf.draggedSegmentNumber = 1

                    newStartRatio = nextStartRatio
                    link.StartOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    link.StartRatio = newStartRatio
                    linkConf.drawSegments(link, linkConf.linkUpdating, linkConf.map_Link_Segment)
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                }
            }
        }

        // last segment
        if (linkConf.draggedSegmentNumber == linkConf.segments!.length - 1 && deltaY != 0) {

            let newRatio = (shapeMouseEvent.Point.Y - linkConf.previousEnd!.Y) / linkConf.previousEnd!.Height

            if (newRatio < 0 || newRatio > 1) {
                let nextStartRatio = (shapeMouseEvent.Point.X - link.End!.X) / link.End!.Width

                if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    if (newRatio < 0) {
                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }
                    if (newRatio > 1) {
                        let targetY_ForVerticalMiddleSegment = shapeMouseEvent.Point.Y
                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }
                    // switch dragged element to previous segment
                    linkConf.draggedSegmentNumber = linkConf.segments!.length - 1

                    newRatio = nextStartRatio
                    link.EndOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    link.EndRatio = newRatio
                    linkConf.drawSegments(link, linkConf.linkUpdating, linkConf.map_Link_Segment)
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newRatio < 0) { newRatio = 0 }
                    if (newRatio > 1) { newRatio = 1 }
                }
            }
            link.EndRatio = newRatio
        }

        // middle segment
        if (linkConf.segments!.length == 3 && linkConf.draggedSegmentNumber == 1 && deltaY != 0) {

            let newCornerOffsetRatio =
                (shapeMouseEvent.Point.Y - link.Start!.Y) / link.Start!.Height

            link.CornerOffsetRatio = newCornerOffsetRatio
        }
    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        // set up the cursor style
        document.body.style.cursor = 'ew-resize'

        let deltaX = (shapeMouseEvent.Point.X - linkConf.PointAtMouseDown!.X)

        // first segment
        if (linkConf.draggedSegmentNumber == 0 && deltaX != 0) {

            let newStartRatio = (shapeMouseEvent.Point.X - linkConf.previousStart!.X) / linkConf.previousStart!.Width

            link.StartRatio = newStartRatio

            // special case where we need to change orientation
            if (newStartRatio < 0 || newStartRatio > 1) {

                // we change orientation (horizontal to vertical), therefore we compute the new start ratio
                let newStartVerticalRatio = (shapeMouseEvent.Point.Y - link.Start!.Y) / link.Start!.Height

                if (newStartVerticalRatio >= 0 && newStartVerticalRatio <= 1) {
                    if (newStartRatio < 0) {
                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetX_ForVerticalMiddleSegment = shapeMouseEvent.Point.X

                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }
                    if (newStartRatio > 1) {
                        let targetX_ForVerticalMiddleSegment = shapeMouseEvent.Point.X
                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }

                    newStartRatio = newStartVerticalRatio
                    link.StartOrientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
                    link.StartRatio = newStartVerticalRatio
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                }
            }

            // in all case, we are finished here
            linkConf.drawSegments(link, linkConf.linkUpdating, linkConf.map_Link_Segment)
            return
        }

        // last segment
        if (linkConf.draggedSegmentNumber == linkConf.segments!.length - 1 && deltaX != 0) {

            let newEndRatio = (shapeMouseEvent.Point.X - linkConf.previousEnd!.X) / linkConf.previousEnd!.Width

            if (newEndRatio < 0 || newEndRatio > 1) {
                let newOrientationRatio = (shapeMouseEvent.Point.Y - link.End!.Y) / link.End!.Height

                if (newOrientationRatio >= 0 && newOrientationRatio <= 1) {
                    if (newEndRatio < 0) {
                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetX_ForVerticalMiddleSegment =
                            Math.max(link.End!.X - offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)
                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }
                    if (newEndRatio > 1) {
                        let targetX_ForVerticalMiddleSegment =
                            Math.max(link.End!.X + link.End!.Width + offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)
                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }

                    newEndRatio = newOrientationRatio
                    link.EndOrientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
                    link.EndRatio = newEndRatio
                    linkConf.drawSegments(link, linkConf.linkUpdating, linkConf.map_Link_Segment)
                    // switch dragged element to previous segment
                    linkConf.draggedSegmentNumber = linkConf.segments!.length - 1
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newEndRatio < 0) { newEndRatio = 0 }
                    if (newEndRatio > 1) { newEndRatio = 1 }
                }
            }
            link.EndRatio = newEndRatio
        }

        // middle segment
        if (linkConf.segments!.length == 3 && linkConf.draggedSegmentNumber == 1 && deltaX != 0) {

            let newCornerOffsetRatio = (shapeMouseEvent.Point.X - linkConf.previousStart!.X) / linkConf.previousStart!.Width

            link.CornerOffsetRatio = newCornerOffsetRatio
        }
    }
    linkConf.drawSegments(link, linkConf.linkUpdating, linkConf.map_Link_Segment)

}