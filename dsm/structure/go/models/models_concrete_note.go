package models

type NoteShape struct {
	Name string

	Note *Note

	RectShape
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

type noteTaskKey struct {
	Note *Note
	Task *Task
}

type NoteTaskShape struct {
	Name string

	Note *Note

	Task *Task

	LinkShape
}

func (s *NoteTaskShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *NoteTaskShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *NoteTaskShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteTaskShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteTaskShape)(nil)
