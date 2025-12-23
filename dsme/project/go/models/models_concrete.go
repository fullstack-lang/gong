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

// ProductShape
type ProductShape struct {
	Name    string
	Product *Product

	IsExpanded bool

	RectShape
}

func newProductShapeToDiagram(product *Product, diagram *Diagram) (shape *ProductShape) {
	shape = new(ProductShape)
	shape.Product = product
	shape.Name = product.GetName() + "-" + diagram.GetName()
	shape.Height = 80
	shape.Width = 200
	shape.X = 100 + rand.Float64()*100.0
	shape.Y = 100 + rand.Float64()*100.0
	diagram.Product_Shapes = append(diagram.Product_Shapes, shape)

	return shape
}

// A ProductCompositionShape is the link between a product
// and its parent product
type ProductCompositionShape struct {
	Name string

	Product *Product

	LinkShape
}

// TaskShape is for both Task
type TaskShape struct {
	Name string
	Task *Task

	IsExpanded bool

	RectShape
}

func newTaskShapeToDiagram(task *Task, diagram *Diagram) (shape *TaskShape) {
	shape = new(TaskShape)
	shape.Task = task
	shape.Name = task.GetName() + "-" + diagram.GetName()
	shape.Height = 80
	shape.Width = 200
	shape.X = 100 + rand.Float64()*100.0
	shape.Y = 100 + rand.Float64()*100.0
	diagram.Task_Shapes = append(diagram.Task_Shapes, shape)

	return shape
}

// A TaskCompositionShape is the link between a task
// and its parent task
type TaskCompositionShape struct {
	Name string

	Task *Task

	LinkShape
}
