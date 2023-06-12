package models

type Presentation struct {
	Color                       string
	FillOpacity                 float64
	Stroke                      string
	StrokeWidth                 float64
	StrokeDashArray             string
	StrokeDashArrayWhenSelected string

	// Transform is a string componding SVG transform
	//
	// for instance : rotate(-10 50 100)
	// translate(-36 45.5)
	// skewX(40)
	// scale(1 0.5)
	Transform string
}
