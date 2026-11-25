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
	if s.DataFlow == nil || (s.DataFlow.EndTask == nil && s.DataFlow.EndParticipant == nil && s.DataFlow.EndDataStore == nil) {
		return nil
	}
	if s.DataFlow.EndTask != nil {
		return s.DataFlow.EndTask
	}
	if s.DataFlow.EndParticipant != nil {
		return s.DataFlow.EndParticipant
	}
	if s.DataFlow.EndDataStore != nil {
		return s.DataFlow.EndDataStore
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	if endTask, ok := abstractElement.(*Task); ok {
		s.DataFlow.EndTask = endTask
		return
	}
	if endExternalParticipant, ok := abstractElement.(*ExternalParticipant); ok {
		s.DataFlow.EndParticipant = endExternalParticipant
		return
	}
	if endDataStore, ok := abstractElement.(*DataStore); ok {
		s.DataFlow.EndDataStore = endDataStore
		return
	}
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || (s.DataFlow.StartTask == nil && s.DataFlow.StartParticipant == nil && s.DataFlow.StartDataStore == nil) {
		return nil
	}
	if s.DataFlow.StartTask != nil {
		return s.DataFlow.StartTask
	}
	if s.DataFlow.StartParticipant != nil {
		return s.DataFlow.StartParticipant
	}
	if s.DataFlow.StartDataStore != nil {
		return s.DataFlow.StartDataStore
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	if startTask, ok := abstractElement.(*Task); ok {
		s.DataFlow.StartTask = startTask
		return
	}
	if startExternalParticipant, ok := abstractElement.(*ExternalParticipant); ok {
		s.DataFlow.StartParticipant = startExternalParticipant
		return
	}
	if startDataStore, ok := abstractElement.(*DataStore); ok {
		s.DataFlow.StartDataStore = startDataStore
		return
	}
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
