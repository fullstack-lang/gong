package models

const GongEnumShapeDefaultWidth = 240.0
const GongEnumShapeDefaultHeigth = 48.0

// GongEnumShape mirrors joint.shapes.uml.Class
// swagger:model GongEnumShape
type GongEnumShape struct {
	Name string

	Position *Position

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:ident
	Identifier string

	GongEnumValueEntrys []*GongEnumValueEntry

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Heigth float64
}

type GongEnumShapeType int

const (
	Int GongEnumShapeType = iota
	String
)
