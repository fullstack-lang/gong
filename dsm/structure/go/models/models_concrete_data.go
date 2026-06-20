package models

// A DataShape is the concrete for a Data that is carried by a DataFlow.
type DataShape struct {
	Name string

	Data     *Data
	DataFlow *DataFlow
}

func (s *DataShape) GetAbstractElement() AbstractType {
	if s.Data == nil {
		return nil // otherwise returns s.Data
	}
	return s.Data
}

func (s *DataShape) SetAbstractElement(abstractElement AbstractType) {
	s.Data = abstractElement.(*Data)
}

var _ ConcreteType = (*DataShape)(nil)
