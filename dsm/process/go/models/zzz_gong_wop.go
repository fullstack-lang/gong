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
}

type Library_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	NbPixPerCharacter float64

	LogoSVGFile string
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
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

type ProcessCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ProcessCompositionShape) CopyBasicFields(to *ProcessCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
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

