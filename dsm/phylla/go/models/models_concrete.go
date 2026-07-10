package models

// AxesShape of a plant, describe the Cartesian reference X & Y vector
//
// each plant has one AxesShape button
type AxesShape struct {
	Name string

	LengthX float64
	LengthY float64

	IsWithHiddenHandle bool
}

func (s *AxesShape) SetIsWithHiddenHandle(isWithHiddenHandle bool) {
	s.IsWithHiddenHandle = isWithHiddenHandle
}

func (s *AxesShape) GetIsWithHiddenHandle() bool {
	return s.IsWithHiddenHandle
}

// RhombusShape of a plant,
type RhombusShape struct {
	Name string

	X, Y float64
}

// PlantCircumferenceShape of a plant,
//
// each plant has one PlantCircumferenceShape
type PlantCircumferenceShape struct {
	Name string

	AngleDegree float64
	Length      float64
}

// GridPathShape of a plant,
//
// each plant has one GridPathShape to visualize the
// N and M steps along the grid to reach the PlantCircumference.
type GridPathShape struct {
	Name string
}

// CircleGridShape of a plant,
type CircleGridShape struct {
	Name string
}

// NextCircleShape of a plant,
type NextCircleShape struct {
	Name string
}

// ExplanationTextShape of a plant,
type ExplanationTextShape struct {
	Name string
}

// InitialRhombusShape of a plant,
type InitialRhombusShape struct {
	Name string

	X, Y float64
}

// RotatedRhombusShape of a plant,
type RotatedRhombusShape struct {
	Name string

	X, Y float64
}

// GrowthCurveRhombusShape of a plant,
type GrowthCurveRhombusShape struct {
	Name string

	X, Y float64
}

// GrowthVectorShape of a plant,
type GrowthVectorShape struct {
	Name string

	X, Y float64
}

// InitialRhombusGridShape of a plant,
type InitialRhombusGridShape struct {
	Name string

	InitialRhombusShapes []*InitialRhombusShape
}

// RotatedRhombusGridShape of a plant,
type RotatedRhombusGridShape struct {
	Name string

	RotatedRhombusShapes []*RotatedRhombusShape
}

// RotatedRhombusGridShape of a plant,
type GrowthCurveRhombusGridShape struct {
	Name string

	GrowthCurveRhombusShapes []*GrowthCurveRhombusShape
}

// PerpendicularVector of a plant,
type PerpendicularVector struct {
	Name string

	StartX, StartY float64
	EndX, EndY float64
}

// PerpendicularVectorGrid of a plant,
type PerpendicularVectorGrid struct {
	Name string

	PerpendicularVectors []*PerpendicularVector
}

// GrowthCurveBezierShape of a plant,
type GrowthCurveBezierShape struct {
	Name string

	StartX, StartY float64
	ControlPointStartX, ControlPointStartY float64

	EndX, EndY float64
	ControlPointEndX, ControlPointEndY float64
}

// GrowthCurveBezierShapeGrid of a plant,
type GrowthCurveBezierShapeGrid struct {
	Name string

	GrowthCurveBezierShapes []*GrowthCurveBezierShape
}
