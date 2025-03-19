package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
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

func (p1 *Presentation) CopyTo(p2 *gongsvg_models.Presentation) {
	p2.Color = p1.Color
	p2.FillOpacity = p1.FillOpacity
	p2.Stroke = p1.Stroke
	p2.StrokeOpacity = p1.StrokeOpacity
	p2.StrokeWidth = p1.StrokeWidth
	p2.StrokeDashArray = p1.StrokeDashArray
	p2.StrokeDashArrayWhenSelected = p1.StrokeDashArrayWhenSelected
	p2.Transform = p1.Transform
}
