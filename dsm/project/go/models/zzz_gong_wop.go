// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Diagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool

	IsChecked bool

	IsEditable_ bool

	ShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64

	IsPBSNodeExpanded bool

	IsWBSNodeExpanded bool

	IsNotesNodeExpanded bool

	IsResourcesNodeExpanded bool
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.ShowPrefix = from.ShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsWBSNodeExpanded = from.IsWBSNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsResourcesNodeExpanded = from.IsResourcesNodeExpanded
}

type Library_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool

	NbPixPerCharacter float64

	LogoSVGFile string
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
}

type Note_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool
}

func (from *Note) CopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
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

type NoteResourceShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteResourceShape) CopyBasicFields(to *NoteResourceShape) {
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

type Product_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool

	Description string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool
}

func (from *Product) CopyBasicFields(to *Product) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.Description = from.Description
	to.IsProducersNodeExpanded = from.IsProducersNodeExpanded
	to.IsConsumersNodeExpanded = from.IsConsumersNodeExpanded
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

type Resource_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool

	Description string
}

func (from *Resource) CopyBasicFields(to *Resource) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.Description = from.Description
}

type ResourceCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ResourceCompositionShape) CopyBasicFields(to *ResourceCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ResourceShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ResourceShape) CopyBasicFields(to *ResourceShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type ResourceTaskShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ResourceTaskShape) CopyBasicFields(to *ResourceTaskShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Task_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsInRenameMode bool

	IsExpanded bool

	Start time.Time

	End time.Time

	Description string

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool

	IsWithCompletion bool

	Completion CompletionEnum
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.Start = from.Start
	to.End = from.End
	to.Description = from.Description
	to.IsInputsNodeExpanded = from.IsInputsNodeExpanded
	to.IsOutputsNodeExpanded = from.IsOutputsNodeExpanded
	to.IsWithCompletion = from.IsWithCompletion
	to.Completion = from.Completion
}

type TaskCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskCompositionShape) CopyBasicFields(to *TaskCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskInputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskInputShape) CopyBasicFields(to *TaskInputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskOutputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskOutputShape) CopyBasicFields(to *TaskOutputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *TaskShape) CopyBasicFields(to *TaskShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

