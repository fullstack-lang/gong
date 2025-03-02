import * as svg from '../../../svg/src/public-api'
import { Segment } from './draw.segments';

export function getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {

    // console.log("getOrientation:", segment.Number)

    if (segment.Orientation == svg.OrientationType.ORIENTATION_HORIZONTAL) {
        return 'horizontal';
    } else if (segment.Orientation == svg.OrientationType.ORIENTATION_VERTICAL) {
        return 'vertical';
    } else {
        return null; // You can return null or another value if the line is not strictly horizontal or vertical
    }
}