// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type CompareAnalysis_WOP struct {
	// insertion point

	Name string

	Mu float64

	Epsilon float64

	ComputedPrefix string

	IsExpanded bool
}

func (from *CompareAnalysis) GongCopyBasicFields(to *CompareAnalysis) {
	// insertion point
	to.Name = from.Name
	to.Mu = from.Mu
	to.Epsilon = from.Epsilon
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Complexity_WOP struct {
	// insertion point

	Name string

	Strength float64

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Complexity) GongCopyBasicFields(to *Complexity) {
	// insertion point
	*to = *from
}

type DiagramFlossEquation_WOP struct {
	// insertion point

	Name string

	Description string

	Scale float64

	FontSize FontSize

	ComputedPrefix string

	IsExpanded bool

	IsChecked bool

	IsEditable_ bool

	IsInDelta3ColumnsMode bool

	AreQuantitativeElementsVisible bool

	AreSubsystemsVisible bool

	AreCommonElementsHidden bool

	AreCPEArrowsVisible bool

	AreColumnTitlesVisible bool

	Width float64

	Height float64

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	IsNotesNodeExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool
}

func (from *DiagramFlossEquation) GongCopyBasicFields(to *DiagramFlossEquation) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.Scale = from.Scale
	to.FontSize = from.FontSize
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsInDelta3ColumnsMode = from.IsInDelta3ColumnsMode
	to.AreQuantitativeElementsVisible = from.AreQuantitativeElementsVisible
	to.AreSubsystemsVisible = from.AreSubsystemsVisible
	to.AreCommonElementsHidden = from.AreCommonElementsHidden
	to.AreCPEArrowsVisible = from.AreCPEArrowsVisible
	to.AreColumnTitlesVisible = from.AreColumnTitlesVisible
	to.Width = from.Width
	to.Height = from.Height
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
}

type Effort_WOP struct {
	// insertion point

	Name string

	Strength float64

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Effort) GongCopyBasicFields(to *Effort) {
	// insertion point
	*to = *from
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

	IsSystemsNodeExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool

	IsCompareAnalysisNodeExpanded bool

	IsNotesNodeExpanded bool

	IsExpandedTmp bool
}

func (from *Library) GongCopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsRootLibrary = from.IsRootLibrary
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.IsSystemsNodeExpanded = from.IsSystemsNodeExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
	to.IsCompareAnalysisNodeExpanded = from.IsCompareAnalysisNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsExpandedTmp = from.IsExpandedTmp
}

type Note_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool
}

func (from *Note) GongCopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
}

type NoteComplexityShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteComplexityShape) GongCopyBasicFields(to *NoteComplexityShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteEffortShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteEffortShape) GongCopyBasicFields(to *NoteEffortShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NotePerformanceShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NotePerformanceShape) GongCopyBasicFields(to *NotePerformanceShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *NoteShape) GongCopyBasicFields(to *NoteShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Performance_WOP struct {
	// insertion point

	Name string

	Strength float64

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Performance) GongCopyBasicFields(to *Performance) {
	// insertion point
	*to = *from
}

type System_WOP struct {
	// insertion point

	Name string

	Description string

	AreCPEsCompoundedFromSubSystems bool

	ComputedPrefix string

	IsExpanded bool

	SVG_Path string

	InverseAppliedScaling float64

	IsSubSystemNodeExpanded bool

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool
}

func (from *System) GongCopyBasicFields(to *System) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.AreCPEsCompoundedFromSubSystems = from.AreCPEsCompoundedFromSubSystems
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.SVG_Path = from.SVG_Path
	to.InverseAppliedScaling = from.InverseAppliedScaling
	to.IsSubSystemNodeExpanded = from.IsSubSystemNodeExpanded
	to.IsComplexitysNodeExpanded = from.IsComplexitysNodeExpanded
	to.IsPerformancesNodeExpanded = from.IsPerformancesNodeExpanded
	to.IsEffortsNodeExpanded = from.IsEffortsNodeExpanded
}

// end of insertion point
