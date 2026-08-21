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

// ComplexityShape
type ComplexityShape struct {
	Name       string
	Complexity *Complexity

	IsExpanded bool

	RectShape
}

func (s *ComplexityShape) GetAbstractElement() AbstractType {
	if s.Complexity == nil {
		return nil
	}
	return s.Complexity
}

func (s *ComplexityShape) SetAbstractElement(abstractElement AbstractType) {
	s.Complexity = abstractElement.(*Complexity)
}

var _ ConcreteType = (*ComplexityShape)(nil)

// PerformanceShape
type PerformanceShape struct {
	Name        string
	Performance *Performance

	IsExpanded bool

	RectShape
}

func (s *PerformanceShape) GetAbstractElement() AbstractType {
	if s.Performance == nil {
		return nil
	}
	return s.Performance
}

func (s *PerformanceShape) SetAbstractElement(abstractElement AbstractType) {
	s.Performance = abstractElement.(*Performance)
}

var _ ConcreteType = (*PerformanceShape)(nil)

// EffortShape
type EffortShape struct {
	Name   string
	Effort *Effort

	IsExpanded bool

	RectShape
}

func (s *EffortShape) GetAbstractElement() AbstractType {
	if s.Effort == nil {
		return nil
	}
	return s.Effort
}

func (s *EffortShape) SetAbstractElement(abstractElement AbstractType) {
	s.Effort = abstractElement.(*Effort)
}

var _ ConcreteType = (*EffortShape)(nil)

