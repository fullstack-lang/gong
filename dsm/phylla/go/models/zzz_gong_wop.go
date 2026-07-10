// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
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

type CircleGridShape_WOP struct {
	// insertion point

	Name string
}

func (from *CircleGridShape) CopyBasicFields(to *CircleGridShape) {
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

type Plant_WOP struct {
	// insertion point

	Name string

	N int

	M int

	RhombusInsideAngle float64

	RhombusSideLength float64

	ShiftToNearestCircle int

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
	to.RhombusInsideAngle = from.RhombusInsideAngle
	to.RhombusSideLength = from.RhombusSideLength
	to.ShiftToNearestCircle = from.ShiftToNearestCircle
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

	IsHiddenGrowthCurveBezierShapeGrid bool

	IsChecked bool

	ComputedPrefix string

	IsExpanded bool
}

func (from *PlantDiagram) CopyBasicFields(to *PlantDiagram) {
	// insertion point
	to.Name = from.Name
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
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
	to.IsHiddenGrowthCurveBezierShapeGrid = from.IsHiddenGrowthCurveBezierShapeGrid
	to.IsChecked = from.IsChecked
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
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

// end of insertion point
