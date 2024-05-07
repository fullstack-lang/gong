import * as gongsvg from '../../../gongsvg/src/public-api'
import { Segment } from './draw.segments';

export function getArcPath(link: gongsvg.Link, segment: Segment, nextSegment: Segment): string {

    const startDegree = 180
    const endDegree = 270
    const startRadians = (startDegree * Math.PI) / 180;
    const endRadians = (endDegree * Math.PI) / 180;
    const startX = segment.EndPoint.X
    const startY = segment.EndPoint.Y
    const endX = nextSegment.StartPoint.X
    const endY = nextSegment.StartPoint.Y
    const largeArcFlag = endDegree - startDegree <= 180 ? 0 : 1;

    // 1 is positive angle direction
    // 0 otherwise
    let sweepFlag = 0
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        let segmentDirection = 0
        if (segment.EndPoint.X > segment.StartPoint.X) {
            segmentDirection = 1
        } else {
            segmentDirection = -1
        }

        let cornerDirection = 0
        if (segment.EndPoint.Y < nextSegment.StartPoint.Y) {
            cornerDirection = 1
        } else {
            cornerDirection = -1
        }

        if (segmentDirection * cornerDirection == 1) {
            sweepFlag = 1
        } else {
            sweepFlag = 0
        }

    }
    if (segment.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
        let segmentDirection = 0
        if (segment.EndPoint.Y > segment.StartPoint.Y) {
            segmentDirection = 1
        } else {
            segmentDirection = -1
        }

        let cornerDirection = 0
        if (segment.EndPoint.X < nextSegment.StartPoint.X) {
            cornerDirection = 1
        } else {
            cornerDirection = -1
        }

        if (segmentDirection * cornerDirection == 1) {
            sweepFlag = 0
        } else {
            sweepFlag = 1
        }
    }
    return `M ${startX} ${startY} A ${link.CornerRadius} ${link.CornerRadius} 0 ${largeArcFlag} ${sweepFlag} ${endX} ${endY}`;
}