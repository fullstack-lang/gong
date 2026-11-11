package models

type PortShape struct {
	Name string

	Port *Port

	IsExpanded bool

	// ParentPartShape scopes this PortShape to the specific PartShape instance
	// in the diagram (enabling independent expansion of multiple Parts sharing the same System Type).
	ParentPartShape *PartShape

	// when a port shape is display the part shape is always
	// present
	owningPartShape *PartShape

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
