// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AllocatedResourceShape_WOP struct {
	// insertion point

	Name string
}

func (from *AllocatedResourceShape) GongCopyBasicFields(to *AllocatedResourceShape) {
	// insertion point
	to.Name = from.Name
}

type AllocatedSystemShape_WOP struct {
	// insertion point

	Name string
}

func (from *AllocatedSystemShape) GongCopyBasicFields(to *AllocatedSystemShape) {
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

func (from *ControlFlow) GongCopyBasicFields(to *ControlFlow) {
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

func (from *ControlFlowShape) GongCopyBasicFields(to *ControlFlowShape) {
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

func (from *Data) GongCopyBasicFields(to *Data) {
	// insertion point
	*to = *from
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

func (from *DataFlow) GongCopyBasicFields(to *DataFlow) {
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

func (from *DataFlowShape) GongCopyBasicFields(to *DataFlowShape) {
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

func (from *DataShape) GongCopyBasicFields(to *DataShape) {
	// insertion point
	to.Name = from.Name
}

type DiagramLayerState_WOP struct {
	// insertion point

	Name string
}

func (from *DiagramLayerState) GongCopyBasicFields(to *DiagramLayerState) {
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

	IsWithDiscretePorts bool

	Width float64

	Height float64

	IsSystemsNodeExpanded bool

	IsPartsNodeExpanded bool

	IsExternalPartsNodeExpanded bool

	IsNotesNodeExpanded bool
}

func (from *DiagramStructure) GongCopyBasicFields(to *DiagramStructure) {
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
	to.IsWithDiscretePorts = from.IsWithDiscretePorts
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

func (from *ExternalPartShape) GongCopyBasicFields(to *ExternalPartShape) {
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

type LayerDefinition_WOP struct {
	// insertion point

	Name string
}

func (from *LayerDefinition) GongCopyBasicFields(to *LayerDefinition) {
	// insertion point
	to.Name = from.Name
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

func (from *Note) GongCopyBasicFields(to *Note) {
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

func (from *NotePartShape) GongCopyBasicFields(to *NotePartShape) {
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

func (from *NotePortShape) GongCopyBasicFields(to *NotePortShape) {
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

func (from *Part) GongCopyBasicFields(to *Part) {
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

func (from *PartAnchoredPath) GongCopyBasicFields(to *PartAnchoredPath) {
	// insertion point
	*to = *from
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

func (from *PartShape) GongCopyBasicFields(to *PartShape) {
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

func (from *Port) GongCopyBasicFields(to *Port) {
	// insertion point
	*to = *from
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

func (from *PortShape) GongCopyBasicFields(to *PortShape) {
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

func (from *Resource) GongCopyBasicFields(to *Resource) {
	// insertion point
	*to = *from
}

type SemanticTag_WOP struct {
	// insertion point

	Name string
}

func (from *SemanticTag) GongCopyBasicFields(to *SemanticTag) {
	// insertion point
	to.Name = from.Name
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

func (from *System) GongCopyBasicFields(to *System) {
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

func (from *SystemShape) GongCopyBasicFields(to *SystemShape) {
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
