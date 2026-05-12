package models

type NoteShape struct {
	Name string

	Note *Note
}

func (s *NoteShape) GetAbstractElement() AbstractType {
	if s.Note == nil {
		return nil // otherwise returns s.Note
	}
	return s.Note
}

func (s *NoteShape) SetAbstractElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ ConcreteType = (*NoteShape)(nil)
