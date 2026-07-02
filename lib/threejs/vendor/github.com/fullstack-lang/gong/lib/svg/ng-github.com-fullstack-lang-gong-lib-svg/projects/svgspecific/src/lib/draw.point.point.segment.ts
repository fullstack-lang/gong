import * as svg from '../../../svg/src/public-api'; // Replace 'gongsvg' with the correct module name
import { Segment, createPoint } from './draw.segments';
import { SvgOrientationType } from './svg-orientation-type';

export function drawPointPointSegment(
    start: svg.Point,
    end: svg.Point,
    orientation: svg.OrientationType,
    cornerRadius: number,
    number: number): Segment {

    let newStart = structuredClone(start)
    let newEnd = structuredClone(end)

    if (orientation == svg.OrientationType.ORIENTATION_HORIZONTAL) {
        if (start.X < end.X) {
            newStart.X = start.X + cornerRadius
            newEnd.X = end.X - cornerRadius
        } else {
            newStart.X = start.X - cornerRadius
            newEnd.X = end.X + cornerRadius
        }
    }
    if (orientation == svg.OrientationType.ORIENTATION_VERTICAL) {
        if (start.Y < end.Y) {
            newStart.Y = start.Y + cornerRadius
            newEnd.Y = end.Y - cornerRadius
        } else {
            newStart.Y = start.Y - cornerRadius
            newEnd.Y = end.Y + cornerRadius
        }
    }

    let segment: Segment =
    {
        StartPoint: newStart,
        EndPoint: newEnd,
        StartPointWithoutRadius: start,
        EndPointWithoutRadius: end,
        Orientation: orientation as unknown as SvgOrientationType,
        Number: number,
        ArrowEndAnchoredText: [],
    }

    return segment
}
