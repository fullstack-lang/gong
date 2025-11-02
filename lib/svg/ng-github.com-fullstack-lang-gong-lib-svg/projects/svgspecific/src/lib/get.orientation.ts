import * as svg from '../../../svg/src/public-api'
import { Segment } from './draw.segments';
import { SvgOrientationType } from './svg-orientation-type';

export function getOrientation(segment: Segment): 'horizontal' | 'vertical' | null {

    // console.log("getOrientation:", segment.Number)

    if (segment.Orientation == SvgOrientationType.ORIENTATION_HORIZONTAL) {
        return 'horizontal';
    } else if (segment.Orientation == SvgOrientationType.ORIENTATION_VERTICAL) {
        return 'vertical';
    } else {
        return null; // You can return null or another value if the line is not strictly horizontal or vertical
    }
}