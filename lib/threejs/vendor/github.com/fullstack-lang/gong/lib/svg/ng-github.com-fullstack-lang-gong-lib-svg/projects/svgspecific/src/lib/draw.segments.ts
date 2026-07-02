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

/**
 * Finds the intersection point of a ray (from p1 to p2) with a rectangle.
 * Assumes p1 (control point) is INSIDE the rectangle.
 * Returns the intersection point on the rect boundary.
 */
function getIntersectionPointFromInside(
    p1: svg.Point, // Point inside rect (Control Point)
    p2: svg.Point, // Point outside rect (Next Point)
    rect: svg.Rect
): svg.Point {
    const dx = p2.X - p1.X;
    const dy = p2.Y - p1.Y;

    if (dx === 0 && dy === 0) {
        return createPoint(p1.X, p1.Y); // Points are the same
    }

    let t = Infinity;

    // Check intersection with each of the 4 edges
    // Left edge (x = rect.X)
    if (dx < 0) {
        const tEdge = (rect.X - p1.X) / dx;
        if (tEdge > 0) t = Math.min(t, tEdge);
    }
    // Right edge (x = rect.X + rect.Width)
    if (dx > 0) {
        const tEdge = (rect.X + rect.Width - p1.X) / dx;
        if (tEdge > 0) t = Math.min(t, tEdge);
    }
    // Top edge (y = rect.Y)
    if (dy < 0) {
        const tEdge = (rect.Y - p1.Y) / dy;
        if (tEdge > 0) t = Math.min(t, tEdge);
    }
    // Bottom edge (y = rect.Y + rect.Height)
    if (dy > 0) {
        const tEdge = (rect.Y + rect.Height - p1.Y) / dy;
        if (tEdge > 0) t = Math.min(t, tEdge);
    }

    if (t === Infinity || t > 1) {
        // This case *shouldn't* happen if p1 is inside and p2 is outside,
        // but as a fallback, we return p2 if it's also inside, or p1 otherwise.
        if (p2.X >= rect.X && p2.X <= rect.X + rect.Width &&
            p2.Y >= rect.Y && p2.Y <= rect.Y + rect.Height) {
            return createPoint(p2.X, p2.Y); // p2 is also inside, just use p2
        }
        // Failsafe, return the control point
        return createPoint(p1.X, p1.Y);
    }

    // Calculate intersection point
    return createPoint(p1.X + dx * t, p1.Y + dy * t);
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

        // START: MODIFIED BLOCK (Added symmetrical logic for EndRect)
        
        // Convert CPs to absolute points first
        const controlPointsAbs: svg.Point[] = []
        for (let controlPoint of link.ControlPoints) {
            controlPointsAbs.push(controlPointToPoint(controlPoint))
        }

        // --- Check START Point ---
        let hideFirstSegment = false;
        if (link.ControlPoints && link.ControlPoints.length > 0 && link.Start) {
            
            const firstCP = controlPointsAbs[0]
            const startRect = link.Start
            const isInside = (
                firstCP.X > startRect.X &&
                firstCP.X < startRect.X + startRect.Width &&
                firstCP.Y > startRect.Y &&
                firstCP.Y < startRect.Y + startRect.Height
            )
            const hasClosestRect = link.ControlPoints[0].ClosestRect &&
                                 link.ControlPoints[0].ClosestRect.ID === startRect.ID;

            // Also ensure there is a point *after* the CP to aim at
            if (isInside && hasClosestRect && (controlPointsAbs.length > 1 || link.End)) {
                hideFirstSegment = true;
            }
        }

        // --- Check END Point ---
        let hideLastSegment = false;
        if (link.ControlPoints && link.ControlPoints.length > 0 && link.End) {

            const lastCPIndex = controlPointsAbs.length - 1
            const lastCP = controlPointsAbs[lastCPIndex]
            const endRect = link.End
            const isInside = (
                lastCP.X > endRect.X &&
                lastCP.X < endRect.X + endRect.Width &&
                lastCP.Y > endRect.Y &&
                lastCP.Y < endRect.Y + endRect.Height
            )
            const hasClosestRect = link.ControlPoints[lastCPIndex].ClosestRect &&
                                 link.ControlPoints[lastCPIndex].ClosestRect.ID === endRect.ID;

            // Also ensure there is a point *before* the CP to aim at
            if (isInside && hasClosestRect && (controlPointsAbs.length > 1 || link.Start)) {
                hideLastSegment = true;
            }
        }

        // Anchor point *never* gets an offset if its segment is hidden
        const startOffset = hideFirstSegment ? 0 : link.StartArrowOffset;
        const endOffset = hideLastSegment ? 0 : link.EndArrowOffset;
        // END: MODIFIED BLOCK
        
        const startPos = getPosition(link.Start, link.StartAnchorType, link.End, startOffset);
        // START: MODIFIED BLOCK
        const endPos = getPosition(link.End, link.EndAnchorType, link.Start, endOffset);
        // END: MODIFIED BLOCK

        allPoints.push(createPoint(startPos[0], startPos[1]))
        // START: MODIFIED BLOCK
        allPoints.push(...controlPointsAbs) // Add the pre-calculated absolute points
        // END: MODIFIED BLOCK
        allPoints.push(createPoint(endPos[0], endPos[1]))


        const cornerRadius = link.CornerRadius
        // START: MODIFIED BLOCK
        const startSegmentIndex = hideFirstSegment ? 1 : 0; // Start from 0 or 1
        // The loop will end *before* this index
        const endSegmentIndex = allPoints.length - (hideLastSegment ? 2 : 1);
        // END: MODIFIED BLOCK


        // If no radius or not enough points for a corner, use the simple logic
        if (cornerRadius === 0 || allPoints.length < 3) {
            
            // START: MODIFIED BLOCK
            // Loop from the first visible segment up to the last visible segment
            for (let i = startSegmentIndex; i < endSegmentIndex; i++) {
            // END: MODIFIED BLOCK
                const p1 = allPoints[i]
                const p2 = allPoints[i + 1]
                const orientation = getLineOrientation(p1.X, p1.Y, p2.X, p2.Y)

                // --- Handle Start Point Offset ---
                let segmentStartPoint = p1;
                let startPointWithoutRadius = p1; // The original point (CP or anchor)
                if (hideFirstSegment && i === startSegmentIndex) { 
                    startPointWithoutRadius = allPoints[1] // The original CP
                    const nextPoint = allPoints[2]
                    const boundaryPoint = getIntersectionPointFromInside(startPointWithoutRadius, nextPoint, link.Start!)
                    const offset = link.StartArrowOffset;
                    const dx = nextPoint.X - startPointWithoutRadius.X;
                    const dy = nextPoint.Y - startPointWithoutRadius.Y;
                    const len = Math.sqrt(dx * dx + dy * dy);
                    if (len > 0) {
                        const ratio = offset / len;
                        segmentStartPoint = createPoint(boundaryPoint.X + dx * ratio, boundaryPoint.Y + dy * ratio);
                    } else {
                        segmentStartPoint = boundaryPoint // 0-length, just use boundary
                    }
                }

                // --- Handle End Point Offset ---
                let segmentEndPoint = p2;
                let endPointWithoutRadius = p2;
                if (hideLastSegment && i === endSegmentIndex - 1) { // If this is the new last segment
                    endPointWithoutRadius = allPoints[allPoints.length - 2]; // The Last CP
                    const prevPoint = allPoints[allPoints.length - 3]; // Point before Last CP
                    const boundaryPoint = getIntersectionPointFromInside(endPointWithoutRadius, prevPoint, link.End!)
                    const offset = link.EndArrowOffset;
                    // Vector from CP *towards* previous point
                    const dx = prevPoint.X - endPointWithoutRadius.X;
                    const dy = prevPoint.Y - endPointWithoutRadius.Y;
                    const len = Math.sqrt(dx * dx + dy * dy);
                    if (len > 0) {
                        const ratio = offset / len;
                        // Apply offset from the BOUNDARY point
                        segmentEndPoint = createPoint(boundaryPoint.X + dx * ratio, boundaryPoint.Y + dy * ratio);
                    } else {
                        segmentEndPoint = boundaryPoint;
                    }
                }
                // END: NEW MODIFICATION

                let segment: Segment = {
                    StartPoint: segmentStartPoint,
                    EndPoint: segmentEndPoint,
                    StartPointWithoutRadius: startPointWithoutRadius,
                    EndPointWithoutRadius: endPointWithoutRadius,
                    Orientation: orientation,
                    Number: i - startSegmentIndex, // Adjust segment number
                    ArrowEndAnchoredText: []
                }
                segments.push(segment)
            }
            return segments
        }

        // --- NEW LOGIC for CornerRadius > 0 ---

        // START: NEW MODIFIED BLOCK
        // Calculate the *actual* start point for the first segment
        let segmentStartPoint: svg.Point;
        let startPointWithoutRadius: svg.Point;
        // The corner loop starts at the 2nd point (if visible) or 3rd (if 1st CP is hidden)
        const startLoopIndex = hideFirstSegment ? 2 : 1;
        // The corner loop ends before the 2nd-to-last (if visible) or 3rd-to-last (if last CP is hidden)
        const endLoopIndex = allPoints.length - (hideLastSegment ? 2 : 1);


        if (hideFirstSegment) {
            startPointWithoutRadius = allPoints[1]; // The original CP
            const nextPoint = allPoints[2] // The next point
            const boundaryPoint = getIntersectionPointFromInside(startPointWithoutRadius, nextPoint, link.Start!)
            const offset = link.StartArrowOffset;
            const dx = nextPoint.X - startPointWithoutRadius.X;
            const dy = nextPoint.Y - startPointWithoutRadius.Y;
            const len = Math.sqrt(dx * dx + dy * dy);
            if (len > 0) {
                const ratio = offset / len;
                segmentStartPoint = createPoint(boundaryPoint.X + dx * ratio, boundaryPoint.Y + dy * ratio);
            } else {
                segmentStartPoint = boundaryPoint; 
            }
        } else {
            // Start from anchor (allPoints[0])
            segmentStartPoint = allPoints[0];
            startPointWithoutRadius = allPoints[0];
        }
        // END: NEW MODIFIED BLOCK


        // Loop over all corners
        // START: MODIFIED BLOCK
        for (let i = startLoopIndex; i < endLoopIndex; i++) {
        // END: MODIFIED BLOCK
            const P_prev = allPoints[i - 1]
            const P_curr = allPoints[i] // This is the vertex
            const P_next = allPoints[i + 1]

            // Calculate vectors *from* the corner
            const v1 = { X: P_prev.X - P_curr.X, Y: P_prev.Y - P_curr.Y }
            const v2 = { X: P_next.X - P_curr.X, Y: P_next.Y - P_curr.Y }

            const lenV1 = Math.sqrt(v1.X * v1.X + v1.Y * v1.Y)
            const lenV2 = Math.sqrt(v2.X * v2.X + v2.Y * v2.Y)

            if (lenV1 === 0 || lenV2 === 0) {
                 // START: MODIFIED BLOCK
                // Handle 0-length segments, especially at the start
                if (i === startLoopIndex && (segmentStartPoint.X !== P_curr.X || segmentStartPoint.Y !== P_curr.Y)) {
                     segments.push({
                        StartPoint: segmentStartPoint,
                        EndPoint: P_curr,
                        StartPointWithoutRadius: startPointWithoutRadius,
                        EndPointWithoutRadius: P_curr,
                        Orientation: getLineOrientation(segmentStartPoint.X, segmentStartPoint.Y, P_curr.X, P_curr.Y),
                        Number: segments.length,
                        ArrowEndAnchoredText: []
                    })
                }
                segmentStartPoint = P_curr
                startPointWithoutRadius = P_curr
                // END: MODIFIED BLOCK
                continue // Points are coincident, skip corner logic
            }

            const v1_unit = { X: v1.X / lenV1, Y: v1.Y / lenV1 }
            const v2_unit = { X: v2.X / lenV2, Y: v2.Y / lenV2 }

            const dot = v1_unit.X * v2_unit.X + v1_unit.Y * v2_unit.Y
            const theta = Math.acos(Math.max(-1, Math.min(1, dot))) // Angle at vertex

            // Calculate cutback distance 'd'
            let d = cornerRadius / Math.tan(theta / 2)

            // START: MODIFIED BLOCK
            let firstSegmentLen = Math.sqrt(Math.pow(P_curr.X - segmentStartPoint.X, 2) + Math.pow(P_curr.Y - segmentStartPoint.Y, 2));
            let d_safe = d;
            if (i === startLoopIndex) {
                // For the first visible segment, check against its actual calculated start
                d_safe = Math.min(d, firstSegmentLen / 2, lenV2 / 2)
            } else {
                 d_safe = Math.min(d, lenV1 / 2, lenV2 / 2)
            }
            
            if (isNaN(d_safe) || d_safe < 0) {
                d_safe = 0 // Handle collinear points
            }
            // END: MODIFIED BLOCK

            const arcStartPoint = createPoint(P_curr.X + v1_unit.X * d_safe, P_curr.Y + v1_unit.Y * d_safe)
            const arcEndPoint = createPoint(P_curr.X + v2_unit.X * d_safe, P_curr.Y + v2_unit.Y * d_safe)

            segments.push({
                StartPoint: segmentStartPoint,
                EndPoint: arcStartPoint,
                StartPointWithoutRadius: startPointWithoutRadius,
                EndPointWithoutRadius: P_curr,
                Orientation: getLineOrientation(segmentStartPoint.X, segmentStartPoint.Y, arcStartPoint.X, arcStartPoint.Y),
                Number: segments.length,
                ArrowEndAnchoredText: []
            })
            segmentStartPoint = arcEndPoint
            startPointWithoutRadius = P_curr // The "withoutRadius" start is the corner itself
        }

        // --- Add the final segment ---
        // START: MODIFIED BLOCK
        // Calculate the *actual* end point for the last segment
        const lastVisiblePoint = allPoints[endLoopIndex];
        let segmentEndPoint = lastVisiblePoint;
        let endPointWithoutRadius = lastVisiblePoint;

        if (hideLastSegment) {
            endPointWithoutRadius = allPoints[allPoints.length - 2]; // The Last CP
            const prevPoint = allPoints[allPoints.length - 3]; // Point before Last CP
            const boundaryPoint = getIntersectionPointFromInside(endPointWithoutRadius, prevPoint, link.End!)
            const offset = link.EndArrowOffset;
            // Vector from CP *towards* previous point
            const dx = prevPoint.X - endPointWithoutRadius.X;
            const dy = prevPoint.Y - endPointWithoutRadius.Y;
            const len = Math.sqrt(dx * dx + dy * dy);
            if (len > 0) {
                const ratio = offset / len;
                // Apply offset from the BOUNDARY point
                segmentEndPoint = createPoint(boundaryPoint.X + dx * ratio, boundaryPoint.Y + dy * ratio);
            } else {
                segmentEndPoint = boundaryPoint;
            }
        }
        
        segments.push({
            StartPoint: segmentStartPoint,
            EndPoint: segmentEndPoint,
            StartPointWithoutRadius: startPointWithoutRadius,
            EndPointWithoutRadius: endPointWithoutRadius,
            Orientation: getLineOrientation(segmentStartPoint.X, segmentStartPoint.Y, segmentEndPoint.X, segmentEndPoint.Y),
            Number: segments.length,
            ArrowEndAnchoredText: []
        })
        // END: MODIFIED BLOCK

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