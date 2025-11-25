package models

type Architecture struct {
	Name string

	StateMachines []*StateMachine
	Roles         []*Role

	NbPixPerCharacter float64
}
