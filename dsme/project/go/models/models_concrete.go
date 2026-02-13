package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Diagram struct {
	Name           string
	IsChecked      bool
	IsEditable_    bool
	IsInRenameMode bool
	ShowPrefix     bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	ExpandableNodeObject

	Product_Shapes              []*ProductShape
	map_Product_ProductShape    map[*Product]*ProductShape
	ProductsWhoseNodeIsExpanded []*Product // to be made private once in production (no need to persist)
	IsPBSNodeExpanded           bool

	ProductComposition_Shapes           []*ProductCompositionShape
	map_Product_ProductCompositionShape map[*Product]*ProductCompositionShape

	IsWBSNodeExpanded bool

	Task_Shapes                    []*TaskShape
	map_Task_TaskShape             map[*Task]*TaskShape
	TasksWhoseNodeIsExpanded       []*Task // to be made private once in production (no need to persist)ExpandableNodeObject
	TasksWhoseInputNodeIsExpanded  []*Task
	TasksWhoseOutputNodeIsExpanded []*Task

	TaskComposition_Shapes        []*TaskCompositionShape
	map_Task_TaskCompositionShape map[*Task]*TaskCompositionShape

	TaskInputShapes         []*TaskInputShape
	map_Task_TaskInputShape map[taskProductKey]*TaskInputShape

	TaskOutputShapes         []*TaskOutputShape
	map_Task_TaskOutputShape map[taskProductKey]*TaskOutputShape

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NoteProductShapes         []*NoteProductShape
	map_Note_NoteProductShape map[noteProductKey]*NoteProductShape

	NoteTaskShapes         []*NoteTaskShape
	map_Note_NoteTaskShape map[noteTaskKey]*NoteTaskShape

	map_Product_Rect map[*Product]*svg.Rect
	map_Task_Rect    map[*Task]*svg.Rect
	map_Note_Rect    map[*Note]*svg.Rect

	map_SvgRect_ProductShape map[*svg.Rect]*ProductShape
	map_SvgRect_TaskShape    map[*svg.Rect]*TaskShape
	map_SvgRect_NoteShape    map[*svg.Rect]*NoteShape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

type ConcreteType interface {
	GongstructIF
	GetAbstractElement() AbstractType
	SetAbstractElement(AbstractType)
}

type AssociationConcreteType interface {
	GongstructIF
	GetAbstractStartElement() AbstractType
	SetAbstractStartElement(AbstractType)
	GetAbstractEndElement() AbstractType
	SetAbstractEndElement(AbstractType)
}

type AssociationConcreteType2[SourceAT AbstractType, TargetAT AbstractType] interface {
	GongstructIF
	GetAbstractStartElement() SourceAT
	SetAbstractStartElement(SourceAT)
	GetAbstractEndElement() TargetAT
	SetAbstractEndElement(TargetAT)
}

// ProductShape
type ProductShape struct {
	Name    string
	Product *Product

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

	IsExpanded bool

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
