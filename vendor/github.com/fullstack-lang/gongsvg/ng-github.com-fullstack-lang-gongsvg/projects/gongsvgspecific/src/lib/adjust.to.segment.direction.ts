import * as gongsvg from '../../../gongsvg/src/public-api'
import { Segment } from './draw.segments';

export function adjustToSegmentDirection(segment: Segment, x: number, y: number): { x: number, y: number } {
    let x_res = 0
    let y_res = 0

    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
        if (segment.EndPoint.X > segment.StartPoint.X) { // 0'
            x_res = x
            y_res = y
        } else { // pi
            x_res = -x
            y_res = y
        }
    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
        if (segment.EndPoint.Y > segment.StartPoint.Y) { // pi/2
            x_res = y
            y_res = x
        } else { // 3*pi/2
            x_res = -y
            y_res = -x
        }
    }

    return { x: x_res, y: y_res }
}