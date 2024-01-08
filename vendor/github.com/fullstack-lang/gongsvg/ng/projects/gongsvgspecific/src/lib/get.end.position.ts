import * as gongsvg from 'gongsvg'
import { Segment, createPoint } from './draw.segments';
import { Coordinate } from './rectangle-event.service';
import { getAnchorPoint } from './get.anchor.point';
import { drawLineFromRectToB } from './draw.line.from.rect.to.point';

export function getEndPosition(rectLinkLink: gongsvg.RectLinkLinkDB, map_Link_Segment: Map<gongsvg.LinkDB, Segment[]>): Coordinate {

    let coordinate: Coordinate = [0, 0]
    if (rectLinkLink.End == undefined || rectLinkLink.Start == undefined) {
        return coordinate
    }

    let sourceSegments = map_Link_Segment.get(rectLinkLink.End)
    if (sourceSegments == undefined) {
        return coordinate
    }

    let target = getAnchorPoint(sourceSegments, rectLinkLink.TargetAnchorPosition)
    if (target == undefined) {
        return coordinate
    }
    let source = drawLineFromRectToB(rectLinkLink.Start, target)

    coordinate[0] = target.X
    coordinate[1] = target.Y

    return coordinate
}