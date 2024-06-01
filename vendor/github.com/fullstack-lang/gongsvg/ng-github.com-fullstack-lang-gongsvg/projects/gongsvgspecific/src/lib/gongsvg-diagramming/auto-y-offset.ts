import * as gongsvg from '../../../../gongsvg/src/public-api'
import { Segment } from '../draw.segments'
import { TextWidthCalculatorComponent } from '../text-width-calculator/text-width-calculator.component'

export function auto_Y_offset(
    link: gongsvg.Link,
    segment: Segment,
    text: gongsvg.LinkAnchoredText,
    draggedSegmentPositionOnArrow: string,
    oneEm: number): number {
    // console.log(getFunctionName(), "text", text.Content)

    let offset = 0
    let offsetSign = 1

    if (!text.AutomaticLayout) {
        return offset
    }

    let orientation: string
    if (draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
        orientation = link.EndOrientation
    } else {
        orientation = link.StartOrientation
    }

    if (draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_START) {
        offsetSign = - offsetSign
    }

    if (orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
        if (segment.EndPoint.Y > segment.StartPoint.Y) {
            offsetSign = -offsetSign
        }
    } else {
        if (draggedSegmentPositionOnArrow == gongsvg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            offsetSign = - offsetSign
        }
        if (text.LinkAnchorType == gongsvg.LinkAnchorType.LINK_RIGHT_OR_BOTTOM) {
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

    return offset * offsetSign
}