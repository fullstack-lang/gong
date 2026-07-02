import * as svg from '../../../svg/src/public-api'
import { Segment } from './draw.segments';
import { SvgOrientationType } from './svg-orientation-type';

export function getArcPath(link: svg.Link, segment: Segment, nextSegment: Segment): string {

    const r = link.CornerRadius
    if (r === 0 || !nextSegment) {
        return ""
    }

    const P1 = segment.EndPoint                // Start of the arc
    const P2 = nextSegment.StartPoint          // End of the arc
    const P_corner = segment.EndPointWithoutRadius // The "true" corner

    // We need to determine the "sweep-flag" (i.e., which way the arc bends)
    // We can use a 2D cross-product of the vectors from the corner

    // Vector from arc-start TO corner
    const v1 = { X: P_corner.X - P1.X, Y: P_corner.Y - P1.Y }
    // Vector from corner TO arc-end
    const v2 = { X: P2.X - P_corner.X, Y: P2.Y - P_corner.Y }

    // Z-component of the cross product
    const cross_z = v1.X * v2.Y - v1.Y * v2.X

    // If cross_z > 0, it's a "left" turn (sweep-flag 1)
    // If cross_z < 0, it's a "right" turn (sweep-flag 0)
    const sweepFlag = (cross_z > 0) ? 1 : 0

    // For a convex polyline, the angle is always < 180, so large-arc-flag is 0
    const largeArcFlag = 0

    // M = MoveTo, A = Arc
    // A [rx, ry, x-axis-rotation, large-arc-flag, sweep-flag, x, y]
    return `M ${P1.X},${P1.Y} A ${r},${r} 0 ${largeArcFlag},${sweepFlag} ${P2.X},${P2.Y}`
}