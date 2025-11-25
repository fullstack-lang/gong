package models

type StateMachine struct {
	Name string

	IsNodeExpanded bool

	States   []*State
	Diagrams []*Diagram

	InitialState *State
}
