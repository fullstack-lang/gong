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

// NoteShape
type NoteShape struct {
	Name string
	Note *Note

	isExpanded bool
	ConcreteTypeFields

	RectShape
}

func (s *NoteShape) GetAbstractElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteShape) SetAbstractElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

type NoteStateShape struct {
	Name string

	Note  *Note
	State *State

	LinkShape
}

type NoteTransitionShape struct {
	Name string

	Note       *Note
	Transition *Transition

	LinkShape
}
