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

	IsChecked bool

	IsEditable_ bool

	IsInRenameMode bool

	ShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	IsExpanded bool

	ComputedPrefix string

	IsPBSNodeExpanded bool

	IsWBSNodeExpanded bool
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsInRenameMode = from.IsInRenameMode
	to.ShowPrefix = from.ShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.IsExpanded = from.IsExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsWBSNodeExpanded = from.IsWBSNodeExpanded
}

type Product_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	ComputedPrefix string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool
}

func (from *Product) CopyBasicFields(to *Product) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.ComputedPrefix = from.ComputedPrefix
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
}

func (from *ProductCompositionShape) CopyBasicFields(to *ProductCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

type ProductShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *ProductShape) CopyBasicFields(to *ProductShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type Project_WOP struct {
	// insertion point

	Name string

	IsPBSNodeExpanded bool

	IsWBSNodeExpanded bool

	IsDiagramsNodeExpanded bool

	IsExpanded bool

	ComputedPrefix string
}

func (from *Project) CopyBasicFields(to *Project) {
	// insertion point
	to.Name = from.Name
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsWBSNodeExpanded = from.IsWBSNodeExpanded
	to.IsDiagramsNodeExpanded = from.IsDiagramsNodeExpanded
	to.IsExpanded = from.IsExpanded
	to.ComputedPrefix = from.ComputedPrefix
}

type Root_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64
}

func (from *Root) CopyBasicFields(to *Root) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
}

type Task_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	ComputedPrefix string

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInputsNodeExpanded = from.IsInputsNodeExpanded
	to.IsOutputsNodeExpanded = from.IsOutputsNodeExpanded
}

type TaskCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64
}

func (from *TaskCompositionShape) CopyBasicFields(to *TaskCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

type TaskInputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64
}

func (from *TaskInputShape) CopyBasicFields(to *TaskInputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

type TaskOutputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64
}

func (from *TaskOutputShape) CopyBasicFields(to *TaskOutputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

type TaskShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *TaskShape) CopyBasicFields(to *TaskShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

