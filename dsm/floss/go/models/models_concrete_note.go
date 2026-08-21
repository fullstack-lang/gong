package models

type NoteShape struct {
	Name string

	Note *Note

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

var _ ConcreteType = (*NoteShape)(nil)

type noteComplexityKey struct {
	Note       *Note
	Complexity *Complexity
}

type NoteComplexityShape struct {
	Name string

	Note       *Note
	Complexity *Complexity

	LinkShape
}

func (s *NoteComplexityShape) GetAbstractEndElement() AbstractType {
	if s.Complexity == nil {
		return nil
	}
	return s.Complexity
}

func (s *NoteComplexityShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Complexity = abstractElement.(*Complexity)
}

func (s *NoteComplexityShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteComplexityShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteComplexityShape)(nil)

type notePerformanceKey struct {
	Note        *Note
	Performance *Performance
}

type NotePerformanceShape struct {
	Name string

	Note        *Note
	Performance *Performance

	LinkShape
}

func (s *NotePerformanceShape) GetAbstractEndElement() AbstractType {
	if s.Performance == nil {
		return nil
	}
	return s.Performance
}

func (s *NotePerformanceShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Performance = abstractElement.(*Performance)
}

func (s *NotePerformanceShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NotePerformanceShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NotePerformanceShape)(nil)

type noteEffortKey struct {
	Note   *Note
	Effort *Effort
}

type NoteEffortShape struct {
	Name string

	Note   *Note
	Effort *Effort

	LinkShape
}

func (s *NoteEffortShape) GetAbstractEndElement() AbstractType {
	if s.Effort == nil {
		return nil
	}
	return s.Effort
}

func (s *NoteEffortShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Effort = abstractElement.(*Effort)
}

func (s *NoteEffortShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteEffortShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteEffortShape)(nil)
