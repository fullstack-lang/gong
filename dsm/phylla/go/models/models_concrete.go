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

// PlantCircumferenceShape of a plant,
//
// each plant has one PlantCircumferenceShape
type PlantCircumferenceShape struct {
	Name string

	AngleDegree float64
	Length      float64

	Shape
}

func (s *PlantCircumferenceShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *PlantCircumferenceShape) GetIsHidden() bool {
	return s.IsHidden
}

// GridPathShape of a plant,
//
// each plant has one GridPathShape to visualize the
// N and M steps along the grid to reach the PlantCircumference.
type GridPathShape struct {
	Name string

	Shape
}

func (s *GridPathShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *GridPathShape) GetIsHidden() bool {
	return s.IsHidden
}

// RhombusGridShape of a plant,
type RhombusGridShape struct {
	Name string

	Shape
}

func (s *RhombusGridShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *RhombusGridShape) GetIsHidden() bool {
	return s.IsHidden
}

// CircleGridShape of a plant,
type CircleGridShape struct {
	Name string

	Shape
}

func (s *CircleGridShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *CircleGridShape) GetIsHidden() bool {
	return s.IsHidden
}

// NextCircleShape of a plant,
type NextCircleShape struct {
	Name string

	Shape
}

func (s *NextCircleShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *NextCircleShape) GetIsHidden() bool {
	return s.IsHidden
}

// ExplanationTextShape of a plant,
type ExplanationTextShape struct {
	Name string

	Shape
}

func (s *ExplanationTextShape) SetIsHidden(isHidden bool) {
	s.IsHidden = isHidden
}

func (s *ExplanationTextShape) GetIsHidden() bool {
	return s.IsHidden
}
