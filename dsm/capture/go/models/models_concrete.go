package models

// ProductShape
type ProductShape struct {
	Name    string
	Product *Deliverable

	IsExpanded bool

	RectShape
}

func (s *ProductShape) GetAbstractElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *ProductShape) SetAbstractElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Deliverable)
}

var _ ConcreteType = (*ProductShape)(nil)

// A ProductCompositionShape is the link between a product
// and its parent product
type ProductCompositionShape struct {
	Name string

	Product *Deliverable

	LinkShape
}

func (s *ProductCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *ProductCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Deliverable)
}

func (s *ProductCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Product == nil || s.Product.parentProduct == nil {
		return nil
	}
	return s.Product.parentProduct
}

func (s *ProductCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element, because it is computed
	// elsewhere
	// s.Product.parentProduct = abstractElement.(*Product)
}

var _ AssociationConcreteType = (*ProductCompositionShape)(nil)

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

// A concernProductKey allows mapping of [ConcernInputShape] within a diagram
type concernProductKey struct {
	Concern *Concern
	Product *Deliverable
}

type noteProductKey struct {
	Note    *Note
	Product *Deliverable
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

	Deliverable *Deliverable // input product

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

	Product *Deliverable

	LinkShape
}

func (s *ConcernOutputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Concern)
}

func (s *ConcernOutputShape) GetAbstractEndElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *ConcernOutputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Deliverable)
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

type NoteProductShape struct {
	Name string

	Note    *Note
	Product *Deliverable

	LinkShape
}

// GetAbstractEndElement implements [AssociationConcreteType].
func (noteproductshape *NoteProductShape) GetAbstractEndElement() AbstractType {
	if noteproductshape.Product == nil {
		return nil
	}
	return noteproductshape.Product
}

// GetAbstractStartElement implements [AssociationConcreteType].
func (noteproductshape *NoteProductShape) GetAbstractStartElement() AbstractType {
	if noteproductshape.Note == nil {
		return nil
	}
	return noteproductshape.Note
}

// SetAbstractEndElement implements [AssociationConcreteType].
func (noteproductshape *NoteProductShape) SetAbstractEndElement(product AbstractType) {
	noteproductshape.Product = product.(*Deliverable)
}

// SetAbstractStartElement implements [AssociationConcreteType].
func (noteproductshape *NoteProductShape) SetAbstractStartElement(note AbstractType) {
	noteproductshape.Note = note.(*Note)
}

var _ AssociationConcreteType = (*NoteProductShape)(nil)

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

var _ AssociationConcreteType = (*ProductCompositionShape)(nil)

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
