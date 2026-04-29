// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
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
}

type Library_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	NbPixPerCharacter float64

	LogoSVGFile string

	IsExpandedTmp bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.IsExpandedTmp = from.IsExpandedTmp
}

type Participant_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string
}

func (from *Participant) CopyBasicFields(to *Participant) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
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
}

func (from *Process) CopyBasicFields(to *Process) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsSubProcessNodeExpanded = from.IsSubProcessNodeExpanded
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
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
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

