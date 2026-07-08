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

	IsWithHiddenHandle bool
}

func (s *AxesShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *AxesShape) GetIsHidden() bool {
	return s.IsHidden
}

func (s *AxesShape) SetIsWithHiddenHandle(isWithHiddenHandle bool) {
	s.IsWithHiddenHandle = isWithHiddenHandle
}

func (s *AxesShape) GetIsWithHiddenHandle() bool {
	return s.IsWithHiddenHandle
}
