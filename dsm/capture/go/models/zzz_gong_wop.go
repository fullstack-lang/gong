// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type AnalysisNeed_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *AnalysisNeed) CopyBasicFields(to *AnalysisNeed) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type Concept_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Concept) CopyBasicFields(to *Concept) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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

func (from *ConceptShape) CopyBasicFields(to *ConceptShape) {
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

	LayoutDirection LayoutDirection

	Description string

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool

	IsWithCompletion bool

	Completion CompletionEnum
}

func (from *Concern) CopyBasicFields(to *Concern) {
	// insertion point
	to.Name = from.Name
	to.IDAirbus = from.IDAirbus
	to.Priority = from.Priority
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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

func (from *ConcernCompositionShape) CopyBasicFields(to *ConcernCompositionShape) {
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

func (from *ConcernInputShape) CopyBasicFields(to *ConcernInputShape) {
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

func (from *ConcernOutputShape) CopyBasicFields(to *ConcernOutputShape) {
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

func (from *ConcernShape) CopyBasicFields(to *ConcernShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Deliverable_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	Description string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool
}

func (from *Deliverable) CopyBasicFields(to *Deliverable) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.Description = from.Description
	to.IsProducersNodeExpanded = from.IsProducersNodeExpanded
	to.IsConsumersNodeExpanded = from.IsConsumersNodeExpanded
}

type Diagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

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
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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
}

type Library_WOP struct {
	// insertion point

	Name string

	IsRootLibrary bool

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	NbPixPerCharacter float64
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.IsRootLibrary = from.IsRootLibrary
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.NbPixPerCharacter = from.NbPixPerCharacter
}

type Note_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Note) CopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type NoteProductShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteProductShape) CopyBasicFields(to *NoteProductShape) {
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

func (from *NoteShape) CopyBasicFields(to *NoteShape) {
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

func (from *NoteStakeholderShape) CopyBasicFields(to *NoteStakeholderShape) {
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

func (from *NoteTaskShape) CopyBasicFields(to *NoteTaskShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ProductCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ProductCompositionShape) CopyBasicFields(to *ProductCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ProductShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ProductShape) CopyBasicFields(to *ProductShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Requirement_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Requirement) CopyBasicFields(to *Requirement) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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

func (from *RequirementShape) CopyBasicFields(to *RequirementShape) {
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

	LayoutDirection LayoutDirection

	Description string
}

func (from *Stakeholder) CopyBasicFields(to *Stakeholder) {
	// insertion point
	to.Name = from.Name
	to.IDAirbus = from.IDAirbus
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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

func (from *StakeholderCompositionShape) CopyBasicFields(to *StakeholderCompositionShape) {
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

func (from *StakeholderConcernShape) CopyBasicFields(to *StakeholderConcernShape) {
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

func (from *StakeholderShape) CopyBasicFields(to *StakeholderShape) {
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

func (from *SupportLevel) CopyBasicFields(to *SupportLevel) {
	// insertion point
	to.Name = from.Name
}

type Tool_WOP struct {
	// insertion point

	Name string
}

func (from *Tool) CopyBasicFields(to *Tool) {
	// insertion point
	to.Name = from.Name
}

// end of insertion point
