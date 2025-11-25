package models

type OrientationType string

// values for EnumType
const (
	ORIENTATION_HORIZONTAL OrientationType = "ORIENTATION_HORIZONTAL"
	ORIENTATION_VERTICAL   OrientationType = "ORIENTATION_VERTICAL"
)

type LinkShape struct {
	StartRatio float64
	EndRatio   float64

	StartOrientation OrientationType
	EndOrientation   OrientationType

	CornerOffsetRatio float64
}

// Setter functions for StateTransitionShape
func (s *LinkShape) SetStartRatio(ratio float64) {
	s.StartRatio = ratio
}

func (s *LinkShape) SetEndRatio(ratio float64) {
	s.EndRatio = ratio
}

func (s *LinkShape) SetStartOrientation(orientation OrientationType) {
	s.StartOrientation = orientation
}

func (s *LinkShape) SetEndOrientation(orientation OrientationType) {
	s.EndOrientation = orientation
}

func (s *LinkShape) SetCornerOffsetRatio(ratio float64) {
	s.CornerOffsetRatio = ratio
}

func (s *LinkShape) GetStartRatio() float64 {
	return s.StartRatio
}

func (s *LinkShape) GetEndRatio() float64 {
	return s.EndRatio
}

func (s *LinkShape) GetStartOrientation() OrientationType {
	return s.StartOrientation
}

func (s *LinkShape) GetEndOrientation() OrientationType {
	return s.EndOrientation
}

func (s *LinkShape) GetCornerOffsetRatio() float64 {
	return s.CornerOffsetRatio
}

// Interface for StateTransitionShape
type LinkShapeInterface interface {
	SetStartRatio(ratio float64)
	SetEndRatio(ratio float64)
	SetStartOrientation(orientation OrientationType)
	SetEndOrientation(orientation OrientationType)
	SetCornerOffsetRatio(ratio float64)

	GetStartRatio() float64
	GetEndRatio() float64
	GetStartOrientation() OrientationType
	GetEndOrientation() OrientationType
	GetCornerOffsetRatio() float64
}
