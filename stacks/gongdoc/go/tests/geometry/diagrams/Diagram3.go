package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/tests/geometry/models"
)

var Diagram3 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 87.000000,
				Y: 154.000000,
			},
			Width:  248.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 166.000000,
				Y: 233.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.PointWithName{}),
			Position: &uml.Position{
				X: 89.000000,
				Y: 320.000000,
			},
			Width:  248.000000,
			Heigth: 48.000000,
		},
	},
}
