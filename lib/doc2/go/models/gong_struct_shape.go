package models

const GongStructShapeDefaultWidth = 240.0
const GongStructShapeDefaultHeight = 48.0

// GongStructShape mirrors joint.shapes.uml.Class
// swagger:model GongStructShape
type GongStructShape struct {
	Name string

	X float64
	Y float64

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:ident
	Identifier string

	// gongdoc can be integrated in a runtime application
	// the application can then set up the number of instances of Struct
	ShowNbInstances bool
	NbInstances     int

	// models of the composition of Field
	AttributeShapes []*AttributeShape

	// models of the composition of Link
	LinkShapes []*LinkShape

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Height float64

	// this is always false in the backend, but it can be set to true by the front end
	// this means it is selected by the user
	IsSelected bool
}
