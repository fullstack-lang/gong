package models

// Position mirrors joint.dia.Point
// swagger:model Position
type Position struct {
	X    float64
	Y    float64
	Name string // temporary
}
