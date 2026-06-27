package models

import (
	"time"
)

type Architecture struct {
	Name string

	StateMachines []*StateMachine
	Roles         []*Role

	NbPixPerCharacter float64
}

type Diagram struct {
	Name           string
	IsChecked      bool
	IsExpanded     bool
	IsEditable_    bool
	isInRenameMode bool

	State_Shapes              []*StateShape
	StatesWhoseNodeIsExpanded []*State

	Transition_Shapes []*Transition_Shape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

type Message struct {
	Name       string
	IsSelected bool

	MessageType *MessageType

	OriginTransition *Transition
}

type MessageType struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	Description string
}

type Object struct {
	Name       string
	State      *State
	IsSelected bool
	Rank       int

	DOF time.Time

	Messages []*Message
}

type Role struct {
	Name string

	// "U" for Unit, "S" for supervisor
	Acronym string

	// Role that inherit the access right of this Role
	RolesWithSamePermissions []*Role
}

type State struct {
	//gong:width 600
	Name string

	Parent *State

	IsDecisionNode bool

	IsFictious bool

	IsEndState bool

	// When there are SubStates, the State is composite
	SubStates []*State

	// Diagrams where a state is present is exported
	// in the XL file
	Diagrams []*Diagram

	Entry      *Action
	Activities []*Activities
	Exit       *Action

	// nodes can be edited
	isInRenameMode bool
}

func (state *State) IsComposite() bool {
	return len(state.SubStates) > 0
}

// Activities action
type Activities struct {
	Name        string
	Criticality Criticality
}

type Criticality string

const (
	CriticalityCritical Criticality = "Critical"
	CriticalityDefault  Criticality = "Default"
)

type Action struct {
	Name        string
	Criticality Criticality
}

type StateMachine struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	States   []*State
	Diagrams []*Diagram

	InitialState *State
}

var _ AbstractType = (*StateMachine)(nil)


// Transition decribes authorized between states
type Transition struct {
	//gong:width 600
	Name string

	//gong:width 600
	Start *State

	//gong:width 600
	End *State

	// For Role Base Access Control
	RolesWithPermissions []*Role

	GeneratedMessages []*MessageType

	Guard *Guard

	// Diagrams where the transition is present
	// it goes against the normalization principle but
	// it is needed in the XL export
	Diagrams []*Diagram

	isInRenameMode bool
}

type Guard struct {
	Name string
}

func (transition *Transition) performTransition(stage *Stage) {

	objectSet := *GetGongstructInstancesSet[Object](stage)
	for object := range objectSet {
		if object.IsSelected {

			object.State = transition.End

			var messageTypeGenere *MessageType
			for _, messageTypeGenere_ := range transition.GeneratedMessages {
				messageTypeGenere = messageTypeGenere_
			}

			if messageTypeGenere != nil {
				message := new(Message).Stage(stage)

				message.Name = time.Now().Format("15:04:05") + " " + transition.Name + " -> " + messageTypeGenere.Name
				message.MessageType = messageTypeGenere
				message.OriginTransition = transition
				object.Messages = append(object.Messages, message)

				// set new message to selected
				for message_ := range *GetGongstructInstancesSet[Message](stage) {
					message_.IsSelected = false
				}
				message.IsSelected = true
			}

			stage.Commit()
		}
	}
}
