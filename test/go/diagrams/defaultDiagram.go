package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/test/go/models"
)

var defaultDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: new(models.AEnumType),
			Position: &uml.Position{
				X: 340.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.ENUM_VAL1,
				},
				{
					Field: models.ENUM_VAL2,
				},
			},
		},
		{
			Struct: &(models.Aclass{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 198.000000,
			Links: []*uml.Link{
				{
					Field: models.Aclass{}.Anarrayofa,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 200.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Aclass{}.Anarrayofb,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 250.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Aclass{}.Anotherarrayofb,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 300.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Aclass{}.Anotherassociationtob_2,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 350.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Aclass{}.Associationtob,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 400.000000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Aclass{}.Cclass,
					Middlevertice: &uml.Vertice{
						X: 290.000000,
						Y: 450.000000,
					},
					Multiplicity: "1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Aclass{}.Aenum,
				},
				{
					Field: models.Aclass{}.Aenum_2,
				},
				{
					Field: models.Aclass{}.Anotherbooleanfield,
				},
				{
					Field: models.Aclass{}.Benum,
				},
				{
					Field: models.Aclass{}.Booleanfield,
				},
				{
					Field: models.Aclass{}.CFloatfield,
				},
				{
					Field: models.Aclass{}.CName,
				},
				{
					Field: models.Aclass{}.Floatfield,
				},
				{
					Field: models.Aclass{}.Intfield,
				},
				{
					Field: models.Aclass{}.Name,
				},
			},
		},
		{
			Struct: new(models.BEnumType),
			Position: &uml.Position{
				X: 40.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.BENUM_VAL1,
				},
				{
					Field: models.BENUM_VAL2,
				},
			},
		},
		{
			Struct: &(models.Bclass{}),
			Position: &uml.Position{
				X: 340.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.Bclass{}.Floatfield,
				},
				{
					Field: models.Bclass{}.Intfield,
				},
				{
					Field: models.Bclass{}.Name,
				},
			},
		},
		{
			Struct: &(models.Cclass{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Cclass{}.CFloatfield,
				},
				{
					Field: models.Cclass{}.CName,
				},
			},
		},
	},
}
