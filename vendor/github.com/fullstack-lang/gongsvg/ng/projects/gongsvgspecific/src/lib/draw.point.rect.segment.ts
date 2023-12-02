import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { Segment, createPoint } from './draw.segments';
import { drawPointPointSegment } from './draw.point.point.segment';
import { swapSegment } from './swap.segment';

export function drawPointRectSegment(
    point: gongsvg.PointDB,
    rect: gongsvg.RectDB,
    direction: gongsvg.OrientationType,
    cornerRadius: number,
    number: number):
    Segment {
    let segment =
        drawPointPointSegment(point, createPoint(0, 0), direction, cornerRadius, number)

    if (direction === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
        let c1_X_withCornerOffset = point.X
        if (point.X < rect.X + 0.5 * rect.Width) {
            if (point.X < rect.X) {
                c1_X_withCornerOffset = c1_X_withCornerOffset + cornerRadius
            } else {
                c1_X_withCornerOffset = c1_X_withCornerOffset - cornerRadius
            }
        } else {
            if (point.X > rect.X + rect.Width) {
                c1_X_withCornerOffset = c1_X_withCornerOffset - cornerRadius
            } else {
                c1_X_withCornerOffset = c1_X_withCornerOffset + cornerRadius
            }
        }

        let startPoint_withHorizontalCornerOffset = createPoint(c1_X_withCornerOffset, point.Y)
        segment.StartPoint = startPoint_withHorizontalCornerOffset

        if (point.X <= rect.X + rect.Width / 2) {
            const endPoint = createPoint(rect.X, point.Y)
            segment.EndPoint = endPoint
        }
        else {
            const endPoint = createPoint(rect.X + rect.Width, point.Y)
            segment.EndPoint = endPoint
        }
    }
    if (direction === gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        let c1_Y_withCornerOffset = point.Y
        if (point.Y < rect.Y + 0.5 * rect.Height) {
            if (point.Y < rect.Y) {
                c1_Y_withCornerOffset = c1_Y_withCornerOffset + cornerRadius
            } else {
                c1_Y_withCornerOffset = c1_Y_withCornerOffset - cornerRadius
            }
        } else {
            if (point.Y < rect.Y + rect.Height) {
                c1_Y_withCornerOffset = c1_Y_withCornerOffset + cornerRadius
            } else {
                c1_Y_withCornerOffset = c1_Y_withCornerOffset - cornerRadius
            }
        }
        const startPoint_withVerticalCornerOffset = createPoint(point.X, c1_Y_withCornerOffset)
        segment.StartPoint = startPoint_withVerticalCornerOffset

        if (point.Y <= rect.Y + rect.Height / 2) {
            const endPoint = createPoint(point.X, rect.Y)
            segment.EndPoint = endPoint
        }
        else {
            const endPoint = createPoint(point.X, rect.Y + rect.Height)
            segment.EndPoint = endPoint
        }
    }
    segment.EndPointWithoutRadius = segment.EndPoint
    if (number == 0) {
        segment = swapSegment(segment)
    }
    return segment
}