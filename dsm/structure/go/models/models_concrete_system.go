package models

// SystemShape
type SystemShape struct {
	Name   string
	System *System

	IsExpanded bool

	RectShape
}

func (s *SystemShape) GetAbstractElement() AbstractType {
	if s.System == nil {
		return nil // otherwise returns s.System returns (*System, nil), not nil
	}
	return s.System
}

func (s *SystemShape) SetAbstractElement(abstractElement AbstractType) {
	s.System = abstractElement.(*System)
}

var _ ConcreteType = (*SystemShape)(nil)
