import * as svg from '../../../svg/src/public-api'
import { Segment, createPoint } from './draw.segments';
import { Coordinate } from './rectangle-event.service';
import { getAnchorPoint } from './get.anchor.point';
import { drawLineFromRectToB } from './draw.line.from.rect.to.point';

export function getStartPosition(rectLinkLink: svg.RectLinkLink, map_Link_Segment: Map<svg.Link, Segment[]>): Coordinate {

    // console.log("getStartPosition:", rectLinkLink.Name)

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

    coordinate[0] = source.X
    coordinate[1] = source.Y

    return coordinate
}