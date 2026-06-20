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

type notePortKey struct {
	Note *Note
	Port *Port
}

type NotePortShape struct {
	Name string

	Note *Note

	Port *Port

	LinkShape
}

func (s *NotePortShape) GetAbstractEndElement() AbstractType {
	if s.Port == nil {
		return nil
	}
	return s.Port
}

func (s *NotePortShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Port = abstractElement.(*Port)
}

func (s *NotePortShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NotePortShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NotePortShape)(nil)
