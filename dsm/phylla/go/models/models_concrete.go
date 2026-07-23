package models

// gong:omit
type RhombusStuff struct {
	Name                           string
	ReferenceRhombus               *RhombusShape
	PlantCircumferenceShape        *PlantCircumferenceShape
	GridPathShape                  *GridPathShape
	InitialRhombusGridShape        *InitialRhombusGridShape
	ExplanationTextShape           *ExplanationTextShape
	RotatedReferenceRhombus        *RhombusShape
	RotatedPlantCircumferenceShape *PlantCircumferenceShape
	RotatedGridPathShape           *GridPathShape
	RotatedRhombusGridShape2       *RotatedRhombusGridShape
	GrowthCurveRhombusGridShape    *GrowthCurveRhombusGridShape
}

// AxesShape of a plant, describe the Cartesian reference X & Y vector
//
// each plant has one AxesShape button
// gong:omit
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
// gong:omit
type RhombusShape struct {
	Name string

	X, Y float64
}

// PlantCircumferenceShape of a plant,
//
// each plant has one PlantCircumferenceShape
// gong:omit
type PlantCircumferenceShape struct {
	Name string

	AngleDegree float64
	Length      float64
}

// GridPathShape of a plant,
//
// each plant has one GridPathShape to visualize the
// N and M steps along the grid to reach the PlantCircumference.
// gong:omit
type GridPathShape struct {
	Name string
}

// CircleGridShape of a plant,
// gong:omit
type CircleGridShape struct {
	Name string
}

// ExplanationTextShape of a plant,
// gong:omit
type ExplanationTextShape struct {
	Name string
}

// InitialRhombusShape of a plant,
// gong:omit
type InitialRhombusShape struct {
	Name string

	X, Y float64
}

// RotatedRhombusShape of a plant,
// gong:omit
type RotatedRhombusShape struct {
	Name string

	X, Y float64
}

// GrowthCurveRhombusShape of a plant,
// gong:omit
type GrowthCurveRhombusShape struct {
	Name string

	X, Y float64
}

// GrowthVectorShape of a plant,
// gong:omit
type GrowthVectorShape struct {
	Name string

	X, Y float64
}

// InitialRhombusGridShape of a plant,
// gong:omit
type InitialRhombusGridShape struct {
	Name string

	InitialRhombusShapes []*InitialRhombusShape
}

// RotatedRhombusGridShape of a plant,
// gong:omit
type RotatedRhombusGridShape struct {
	Name string

	RotatedRhombusShapes []*RotatedRhombusShape
}

// RotatedRhombusGridShape of a plant,
// gong:omit
type GrowthCurveRhombusGridShape struct {
	Name string

	GrowthCurveRhombusShapes []*GrowthCurveRhombusShape
}

// PerpendicularVector of a plant,
// gong:omit
type PerpendicularVector struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// PerpendicularVectorGrid of a plant,
// gong:omit
type PerpendicularVectorGrid struct {
	Name string

	PerpendicularVectors []*PerpendicularVector
}

// PerpendicularVectorHalfway of a plant,
// gong:omit
type PerpendicularVectorHalfway struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// PerpendicularVectorGridHalfway of a plant,
// gong:omit
type PerpendicularVectorGridHalfway struct {
	Name string

	PerpendicularVectorHalfways []*PerpendicularVectorHalfway
}

// StartArcShape of a plant,
// gong:omit
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
// gong:omit
type StartArcShapeGrid struct {
	Name string

	StartArcShapes []*StartArcShape
}

// TopStartArcShape of a plant,
// gong:omit
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
// gong:omit
type TopStartArcShapeGrid struct {
	Name string

	TopStartArcShapes []*TopStartArcShape
}

// EndArcShape of a plant,
// gong:omit
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
// gong:omit
type EndArcShapeGrid struct {
	Name string

	EndArcShapes []*EndArcShape
}

// TopEndArcShape of a plant,
// gong:omit
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
// gong:omit
type TopEndArcShapeGrid struct {
	Name string

	TopEndArcShapes []*TopEndArcShape
}

// ArcNormalVectorShape of a plant,
// gong:omit
type ArcNormalVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// ArcNormalVectorShapeGrid of a plant,
// gong:omit
type ArcNormalVectorShapeGrid struct {
	Name string

	ArcNormalVectorShapes []*ArcNormalVectorShape
}

// BaseVectorShape of a plant,
// gong:omit
type BaseVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// BaseVectorShapeGrid of a plant,
// gong:omit
type BaseVectorShapeGrid struct {
	Name string

	BaseVectorShapes []*BaseVectorShape
}

// ShiftedBottomTopStartArcShapeGrid of a plant,
// gong:omit
type ShiftedBottomTopStartArcShapeGrid struct {
	Name string

	ShiftedBottomTopStartArcShapes []*ShiftedBottomTopStartArcShape
}

// gong:omit
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
// gong:omit
type MidArcVectorShapeGrid struct {
	Name string

	MidArcVectorShapes []*MidArcVectorShape
}

// gong:omit
type MidArcVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// TopMidArcVectorShapeGrid of a plant,
// gong:omit
type TopMidArcVectorShapeGrid struct {
	Name string

	TopMidArcVectorShapes []*TopMidArcVectorShape
}

// gong:omit
type TopMidArcVectorShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// ShiftedLeftStackOfNormalVector of a plant,
// gong:omit
type ShiftedLeftStackOfNormalVector struct {
	Name string

	ShiftedLeftStackNormalVectors []*ShiftedLeftStackNormalVector
}

// gong:omit
type ShiftedLeftStackNormalVector struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// StackOfRotatedGrowthCurve2D of a plant,
// gong:omit
type StackOfRotatedGrowthCurve2D struct {
	Name string

	StackRotatedGrowthCurve2DStartArcShapes []*StackRotatedGrowthCurve2DStartArcShape
	StackRotatedGrowthCurve2DEndArcShapes   []*StackRotatedGrowthCurve2DEndArcShape
}

// gong:omit
type StackRotatedGrowthCurve2DStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// gong:omit
type StackRotatedGrowthCurve2DEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// TopStackOfRotatedGrowthCurve2D of a plant,
// gong:omit
type TopStackOfRotatedGrowthCurve2D struct {
	Name string

	TopStackOfRotatedGrowthCurve2DStartArcShapes []*TopStackOfRotatedGrowthCurve2DStartArcShape
	TopStackOfRotatedGrowthCurve2DEndArcShapes   []*TopStackOfRotatedGrowthCurve2DEndArcShape
}

// ShiftedLeftStackOfGrowthCurve of a plant,
// gong:omit
type ShiftedLeftStackOfGrowthCurve struct {
	Name string

	ShiftedLeftStackGrowthCurveStartArcShapes []*ShiftedLeftStackGrowthCurveStartArcShape
	ShiftedLeftStackGrowthCurveEndArcShapes   []*ShiftedLeftStackGrowthCurveEndArcShape
}

// gong:omit
type TopStackOfRotatedGrowthCurve2DStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// gong:omit
type TopStackOfRotatedGrowthCurve2DEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// gong:omit
type ShiftedLeftStackGrowthCurveStartArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// gong:omit
type ShiftedLeftStackGrowthCurveEndArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
	RadiusX, RadiusY float64
}

// GrowthCurve2D of a plant,
// gong:omit
type GrowthCurve2D struct {
	Name string

	StartHalfwayArcShapeGrid *StartHalfwayArcShapeGrid
	EndHalfwayArcShapeGrid   *EndHalfwayArcShapeGrid
}

// TopGrowthCurve2D of a plant,
// gong:omit
type TopGrowthCurve2D struct {
	Name string

	TopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
	TopEndHalfwayArcShapeGrid   *TopEndHalfwayArcShapeGrid
}

// gong:omit
type StartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type StartHalfwayArcShapeGrid struct {
	Name                  string
	StartHalfwayArcShapes []*StartHalfwayArcShape
}

// gong:omit
type EndHalfwayArcShapeGrid struct {
	Name                string
	EndHalfwayArcShapes []*EndHalfwayArcShape
}

// gong:omit
type EndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type TopStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type TopStartHalfwayArcShapeGrid struct {
	Name                     string
	TopStartHalfwayArcShapes []*TopStartHalfwayArcShape
}

// gong:omit
type TopEndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type TopEndHalfwayArcShapeGrid struct {
	Name                   string
	TopEndHalfwayArcShapes []*TopEndHalfwayArcShape
}

// TorusStackShape of a plant,
// gong:omit
type TorusStackShape struct {
	Name string
}

// VerticalTorusStackShape of a plant,
// gong:omit
type VerticalTorusStackShape struct {
	Name string
}

// PartiallyRotatedTorusShape of a plant,
// gong:omit
type PartiallyRotatedTorusShape struct {
	Name string
}

// StackOfPartiallyRotatedTorusShape of a plant,
// gong:omit
type StackOfPartiallyRotatedTorusShape struct {
	Name string
}

// StackOfGrowthCurve2D of a plant,
// gong:omit
type StackOfGrowthCurve2D struct {
	Name string

	StackGrowthCurve2DStartHalfwayArcShapes []*StackGrowthCurve2DStartHalfwayArcShape
	StackGrowthCurve2DEndHalfwayArcShapes   []*StackGrowthCurve2DEndHalfwayArcShape
}

// TopStackOfGrowthCurve2D of a plant,
// gong:omit
type TopStackOfGrowthCurve2D struct {
	Name string

	TopStackGrowthCurve2DStartHalfwayArcShapes []*TopStackGrowthCurve2DStartHalfwayArcShape
	TopStackGrowthCurve2DEndHalfwayArcShapes   []*TopStackGrowthCurve2DEndHalfwayArcShape
}

// gong:omit
type StackGrowthCurve2DStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type StackGrowthCurve2DEndHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
type TopStackGrowthCurve2DStartHalfwayArcShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64

	RadiusX, RadiusY float64
	XAxisRotation    float64
	LargeArcFlag     bool
	SweepFlag        bool
}

// gong:omit
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
// gong:omit
type StackOfGrowthCurve2DRibbon struct {
	Name string

	StackGrowthCurve2DRibbonStartShapes []*StackGrowthCurve2DRibbonStartShape
	StackGrowthCurve2DRibbonEndShapes   []*StackGrowthCurve2DRibbonEndShape
}

// gong:omit
type StackGrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// gong:omit
type StackGrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// StackOfRotatedGrowthCurve2DRibbon of a plant,
// gong:omit
type StackOfRotatedGrowthCurve2DRibbon struct {
	Name string

	StackRotatedGrowthCurve2DRibbonStartShapes []*StackRotatedGrowthCurve2DRibbonStartShape
	StackRotatedGrowthCurve2DRibbonEndShapes   []*StackRotatedGrowthCurve2DRibbonEndShape
}

// gong:omit
type StackRotatedGrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// gong:omit
type StackRotatedGrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// GrowthCurve2DRibbon of a plant,
// gong:omit
type GrowthCurve2DRibbon struct {
	Name string

	GrowthCurve2DRibbonStartShapes []*GrowthCurve2DRibbonStartShape
	GrowthCurve2DRibbonEndShapes   []*GrowthCurve2DRibbonEndShape
}

// gong:omit
type GrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// gong:omit
type GrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// ShiftedRightGrowthCurve2DRibbon of a plant,
// gong:omit
type ShiftedRightGrowthCurve2DRibbon struct {
	Name string

	ShiftedRightGrowthCurve2DRibbonStartShapes []*ShiftedRightGrowthCurve2DRibbonStartShape
	ShiftedRightGrowthCurve2DRibbonEndShapes   []*ShiftedRightGrowthCurve2DRibbonEndShape
}

// gong:omit
type ShiftedRightGrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// gong:omit
type ShiftedRightGrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// PartiallyGrowthCurve2DRibbon of a plant,
// gong:omit
type PartiallyGrowthCurve2DRibbon struct {
	Name string

	PartiallyGrowthCurve2DRibbonStartShapes []*PartiallyGrowthCurve2DRibbonStartShape
	PartiallyGrowthCurve2DRibbonEndShapes   []*PartiallyGrowthCurve2DRibbonEndShape
}

// gong:omit
type PartiallyGrowthCurve2DRibbonStartShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// gong:omit
type PartiallyGrowthCurve2DRibbonEndShape struct {
	Name string

	BottomStartX, BottomStartY   float64
	BottomEndX, BottomEndY       float64
	BottomRadiusX, BottomRadiusY float64
	BottomXAxisRotation          float64
	BottomLargeArcFlag           bool
	BottomSweepFlag              bool

	TopStartX, TopStartY   float64
	TopEndX, TopEndY       float64
	TopRadiusX, TopRadiusY float64
	TopXAxisRotation       float64
	TopLargeArcFlag        bool
	TopSweepFlag           bool
}

// PartiallyGrowthCurve2DTrajectory of a plant,
// gong:omit
type PartiallyGrowthCurve2DTrajectory struct {
	Name string

	PartiallyGrowthCurve2DTrajectoryShapes []*PartiallyGrowthCurve2DTrajectoryShape
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP1P2 struct {
	Name string

	P1PointShapes    []*PartiallyGrowthCurve2DTrajectoryP1PointShape
	P2PointShapes    []*PartiallyGrowthCurve2DTrajectoryP2PointShape
	P1CurveShapes    []*PartiallyGrowthCurve2DTrajectoryP1CurveShape
	P2CurveShapes    []*PartiallyGrowthCurve2DTrajectoryP2CurveShape
	P1P2PairLineShapes []*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP1PointShape struct {
	Name string

	X, Y float64
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP2PointShape struct {
	Name string

	X, Y float64
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP1CurveShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP2CurveShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// gong:omit
type PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape struct {
	Name string

	StartX, StartY float64
	EndX, EndY     float64
}

// gong:omit
type PxShape struct {
	Name string

	X, Y float64
}

// gong:omit
type ChosenP1P2PairShape struct {
	Name string

	P1X, P1Y float64
	P2X, P2Y float64
}




