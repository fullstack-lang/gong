package models

type DataShape struct {
	Name string

	Data *Data
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
