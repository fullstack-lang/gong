import * as svg from '../../../svg/src/public-api'; // Replace 'gongsvg' with the correct module name
import { drawPointRectSegment } from './draw.point.rect.segment'
import { drawPointPointSegment } from './draw.point.point.segment';
import { SvgOrientationType } from './svg-orientation-type';
import { getPosition } from './get-position';
import { getLineOrientation } from './gete-line-orientation';
import { controlPointToPoint } from './control-points';

export type SegmentsParams = {
    StartRect: svg.Rect
    EndRect: svg.Rect
    StartDirection: svg.OrientationType
    EndDirection: svg.OrientationType
    StartRatio: number
    EndRatio: number
    CornerOffsetRatio: number
    CornerRadius: number
}

export type Offset = {
    X_Offset: number
    Y_Offset: number
}

export type Segment = {
    StartPoint: svg.Point,
    EndPoint: svg.Point
    StartPointWithoutRadius: svg.Point,
    EndPointWithoutRadius: svg.Point
    Orientation: SvgOrientationType
    Number: number
    ArrowEndAnchoredText: Array<Offset>
}

export function createPoint(x: number, y: number): svg.Point {

    let point = new svg.Point
    point.X = x
    point.Y = y
    return point
}

export function drawSegmentsFromLink(link: svg.Link): Segment[] {

    if (link.Type === svg.LinkType.LINK_TYPE_FLOATING_ORTHOGONAL) {
        let segmentsParams = {
            StartRect: link.Start!,
            EndRect: link.End!,
            StartDirection: link.StartOrientation! as svg.OrientationType,
            EndDirection: link.EndOrientation! as svg.OrientationType,
            StartRatio: link.StartRatio,
            EndRatio: link.EndRatio,
            CornerOffsetRatio: link.CornerOffsetRatio,
            CornerRadius: link.CornerRadius,
        }
        return drawSegments(segmentsParams)
    }

    if (link.Type === svg.LinkType.LINK_TYPE_LINE_WITH_CONTROL_POINTS) {
        const segments: Segment[] = []

        // Get all points in the polyline: Start Anchor, Control Points, End Anchor
        const allPoints: svg.Point[] = []

        const startPos = getPosition(link.Start, link.StartAnchorType, link.End, link.StartArrowOffset);
        const endPos = getPosition(link.End, link.EndAnchorType, link.Start, link.EndArrowOffset);

        allPoints.push(createPoint(startPos[0], startPos[1]))
        for (let controlPoint of link.ControlPoints) {
            allPoints.push( controlPointToPoint(controlPoint))
        }
        allPoints.push(createPoint(endPos[0], endPos[1]))

        const cornerRadius = link.CornerRadius

        // If no radius or not enough points for a corner, use the simple logic
        if (cornerRadius === 0 || allPoints.length < 3) {
            for (let i = 0; i < allPoints.length - 1; i++) {
                const p1 = allPoints[i]
                const p2 = allPoints[i + 1]
                const orientation = getLineOrientation(p1.X, p1.Y, p2.X, p2.Y)

                let segment: Segment = {
                    StartPoint: p1,
                    EndPoint: p2,
                    StartPointWithoutRadius: p1,
                    EndPointWithoutRadius: p2,
                    Orientation: orientation,
                    Number: i,
                    ArrowEndAnchoredText: []
                }
                segments.push(segment)
            }
            return segments
        }

        // --- NEW LOGIC for CornerRadius > 0 ---

        let segmentStartPoint = allPoints[0] // Start of the *current* line segment

        // Loop over all corners (all points except the very first and very last)
        for (let i = 1; i < allPoints.length - 1; i++) {
            const P_prev = allPoints[i - 1]
            const P_curr = allPoints[i] // This is the vertex
            const P_next = allPoints[i + 1]

            // Calculate vectors *from* the corner
            const v1 = { X: P_prev.X - P_curr.X, Y: P_prev.Y - P_curr.Y }
            const v2 = { X: P_next.X - P_curr.X, Y: P_next.Y - P_curr.Y }

            const lenV1 = Math.sqrt(v1.X * v1.X + v1.Y * v1.Y)
            const lenV2 = Math.sqrt(v2.X * v2.X + v2.Y * v2.Y)

            if (lenV1 === 0 || lenV2 === 0) {
                continue // Points are coincident, skip
            }

            const v1_unit = { X: v1.X / lenV1, Y: v1.Y / lenV1 }
            const v2_unit = { X: v2.X / lenV2, Y: v2.Y / lenV2 }

            const dot = v1_unit.X * v2_unit.X + v1_unit.Y * v2_unit.Y
            const theta = Math.acos(Math.max(-1, Math.min(1, dot))) // Angle at vertex

            // Calculate cutback distance 'd'
            let d = cornerRadius / Math.tan(theta / 2)

            // Safety check: Don't cut back more than half the segment length
            d = Math.min(d, lenV1 / 2, lenV2 / 2)
            if (isNaN(d)) {
                d = 0 // Handle collinear points
            }

            // Calculate the end of the first line (start of arc)
            const arcStartPoint = createPoint(P_curr.X + v1_unit.X * d, P_curr.Y + v1_unit.Y * d)
            // Calculate the start of the next line (end of arc)
            const arcEndPoint = createPoint(P_curr.X + v2_unit.X * d, P_curr.Y + v2_unit.Y * d)

            // Create the segment leading *up to* the corner
            segments.push({
                StartPoint: segmentStartPoint,
                EndPoint: arcStartPoint,
                StartPointWithoutRadius: allPoints[i - 1],
                EndPointWithoutRadius: P_curr,
                Orientation: getLineOrientation(segmentStartPoint.X, segmentStartPoint.Y, arcStartPoint.X, arcStartPoint.Y),
                Number: i - 1,
                ArrowEndAnchoredText: []
            })

            // The start of the *next* segment is the end of this arc
            segmentStartPoint = arcEndPoint
        }

        // Add the final segment (from the last corner's arc-end to the final point)
        const lastPoint = allPoints[allPoints.length - 1]
        segments.push({
            StartPoint: segmentStartPoint,
            EndPoint: lastPoint,
            StartPointWithoutRadius: allPoints[allPoints.length - 2],
            EndPointWithoutRadius: lastPoint,
            Orientation: getLineOrientation(segmentStartPoint.X, segmentStartPoint.Y, lastPoint.X, lastPoint.Y),
            Number: segments.length,
            ArrowEndAnchoredText: []
        })

        return segments
    }

    return []
}

export function drawSegments(params: SegmentsParams): Segment[] {
    const {
        StartRect: StartRect,
        EndRect: EndRect,
        StartDirection: StartDirection,
        EndDirection: EndDirection,
        StartRatio: StartRatio,
        EndRatio: EndRatio,
        CornerOffsetRatio: CornerOffsetRatio,
        CornerRadius: CornerRadius,
    } = params;
    const segments: Segment[] = []

    if (StartDirection === svg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === svg.OrientationType.ORIENTATION_VERTICAL) {

        const startY = StartRect.Y + StartRatio * StartRect.Height
        const c1_X = EndRect.X + EndRatio * EndRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)
        const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        const secondSegment = drawPointRectSegment(c1, EndRect, EndDirection, CornerRadius, 1)
        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === svg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === svg.OrientationType.ORIENTATION_HORIZONTAL) {

        const startX = StartRect.X + StartRatio * StartRect.Width
        const c1_X = startX
        const c1_Y = EndRect.Y + EndRatio * EndRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        const secondSegment = drawPointRectSegment(c1, EndRect, EndDirection, CornerRadius, 1)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === svg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === svg.OrientationType.ORIENTATION_HORIZONTAL) {

        const c1_X = StartRect.X + CornerOffsetRatio * StartRect.Width
        const c1_Y = StartRect.Y + StartRatio * StartRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = c1_X
        const c2_Y = EndRect.Y + EndRatio * EndRect.Height
        let c2 = createPoint(c2_X, c2_Y)

        // reduce second segment if start and end are aligned
        if (Math.abs(c1_Y - c2_Y) <= 2 * CornerRadius) {
            c2 = createPoint(c2_X, c1_Y)
            const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, 0, 0)
            const secondSegment = drawPointPointSegment(c1, c2, svg.OrientationType.ORIENTATION_HORIZONTAL, 0, 1)
            const thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, 0, 2)
            segments.push(firstSegment, secondSegment, thirdSegment)
        } else {
            const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
            const secondSegment = drawPointPointSegment(c1, c2, svg.OrientationType.ORIENTATION_VERTICAL, CornerRadius, 1)
            const thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, CornerRadius, 2)
            segments.push(firstSegment, secondSegment, thirdSegment)
        }
    }

    if (StartDirection === svg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === svg.OrientationType.ORIENTATION_VERTICAL) {

        const c1_X = StartRect.X + StartRatio * StartRect.Width
        const c1_Y = StartRect.Y + CornerOffsetRatio * StartRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = EndRect.X + EndRatio * EndRect.Width
        const c2_Y = c1_Y
        let c2 = createPoint(c2_X, c2_Y)

        // reduce second segment if start and end are aligned
        if (Math.abs(c1_X - c2_X) <= 2 * CornerRadius) {
            c2 = createPoint(c1_X, c2_Y)
            const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, 0, 0)
            const secondSegment = drawPointPointSegment(c1, c2, svg.OrientationType.ORIENTATION_HORIZONTAL, 0, 1)
            const thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, 0, 2)
            segments.push(firstSegment, secondSegment, thirdSegment)
        } else {
            const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
            const secondSegment = drawPointPointSegment(c1, c2, svg.OrientationType.ORIENTATION_HORIZONTAL, CornerRadius, 1)
            const thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, CornerRadius, 2)
            segments.push(firstSegment, secondSegment, thirdSegment)
        }
    }

    return segments;
}