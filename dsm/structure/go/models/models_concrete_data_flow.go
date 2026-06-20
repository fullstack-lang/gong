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
	if s.DataFlow == nil || (s.DataFlow.EndPort == nil && s.DataFlow.EndExternalPart == nil) {
		return nil
	}
	if s.DataFlow.EndPort != nil {
		return s.DataFlow.EndPort
	}
	if s.DataFlow.EndExternalPart != nil {
		return s.DataFlow.EndExternalPart
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	if endPort, ok := abstractElement.(*Port); ok {
		s.DataFlow.EndPort = endPort
		return
	}
	if endExternalPart, ok := abstractElement.(*Part); ok {
		s.DataFlow.EndExternalPart = endExternalPart
		return
	}
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || (s.DataFlow.StartPort == nil && s.DataFlow.StartExternalPart == nil) {
		return nil
	}
	if s.DataFlow.StartPort != nil {
		return s.DataFlow.StartPort
	}
	if s.DataFlow.StartExternalPart != nil {
		return s.DataFlow.StartExternalPart
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	if startPort, ok := abstractElement.(*Port); ok {
		s.DataFlow.StartPort = startPort
		return
	}
	if startExternalPart, ok := abstractElement.(*Part); ok {
		s.DataFlow.StartExternalPart = startExternalPart
		return
	}
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
