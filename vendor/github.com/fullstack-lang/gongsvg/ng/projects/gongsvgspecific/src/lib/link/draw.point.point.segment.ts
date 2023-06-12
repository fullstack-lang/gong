import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { Segment, createPoint } from './draw.segments';

export function drawPointPointSegment(
    start: gongsvg.PointDB,
    end: gongsvg.PointDB,
    orientation: gongsvg.OrientationType,
    cornerRadius: number,
    number: number): Segment {

    let newStart = structuredClone(start)
    let newEnd = structuredClone(end)

    if (orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
        if (start.X < end.X) {
            newStart.X = start.X + cornerRadius
            newEnd.X = end.X - cornerRadius
        } else {
            newStart.X = start.X - cornerRadius
            newEnd.X = end.X + cornerRadius
        }
    }
    if (orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
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
        Orientation: orientation,
        Number: number,
        ArrowEndAnchoredText: [],
    }

    return segment
}
