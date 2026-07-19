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
	EndX, EndY     float64
}

// PerpendicularVectorGrid of a plant,
type PerpendicularVectorGrid struct {
	Name string

	PerpendicularVectors []*PerpendicularVector
}

// PerpendicularVectorHalfway of a plant,
type PerpendicularVectorHalfway struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// PerpendicularVectorGridHalfway of a plant,
type PerpendicularVectorGridHalfway struct {
	Name string

	PerpendicularVectorHalfways []*PerpendicularVectorHalfway
}

// StartArcShape of a plant,
type StartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// StartArcShapeGrid of a plant,
type StartArcShapeGrid struct {
	Name string

	StartArcShapes []*StartArcShape
}

// TopStartArcShape of a plant,
type TopStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// TopStartArcShapeGrid of a plant,
type TopStartArcShapeGrid struct {
	Name string

	TopStartArcShapes []*TopStartArcShape
}

// EndArcShape of a plant,
type EndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// EndArcShapeGrid of a plant,
type EndArcShapeGrid struct {
	Name string

	EndArcShapes []*EndArcShape
}

// TopEndArcShape of a plant,
type TopEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// TopEndArcShapeGrid of a plant,
type TopEndArcShapeGrid struct {
	Name string

	TopEndArcShapes []*TopEndArcShape
}

// ArcNormalVectorShape of a plant,
type ArcNormalVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// ArcNormalVectorShapeGrid of a plant,
type ArcNormalVectorShapeGrid struct {
	Name string

	ArcNormalVectorShapes []*ArcNormalVectorShape
}

// BaseVectorShape of a plant,
type BaseVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// BaseVectorShapeGrid of a plant,
type BaseVectorShapeGrid struct {
	Name string

	BaseVectorShapes []*BaseVectorShape
}

// ShiftedBottomTopStartArcShapeGrid of a plant,
type ShiftedBottomTopStartArcShapeGrid struct {
	Name string

	ShiftedBottomTopStartArcShapes []*ShiftedBottomTopStartArcShape
}

type ShiftedBottomTopStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// MidArcVectorShapeGrid of a plant,
type MidArcVectorShapeGrid struct {
	Name string

	MidArcVectorShapes []*MidArcVectorShape
}

type MidArcVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// TopMidArcVectorShapeGrid of a plant,
type TopMidArcVectorShapeGrid struct {
	Name string

	TopMidArcVectorShapes []*TopMidArcVectorShape
}

type TopMidArcVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// ShiftedLeftStackOfNormalVector of a plant,
type ShiftedLeftStackOfNormalVector struct {
	Name string

	ShiftedLeftStackNormalVectors []*ShiftedLeftStackNormalVector
}

type ShiftedLeftStackNormalVector struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// StackOfGrowthCurve of a plant,
type StackOfGrowthCurve struct {
	Name string

	StackGrowthCurveStartArcShapes []*StackGrowthCurveStartArcShape
	StackGrowthCurveEndArcShapes   []*StackGrowthCurveEndArcShape
}

type StackGrowthCurveStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

type StackGrowthCurveEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// TopStackOfGrowthCurve of a plant,
type TopStackOfGrowthCurve struct {
	Name string

	TopStackGrowthCurveStartArcShapes []*TopStackGrowthCurveStartArcShape
	TopStackGrowthCurveEndArcShapes   []*TopStackGrowthCurveEndArcShape
}

// ShiftedLeftStackOfGrowthCurve of a plant,
type ShiftedLeftStackOfGrowthCurve struct {
	Name string

	ShiftedLeftStackGrowthCurveStartArcShapes []*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveEndArcShapes   []*ShiftedLeftStackGrowthCurveEndArcShape
}
type TopStackGrowthCurveStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

type TopStackGrowthCurveEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

type ShiftedLeftStackGrowthCurveStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

type ShiftedLeftStackGrowthCurveEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

type Rendered3DShape struct {
	Name string

	ViewX, ViewY, ViewZ       float64
	TargetX, TargetY, TargetZ float64
	Fov                       float64
}

// GrowthCurve2D of a plant,
type GrowthCurve2D struct {
	Name string

	StartHalfwayArcShapeGrid *StartHalfwayArcShapeGrid
	EndHalfwayArcShapeGrid   *EndHalfwayArcShapeGrid
}

// TopGrowthCurve2D of a plant,
type TopGrowthCurve2D struct {
	Name string

	TopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
	TopEndHalfwayArcShapeGrid   *TopEndHalfwayArcShapeGrid
}

type StartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type StartHalfwayArcShapeGrid struct {
	Name                  string
	StartHalfwayArcShapes []*StartHalfwayArcShape
}

type EndHalfwayArcShapeGrid struct {
	Name                string
	EndHalfwayArcShapes []*EndHalfwayArcShape
}

type EndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type TopStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type TopStartHalfwayArcShapeGrid struct {
	Name                     string
	TopStartHalfwayArcShapes []*TopStartHalfwayArcShape
}

type TopEndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type TopEndHalfwayArcShapeGrid struct {
	Name                   string
	TopEndHalfwayArcShapes []*TopEndHalfwayArcShape
}

// TorusStackShape of a plant,
type TorusStackShape struct {
	Name string
}

// StackOfGrowthCurve2D of a plant,
type StackOfGrowthCurve2D struct {
	Name string

	StackGrowthCurve2DStartHalfwayArcShapes []*StackGrowthCurve2DStartHalfwayArcShape
	StackGrowthCurve2DEndHalfwayArcShapes   []*StackGrowthCurve2DEndHalfwayArcShape
}

// TopStackOfGrowthCurve2D of a plant,
type TopStackOfGrowthCurve2D struct {
	Name string

	TopStackGrowthCurve2DStartHalfwayArcShapes []*TopStackGrowthCurve2DStartHalfwayArcShape
	TopStackGrowthCurve2DEndHalfwayArcShapes   []*TopStackGrowthCurve2DEndHalfwayArcShape
}

type StackGrowthCurve2DStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type StackGrowthCurve2DEndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type TopStackGrowthCurve2DStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

type TopStackGrowthCurve2DEndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// StackOfGrowthCurve2DRibbon of a plant,
type StackOfGrowthCurve2DRibbon struct {
	Name string

	StackGrowthCurve2DRibbonStartShapes []*StackGrowthCurve2DRibbonStartShape
	StackGrowthCurve2DRibbonEndShapes   []*StackGrowthCurve2DRibbonEndShape
}

type StackGrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY float64
	BottomEndX, BottomEndY     float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation    float64
	BottomLargeArcFlag     bool
	BottomSweepFlag        bool

	TopStartX, TopStartY float64
	TopEndX, TopEndY     float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation    float64
	TopLargeArcFlag     bool
	TopSweepFlag        bool
}

type StackGrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY float64
	BottomEndX, BottomEndY     float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation    float64
	BottomLargeArcFlag     bool
	BottomSweepFlag        bool

	TopStartX, TopStartY float64
	TopEndX, TopEndY     float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation    float64
	TopLargeArcFlag     bool
	TopSweepFlag        bool
}
