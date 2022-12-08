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
				X: 52.000000,
				Y: 78.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Astruct{}.Anotherarrayofb,
					Middlevertice: &uml.Vertice{
						X: 551.000000,
						Y: 23.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Astruct{}.Associationtob,
					Middlevertice: &uml.Vertice{
						X: 291.000000,
						Y: 233.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Astruct{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Bstruct{}),
			Position: &uml.Position{
				X: 460.000000,
				Y: 130.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
	Notes: []*uml.NoteShape{
		{
			Name: `note on the organization between line and points`,
			Body: `
github.com/fullstack-lang/test/go/models is a model that showcases gong possibilites

Astruct is associated with Bstruct via 2 kinds of association
- a N-0..1 association (directly)
- a N-M association via AstructBstructUse
`,
			X:      130.000000,
			Y:      260.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
