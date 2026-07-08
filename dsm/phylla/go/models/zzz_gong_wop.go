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

type GrowthVectorShape_WOP struct {
	// insertion point

	Name string

	AngleDegree float64

	Length float64

	IsHidden bool
}

func (from *GrowthVectorShape) CopyBasicFields(to *GrowthVectorShape) {
	// insertion point
	to.Name = from.Name
	to.AngleDegree = from.AngleDegree
	to.Length = from.Length
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

type Plant_WOP struct {
	// insertion point

	Name string

	N int

	M int

	Z int

	RhombusInsideAngle float64

	RhombusSideLength float64

	ShiftToNearestCircle int

	ComputedPrefix string

	IsExpanded bool

	IsPlantDiagramsNodeExpanded bool
}

func (from *Plant) CopyBasicFields(to *Plant) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.Z = from.Z
	to.RhombusInsideAngle = from.RhombusInsideAngle
	to.RhombusSideLength = from.RhombusSideLength
	to.ShiftToNearestCircle = from.ShiftToNearestCircle
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsPlantDiagramsNodeExpanded = from.IsPlantDiagramsNodeExpanded
}

type PlantDiagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	OriginX float64

	OriginY float64

	IsChecked bool
}

func (from *PlantDiagram) CopyBasicFields(to *PlantDiagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.IsChecked = from.IsChecked
}

type ReferenceRhombus_WOP struct {
	// insertion point

	Name string

	IsHidden bool
}

func (from *ReferenceRhombus) CopyBasicFields(to *ReferenceRhombus) {
	// insertion point
	to.Name = from.Name
	to.IsHidden = from.IsHidden
}

// end of insertion point
