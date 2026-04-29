package models

type ControlFlowShape struct {
	Name string

	ControlFlow *ControlFlow

	LinkShape
}

func (s *ControlFlowShape) GetAbstractElement() AbstractType {
	return s.ControlFlow
}

func (s *ControlFlowShape) SetAbstractElement(a AbstractType) {
	s.ControlFlow = a.(*ControlFlow)
}
func (s *ControlFlowShape) GetAbstractEndElement() AbstractType {
	if s.ControlFlow == nil || s.ControlFlow.End == nil {
		return nil
	}
	return s.ControlFlow.End
}

func (s *ControlFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.ControlFlow.End = abstractElement.(*Task)
}

func (s *ControlFlowShape) GetAbstractStartElement() AbstractType {
	if s.ControlFlow == nil || s.ControlFlow.Start == nil {
		return nil
	}
	return s.ControlFlow.Start
}

func (s *ControlFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.ControlFlow.Start = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*ControlFlowShape)(nil)
