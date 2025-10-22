package models

type Circle struct {
	Name           string
	CX, CY, Radius float64

	Presentation
	ShapeConditions

	Animations []*Animate
}
