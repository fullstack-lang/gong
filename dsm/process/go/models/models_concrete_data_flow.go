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
	if s.DataFlow == nil || (s.DataFlow.EndTask == nil && s.DataFlow.EndExternalParticipant == nil) {
		return nil
	}
	if s.DataFlow.EndTask != nil {
		return s.DataFlow.EndTask
	}
	if s.DataFlow.EndExternalParticipant != nil {
		return s.DataFlow.EndExternalParticipant
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	if endTask, ok := abstractElement.(*Task); ok {
		s.DataFlow.EndTask = endTask
		return
	}
	if endExternalParticipant, ok := abstractElement.(*Participant); ok {
		s.DataFlow.EndExternalParticipant = endExternalParticipant
		return
	}
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || (s.DataFlow.StartTask == nil && s.DataFlow.StartExternalParticipant == nil) {
		return nil
	}
	if s.DataFlow.StartTask != nil {
		return s.DataFlow.StartTask
	}
	if s.DataFlow.StartExternalParticipant != nil {
		return s.DataFlow.StartExternalParticipant
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	if startTask, ok := abstractElement.(*Task); ok {
		s.DataFlow.StartTask = startTask
		return
	}
	if startExternalParticipant, ok := abstractElement.(*Participant); ok {
		s.DataFlow.StartExternalParticipant = startExternalParticipant
		return
	}
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
