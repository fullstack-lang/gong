package models

// ProcessShape
type ProcessShape struct {
	Name    string
	Process *Process

	IsExpanded bool

	RectShape
}

func (s *ProcessShape) GetAbstractElement() AbstractType {
	if s.Process == nil {
		return nil // otherwise returns s.Process returns (*Process, nil), not nil
	}
	return s.Process
}

func (s *ProcessShape) SetAbstractElement(abstractElement AbstractType) {
	s.Process = abstractElement.(*Process)
}

var _ ConcreteType = (*ProcessShape)(nil)

// A ProcessCompositionShape is the link between a Process
// and its parent Process
type ProcessCompositionShape struct {
	Name string

	Process *Process

	LinkShape
}

func (s *ProcessCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Process == nil {
		return nil
	}
	return s.Process
}

func (s *ProcessCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Process = abstractElement.(*Process)
}

func (s *ProcessCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Process == nil || s.Process.parentProcess == nil {
		return nil
	}
	return s.Process.parentProcess
}

func (s *ProcessCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element, because it is computed
	// elsewhere
	// s.Process.parentProcess = abstractElement.(*Process)
}

var _ AssociationConcreteType = (*ProcessCompositionShape)(nil)
