package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

var defaultDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Agent{}),
			Position: &uml.Position{
				X: 340.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Agent{}.Engine,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Agent{}.TechName,
				},
			},
		},
		{
			Struct: new(models.ControlMode),
			Position: &uml.Position{
				X: 340.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.AUTONOMOUS,
				},
				{
					Field: models.CLIENT_CONTROL,
				},
			},
		},
		{
			Struct: &(models.Engine{}),
			Position: &uml.Position{
				X: 640.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 183.000000,
			Fields: []*uml.Field{
				{
					Field: models.Engine{}.ControlMode,
				},
				{
					Field: models.Engine{}.Fired,
				},
				{
					Field: models.Engine{}.LastEvent,
				},
				{
					Field: models.Engine{}.LastEventAgent,
				},
				{
					Field: models.Engine{}.Name,
				},
				{
					Field: models.Engine{}.SecondsSinceStart,
				},
				{
					Field: models.Engine{}.Simulation,
				},
				{
					Field: models.Engine{}.Speed,
				},
				{
					Field: models.Engine{}.State,
				},
			},
		},
		{
			Struct: new(models.EngineState),
			Position: &uml.Position{
				X: 640.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.OVER,
				},
				{
					Field: models.PAUSED,
				},
				{
					Field: models.RUNNING,
				},
			},
		},
	},
}
