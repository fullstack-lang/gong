import * as svg from '../../../svg/src/public-api'
import { createPoint } from './draw.segments'

export function mouseCoordInComponentRef(
    event: MouseEvent,
    zoom: number,      // Added parameter
    shiftX: number,    // Added parameter
    shiftY: number     // Added parameter
): svg.Point {

    const targetElement = event.target as SVGElement;
    // Try to find the owner SVG element reliably
    const svgElement = targetElement.ownerSVGElement;

    if (svgElement) {
        const point = svgElement.createSVGPoint();

        point.x = event.clientX;
        point.y = event.clientY;

        const SVGmatrix = svgElement.getScreenCTM();

        if (SVGmatrix) {
            // Transform screen coordinates to base SVG coordinates
            const localPoint = point.matrixTransform(SVGmatrix.inverse());

            // Apply inverse of the zoom and shift applied to the content group
            const adjustedX = (localPoint.x + shiftX) / zoom;
            const adjustedY = (localPoint.y + shiftY) / zoom;

            let res = createPoint(adjustedX, adjustedY);
            return res;
        } else {
            // Fallback if CTM is null (should be rare for SVG elements)
            console.error("Could not get Screen CTM for SVG element.");
            // Return the original point (screen coordinates relative to SVG viewport).
            // This won't account for zoom/shift or the SVG's internal coordinate system.
            // It's a fallback for an error condition.
            let res = createPoint(point.x, point.y); // <--- FIXED: Use 'point' instead of 'localPoint'
            return res;
        }

    } else {
        // Fallback for non-SVG elements or if ownerSVGElement is null
        // This calculation is relative to the target element's bounding box,
        // NOT the SVG coordinate system. Applying zoom/shift here is complex
        // and likely incorrect without more context about where this element sits.
        // It's best to ensure interactions happen on SVG elements within the main <g>.
        console.warn("Event target is not an SVG element or ownerSVGElement is null. Coordinates might be inaccurate relative to SVG content.");
        var boundingClientRect = targetElement.getBoundingClientRect();
        var offsetX = event.clientX - boundingClientRect.left;
        var offsetY = event.clientY - boundingClientRect.top;

        // Returning offsetX/Y directly. These are NOT in the transformed SVG space.
        let res = createPoint(offsetX, offsetY);
        return res;
    }
}
