import * as svg from '../../../svg/src/public-api'; // Replace 'gongsvg' with the correct module name
import { drawPointRectSegment } from './draw.point.rect.segment'
import { drawPointPointSegment } from './draw.point.point.segment';
import { SvgOrientationType } from './svg-orientation-type';
import { getPosition } from './get-position';
import { getLineOrientation } from './gete-line-orientation';

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

        const startPos = getPosition(link.Start, link.StartAnchorType, link.End);
        const endPos = getPosition(link.End, link.EndAnchorType, link.Start);

        allPoints.push(createPoint(startPos[0], startPos[1]))
        allPoints.push(...link.ControlPoints)
        allPoints.push(createPoint(endPos[0], endPos[1]))

        // Create one segment for each pair of points
        for (let i = 0; i < allPoints.length - 1; i++) {
            const p1 = allPoints[i]
            const p2 = allPoints[i + 1]
            const orientation = getLineOrientation(p1.X, p1.Y, p2.X, p2.Y)

            let segment: Segment = {
                StartPoint: p1,
                EndPoint: p2,
                StartPointWithoutRadius: p1, // No radius for this link type
                EndPointWithoutRadius: p2,   // No radius for this link type
                Orientation: orientation,
                Number: i,
                ArrowEndAnchoredText: []
            }
            segments.push(segment)
        }
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