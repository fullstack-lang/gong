package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/tests/geometry/models"
)

var Diagram4 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 180.000000,
				Y: 160.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 227.000000,
				Y: 266.000000,
			},
			Width:  240.000000,
			Heigth: 68.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
			},
		},
	},
}
