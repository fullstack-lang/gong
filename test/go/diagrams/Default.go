package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/test/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Astruct{}),
			Position: &uml.Position{
				X: 82.000000,
				Y: 58.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Astruct{}.AnAstruct,
					Middlevertice: &uml.Vertice{
						X: 442.000000,
						Y: 89.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Astruct{}.Associationtob,
					Middlevertice: &uml.Vertice{
						X: 556.000000,
						Y: 150.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
		},
		{
			Struct: &(models.AstructBstruct2Use{}),
			Position: &uml.Position{
				X: -395.000000,
				Y: 287.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.AstructBstructUse{}),
			Position: &uml.Position{
				X: 370.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Bstruct{}),
			Position: &uml.Position{
				X: 770.000000,
				Y: 60.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Cstruct{}),
			Position: &uml.Position{
				X: 72.000000,
				Y: 191.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
