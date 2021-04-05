package diagrams

import (
	uml "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gong/examples/laundromat/go/models"
)

var defaultDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Machine{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 153.000000,
			Fields: []*uml.Field{
				{
					Field: models.Machine{}.Agent,
				},
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
			Struct: new(models.MachineStateEnum),
			Position: &uml.Position{
				X: 40.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Fields: []*uml.Field{
				{
					Field: models.MACHINE_DOOR_CLOSED_IDLE,
				},
				{
					Field: models.MACHINE_DOOR_CLOSED_RUNNING,
				},
				{
					Field: models.MACHINE_DOOR_OPEN,
				},
			},
		},
		{
			Struct: &(models.Washer{}),
			Position: &uml.Position{
				X: 340.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Links: []*uml.Link{
				{
					Field: models.Washer{}.Machine,
					Middlevertice: &uml.Vertice{
						X: 590.000000,
						Y: 200.000000,
					},
					Multiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Washer{}.Agent,
				},
				{
					Field: models.Washer{}.CleanedLaundryWeight,
				},
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
		{
			Struct: new(models.WasherStateEnum),
			Position: &uml.Position{
				X: 340.000000,
				Y: 500.000000,
			},
			Width:  240.000000,
			Heigth: 153.000000,
			Fields: []*uml.Field{
				{
					Field: models.WASHER_CLOSE_DOOR,
				},
				{
					Field: models.WASHER_IDLE,
				},
				{
					Field: models.WASHER_LOAD_DRUM,
				},
				{
					Field: models.WASHER_OPEN_DOOR,
				},
				{
					Field: models.WASHER_START_PROGRAM,
				},
				{
					Field: models.WASHER_UNLOAD_DRUM,
				},
				{
					Field: models.WASHER_WAIT_PROGRAM_END,
				},
			},
		},
	},
}
