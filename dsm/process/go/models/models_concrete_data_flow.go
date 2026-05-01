package models

type DataFlowShape struct {
	Name string

	DataFlow *DataFlow

	LinkShape
}

func (s *DataFlowShape) GetAbstractElement() AbstractType {
	// very important to return nil, otherwisse, it will be a typed nil and the check of nil will not work
	if s.DataFlow == nil {
		return nil
	}
	return s.DataFlow
}

func (s *DataFlowShape) SetAbstractElement(a AbstractType) {
	s.DataFlow = a.(*DataFlow)
}
func (s *DataFlowShape) GetAbstractEndElement() AbstractType {
	if s.DataFlow == nil || s.DataFlow.End == nil {
		return nil
	}
	return s.DataFlow.End
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.DataFlow.End = abstractElement.(*Task)
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || s.DataFlow.Start == nil {
		return nil
	}
	return s.DataFlow.Start
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.DataFlow.Start = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
