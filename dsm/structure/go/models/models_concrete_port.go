package models

type PortShape struct {
	Name string

	Port *Port

	IsExpanded bool

	RectShape
}

func (s *PortShape) GetAbstractElement() AbstractType {
	if s.Port == nil {
		return nil // otherwise returns s.Port
	}
	return s.Port
}

func (s *PortShape) SetAbstractElement(abstractElement AbstractType) {
	s.Port = abstractElement.(*Port)
}

var _ ConcreteType = (*PortShape)(nil)
