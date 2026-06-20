package models

type ParticipantShape struct {
	Name        string
	Participant *Participant

	IsExpanded bool

	RectShape

	// Every participant defaults to a weight of 1.0
	// if a user wants a participant to be 50% wider, its weight becomes 1.5.
	WidthWeight float64
}

func (s *ParticipantShape) GetAbstractElement() AbstractType {
	if s.Participant == nil {
		return nil // otherwise returns s.Participant
	}
	return s.Participant
}

func (s *ParticipantShape) SetAbstractElement(abstractElement AbstractType) {
	s.Participant = abstractElement.(*Participant)
}

var _ ConcreteType = (*ParticipantShape)(nil)

type ExternalParticipantShape struct {
	Name        string
	Participant *Participant

	IsExpanded bool

	// this is the rect shape for the title
	RectShape

	// there is an additionnal rect for the below tail
	TailHeigth float64
}

func (s *ExternalParticipantShape) GetAbstractElement() AbstractType {
	if s.Participant == nil {
		return nil // otherwise returns s.Participant
	}
	return s.Participant
}

func (s *ExternalParticipantShape) SetAbstractElement(abstractElement AbstractType) {
	s.Participant = abstractElement.(*Participant)
}

var _ ConcreteType = (*ExternalParticipantShape)(nil)
