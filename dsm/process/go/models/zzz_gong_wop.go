// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type ControlFlow_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string
}

func (from *ControlFlow) CopyBasicFields(to *ControlFlow) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
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

	ComputedPrefix string
}

func (from *Data) CopyBasicFields(to *Data) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
}

type DataFlow_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsDatasNodeExpanded bool
}

func (from *DataFlow) CopyBasicFields(to *DataFlow) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsDatasNodeExpanded = from.IsDatasNodeExpanded
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

type DiagramProcess_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsChecked bool

	IsEditable_ bool

	IsShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64

	IsProcesssNodeExpanded bool

	IsParticipantsNodeExpanded bool

	IsExternalParticipantsNodeExpanded bool
}

func (from *DiagramProcess) CopyBasicFields(to *DiagramProcess) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
	to.IsProcesssNodeExpanded = from.IsProcesssNodeExpanded
	to.IsParticipantsNodeExpanded = from.IsParticipantsNodeExpanded
	to.IsExternalParticipantsNodeExpanded = from.IsExternalParticipantsNodeExpanded
}

type ExternalParticipantShape_WOP struct {
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

func (from *ExternalParticipantShape) CopyBasicFields(to *ExternalParticipantShape) {
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

	ComputedPrefix string

	IsSubLibrariesNodeExpanded bool

	NbPixPerCharacter float64

	LogoSVGFile string

	IsProcessesNodeExpanded bool

	IsDataFlowsNodeExpanded bool

	IsDatasNodeExpanded bool

	IsExpandedTmp bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.IsProcessesNodeExpanded = from.IsProcessesNodeExpanded
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
	to.IsDatasNodeExpanded = from.IsDatasNodeExpanded
	to.IsExpandedTmp = from.IsExpandedTmp
}

type Participant_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsTasksNodeExpanded bool

	IsControlFlowsNodeExpanded bool

	IsDataFlowsNodeExpanded bool
}

func (from *Participant) CopyBasicFields(to *Participant) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsTasksNodeExpanded = from.IsTasksNodeExpanded
	to.IsControlFlowsNodeExpanded = from.IsControlFlowsNodeExpanded
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
}

type ParticipantShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ParticipantShape) CopyBasicFields(to *ParticipantShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Process_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsSubProcessNodeExpanded bool

	IsDataFlowsNodeExpanded bool
}

func (from *Process) CopyBasicFields(to *Process) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsSubProcessNodeExpanded = from.IsSubProcessNodeExpanded
	to.IsDataFlowsNodeExpanded = from.IsDataFlowsNodeExpanded
}

type ProcessShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ProcessShape) CopyBasicFields(to *ProcessShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Task_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsStartTask bool

	IsEndTask bool
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsStartTask = from.IsStartTask
	to.IsEndTask = from.IsEndTask
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

