package models

type SVGImplInterface interface {

	// SVGUpdated function is called each time a SVG is modified
	SVGUpdated(updatedSVG *SVG)
}
