import * as gongsvg from '../../../gongsvg/src/public-api'
import { Segment } from './draw.segments'
import { getFunctionName } from './gongsvg-diagramming/get.function.name'



// computeLinkFromMouseEvent takes a mouse event
// 1. update the link conf
// 2. call drawSegment
export function updateLinkFromCursor(
    draggedLink: gongsvg.Link,
    draggedSegmentNumber: number,
    segments: Segment[],

    // for change detection, we need to store start and end rect
    // TODO : understand why we need this previous start
    previousStart: gongsvg.Rect | undefined,
    previousEnd: gongsvg.Rect | undefined,

    pointAtMouseDown: gongsvg.Point,
    pointAtMouseMove: gongsvg.Point,

) {



    // offSetForNewMidlleSegment denotes the standard distance between
    // a rect and the middle segment that is created when going from a
    // two segment link to a three segment link
    const offSetForNewMidlleSegment = 100

    // in order to have middle segment always visible
    const offsetFromBorderForNewMidlleSegment = 50

    if (draggedSegmentNumber >= segments!.length) {
        draggedSegmentNumber = segments!.length - 1
    }

    let segment = segments![draggedSegmentNumber]
    let link = draggedLink
    if (link == undefined) {
        return
    }

    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        // set up the cursor style
        document.body.style.cursor = 'ns-resize'

        let deltaY = pointAtMouseMove.Y - pointAtMouseDown!.Y
        if (draggedSegmentNumber == 0 && deltaY != 0) {

            let newStartRatio = (pointAtMouseMove.Y - previousStart!.Y) / previousStart!.Height
            link.StartRatio = newStartRatio

            if (newStartRatio < 0 || newStartRatio > 1) {

                let nextStartRatio = (pointAtMouseMove.X - link.Start!.X) / link.Start!.Width

                // we swap first segment from horizontal to vertical only if the new ratio
                // is within 0 and 1 
                if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    // the cursor is at the left of the start rectangle
                    if (newStartRatio < 0) {

                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetY_ForVerticalMiddleSegment = pointAtMouseMove.Y
                        // Math.max(link.Start!.Y - offSetForNewMidlleSegment, offsetFromBorderForNewMidlleSegment)

                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }

                    // 
                    if (newStartRatio > 1) {
                        let targetY_ForVerticalMiddleSegment = pointAtMouseMove.Y

                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height

                        console.log("link.CornerOffsetRatio", link.CornerOffsetRatio)
                    }
                    // switch dragged element to next segment
                    draggedSegmentNumber = 1

                    newStartRatio = nextStartRatio
                    link.StartOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    link.StartRatio = newStartRatio
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                }
            }
        }

        // last segment
        if (draggedSegmentNumber == segments!.length - 1 && deltaY != 0) {

            // console.log(getFunctionName(), previousEnd?.Y)
            let newRatio = (pointAtMouseMove.Y - previousEnd!.Y) / previousEnd!.Height

            if (newRatio < 0 || newRatio > 1) {
                let nextStartRatio = (pointAtMouseMove.X - link.End!.X) / link.End!.Width

                if (nextStartRatio >= 0 && nextStartRatio <= 1) {

                    if (newRatio < 0) {
                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetY_ForVerticalMiddleSegment = pointAtMouseMove.Y
                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }
                    if (newRatio > 1) {
                        let targetY_ForVerticalMiddleSegment = pointAtMouseMove.Y
                        link.CornerOffsetRatio = (targetY_ForVerticalMiddleSegment - link.Start!.Y) / link.Start!.Height
                    }
                    // switch dragged element to previous segment
                    draggedSegmentNumber = segments!.length - 1

                    newRatio = nextStartRatio
                    link.EndOrientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
                    link.EndRatio = newRatio
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newRatio < 0) { newRatio = 0 }
                    if (newRatio > 1) { newRatio = 1 }
                }
            }
            link.EndRatio = newRatio
        }

        // middle segment
        if (segments!.length == 3 && draggedSegmentNumber == 1 && deltaY != 0) {

            let newCornerOffsetRatio =
                (pointAtMouseMove.Y - link.Start!.Y) / link.Start!.Height

            link.CornerOffsetRatio = newCornerOffsetRatio
        }
    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {



        // set up the cursor style
        document.body.style.cursor = 'ew-resize'

        let deltaX = (pointAtMouseMove.X - pointAtMouseDown!.X)

        // first segment
        if (draggedSegmentNumber == 0 && deltaX != 0) {

            let newStartRatio = (pointAtMouseMove.X - previousStart!.X) / previousStart!.Width

            link.StartRatio = newStartRatio

            // console.log("updateLinkFromCursor", draggedSegmentNumber, "newStartRatio", newStartRatio)

            // special case where we need to change orientation
            if (newStartRatio < 0 || newStartRatio > 1) {

                // we change orientation (horizontal to vertical), therefore we compute the new start ratio
                let newStartVerticalRatio = (pointAtMouseMove.Y - link.Start!.Y) / link.Start!.Height

                if (newStartVerticalRatio >= 0 && newStartVerticalRatio <= 1) {
                    if (newStartRatio < 0) {
                        // we compute the CornerOffsetRatio in order to have the vertical middle segment
                        // at the left of the End Rect
                        let targetX_ForVerticalMiddleSegment = pointAtMouseMove.X

                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }
                    if (newStartRatio > 1) {
                        let targetX_ForVerticalMiddleSegment = pointAtMouseMove.X
                        link.CornerOffsetRatio = (targetX_ForVerticalMiddleSegment - link.Start!.X) / link.Start!.Width
                    }

                    newStartRatio = newStartVerticalRatio
                    link.StartOrientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
                    link.StartRatio = newStartVerticalRatio
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newStartRatio < 0) { newStartRatio = 0 }
                    if (newStartRatio > 1) { newStartRatio = 1 }
                    link.StartRatio = newStartRatio
                }
            }

            // in all case, we are finished here
            return
        }

        // last segment
        if (draggedSegmentNumber == segments!.length - 1 && deltaX != 0) {

            let newEndRatio = (pointAtMouseMove.X - previousEnd!.X) / previousEnd!.Width

            if (newEndRatio < 0 || newEndRatio > 1) {
                let newOrientationRatio = (pointAtMouseMove.Y - link.End!.Y) / link.End!.Height

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
                    // switch dragged element to previous segment
                    draggedSegmentNumber = segments!.length - 1
                } else {
                    document.body.style.cursor = 'not-allowed'
                    if (newEndRatio < 0) { newEndRatio = 0 }
                    if (newEndRatio > 1) { newEndRatio = 1 }
                }
            }
            link.EndRatio = newEndRatio
        }

        // middle segment
        if (segments!.length == 3 && draggedSegmentNumber == 1 && deltaX != 0) {

            let newCornerOffsetRatio = (pointAtMouseMove.X - previousStart!.X) / previousStart!.Width

            link.CornerOffsetRatio = newCornerOffsetRatio
        }
    }
}