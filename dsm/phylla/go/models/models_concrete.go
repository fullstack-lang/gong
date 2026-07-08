package models

type Shape struct {
	IsHidden bool
}

// AxesShape of a plant
//
// each plant has one AxesShape button
type AxesShape struct {
	Name string

	LengthX float64
	LengthY float64

	Shape
}

func (s *AxesShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *AxesShape) GetIsHidden() bool {
	return s.IsHidden
}
