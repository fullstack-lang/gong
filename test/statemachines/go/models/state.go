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

	Entry      *Action
	Activities []*Activities
	Exit       *Action
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
