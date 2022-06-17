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
				X: 70.000000,
				Y: 60.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
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
					Field: models.Astruct{}.Anarrayofb,
					Middlevertice: &uml.Vertice{
						X: 550.000000,
						Y: 161.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Astruct{}.Anarrayofb2Use,
					Middlevertice: &uml.Vertice{
						X: 410.000000,
						Y: 331.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Astruct{}.Aenum,
				},
				{
					Field: models.Astruct{}.Anotherbooleanfield,
				},
				{
					Field: models.Astruct{}.Benum,
				},
				{
					Field: models.Astruct{}.Booleanfield,
				},
				{
					Field: models.Astruct{}.CEnum,
				},
			},
		},
		{
			Struct: &(models.AstructBstruct2Use{}),
			Position: &uml.Position{
				X: 650.000000,
				Y: 200.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
		{
			Struct: &(models.Bstruct{}),
			Position: &uml.Position{
				X: 650.000000,
				Y: 60.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
