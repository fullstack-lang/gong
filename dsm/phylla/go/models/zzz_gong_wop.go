// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
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

	InsideAngle float64

	ShiftToNearestCircle int

	SideLength float64

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
	to.InsideAngle = from.InsideAngle
	to.ShiftToNearestCircle = from.ShiftToNearestCircle
	to.SideLength = from.SideLength
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsPlantDiagramsNodeExpanded = from.IsPlantDiagramsNodeExpanded
}

type PlantDiagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *PlantDiagram) CopyBasicFields(to *PlantDiagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

// end of insertion point
