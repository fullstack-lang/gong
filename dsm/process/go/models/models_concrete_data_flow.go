package models

type DataFlowShape struct {
	Name string

	DataFlow *DataFlow

	LinkShape

	dataShapes []*DataShape
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
	if s.DataFlow == nil || s.DataFlow.EndTask == nil {
		return nil
	}
	return s.DataFlow.EndTask
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.DataFlow.EndTask = abstractElement.(*Task)
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || s.DataFlow.StartTask == nil {
		return nil
	}
	return s.DataFlow.StartTask
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.DataFlow.StartTask = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
