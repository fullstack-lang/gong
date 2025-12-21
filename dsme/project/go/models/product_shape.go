package models

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
	shape.X = 100
	shape.Y = 100
	diagram.Product_Shapes = append(diagram.Product_Shapes, shape)

	return shape
}
