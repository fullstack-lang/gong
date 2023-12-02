import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { drawPointRectSegment } from './draw.point.rect.segment'
import { drawPointPointSegment } from './draw.point.point.segment';

export type SegmentsParams = {
    StartRect: gongsvg.RectDB
    EndRect: gongsvg.RectDB
    StartDirection: gongsvg.OrientationType
    EndDirection: gongsvg.OrientationType
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
    StartPoint: gongsvg.PointDB,
    EndPoint: gongsvg.PointDB
    StartPointWithoutRadius: gongsvg.PointDB,
    EndPointWithoutRadius: gongsvg.PointDB
    Orientation: gongsvg.OrientationType
    Number: number
    ArrowEndAnchoredText: Array<Offset>
}

export function createPoint(x: number, y: number): gongsvg.PointDB {

    let point = new gongsvg.PointDB
    point.X = x
    point.Y = y
    return point
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

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        const startY = StartRect.Y + StartRatio * StartRect.Height
        const c1_X = EndRect.X + EndRatio * EndRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)
        const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        const secondSegment = drawPointRectSegment(c1, EndRect, EndDirection, CornerRadius, 1)
        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        const startX = StartRect.X + StartRatio * StartRect.Width
        const c1_X = startX
        const c1_Y = EndRect.Y + EndRatio * EndRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        const secondSegment = drawPointRectSegment(c1, EndRect, EndDirection, CornerRadius, 1)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        const c1_X = StartRect.X + CornerOffsetRatio * StartRect.Width
        const c1_Y = StartRect.Y + StartRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = c1_X
        const c2_Y = EndRect.Y + EndRatio * EndRect.Height

        let c2 = createPoint(c2_X, c2_Y)

        let firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        let secondSegment = drawPointPointSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_VERTICAL, CornerRadius, 1)
        let thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, CornerRadius, 2)

        // reduce second segment if start and end are aligned
        if (Math.abs(c1_Y - c2_Y) <= 2 * CornerRadius) {
            c2 = createPoint(c2_X, c1_Y)
            firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, 0, 0)
            secondSegment = drawPointPointSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_HORIZONTAL, 0, 1)
            thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, 0, 2)
        }


        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        const c1_X = StartRect.X + StartRatio * StartRect.Width
        const c1_Y = StartRect.Y + CornerOffsetRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = EndRect.X + EndRatio * EndRect.Width
        const c2_Y = c1_Y

        let c2 = createPoint(c2_X, c2_Y)

        if (Math.abs(c1_X - c2_Y) <= CornerRadius) {

        }

        let firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, CornerRadius, 0)
        let secondSegment = drawPointPointSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_HORIZONTAL, CornerRadius, 1)
        let thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, CornerRadius, 2)

        // reduce second segment if start and end are aligned
        if (Math.abs(c1_X - c2_X) <= 2 * CornerRadius) {
            c2 = createPoint(c1_X, c2_Y)
            firstSegment = drawPointRectSegment(c1, StartRect, StartDirection, 0, 0)
            secondSegment = drawPointPointSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_HORIZONTAL, 0, 1)
            thirdSegment = drawPointRectSegment(c2, EndRect, EndDirection, 0, 2)
        }

        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    return segments;
}