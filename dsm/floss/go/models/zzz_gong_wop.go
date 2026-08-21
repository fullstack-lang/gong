// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Complexity_WOP struct {
	// insertion point

	Name string

	Strength float64

	ComputedPrefix string

	IsExpanded bool
}

func (from *Complexity) CopyBasicFields(to *Complexity) {
	// insertion point
	to.Name = from.Name
	to.Strength = from.Strength
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type DiagramFloss_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	IsChecked bool

	IsEditable_ bool

	IsShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64

	IsSystemsNodeExpanded bool
}

func (from *DiagramFloss) CopyBasicFields(to *DiagramFloss) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
	to.IsSystemsNodeExpanded = from.IsSystemsNodeExpanded
}

type Effort_WOP struct {
	// insertion point

	Name string

	Strength float64

	ComputedPrefix string

	IsExpanded bool
}

func (from *Effort) CopyBasicFields(to *Effort) {
	// insertion point
	to.Name = from.Name
	to.Strength = from.Strength
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Library_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	IsRootLibrary bool

	IsSubLibrariesNodeExpanded bool

	NbPixPerCharacter float64

	LogoSVGFile string

	IsSystemesNodeExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool

	IsExpandedTmp bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsRootLibrary = from.IsRootLibrary
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.IsSystemesNodeExpanded = from.IsSystemesNodeExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
	to.IsExpandedTmp = from.IsExpandedTmp
}

type Performance_WOP struct {
	// insertion point

	Name string

	Strength float64

	ComputedPrefix string

	IsExpanded bool
}

func (from *Performance) CopyBasicFields(to *Performance) {
	// insertion point
	to.Name = from.Name
	to.Strength = from.Strength
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type System_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	SVG_Path string

	InverseAppliedScaling float64

	IsSubSystemNodeExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool
}

func (from *System) CopyBasicFields(to *System) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.SVG_Path = from.SVG_Path
	to.InverseAppliedScaling = from.InverseAppliedScaling
	to.IsSubSystemNodeExpanded = from.IsSubSystemNodeExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
}

type SystemShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *SystemShape) CopyBasicFields(to *SystemShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

// end of insertion point
