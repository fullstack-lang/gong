import * as svg from '../../../svg/src/public-api'
import { SvgOrientationType } from './svg-orientation-type';

export function getLineOrientation(p1x: number, p1y: number, p2x: number, p2y: number): SvgOrientationType {
    if (Math.abs(p1x - p2x) < 0.01) { // Use a small epsilon for float comparison
        return SvgOrientationType.ORIENTATION_VERTICAL;
    }
    if (Math.abs(p1y - p2y) < 0.01) { // Use a small epsilon for float comparison
        return SvgOrientationType.ORIENTATION_HORIZONTAL;
    }
    return SvgOrientationType.ORIENTATION_DIAGONAL;
}