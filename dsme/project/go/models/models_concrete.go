package models

import "math/rand"

type Diagram struct {
	Name           string
	IsChecked      bool
	IsEditable_    bool
	IsInRenameMode bool

	Product_Shapes              []*ProductShape
	map_Product_ProductShape    map[*Product]*ProductShape
	ProductsWhoseNodeIsExpanded []*Product // to be made private once in production (no need to persist)
	IsPBSNodeExpanded           bool
	IsWBSNodeExpanded           bool

	Composition_Shapes           []*CompositionShape
	map_Product_CompositionShape map[*Product]*CompositionShape

	ExpandableNodeObject
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

// ProductShape is for both Product and Task
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

// A CompositionShape is the link between a product
// and its parent product
type CompositionShape struct {
	Name string

	Product *Product

	LinkShape
}
