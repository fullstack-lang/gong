// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Action_WOP struct {
	// insertion point

	Name string

	Criticality Criticality
}

func (from *Action) CopyBasicFields(to *Action) {
	// insertion point
	to.Name = from.Name
	to.Criticality = from.Criticality
}

type Activities_WOP struct {
	// insertion point

	Name string

	Criticality Criticality
}

func (from *Activities) CopyBasicFields(to *Activities) {
	// insertion point
	to.Name = from.Name
	to.Criticality = from.Criticality
}

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

	IsStatesNodeExpanded bool

	IsNotesNodeExpanded bool
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.IsChecked = from.IsChecked
	to.IsExpanded = from.IsExpanded
	to.IsEditable_ = from.IsEditable_
	to.IsStatesNodeExpanded = from.IsStatesNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
}

type Guard_WOP struct {
	// insertion point

	Name string
}

func (from *Guard) CopyBasicFields(to *Guard) {
	// insertion point
	to.Name = from.Name
}

type Kill_WOP struct {
	// insertion point

	Name string
}

func (from *Kill) CopyBasicFields(to *Kill) {
	// insertion point
	to.Name = from.Name
}

type Library_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64

	LogoSVGFile string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsRootLibrary bool

	IsStateMachinesNodeExpanded bool

	IsNotesNodeExpanded bool

	IsSubLibrariesNodeExpanded bool

	IsExpandedTmp bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsRootLibrary = from.IsRootLibrary
	to.IsStateMachinesNodeExpanded = from.IsStateMachinesNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.IsExpandedTmp = from.IsExpandedTmp
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

type NoteShape_WOP struct {
	// insertion point

	Name string

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *NoteShape) CopyBasicFields(to *NoteShape) {
	// insertion point
	to.Name = from.Name
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type NoteStateShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteStateShape) CopyBasicFields(to *NoteStateShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteTransitionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteTransitionShape) CopyBasicFields(to *NoteTransitionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
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

	IsFictious bool

	IsEndState bool
}

func (from *State) CopyBasicFields(to *State) {
	// insertion point
	to.Name = from.Name
	to.IsDecisionNode = from.IsDecisionNode
	to.IsFictious = from.IsFictious
	to.IsEndState = from.IsEndState
}

type StateMachine_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *StateMachine) CopyBasicFields(to *StateMachine) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type StateShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *StateShape) CopyBasicFields(to *StateShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
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

	IsHidden bool
}

func (from *Transition_Shape) CopyBasicFields(to *Transition_Shape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

// end of insertion point
