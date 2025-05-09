import * as svg from '../../../svg/src/public-api'
import { Segment } from './draw.segments'
import { TextWidthCalculatorComponent } from './calc/calc.component'

export function auto_Y_offset(
    link: svg.Link,
    segment: Segment,
    text: svg.LinkAnchoredText,
    draggedSegmentPositionOnArrow: string,
    oneEm: number): number {
    // console.log(getFunctionName(), "text", text.Content)

    let offset = 0
    let offsetSign = 1

    if (!text.AutomaticLayout) {
        return offset
    }

    let orientation: string
    if (draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        orientation = link.EndOrientation
    } else {
        orientation = link.StartOrientation
    }

    if (draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        offsetSign = - offsetSign
    }

    if (orientation == svg.OrientationType.ORIENTATION_VERTICAL) {
        if (segment.EndPoint.Y > segment.StartPoint.Y) {
            offsetSign = -offsetSign
        }
    } else {
        if (draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            offsetSign = - offsetSign
        }
        if (text.LinkAnchorType == svg.LinkAnchorType.LINK_RIGHT_OR_BOTTOM) {
            offsetSign = -offsetSign
        }
    }

    if (link.HasEndArrow) {
        offset += link.EndArrowSize
    }
    if (offsetSign == 1) {
        offset += oneEm
    } else {
        offset += oneEm * 0.4
    }

    return offset * offsetSign + text.Y_Offset
}