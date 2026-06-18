package models

// DeliverableShape
type DeliverableShape struct {
	Name    string
	Deliverable *Deliverable

	IsExpanded bool

	RectShape
}

func (s *DeliverableShape) GetAbstractElement() AbstractType {
	if s.Deliverable == nil {
		return nil
	}
	return s.Deliverable
}

func (s *DeliverableShape) SetAbstractElement(abstractElement AbstractType) {
	s.Deliverable = abstractElement.(*Deliverable)
}

var _ ConcreteType = (*DeliverableShape)(nil)

// A DeliverableCompositionShape is the link between a deliverable
// and its parent deliverable
type DeliverableCompositionShape struct {
	Name string

	Deliverable *Deliverable

	LinkShape
}

func (s *DeliverableCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Deliverable == nil {
		return nil
	}
	return s.Deliverable
}

func (s *DeliverableCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Deliverable = abstractElement.(*Deliverable)
}

func (s *DeliverableCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Deliverable == nil || s.Deliverable.parentDeliverable == nil {
		return nil
	}
	return s.Deliverable.parentDeliverable
}

func (s *DeliverableCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element, because it is computed
	// elsewhere
	// s.Deliverable.parentDeliverable = abstractElement.(*Deliverable)
}

var _ AssociationConcreteType = (*DeliverableCompositionShape)(nil)

// ConcernShape is for both Task
type ConcernShape struct {
	Name    string
	Concern *Concern

	IsExpanded bool

	RectShape
}

func (s *ConcernShape) GetAbstractElement() AbstractType {
	if s.Concern == nil {
		return nil // Explicitly return interface nil, otherwise returns (*Task, nil)
	}
	return s.Concern
}

func (s *ConcernShape) SetAbstractElement(abstractElement AbstractType) {
	s.Concern = abstractElement.(*Concern)
}

var _ ConcreteType = (*ConcernShape)(nil)

// A ConcernCompositionShape is the link between a task
// and its parent task
type ConcernCompositionShape struct {
	Name string

	Concern *Concern

	LinkShape
}

func (s *ConcernCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Concern == nil {
		return nil
	}
	return s.Concern
}

func (s *ConcernCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Concern = abstractElement.(*Concern)
}

func (s *ConcernCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Concern == nil || s.Concern.parentConcern == nil {
		return nil
	}
	return s.Concern.parentConcern
}

func (s *ConcernCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element because it is computed
	// elsewhere
	// s.Task.parentTask = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*ConcernCompositionShape)(nil)

// A concernDeliverableKey allows mapping of [ConcernInputShape] within a diagram
type concernDeliverableKey struct {
	Concern *Concern
	Deliverable *Deliverable
}

type noteDeliverableKey struct {
	Note    *Note
	Deliverable *Deliverable
}

type noteTaskKey struct {
	Note *Note
	Task *Concern
}

type noteResourceKey struct {
	Note        *Note
	Stakeholder *Stakeholder
}

type stakeholderConcernKey struct {
	Stakeholder *Stakeholder
	Concern     *Concern
}

type ConcernInputShape struct {
	Name string

	Deliverable *Deliverable // input deliverable

	Concern *Concern

	LinkShape
}

func (s *ConcernInputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Concern = abstractElement.(*Concern)
}

func (s *ConcernInputShape) GetAbstractEndElement() AbstractType {
	if s.Concern == nil {
		return nil
	}
	return s.Concern
}

func (s *ConcernInputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Deliverable = abstractElement.(*Deliverable)
}

func (s *ConcernInputShape) GetAbstractStartElement() AbstractType {
	if s.Deliverable == nil {
		return nil
	}
	return s.Deliverable
}

var _ AssociationConcreteType = (*ConcernInputShape)(nil)

type ConcernOutputShape struct {
	Name string

	Task *Concern

	Deliverable *Deliverable

	LinkShape
}

func (s *ConcernOutputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Concern)
}

func (s *ConcernOutputShape) GetAbstractEndElement() AbstractType {
	if s.Deliverable == nil {
		return nil
	}
	return s.Deliverable
}

func (s *ConcernOutputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Deliverable = abstractElement.(*Deliverable)
}

func (s *ConcernOutputShape) GetAbstractStartElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

var _ AssociationConcreteType = (*ConcernOutputShape)(nil)

// NoteShape
type NoteShape struct {
	Name string
	Note *Note

	IsExpanded bool

	RectShape
}

func (s *NoteShape) GetAbstractElement() AbstractType {
	if s.Note == nil {
		return nil // otherwise returns s.Note returns (*Note, nil), not nil
	}
	return s.Note
}

func (s *NoteShape) SetAbstractElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ ConcreteType = (*NoteShape)(nil)

type NoteDeliverableShape struct {
	Name string

	Note    *Note
	Deliverable *Deliverable

	LinkShape
}

// GetAbstractEndElement implements [AssociationConcreteType].
func (notedeliverableshape *NoteDeliverableShape) GetAbstractEndElement() AbstractType {
	if notedeliverableshape.Deliverable == nil {
		return nil
	}
	return notedeliverableshape.Deliverable
}

// GetAbstractStartElement implements [AssociationConcreteType].
func (notedeliverableshape *NoteDeliverableShape) GetAbstractStartElement() AbstractType {
	if notedeliverableshape.Note == nil {
		return nil
	}
	return notedeliverableshape.Note
}

// SetAbstractEndElement implements [AssociationConcreteType].
func (notedeliverableshape *NoteDeliverableShape) SetAbstractEndElement(deliverable AbstractType) {
	notedeliverableshape.Deliverable = deliverable.(*Deliverable)
}

// SetAbstractStartElement implements [AssociationConcreteType].
func (notedeliverableshape *NoteDeliverableShape) SetAbstractStartElement(note AbstractType) {
	notedeliverableshape.Note = note.(*Note)
}

var _ AssociationConcreteType = (*NoteDeliverableShape)(nil)

type NoteTaskShape struct {
	Name string

	Note *Note

	Task *Concern

	LinkShape
}

func (s *NoteTaskShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *NoteTaskShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Concern)
}

func (s *NoteTaskShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteTaskShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteTaskShape)(nil)

type NoteStakeholderShape struct {
	Name string

	Note *Note

	Stakeholder *Stakeholder

	LinkShape
}

func (s *NoteStakeholderShape) GetAbstractEndElement() AbstractType {
	if s.Stakeholder == nil {
		return nil
	}
	return s.Stakeholder
}

func (s *NoteStakeholderShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Stakeholder = abstractElement.(*Stakeholder)
}

func (s *NoteStakeholderShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteStakeholderShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteStakeholderShape)(nil)

// StakeholderShape
type StakeholderShape struct {
	Name        string
	Stakeholder *Stakeholder

	IsExpanded bool

	RectShape
}

func (s *StakeholderShape) GetAbstractElement() AbstractType {
	if s.Stakeholder == nil {
		return nil // otherwise returns s.Resource returns (*Resource, nil), not nil
	}
	return s.Stakeholder
}

func (s *StakeholderShape) SetAbstractElement(abstractElement AbstractType) {
	s.Stakeholder = abstractElement.(*Stakeholder)
}

var _ ConcreteType = (*StakeholderShape)(nil)

// A StakeholderCompositionShape is the link between a Resource
// and its parent Resource
type StakeholderCompositionShape struct {
	Name string

	Stakeholder *Stakeholder

	LinkShape
}

func (s *StakeholderCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Stakeholder == nil {
		return nil
	}
	return s.Stakeholder
}

func (s *StakeholderCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Stakeholder = abstractElement.(*Stakeholder)
}

func (s *StakeholderCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Stakeholder == nil || s.Stakeholder.parentStakeholder == nil {
		return nil
	}
	return s.Stakeholder.parentStakeholder
}

func (s *StakeholderCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element, because it is computed
	// elsewhere
	// s.Resource.parentResource = abstractElement.(*Resource)
}

var _ AssociationConcreteType = (*DeliverableCompositionShape)(nil)

type StakeholderConcernShape struct {
	Name string

	Stakeholder *Stakeholder

	Concern *Concern

	LinkShape
}

func (s *StakeholderConcernShape) GetAbstractEndElement() AbstractType {
	if s.Concern == nil {
		return nil
	}
	return s.Concern
}

func (s *StakeholderConcernShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Concern = abstractElement.(*Concern)
}

func (s *StakeholderConcernShape) GetAbstractStartElement() AbstractType {
	if s.Stakeholder == nil {
		return nil
	}
	return s.Stakeholder
}

func (s *StakeholderConcernShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Stakeholder = abstractElement.(*Stakeholder)
}

var _ AssociationConcreteType = (*StakeholderConcernShape)(nil)

// RequirementShape
type RequirementShape struct {
	Name        string
	Requirement *Requirement

	IsExpanded bool

	RectShape
}

func (s *RequirementShape) GetAbstractElement() AbstractType {
	if s.Requirement == nil {
		return nil
	}
	return s.Requirement
}

func (s *RequirementShape) SetAbstractElement(abstractElement AbstractType) {
	s.Requirement = abstractElement.(*Requirement)
}

var _ ConcreteType = (*RequirementShape)(nil)

// ConceptShape
type ConceptShape struct {
	Name    string
	Concept *Concept

	IsExpanded bool

	RectShape
}

func (s *ConceptShape) GetAbstractElement() AbstractType {
	if s.Concept == nil {
		return nil
	}
	return s.Concept
}

func (s *ConceptShape) SetAbstractElement(abstractElement AbstractType) {
	s.Concept = abstractElement.(*Concept)
}

var _ ConcreteType = (*ConceptShape)(nil)

type deliverableConceptKey struct {
	Deliverable *Deliverable
	Concept     *Concept
}

type DeliverableConceptShape struct {
	Name string

	Deliverable *Deliverable
	Concept     *Concept

	LinkShape
}

func (s *DeliverableConceptShape) GetAbstractEndElement() AbstractType {
	if s.Concept == nil {
		return nil
	}
	return s.Concept
}

func (s *DeliverableConceptShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Concept = abstractElement.(*Concept)
}

func (s *DeliverableConceptShape) GetAbstractStartElement() AbstractType {
	if s.Deliverable == nil {
		return nil
	}
	return s.Deliverable
}

func (s *DeliverableConceptShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Deliverable = abstractElement.(*Deliverable)
}

var _ AssociationConcreteType = (*DeliverableConceptShape)(nil)
