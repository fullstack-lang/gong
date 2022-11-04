package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/test/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: &(models.Astruct{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 20.000000,
			},
			Width:  240.000000,
			Heigth: 243.000000,
			Links: []*uml.Link{
				{
					Field: models.Astruct{}.Associationtob,
					Middlevertice: &uml.Vertice{
						X: 605.000000,
						Y: 141.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Astruct{}.Aenum,
				},
				{
					Field: models.Astruct{}.Aenum_2,
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
				{
					Field: models.Astruct{}.CFloatfield,
				},
				{
					Field: models.Astruct{}.CName,
				},
				{
					Field: models.Astruct{}.Duration1,
				},
				{
					Field: models.Astruct{}.Floatfield,
				},
				{
					Field: models.Astruct{}.Intfield,
				},
				{
					Field: models.Astruct{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Bstruct{}),
			Position: &uml.Position{
				X: 470.000000,
				Y: 20.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Bstruct{}.Name,
				},
			},
		},
	},
}
