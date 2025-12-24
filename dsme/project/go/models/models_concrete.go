package models

import "math/rand"

type Diagram struct {
	Name           string
	IsChecked      bool
	IsEditable_    bool
	IsInRenameMode bool

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

// ProductShape
type ProductShape struct {
	Name    string
	Product *Product

	IsExpanded bool

	RectShape
}

func (s *ProductShape) GetAbstractElement() AbstractType {
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
	return s.Product
}

func (s *ProductCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *ProductCompositionShape) GetAbstractStartElement() AbstractType {
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
	return s.Task
}

func (s *TaskCompositionShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *TaskCompositionShape) GetAbstractStartElement() AbstractType {
	return s.Task.parentTask
}

func (s *TaskCompositionShape) SetAbstractStartElement(abstractElement AbstractType) {
	// the parent element shall not be set by the concrete element because it is computed
	// elsewhere
	// s.Task.parentTask = abstractElement.(*Task)
}

var _ AssociationConcreteType = (*TaskCompositionShape)(nil)

func newShapeToDiagram[AT AbstractType, CT interface {
	*S
	RectShapeInterface
	ConcreteType
}, S Gongstruct](
	abstractElement AT,
	diagram *Diagram,
	shapes *[]CT,
	stage *Stage,
) CT {
	shape := CT(new(S))
	shape.StageVoid(stage)
	shape.SetAbstractElement(abstractElement)
	shape.SetName(abstractElement.GetName() + "-" + diagram.GetName())
	shape.SetHeight(80)
	shape.SetWidth(200)
	shape.SetX(100 + rand.Float64()*100.0)
	shape.SetY(100 + rand.Float64()*100.0)
	*shapes = append(*shapes, shape)

	return shape
}

// A taskProductKey allows mapping of [TaskInputShape] within a diagram
type taskProductKey struct {
	Task    *Task
	Product *Product
}

type TaskInputShape struct {
	Name string

	Task *Task

	Product *Product

	LinkShape
}

func (s *TaskInputShape) SetAbstractStartElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

func (s *TaskInputShape) GetAbstractEndElement() AbstractType {
	return s.Task
}

func (s *TaskInputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *TaskInputShape) GetAbstractStartElement() AbstractType {
	return s.Task
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
	return s.Task
}

func (s *TaskOutputShape) SetAbstractEndElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *TaskOutputShape) GetAbstractStartElement() AbstractType {
	return s.Task
}

var _ AssociationConcreteType = (*TaskOutputShape)(nil)
