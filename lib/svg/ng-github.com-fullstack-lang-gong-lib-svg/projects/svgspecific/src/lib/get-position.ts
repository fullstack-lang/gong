import * as svg from '../../../svg/src/public-api'
import { createPoint } from './draw.segments';
import { Coordinate } from './rectangle-event.service';
import { drawLineFromRectToB } from './draw.line.from.rect.to.point';

export function getPosition(
    startRect: svg.Rect | undefined,
    position: string | undefined,
    endRect?: svg.Rect | undefined
): Coordinate {

    let coordinate: Coordinate = [0, 0]

    if (startRect == undefined || position == undefined) {
        return coordinate
    }

    switch (position) {
        case svg.AnchorType.ANCHOR_BOTTOM:
            coordinate = [startRect.X + startRect.Width / 2, startRect.Y + startRect.Height]
            break;
        case svg.AnchorType.ANCHOR_TOP:
            coordinate = [startRect.X + startRect.Width / 2, startRect.Y]
            break;
        case svg.AnchorType.ANCHOR_LEFT:
            coordinate = [startRect.X, startRect.Y + startRect.Height / 2]
            break;
        case svg.AnchorType.ANCHOR_RIGHT:
            coordinate = [startRect.X + startRect.Width, startRect.Y + startRect.Height / 2]
            break;
        case svg.AnchorType.ANCHOR_CENTER:
            if (endRect == undefined) {
                coordinate = [startRect.X + startRect.Width / 2, startRect.Y + startRect.Height / 2]
            } else {
                let endRectCenter = createPoint(endRect.X + endRect.Width / 2, endRect.Y + endRect.Height / 2)
                let borderPoint = drawLineFromRectToB(startRect, endRectCenter)
                coordinate = [borderPoint.X, borderPoint.Y]
            }
            break;
    }

    return coordinate
}