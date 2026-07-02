package models

import (
	"fmt"
	"strings"
)

type Presentation struct {
	Color                       string
	FillOpacity                 float64
	Stroke                      string
	StrokeOpacity               float64
	StrokeWidth                 float64
	StrokeDashArray             string
	StrokeDashArrayWhenSelected string

	// Transform is a string componding SVG transform
	//
	// for instance : rotate(-10 50 100)
	// translate(-36 45.5)
	// skewX(40)
	// scale(1 0.5)
	// gong:text gong:width 600 gong:height 400
	Transform string
}

func (p *Presentation) WriteSVG(sb *strings.Builder) {

	sb.WriteString(
		fmt.Sprintf(
			`
			fill="%s"
            fill-opacity="%s"
            stroke="%s"
            stroke-opacity="%s"
            stroke-width="%s"
            stroke-dasharray="%s"
            transform="%s"`,
			p.Color,
			fmt.Sprintf("%f", p.FillOpacity),
			p.Stroke,
			fmt.Sprintf("%f", p.StrokeOpacity),
			fmt.Sprintf("%f", p.StrokeWidth),
			p.StrokeDashArray,
			p.Transform))
}
