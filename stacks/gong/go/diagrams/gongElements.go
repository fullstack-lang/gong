package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/stacks/gong/go/models"
)

var gongElements uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.GongBasicField{}),
			Position: &uml.Position{
				X: 370.000000,
				Y: 180.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.GongBasicField{}.GongEnum,
					Middlevertice: &uml.Vertice{
						X: 840.000000,
						Y: 160.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.GongBasicField{}.Name,
				},
				{
					Field: models.GongBasicField{}.Type,
				},
			},
		},
		{
			Struct: &(models.GongEnum{}),
			Position: &uml.Position{
				X: 710.000000,
				Y: 230.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.GongEnum{}.GongEnumValues,
					Middlevertice: &uml.Vertice{
						X: 830.000000,
						Y: 310.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.GongEnum{}.Name,
				},
			},
		},
		{
			Struct: &(models.GongEnumValue{}),
			Position: &uml.Position{
				X: 700.000000,
				Y: 360.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.GongEnumValue{}.Name,
				},
				{
					Field: models.GongEnumValue{}.Value,
				},
			},
		},
		{
			Struct: &(models.GongStruct{}),
			Position: &uml.Position{
				X: 30.000000,
				Y: 30.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.GongStruct{}.GongBasicFields,
					Middlevertice: &uml.Vertice{
						X: 180.000000,
						Y: 230.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.GongStruct{}.PointerToGongStructFields,
					Middlevertice: &uml.Vertice{
						X: 150.000000,
						Y: 320.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.GongStruct{}.SliceOfPointerToGongStructFields,
					Middlevertice: &uml.Vertice{
						X: 60.000000,
						Y: 410.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.GongStruct{}.Fields,
				},
				{
					Field: models.GongStruct{}.Name,
				},
			},
		},
		{
			Struct: &(models.ModelPkg{}),
			Position: &uml.Position{
				X: 420.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.ModelPkg{}.GongEnums,
				},
				{
					Field: models.ModelPkg{}.GongStructs,
				},
				{
					Field: models.ModelPkg{}.Name,
				},
				{
					Field: models.ModelPkg{}.PkgPath,
				},
			},
		},
		{
			Struct: &(models.PointerToGongStructField{}),
			Position: &uml.Position{
				X: 370.000000,
				Y: 290.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.PointerToGongStructField{}.GongStruct,
					Middlevertice: &uml.Vertice{
						X: 620.000000,
						Y: 120.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.PointerToGongStructField{}.Name,
				},
			},
		},
		{
			Struct: &(models.SliceOfPointerToGongStructField{}),
			Position: &uml.Position{
				X: 370.000000,
				Y: 380.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.SliceOfPointerToGongStructField{}.GongStruct,
					Middlevertice: &uml.Vertice{
						X: 810.000000,
						Y: 70.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.SliceOfPointerToGongStructField{}.Name,
				},
			},
		},
	},
}
