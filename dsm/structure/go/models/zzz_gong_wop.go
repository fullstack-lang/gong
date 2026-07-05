// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type AllocatedResourceShape_WOP struct {
	// insertion point

	Name string
}

func (from *AllocatedResourceShape) CopyBasicFields(to *AllocatedResourceShape) {
	// insertion point
	to.Name = from.Name
}

type AllocatedSystemShape_WOP struct {
	// insertion point

	Name string
}

func (from *AllocatedSystemShape) CopyBasicFields(to *AllocatedSystemShape) {
	// insertion point
	to.Name = from.Name
}

type ControlFlow_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *ControlFlow) CopyBasicFields(to *ControlFlow) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ControlFlowShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ControlFlowShape) CopyBasicFields(to *ControlFlowShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Data_WOP struct {
	// insertion point

	Name string

	Acronym string

	Description string

	ComputedPrefix string

	IsExpanded bool

	SVG_Path string

	InverseAppliedScaling float64
}

func (from *Data) CopyBasicFields(to *Data) {
	// insertion point
	to.Name = from.Name
	to.Acronym = from.Acronym
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.SVG_Path = from.SVG_Path
	to.InverseAppliedScaling = from.InverseAppliedScaling
}

type DataFlow_WOP struct {
	// insertion point

	Name string

	Description string

	Direction DataFlowDirection

	IsDatasNodeExpanded bool

	ComputedPrefix string

	IsExpanded bool

	Type DataFlowType
}

func (from *DataFlow) CopyBasicFields(to *DataFlow) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.Direction = from.Direction
	to.IsDatasNodeExpanded = from.IsDatasNodeExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.Type = from.Type
}

type DataFlowShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *DataFlowShape) CopyBasicFields(to *DataFlowShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type DataShape_WOP struct {
	// insertion point

	Name string
}

func (from *DataShape) CopyBasicFields(to *DataShape) {
	// insertion point
	to.Name = from.Name
}

type DiagramStructure_WOP struct {
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

	IsPartsNodeExpanded bool

	IsExternalPartsNodeExpanded bool

	IsNotesNodeExpanded bool
}

func (from *DiagramStructure) CopyBasicFields(to *DiagramStructure) {
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
	to.IsPartsNodeExpanded = from.IsPartsNodeExpanded
	to.IsExternalPartsNodeExpanded = from.IsExternalPartsNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
}

type ExternalPartShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool

	TailHeigth float64
}

func (from *ExternalPartShape) CopyBasicFields(to *ExternalPartShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
	to.TailHeigth = from.TailHeigth
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

	IsDataFlowsNodeExpanded bool

	IsDatasNodeExpanded bool

	IsResourcesNodeExpanded bool

	IsNotesNodeExpanded bool

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
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
	to.IsDatasNodeExpanded = from.IsDatasNodeExpanded
	to.IsResourcesNodeExpanded = from.IsResourcesNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsExpandedTmp = from.IsExpandedTmp
}

type Note_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	IsPartsNodeExpanded bool

	IsPortsNodeExpanded bool
}

func (from *Note) CopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsPartsNodeExpanded = from.IsPartsNodeExpanded
	to.IsPortsNodeExpanded = from.IsPortsNodeExpanded
}

type NotePartShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NotePartShape) CopyBasicFields(to *NotePartShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NotePortShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NotePortShape) CopyBasicFields(to *NotePortShape) {
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

func (from *NoteShape) CopyBasicFields(to *NoteShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Part_WOP struct {
	// insertion point

	Name string

	Description string

	IsPartNameNotSystemName bool

	IsControlFlowsNodeExpanded bool

	IsDataFlowsNodeExpanded bool

	ComputedPrefix string

	IsExpanded bool

	IsPortsNodeExpanded bool
}

func (from *Part) CopyBasicFields(to *Part) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsPartNameNotSystemName = from.IsPartNameNotSystemName
	to.IsControlFlowsNodeExpanded = from.IsControlFlowsNodeExpanded
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsPortsNodeExpanded = from.IsPortsNodeExpanded
}

type PartAnchoredPath_WOP struct {
	// insertion point

	Name string

	Definition string

	X_Offset float64

	Y_Offset float64

	RectAnchorType RectAnchorType

	ScalePropotionnally bool

	AppliedScaling float64

	Color string

	FillOpacity float64

	Stroke string

	StrokeOpacity float64

	StrokeWidth float64

	StrokeDashArray string

	StrokeDashArrayWhenSelected string

	Transform string
}

func (from *PartAnchoredPath) CopyBasicFields(to *PartAnchoredPath) {
	// insertion point
	to.Name = from.Name
	to.Definition = from.Definition
	to.X_Offset = from.X_Offset
	to.Y_Offset = from.Y_Offset
	to.RectAnchorType = from.RectAnchorType
	to.ScalePropotionnally = from.ScalePropotionnally
	to.AppliedScaling = from.AppliedScaling
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
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
}

type Port_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Port) CopyBasicFields(to *Port) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type PortShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *PortShape) CopyBasicFields(to *PortShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Resource_WOP struct {
	// insertion point

	Name string

	Acronym string

	Description string

	ComputedPrefix string

	IsExpanded bool

	SVG_Path string

	InverseAppliedScaling float64
}

func (from *Resource) CopyBasicFields(to *Resource) {
	// insertion point
	to.Name = from.Name
	to.Acronym = from.Acronym
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.SVG_Path = from.SVG_Path
	to.InverseAppliedScaling = from.InverseAppliedScaling
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

	IsDataFlowsNodeExpanded bool
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
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
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
