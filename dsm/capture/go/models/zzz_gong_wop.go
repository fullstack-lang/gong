// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AnalysisNeed_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *AnalysisNeed) GongCopyBasicFields(to *AnalysisNeed) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Concept_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Concept) GongCopyBasicFields(to *Concept) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ConceptShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ConceptShape) GongCopyBasicFields(to *ConceptShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Concern_WOP struct {
	// insertion point

	Name string

	IDAirbus string

	Priority Priority

	ComputedPrefix string

	IsExpanded bool

	Description string

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool

	IsWithCompletion bool

	Completion CompletionEnum
}

func (from *Concern) GongCopyBasicFields(to *Concern) {
	// insertion point
	to.Name = from.Name
	to.IDAirbus = from.IDAirbus
	to.Priority = from.Priority
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.Description = from.Description
	to.IsInputsNodeExpanded = from.IsInputsNodeExpanded
	to.IsOutputsNodeExpanded = from.IsOutputsNodeExpanded
	to.IsWithCompletion = from.IsWithCompletion
	to.Completion = from.Completion
}

type ConcernCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ConcernCompositionShape) GongCopyBasicFields(to *ConcernCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ConcernInputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ConcernInputShape) GongCopyBasicFields(to *ConcernInputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ConcernOutputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ConcernOutputShape) GongCopyBasicFields(to *ConcernOutputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ConcernShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ConcernShape) GongCopyBasicFields(to *ConcernShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type ControlPointShape_WOP struct {
	// insertion point

	Name string

	X_Relative float64

	Y_Relative float64

	IsStartShapeTheClosestShape bool
}

func (from *ControlPointShape) GongCopyBasicFields(to *ControlPointShape) {
	// insertion point
	to.Name = from.Name
	to.X_Relative = from.X_Relative
	to.Y_Relative = from.Y_Relative
	to.IsStartShapeTheClosestShape = from.IsStartShapeTheClosestShape
}

type Deliverable_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	Description string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool
}

func (from *Deliverable) GongCopyBasicFields(to *Deliverable) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.Description = from.Description
	to.IsProducersNodeExpanded = from.IsProducersNodeExpanded
	to.IsConsumersNodeExpanded = from.IsConsumersNodeExpanded
}

type DeliverableCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *DeliverableCompositionShape) GongCopyBasicFields(to *DeliverableCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type DeliverableConceptShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *DeliverableConceptShape) GongCopyBasicFields(to *DeliverableConceptShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type DeliverableShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *DeliverableShape) GongCopyBasicFields(to *DeliverableShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Diagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	IsChecked bool

	IsEditable_ bool

	ShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64

	IsRequirementsNodeExpanded bool

	IsConceptsNodeExpanded bool

	IsPBSNodeExpanded bool

	IsConcernsNodeExpanded bool

	IsNotesNodeExpanded bool

	IsStakeholdersNodeExpanded bool

	IsDiagramsNodeExpanded bool
}

func (from *Diagram) GongCopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.ShowPrefix = from.ShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
	to.IsRequirementsNodeExpanded = from.IsRequirementsNodeExpanded
	to.IsConceptsNodeExpanded = from.IsConceptsNodeExpanded
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsConcernsNodeExpanded = from.IsConcernsNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsStakeholdersNodeExpanded = from.IsStakeholdersNodeExpanded
	to.IsDiagramsNodeExpanded = from.IsDiagramsNodeExpanded
}

type DiagramShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *DiagramShape) GongCopyBasicFields(to *DiagramShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Library_WOP struct {
	// insertion point

	Name string

	IsRootLibrary bool

	ComputedPrefix string

	IsExpanded bool

	NbPixPerCharacter float64
}

func (from *Library) GongCopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.IsRootLibrary = from.IsRootLibrary
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
}

type Note_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Note) GongCopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type NoteDeliverableShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteDeliverableShape) GongCopyBasicFields(to *NoteDeliverableShape) {
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

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *NoteShape) GongCopyBasicFields(to *NoteShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type NoteStakeholderShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteStakeholderShape) GongCopyBasicFields(to *NoteStakeholderShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteTaskShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteTaskShape) GongCopyBasicFields(to *NoteTaskShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Requirement_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Requirement) GongCopyBasicFields(to *Requirement) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type RequirementShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *RequirementShape) GongCopyBasicFields(to *RequirementShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Stakeholder_WOP struct {
	// insertion point

	Name string

	IDAirbus string

	ComputedPrefix string

	IsExpanded bool

	Description string
}

func (from *Stakeholder) GongCopyBasicFields(to *Stakeholder) {
	// insertion point
	to.Name = from.Name
	to.IDAirbus = from.IDAirbus
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.Description = from.Description
}

type StakeholderCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *StakeholderCompositionShape) GongCopyBasicFields(to *StakeholderCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type StakeholderConcernShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *StakeholderConcernShape) GongCopyBasicFields(to *StakeholderConcernShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type StakeholderShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *StakeholderShape) GongCopyBasicFields(to *StakeholderShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type SupportLevel_WOP struct {
	// insertion point

	Name string
}

func (from *SupportLevel) GongCopyBasicFields(to *SupportLevel) {
	// insertion point
	to.Name = from.Name
}

type Tool_WOP struct {
	// insertion point

	Name string
}

func (from *Tool) GongCopyBasicFields(to *Tool) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
