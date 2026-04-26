package models

type ParticipantShape struct {
	Name        string
	Participant *Participant

	IsExpanded bool

	RectShape
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
