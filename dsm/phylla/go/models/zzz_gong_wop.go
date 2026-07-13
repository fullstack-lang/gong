// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type ArcNormalVectorShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64
}

func (from *ArcNormalVectorShape) CopyBasicFields(to *ArcNormalVectorShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
}

type ArcNormalVectorShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *ArcNormalVectorShapeGrid) CopyBasicFields(to *ArcNormalVectorShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type AxesShape_WOP struct {
	// insertion point

	Name string

	LengthX float64

	LengthY float64

	IsWithHiddenHandle bool
}

func (from *AxesShape) CopyBasicFields(to *AxesShape) {
	// insertion point
	to.Name = from.Name
	to.LengthX = from.LengthX
	to.LengthY = from.LengthY
	to.IsWithHiddenHandle = from.IsWithHiddenHandle
}

type BaseVectorShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64
}

func (from *BaseVectorShape) CopyBasicFields(to *BaseVectorShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
}

type BaseVectorShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *BaseVectorShapeGrid) CopyBasicFields(to *BaseVectorShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type CircleGridShape_WOP struct {
	// insertion point

	Name string
}

func (from *CircleGridShape) CopyBasicFields(to *CircleGridShape) {
	// insertion point
	to.Name = from.Name
}

type EndArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *EndArcShape) CopyBasicFields(to *EndArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type EndArcShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *EndArcShapeGrid) CopyBasicFields(to *EndArcShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type ExplanationTextShape_WOP struct {
	// insertion point

	Name string
}

func (from *ExplanationTextShape) CopyBasicFields(to *ExplanationTextShape) {
	// insertion point
	to.Name = from.Name
}

type GridPathShape_WOP struct {
	// insertion point

	Name string
}

func (from *GridPathShape) CopyBasicFields(to *GridPathShape) {
	// insertion point
	to.Name = from.Name
}

type GrowthCurve2D_WOP struct {
	// insertion point

	Name string
}

func (from *GrowthCurve2D) CopyBasicFields(to *GrowthCurve2D) {
	// insertion point
	to.Name = from.Name
}

type GrowthCurveBezierShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	ControlPointStartX float64

	ControlPointStartY float64

	EndX float64

	EndY float64

	ControlPointEndX float64

	ControlPointEndY float64
}

func (from *GrowthCurveBezierShape) CopyBasicFields(to *GrowthCurveBezierShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.ControlPointStartX = from.ControlPointStartX
	to.ControlPointStartY = from.ControlPointStartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.ControlPointEndX = from.ControlPointEndX
	to.ControlPointEndY = from.ControlPointEndY
}

type GrowthCurveBezierShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *GrowthCurveBezierShapeGrid) CopyBasicFields(to *GrowthCurveBezierShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type GrowthCurveRhombusGridShape_WOP struct {
	// insertion point

	Name string
}

func (from *GrowthCurveRhombusGridShape) CopyBasicFields(to *GrowthCurveRhombusGridShape) {
	// insertion point
	to.Name = from.Name
}

type GrowthCurveRhombusShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *GrowthCurveRhombusShape) CopyBasicFields(to *GrowthCurveRhombusShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type GrowthVectorShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *GrowthVectorShape) CopyBasicFields(to *GrowthVectorShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type InitialRhombusGridShape_WOP struct {
	// insertion point

	Name string
}

func (from *InitialRhombusGridShape) CopyBasicFields(to *InitialRhombusGridShape) {
	// insertion point
	to.Name = from.Name
}

type InitialRhombusShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *InitialRhombusShape) CopyBasicFields(to *InitialRhombusShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type Library_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64

	LogoSVGFile string

	ComputedPrefix string

	IsExpanded bool

	IsRootLibrary bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsRootLibrary = from.IsRootLibrary
}

type NextCircleShape_WOP struct {
	// insertion point

	Name string
}

func (from *NextCircleShape) CopyBasicFields(to *NextCircleShape) {
	// insertion point
	to.Name = from.Name
}

type PerpendicularVector_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64
}

func (from *PerpendicularVector) CopyBasicFields(to *PerpendicularVector) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
}

type PerpendicularVectorGrid_WOP struct {
	// insertion point

	Name string
}

func (from *PerpendicularVectorGrid) CopyBasicFields(to *PerpendicularVectorGrid) {
	// insertion point
	to.Name = from.Name
}

type PerpendicularVectorGridHalfway_WOP struct {
	// insertion point

	Name string
}

func (from *PerpendicularVectorGridHalfway) CopyBasicFields(to *PerpendicularVectorGridHalfway) {
	// insertion point
	to.Name = from.Name
}

type PerpendicularVectorHalfway_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64
}

func (from *PerpendicularVectorHalfway) CopyBasicFields(to *PerpendicularVectorHalfway) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
}

type Plant_WOP struct {
	// insertion point

	Name string

	N int

	M int

	StackHeight int

	RhombusInsideAngle float64

	VerticalThickness float64

	RadialThickness float64

	RhombusSideLength float64

	ComputedPrefix string

	IsExpanded bool

	IsSelected bool

	IsPlantDiagramsNodeExpanded bool
}

func (from *Plant) CopyBasicFields(to *Plant) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.StackHeight = from.StackHeight
	to.RhombusInsideAngle = from.RhombusInsideAngle
	to.VerticalThickness = from.VerticalThickness
	to.RadialThickness = from.RadialThickness
	to.RhombusSideLength = from.RhombusSideLength
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsSelected = from.IsSelected
	to.IsPlantDiagramsNodeExpanded = from.IsPlantDiagramsNodeExpanded
}

type PlantCircumferenceShape_WOP struct {
	// insertion point

	Name string

	AngleDegree float64

	Length float64
}

func (from *PlantCircumferenceShape) CopyBasicFields(to *PlantCircumferenceShape) {
	// insertion point
	to.Name = from.Name
	to.AngleDegree = from.AngleDegree
	to.Length = from.Length
}

type PlantDiagram_WOP struct {
	// insertion point

	Name string

	OriginX float64

	OriginY float64

	IsRhombusNodesExpanded bool

	IsHiddenAxesShape bool

	IsHiddenReferenceRhombus bool

	IsHiddenPlantCircumferenceShape bool

	IsHiddenGridPathShape bool

	IsHiddenRhombusGridShape bool

	IsHiddenExplanationTextShape bool

	IsHiddenRotatedReferenceRhombus bool

	IsHiddenRotatedPlantCircumferenceShape bool

	IsHiddenRotatedGridPathShape bool

	IsHiddenRotatedRhombusGridShape bool

	IsHiddenGrowthPathRhombusGridShape bool

	IsHiddenGrowthVectorShape bool

	IsHiddenPerpendicularVectorGrid bool

	IsHiddenPerpendicularVectorGridHalfway bool

	IsHiddenBaseVectorShapeGrid bool

	IsHiddenArcNormalVectorShapeGrid bool

	IsHiddenStartArcShapeGrid bool

	IsHiddenTopStartArcShapeGrid bool

	IsHiddenEndArcShapeGrid bool

	IsHiddenTopEndArcShapeGrid bool

	IsHiddenBottomStartArcShapeGrid bool

	IsHiddenBottomEndArcShapeGrid bool

	IsHiddenGrowthCurveBezierShapeGrid bool

	IsHiddenStackOfGrowthCurve bool

	IsHiddenTopStackOfGrowthCurve bool

	IsHiddenBottomStackOfGrowthCurve bool

	IsHiddenGrowthCurve2D bool

	IsHiddenTopGrowthCurve2D bool

	IsChecked bool

	ComputedPrefix string

	IsExpanded bool
}

func (from *PlantDiagram) CopyBasicFields(to *PlantDiagram) {
	// insertion point
	to.Name = from.Name
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.IsRhombusNodesExpanded = from.IsRhombusNodesExpanded
	to.IsHiddenAxesShape = from.IsHiddenAxesShape
	to.IsHiddenReferenceRhombus = from.IsHiddenReferenceRhombus
	to.IsHiddenPlantCircumferenceShape = from.IsHiddenPlantCircumferenceShape
	to.IsHiddenGridPathShape = from.IsHiddenGridPathShape
	to.IsHiddenRhombusGridShape = from.IsHiddenRhombusGridShape
	to.IsHiddenExplanationTextShape = from.IsHiddenExplanationTextShape
	to.IsHiddenRotatedReferenceRhombus = from.IsHiddenRotatedReferenceRhombus
	to.IsHiddenRotatedPlantCircumferenceShape = from.IsHiddenRotatedPlantCircumferenceShape
	to.IsHiddenRotatedGridPathShape = from.IsHiddenRotatedGridPathShape
	to.IsHiddenRotatedRhombusGridShape = from.IsHiddenRotatedRhombusGridShape
	to.IsHiddenGrowthPathRhombusGridShape = from.IsHiddenGrowthPathRhombusGridShape
	to.IsHiddenGrowthVectorShape = from.IsHiddenGrowthVectorShape
	to.IsHiddenPerpendicularVectorGrid = from.IsHiddenPerpendicularVectorGrid
	to.IsHiddenPerpendicularVectorGridHalfway = from.IsHiddenPerpendicularVectorGridHalfway
	to.IsHiddenBaseVectorShapeGrid = from.IsHiddenBaseVectorShapeGrid
	to.IsHiddenArcNormalVectorShapeGrid = from.IsHiddenArcNormalVectorShapeGrid
	to.IsHiddenStartArcShapeGrid = from.IsHiddenStartArcShapeGrid
	to.IsHiddenTopStartArcShapeGrid = from.IsHiddenTopStartArcShapeGrid
	to.IsHiddenEndArcShapeGrid = from.IsHiddenEndArcShapeGrid
	to.IsHiddenTopEndArcShapeGrid = from.IsHiddenTopEndArcShapeGrid
	to.IsHiddenBottomStartArcShapeGrid = from.IsHiddenBottomStartArcShapeGrid
	to.IsHiddenBottomEndArcShapeGrid = from.IsHiddenBottomEndArcShapeGrid
	to.IsHiddenGrowthCurveBezierShapeGrid = from.IsHiddenGrowthCurveBezierShapeGrid
	to.IsHiddenStackOfGrowthCurve = from.IsHiddenStackOfGrowthCurve
	to.IsHiddenTopStackOfGrowthCurve = from.IsHiddenTopStackOfGrowthCurve
	to.IsHiddenBottomStackOfGrowthCurve = from.IsHiddenBottomStackOfGrowthCurve
	to.IsHiddenGrowthCurve2D = from.IsHiddenGrowthCurve2D
	to.IsHiddenTopGrowthCurve2D = from.IsHiddenTopGrowthCurve2D
	to.IsChecked = from.IsChecked
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Rendered3DShape_WOP struct {
	// insertion point

	Name string

	ViewX float64

	ViewY float64

	ViewZ float64

	TargetX float64

	TargetY float64

	TargetZ float64

	Fov float64
}

func (from *Rendered3DShape) CopyBasicFields(to *Rendered3DShape) {
	// insertion point
	to.Name = from.Name
	to.ViewX = from.ViewX
	to.ViewY = from.ViewY
	to.ViewZ = from.ViewZ
	to.TargetX = from.TargetX
	to.TargetY = from.TargetY
	to.TargetZ = from.TargetZ
	to.Fov = from.Fov
}

type RhombusShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *RhombusShape) CopyBasicFields(to *RhombusShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type RotatedRhombusGridShape_WOP struct {
	// insertion point

	Name string
}

func (from *RotatedRhombusGridShape) CopyBasicFields(to *RotatedRhombusGridShape) {
	// insertion point
	to.Name = from.Name
}

type RotatedRhombusShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *RotatedRhombusShape) CopyBasicFields(to *RotatedRhombusShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type StackGrowthCurveEndArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *StackGrowthCurveEndArcShape) CopyBasicFields(to *StackGrowthCurveEndArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type StackGrowthCurveStartArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *StackGrowthCurveStartArcShape) CopyBasicFields(to *StackGrowthCurveStartArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type StackOfGrowthCurve_WOP struct {
	// insertion point

	Name string
}

func (from *StackOfGrowthCurve) CopyBasicFields(to *StackOfGrowthCurve) {
	// insertion point
	to.Name = from.Name
}

type StartArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *StartArcShape) CopyBasicFields(to *StartArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type StartArcShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *StartArcShapeGrid) CopyBasicFields(to *StartArcShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type TopEndArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *TopEndArcShape) CopyBasicFields(to *TopEndArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type TopEndArcShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *TopEndArcShapeGrid) CopyBasicFields(to *TopEndArcShapeGrid) {
	// insertion point
	to.Name = from.Name
}

type TopGrowthCurve2D_WOP struct {
	// insertion point

	Name string
}

func (from *TopGrowthCurve2D) CopyBasicFields(to *TopGrowthCurve2D) {
	// insertion point
	to.Name = from.Name
}

type TopStackGrowthCurveEndArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *TopStackGrowthCurveEndArcShape) CopyBasicFields(to *TopStackGrowthCurveEndArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type TopStackGrowthCurveStartArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *TopStackGrowthCurveStartArcShape) CopyBasicFields(to *TopStackGrowthCurveStartArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type TopStackOfGrowthCurve_WOP struct {
	// insertion point

	Name string
}

func (from *TopStackOfGrowthCurve) CopyBasicFields(to *TopStackOfGrowthCurve) {
	// insertion point
	to.Name = from.Name
}

type TopStartArcShape_WOP struct {
	// insertion point

	Name string

	StartX float64

	StartY float64

	EndX float64

	EndY float64

	XAxisRotation float64

	LargeArcFlag bool

	SweepFlag bool

	RadiusX float64

	RadiusY float64
}

func (from *TopStartArcShape) CopyBasicFields(to *TopStartArcShape) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.StartY = from.StartY
	to.EndX = from.EndX
	to.EndY = from.EndY
	to.XAxisRotation = from.XAxisRotation
	to.LargeArcFlag = from.LargeArcFlag
	to.SweepFlag = from.SweepFlag
	to.RadiusX = from.RadiusX
	to.RadiusY = from.RadiusY
}

type TopStartArcShapeGrid_WOP struct {
	// insertion point

	Name string
}

func (from *TopStartArcShapeGrid) CopyBasicFields(to *TopStartArcShapeGrid) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
