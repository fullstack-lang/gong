package models

type Shape struct {
	IsHidden bool
}

// AxesShape of a plant, describe the Cartesian reference X & Y vector
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

// ReferenceRhombus of a plant,
//
// each plant has one ReferenceRhombus
// the direction and distance of new cell growth.
type ReferenceRhombus struct {
	Name string

	Shape
}

func (s *ReferenceRhombus) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *ReferenceRhombus) GetIsHidden() bool {
	return s.IsHidden
}

// GrowthVectorShape of a plant,
//
// each plant has one GrowthVectorShape
// the direction and distance of new cell growth.
type GrowthVectorShape struct {
	Name string

	AngleDegree float64
	Length      float64

	Shape
}

func (s *GrowthVectorShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *GrowthVectorShape) GetIsHidden() bool {
	return s.IsHidden
}
