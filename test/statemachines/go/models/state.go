package models

type State struct {
	//gong:width 600
	Name string

	Parent *State

	IsDecisionNode bool

	IsFictif bool

	IsEndState bool

	// When there are SubStates, the State is composite
	SubStates []*State

	// Diagrams where a state is present is exported
	// in the XL file
	Diagrams []*Diagram

	DoActions []*DoAction
}

func (state *State) IsComposite() bool {
	return len(state.SubStates) > 0
}

// DoAction action
type DoAction struct {
	Name        string
	Criticality Criticality
}

type Criticality string

const (
	DoActionCritical Criticality = "Critical"
	DoActionDefault  Criticality = "Default"
)
