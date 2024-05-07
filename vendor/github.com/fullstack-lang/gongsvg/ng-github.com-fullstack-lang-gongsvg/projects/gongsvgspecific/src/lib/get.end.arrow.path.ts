import * as gongsvg from '../../../gongsvg/src/public-api'
import { Segment } from './draw.segments';

export function getEndArrowPath(link: gongsvg.Link, segment: Segment, arrowSize: number): string {
    const ratio = 0.707106781 / 2 // (1/sqrt(2)) / 2

    let firstStartX = segment.EndPoint.X
    let firstStartY = segment.EndPoint.Y

    let secondStartX = segment.EndPoint.X
    let secondStartY = segment.EndPoint.Y

    let firstTipX = segment.EndPoint.X
    let firstTipY = segment.EndPoint.Y
    let secondTipX = segment.EndPoint.X
    let secondTipY = segment.EndPoint.Y

    {
        let { x, y } = rotateToSegmentDirection(segment, - arrowSize, - arrowSize)

        firstTipX += x
        firstTipY += y
    }
    {
        let { x, y } = rotateToSegmentDirection(segment, link.StrokeWidth * ratio, link.StrokeWidth * ratio)
        firstStartX += x
        firstStartY += y
    }
    {
        let { x, y } = rotateToSegmentDirection(segment, - arrowSize, arrowSize)

        secondTipX += x
        secondTipY += y
    }
    {
        let { x, y } = rotateToSegmentDirection(segment, link.StrokeWidth * ratio, - link.StrokeWidth * ratio)

        secondStartX += x
        secondStartY += y
    }

    let path = `M ${firstStartX} ${firstStartY} L ${firstTipX} ${firstTipY} M ${secondStartX} ${secondStartY} L ${secondTipX} ${secondTipY}`

    return path
}

export function rotateToSegmentDirection(segment: Segment, x: number, y: number): { x: number, y: number } {
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