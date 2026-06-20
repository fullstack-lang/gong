package models

type PartShape struct {
	Name        string
	Part *Part

	IsExpanded bool

	RectShape

	// Every part defaults to a weight of 1.0
	// if a user wants a part to be 50% wider, its weight becomes 1.5.
	WidthWeight float64
}

func (s *PartShape) GetAbstractElement() AbstractType {
	if s.Part == nil {
		return nil // otherwise returns s.Part
	}
	return s.Part
}

func (s *PartShape) SetAbstractElement(abstractElement AbstractType) {
	s.Part = abstractElement.(*Part)
}

var _ ConcreteType = (*PartShape)(nil)

type ExternalPartShape struct {
	Name        string
	Part *Part

	IsExpanded bool

	// this is the rect shape for the title
	RectShape

	// there is an additionnal rect for the below tail
	TailHeigth float64
}

func (s *ExternalPartShape) GetAbstractElement() AbstractType {
	if s.Part == nil {
		return nil // otherwise returns s.Part
	}
	return s.Part
}

func (s *ExternalPartShape) SetAbstractElement(abstractElement AbstractType) {
	s.Part = abstractElement.(*Part)
}

var _ ConcreteType = (*ExternalPartShape)(nil)
