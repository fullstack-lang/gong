package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/examples/laundromat/go/models"
)

var Diagramme_UML_Laundromat uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Machine{}),
			Position: &uml.Position{
				X: 570.000000,
				Y: 100.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Fields: []*uml.Field{
				{
					Field: models.Machine{}.Cleanedlaundry,
				},
				{
					Field: models.Machine{}.DrumLoad,
				},
				{
					Field: models.Machine{}.Name,
				},
				{
					Field: models.Machine{}.RemainingTime,
				},
				{
					Field: models.Machine{}.RemainingTimeMinutes,
				},
				{
					Field: models.Machine{}.State,
				},
			},
		},
		{
			Struct: &(models.Washer{}),
			Position: &uml.Position{
				X: 60.000000,
				Y: 100.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.Washer{}.Machine,
					Middlevertice: &uml.Vertice{
						X: 420.000000,
						Y: 70.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Washer{}.DirtyLaundryWeight,
				},
				{
					Field: models.Washer{}.Name,
				},
				{
					Field: models.Washer{}.State,
				},
			},
		},
	},
}
