
export function getTextSize(
    textContent: string,
    fontSize: string,
    strokeWidth: string,
    fontFamily: string
): { width: number; height: number } {
    // Create an off-screen SVG element
    const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("width", "0");
    svg.setAttribute("height", "0");
    document.body.appendChild(svg);

    // Add the text element with the given stroke width
    const text = document.createElementNS("http://www.w3.org/2000/svg", "text");
    text.setAttribute("x", "0");
    text.setAttribute("y", "0");
    text.setAttribute("font-size", fontSize);
    text.setAttribute("stroke-width", strokeWidth);
    text.setAttribute("font-family", fontFamily);
    text.textContent = textContent;
    svg.appendChild(text);

    // Get the bounding box of the text element
    const bbox = text.getBBox();

    // Convert the width and height from SVG user units to CSS pixels
    const ctm = svg.getScreenCTM();
    const widthInCSSPixels = bbox.width * ctm!.a;
    const heightInCSSPixels = bbox.height * ctm!.d;

    // Remove the off-screen SVG element
    document.body.removeChild(svg);

    // Return the width and height in CSS pixels
    return { width: widthInCSSPixels, height: heightInCSSPixels };
}
export function getComputedFontSizeAndFontFamily(element: HTMLElement) {
    const computedStyle = window.getComputedStyle(element);
    const fontSize = computedStyle.getPropertyValue("font-size");
    const fontFamily = computedStyle.getPropertyValue("font-family");
    return { fontSize, fontFamily };
}