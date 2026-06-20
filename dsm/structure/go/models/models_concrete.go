package models

type PartShape struct {
	Name string

	Part *Part

	IsExpanded bool

	RectShape

	// Every part defaults to a weight of 1.0
	WidthWeight float64

	ConcreteTypeFields
}

func (s *PartShape) GetAbstractElement() AbstractType {
	if s.Part == nil {
		return nil
	}
	return s.Part
}

func (s *PartShape) SetAbstractElement(abstractElement AbstractType) {
	s.Part = abstractElement.(*Part)
}

var _ ConcreteType = (*PartShape)(nil)
var _ LayoutConcreteType = (*PartShape)(nil)

type SystemShape struct {
	Name   string
	System *System

	IsExpanded bool

	RectShape

	ConcreteTypeFields
}

func (s *SystemShape) GetAbstractElement() AbstractType {
	if s.System == nil {
		return nil
	}
	return s.System
}

func (s *SystemShape) SetAbstractElement(abstractElement AbstractType) {
	s.System = abstractElement.(*System)
}

var _ ConcreteType = (*SystemShape)(nil)
var _ LayoutConcreteType = (*SystemShape)(nil)

type LinkAssociationShape struct {
	Name string

	Link *Link

	LinkShape
}

func (s *LinkAssociationShape) GetAbstractElement() AbstractType {
	if s.Link == nil {
		return nil
	}
	return s.Link
}

func (s *LinkAssociationShape) SetAbstractElement(a AbstractType) {
	s.Link = a.(*Link)
}

func (s *LinkAssociationShape) GetAbstractEndElement() AbstractType {
	if s.Link == nil || s.Link.Target == nil {
		return nil
	}
	return s.Link.Target
}

func (s *LinkAssociationShape) SetAbstractEndElement(abstractElement AbstractType) {
	if target, ok := abstractElement.(*Part); ok {
		s.Link.Target = target
	}
}

func (s *LinkAssociationShape) GetAbstractStartElement() AbstractType {
	if s.Link == nil || s.Link.Source == nil {
		return nil
	}
	return s.Link.Source
}

func (s *LinkAssociationShape) SetAbstractStartElement(abstractElement AbstractType) {
	if source, ok := abstractElement.(*Part); ok {
		s.Link.Source = source
	}
}

var _ AssociationConcreteType = (*LinkAssociationShape)(nil)

var _ ConcreteType = (*LinkAssociationShape)(nil)
