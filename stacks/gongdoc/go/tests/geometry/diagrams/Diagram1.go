package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/tests/geometry/models"
)

var Diagram1 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 80.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 380.000000,
						Y: 250.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 360.000000,
						Y: 10.000000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 90.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
				{
					Field: models.Point{}.Y,
				},
			},
		},
		{
			Struct: &(models.PointWithName{}),
			Position: &uml.Position{
				X: 800.000000,
				Y: 90.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointWithName{}.Point,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 200.000000,
					},
					Multiplicity: "1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointWithName{}.Name,
				},
			},
		},
	},
}
