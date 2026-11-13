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

type noteDataFlowKey struct {
	Note     *Note
	DataFlow *DataFlow
}

type NoteDataFlowShape struct {
	Name string

	Note *Note

	DataFlow *DataFlow

	LinkShape
}

func (s *NoteDataFlowShape) GetAbstractEndElement() AbstractType {
	if s.DataFlow == nil {
		return nil
	}
	return s.DataFlow
}

func (s *NoteDataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.DataFlow = abstractElement.(*DataFlow)
}

func (s *NoteDataFlowShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteDataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteDataFlowShape)(nil)

type noteControlFlowKey struct {
	Note        *Note
	ControlFlow *ControlFlow
}

type NoteControlFlowShape struct {
	Name string

	Note *Note

	ControlFlow *ControlFlow

	LinkShape
}

func (s *NoteControlFlowShape) GetAbstractEndElement() AbstractType {
	if s.ControlFlow == nil {
		return nil
	}
	return s.ControlFlow
}

func (s *NoteControlFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.ControlFlow = abstractElement.(*ControlFlow)
}

func (s *NoteControlFlowShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteControlFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteControlFlowShape)(nil)
