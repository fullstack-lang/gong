import * as svg from '../../../svg/src/public-api'
import { Segment } from './draw.segments'
import { TextWidthCalculatorComponent } from './text-width-calculator/text-width-calculator.component'


export function auto_X_offset(
    link: svg.Link,
    segment: Segment,
    text: svg.LinkAnchoredText,
    line: string,
    draggedSegmentPositionOnArrow: string,
    map_text_textWidth: Map<string, number>,
    oneEm: number,
    textWidthCalculator: TextWidthCalculatorComponent,
): number {
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

    if (orientation == svg.OrientationType.ORIENTATION_VERTICAL) {
        if (text.LinkAnchorType == svg.LinkAnchorType.LINK_RIGHT_OR_BOTTOM) {
            offset += 16
        }

        if (text.LinkAnchorType == svg.LinkAnchorType.LINK_LEFT_OR_TOP) {
            let _width = map_text_textWidth.get(line)
            if (_width != undefined) {
                offset -= oneEm
                offset -= _width
                // console.log("cache hit")
            } else {
                if (textWidthCalculator != undefined) {
                    const width = 1.1 * textWidthCalculator.measureTextWidth(line);
                    // console.log(`Width of the text, "` + line + `" : ${width}px`);
                    // const height = textWidthCalculator.measureTextHeight(line);
                    // console.log(`Width of the text, "` + line + `" : ${height}px`);
                    offset -= oneEm
                    offset -= width
                    map_text_textWidth.set(line, width)
                }
            }
        }
    } else { // ORIENTATION_HORIZONTAL

        let onTheRight = segment.EndPoint.X > segment.StartPoint.X
        if (draggedSegmentPositionOnArrow == svg.PositionOnArrowType.POSITION_ON_ARROW_END) {
            onTheRight = !onTheRight
        }

        if (onTheRight) {
            offset += 16
        }
        else {
            let _width = map_text_textWidth.get(line)
            if (_width != undefined) {
                offset -= oneEm
                offset -= _width
                // console.log("cache hit")
            } else {
                if (textWidthCalculator != undefined) {
                    const width = 1.1 * textWidthCalculator.measureTextWidth(line);
                    // console.log(`Width of the text, "` + line + `" : ${width}px`);
                    // const height =textWidthCalculator.measureTextHeight(line);
                    // console.log(`Width of the text, "` + line + `" : ${height}px`);
                    offset -= oneEm
                    offset -= width
                    map_text_textWidth.set(line, width)
                }
            }
        }
    }

    return offset * offsetSign + text.X_Offset
}