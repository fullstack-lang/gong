package models

// ProductShape
type ProductShape struct {
	Name    string
	Product *Product

	isExpanded bool

	RectShape
}

func (s *ProductShape) GetAbstractElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *ProductShape) SetAbstractElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

var _ ConcreteType = (*ProductShape)(nil)

// A ProductCompositionShape is the link between a product
// and its parent product
type ProductCompositionShape struct {
	Name string

	Product *Product

	LinkShape
}

func (s *ProductCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *ProductCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
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

// TaskShape is for both Task
type TaskShape struct {
	Name string
	Task *Task

	isExpanded bool

	IsShowDate bool

	RectShape
}

func (s *TaskShape) GetAbstractElement() AbstractType {
	if s.Task == nil {
		return nil // Explicitly return interface nil, otherwise returns (*Task, nil)
	}
	return s.Task
}

func (s *TaskShape) SetAbstractElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

var _ ConcreteType = (*TaskShape)(nil)

// A TaskCompositionShape is the link between a task
// and its parent task
type TaskCompositionShape struct {
	Name string

	Task *Task

	LinkShape
}

func (s *TaskCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *TaskCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *TaskCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Task == nil || s.Task.parentTask == nil {
		return nil
	}
	return s.Task.parentTask
}

func (s *TaskCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element because it is computed
	// elsewhere
	// s.Task.parentTask = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*TaskCompositionShape)(nil)

// TaskGroupShape
type TaskGroupShape struct {
	Name      string
	TaskGroup *TaskGroup

	isExpanded bool

	RectShape
}

func (s *TaskGroupShape) GetAbstractElement() AbstractType {
	if s.TaskGroup == nil {
		return nil // Explicitly return interface nil
	}
	return s.TaskGroup
}

func (s *TaskGroupShape) SetAbstractElement(abstractElement AbstractType) {
	s.TaskGroup = abstractElement.(*TaskGroup)
}

var _ ConcreteType = (*TaskGroupShape)(nil)

// A taskProductKey allows mapping of [TaskInputShape] within a diagram
type taskProductKey struct {
	Task    *Task
	Product *Product
}

type noteProductKey struct {
	Note    *Note
	Product *Product
}

type noteTaskKey struct {
	Note *Note
	Task *Task
}

type noteResourceKey struct {
	Note     *Note
	Resource *Resource
}

type resourceTaskKey struct {
	Resource *Resource
	Task     *Task
}

type TaskInputShape struct {
	Name string

	Product *Product // input product

	Task *Task

	LinkShape
}

func (s *TaskInputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *TaskInputShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *TaskInputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *TaskInputShape) GetAbstractStartElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

var _ AssociationConcreteType = (*TaskInputShape)(nil)

type TaskOutputShape struct {
	Name string

	Task *Task

	Product *Product

	LinkShape
}

func (s *TaskOutputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *TaskOutputShape) GetAbstractEndElement() AbstractType {
	if s.Product == nil {
		return nil
	}
	return s.Product
}

func (s *TaskOutputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *TaskOutputShape) GetAbstractStartElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

var _ AssociationConcreteType = (*TaskOutputShape)(nil)

// NoteShape
type NoteShape struct {
	Name string
	Note *Note

	isExpanded bool

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
	Product *Product

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
	noteproductshape.Product = product.(*Product)
}

// SetAbstractStartElement implements [AssociationConcreteType].
func (noteproductshape *NoteProductShape) SetAbstractStartElement(note AbstractType) {
	noteproductshape.Note = note.(*Note)
}

var _ AssociationConcreteType = (*NoteProductShape)(nil)

type NoteTaskShape struct {
	Name string

	Note *Note

	Task *Task

	LinkShape
}

func (s *NoteTaskShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *NoteTaskShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
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

type NoteResourceShape struct {
	Name string

	Note *Note

	Resource *Resource

	LinkShape
}

func (s *NoteResourceShape) GetAbstractEndElement() AbstractType {
	if s.Resource == nil {
		return nil
	}
	return s.Resource
}

func (s *NoteResourceShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Resource = abstractElement.(*Resource)
}

func (s *NoteResourceShape) GetAbstractStartElement() AbstractType {
	if s.Note == nil {
		return nil
	}
	return s.Note
}

func (s *NoteResourceShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Note = abstractElement.(*Note)
}

var _ AssociationConcreteType = (*NoteResourceShape)(nil)

// ResourceShape
type ResourceShape struct {
	Name     string
	Resource *Resource

	isExpanded bool

	RectShape
}

func (s *ResourceShape) GetAbstractElement() AbstractType {
	if s.Resource == nil {
		return nil // otherwise returns s.Resource returns (*Resource, nil), not nil
	}
	return s.Resource
}

func (s *ResourceShape) SetAbstractElement(abstractElement AbstractType) {
	s.Resource = abstractElement.(*Resource)
}

var _ ConcreteType = (*ResourceShape)(nil)

// A ResourceCompositionShape is the link between a Resource
// and its parent Resource
type ResourceCompositionShape struct {
	Name string

	Resource *Resource

	LinkShape
}

func (s *ResourceCompositionShape) GetAbstractEndElement() AbstractType {
	if s.Resource == nil {
		return nil
	}
	return s.Resource
}

func (s *ResourceCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Resource = abstractElement.(*Resource)
}

func (s *ResourceCompositionShape) GetAbstractStartElement() AbstractType {
	if s.Resource == nil || s.Resource.parentResource == nil {
		return nil
	}
	return s.Resource.parentResource
}

func (s *ResourceCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element, because it is computed
	// elsewhere
	// s.Resource.parentResource = abstractElement.(*Resource)
}

var _ AssociationConcreteType = (*ProductCompositionShape)(nil)

type ResourceTaskShape struct {
	Name string

	Resource *Resource

	Task *Task

	LinkShape
}

func (s *ResourceTaskShape) GetAbstractEndElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

func (s *ResourceTaskShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *ResourceTaskShape) GetAbstractStartElement() AbstractType {
	if s.Resource == nil {
		return nil
	}
	return s.Resource
}

func (s *ResourceTaskShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Resource = abstractElement.(*Resource)
}

var _ AssociationConcreteType = (*ResourceTaskShape)(nil)
