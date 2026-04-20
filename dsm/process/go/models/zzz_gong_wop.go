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

	IsShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsInRenameMode = from.IsInRenameMode
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
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

type Process_WOP struct {
	// insertion point

	Name string
}

func (from *Process) CopyBasicFields(to *Process) {
	// insertion point
	to.Name = from.Name
}

