// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type DiagramStructure_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsChecked bool

	IsEditable_ bool

	IsShowPrefix bool

	Width float64

	Height float64

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	IsPartsNodeExpanded bool
}

func (from *DiagramStructure) CopyBasicFields(to *DiagramStructure) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.Width = from.Width
	to.Height = from.Height
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.IsPartsNodeExpanded = from.IsPartsNodeExpanded
}

type Library_WOP struct {
	// insertion point

	Name string

	IsSubLibrariesNodeExpanded bool

	NbPixPerCharacter float64

	LogoSVGFile string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsRootLibrary bool

	IsSystemsNodeExpanded bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsRootLibrary = from.IsRootLibrary
	to.IsSystemsNodeExpanded = from.IsSystemsNodeExpanded
}

type Link_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Link) CopyBasicFields(to *Link) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type LinkAssociationShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *LinkAssociationShape) CopyBasicFields(to *LinkAssociationShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Part_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Part) CopyBasicFields(to *Part) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type PartShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool

	WidthWeight float64

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection
}

func (from *PartShape) CopyBasicFields(to *PartShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
	to.WidthWeight = from.WidthWeight
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
}

type System_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsPartsNodeExpanded bool

	IsSubSystemsNodeExpanded bool

	IsLinksNodeExpanded bool

	IsDiagramStructuresNodeExpanded bool
}

func (from *System) CopyBasicFields(to *System) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsPartsNodeExpanded = from.IsPartsNodeExpanded
	to.IsSubSystemsNodeExpanded = from.IsSubSystemsNodeExpanded
	to.IsLinksNodeExpanded = from.IsLinksNodeExpanded
	to.IsDiagramStructuresNodeExpanded = from.IsDiagramStructuresNodeExpanded
}

// end of insertion point
