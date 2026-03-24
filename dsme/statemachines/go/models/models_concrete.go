package models

type StateShape struct {
	Name  string
	State *State

	RectShape
}

func newStateShapeToDiagram(state *State, diagram *Diagram) *StateShape {
	stateShape := new(StateShape)
	stateShape.State = state
	stateShape.Name = state.GetName() + "-" + diagram.GetName()
	stateShape.Height = 80
	stateShape.Width = 200
	stateShape.X = 100
	stateShape.Y = 100
	diagram.State_Shapes = append(diagram.State_Shapes, stateShape)

	return stateShape
}

type Transition_Shape struct {
	Name string

	Transition *Transition

	LinkShape
}
