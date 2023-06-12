import * as gongsvg from 'gongsvg'
import { createPoint } from './link/draw.segments'

export function mouseCoordInComponentRef(event: MouseEvent): gongsvg.PointDB {

    const targetElement = event.target as SVGElement
    const svg = targetElement.ownerSVGElement

    if (svg) {
        const point = svg.createSVGPoint()

        point.x = event.clientX;
        point.y = event.clientY;

        // console.log(point.x, point.y); // window x, y

        const SVGmatrix = svg.getScreenCTM()

        const localPoint = point.matrixTransform(SVGmatrix!.inverse());
        let res = createPoint(localPoint.x, localPoint.y)
        return res
    } else {
        var boundingClientRect = targetElement.getBoundingClientRect();
        var offsetX = event.clientX - boundingClientRect.left;
        var offsetY = event.clientY - boundingClientRect.top;

        // console.log("event.client X", event.clientX, "Y", event.clientY)
        // console.log("window.scroll X", window.scrollX, "Y", window.scrollY)
        // console.log("event.page X", event.pageX, "Y", event.pageY)
        // console.log("offset X", offsetX, "Y", offsetY)
        // console.log("")

        let res = createPoint(offsetX, offsetY)
        return res
    }
}