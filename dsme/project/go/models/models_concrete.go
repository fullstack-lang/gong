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

	Task_Shapes              []*TaskShape
	map_Task_TaskShape       map[*Task]*TaskShape
	TasksWhoseNodeIsExpanded []*Task // to be made private once in production (no need to persist)ExpandableNodeObject

	TaskComposition_Shapes        []*TaskCompositionShape
	map_Task_TaskCompositionShape map[*Task]*TaskCompositionShape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

type ConcreteType interface {
	GongstructIF
	GetAbstractElement() AbstractType
	SetAbstractElement(AbstractType)
}

type CompositionConcreteType interface {
	GongstructIF
	GetAbstractParentElement() AbstractType
	SetAbstractParentElement(AbstractType)
	GetAbstractChildElement() AbstractType
	SetAbstractChildElement(AbstractType)
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

func newProductShapeToDiagram(product *Product, diagram *Diagram) (shape *ProductShape) {
	return newShapeToDiagram(
		product,
		diagram,
		&diagram.Product_Shapes,
	)
}

// A ProductCompositionShape is the link between a product
// and its parent product
type ProductCompositionShape struct {
	Name string

	Product *Product

	LinkShape
}

func (s *ProductCompositionShape) GetAbstractChildElement() AbstractType {
	return s.Product
}

func (s *ProductCompositionShape) SetAbstractChildElement(abstractElement AbstractType) {
	s.Product = abstractElement.(*Product)
}

func (s *ProductCompositionShape) GetAbstractParentElement() AbstractType {
	return s.Product.parentProduct
}

func (s *ProductCompositionShape) SetAbstractParentElement(abstractElement AbstractType) {
	s.Product.parentProduct = abstractElement.(*Product)
}

var _ CompositionConcreteType = (*ProductCompositionShape)(nil)

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

func newTaskShapeToDiagram(task *Task, diagram *Diagram) (shape *TaskShape) {
	return newShapeToDiagram(
		task,
		diagram,
		&diagram.Task_Shapes,
	)
}

// A TaskCompositionShape is the link between a task
// and its parent task
type TaskCompositionShape struct {
	Name string

	Task *Task

	LinkShape
}

func newShapeToDiagram[AT AbstractType, CT interface {
	*S
	RectShapeInterface
	ConcreteType
}, S Gongstruct](
	abstractElement AT,
	diagram *Diagram,
	shapes *[]CT,
) CT {
	shape := CT(new(S))
	shape.SetAbstractElement(abstractElement)
	shape.SetName(abstractElement.GetName() + "-" + diagram.GetName())
	shape.SetHeight(80)
	shape.SetWidth(200)
	shape.SetX(100 + rand.Float64()*100.0)
	shape.SetY(100 + rand.Float64()*100.0)
	*shapes = append(*shapes, shape)

	return shape
}
