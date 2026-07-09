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

	IsHidden bool

	IsWithHiddenHandle bool
}

func (from *AxesShape) CopyBasicFields(to *AxesShape) {
	// insertion point
	to.Name = from.Name
	to.LengthX = from.LengthX
	to.LengthY = from.LengthY
	to.IsHidden = from.IsHidden
	to.IsWithHiddenHandle = from.IsWithHiddenHandle
}

type CircleGridShape_WOP struct {
	// insertion point

	Name string

	IsHidden bool
}

func (from *CircleGridShape) CopyBasicFields(to *CircleGridShape) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
}

type ExplanationTextShape_WOP struct {
	// insertion point

	Name string

	IsHidden bool
}

func (from *ExplanationTextShape) CopyBasicFields(to *ExplanationTextShape) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
}

type GridPathShape_WOP struct {
	// insertion point

	Name string

	IsHidden bool
}

func (from *GridPathShape) CopyBasicFields(to *GridPathShape) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
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

	IsHidden bool
}

func (from *NextCircleShape) CopyBasicFields(to *NextCircleShape) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
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

	IsHidden bool
}

func (from *PlantCircumferenceShape) CopyBasicFields(to *PlantCircumferenceShape) {
	// insertion point
	to.Name = from.Name
	to.AngleDegree = from.AngleDegree
	to.Length = from.Length
	to.IsHidden = from.IsHidden
}

type PlantDiagram_WOP struct {
	// insertion point

	Name string

	OriginX float64

	OriginY float64

	IsChecked bool

	ComputedPrefix string

	IsExpanded bool
}

func (from *PlantDiagram) CopyBasicFields(to *PlantDiagram) {
	// insertion point
	to.Name = from.Name
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.IsChecked = from.IsChecked
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type RhombusGridShape_WOP struct {
	// insertion point

	Name string

	IsHidden bool
}

func (from *RhombusGridShape) CopyBasicFields(to *RhombusGridShape) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
}

type RhombusShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	IsHidden bool
}

func (from *RhombusShape) CopyBasicFields(to *RhombusShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.IsHidden = from.IsHidden
}

// end of insertion point
