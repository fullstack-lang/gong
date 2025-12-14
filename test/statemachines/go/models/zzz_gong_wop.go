// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Architecture_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64
}

func (from *Architecture) CopyBasicFields(to *Architecture) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
}

type Diagram_WOP struct {
	// insertion point

	Name string

	IsChecked bool

	IsExpanded bool

	IsEditable_ bool

	IsInRenameMode bool
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.IsChecked = from.IsChecked
	to.IsExpanded = from.IsExpanded
	to.IsEditable_ = from.IsEditable_
	to.IsInRenameMode = from.IsInRenameMode
}

type DoAction_WOP struct {
	// insertion point

	Name string

	Criticality Criticality
}

func (from *DoAction) CopyBasicFields(to *DoAction) {
	// insertion point
	to.Name = from.Name
	to.Criticality = from.Criticality
}

type Kill_WOP struct {
	// insertion point

	Name string
}

func (from *Kill) CopyBasicFields(to *Kill) {
	// insertion point
	to.Name = from.Name
}

type Message_WOP struct {
	// insertion point

	Name string

	IsSelected bool
}

func (from *Message) CopyBasicFields(to *Message) {
	// insertion point
	to.Name = from.Name
	to.IsSelected = from.IsSelected
}

type MessageType_WOP struct {
	// insertion point

	Name string

	Description string
}

func (from *MessageType) CopyBasicFields(to *MessageType) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
}

type Object_WOP struct {
	// insertion point

	Name string

	IsSelected bool

	Rank int

	DOF time.Time
}

func (from *Object) CopyBasicFields(to *Object) {
	// insertion point
	to.Name = from.Name
	to.IsSelected = from.IsSelected
	to.Rank = from.Rank
	to.DOF = from.DOF
}

type Role_WOP struct {
	// insertion point

	Name string

	Acronym string
}

func (from *Role) CopyBasicFields(to *Role) {
	// insertion point
	to.Name = from.Name
	to.Acronym = from.Acronym
}

type State_WOP struct {
	// insertion point

	Name string

	IsDecisionNode bool

	IsFictif bool

	IsEndState bool
}

func (from *State) CopyBasicFields(to *State) {
	// insertion point
	to.Name = from.Name
	to.IsDecisionNode = from.IsDecisionNode
	to.IsFictif = from.IsFictif
	to.IsEndState = from.IsEndState
}

type StateMachine_WOP struct {
	// insertion point

	Name string

	IsNodeExpanded bool
}

func (from *StateMachine) CopyBasicFields(to *StateMachine) {
	// insertion point
	to.Name = from.Name
	to.IsNodeExpanded = from.IsNodeExpanded
}

type StateShape_WOP struct {
	// insertion point

	Name string

	IsExpanded bool

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *StateShape) CopyBasicFields(to *StateShape) {
	// insertion point
	to.Name = from.Name
	to.IsExpanded = from.IsExpanded
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type Transition_WOP struct {
	// insertion point

	Name string
}

func (from *Transition) CopyBasicFields(to *Transition) {
	// insertion point
	to.Name = from.Name
}

type Transition_Shape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64
}

func (from *Transition_Shape) CopyBasicFields(to *Transition_Shape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
}

