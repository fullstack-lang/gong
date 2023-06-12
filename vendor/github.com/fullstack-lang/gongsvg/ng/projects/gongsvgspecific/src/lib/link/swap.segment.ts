import { Segment, createPoint } from './draw.segments';

export function swapSegment(segment: Segment): Segment {

    let res = structuredClone(segment)

    res.StartPoint = segment.EndPoint
    res.StartPointWithoutRadius = segment.EndPointWithoutRadius
    res.EndPoint = segment.StartPoint
    res.EndPointWithoutRadius = segment.StartPointWithoutRadius

    return res

}