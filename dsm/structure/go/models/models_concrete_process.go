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
