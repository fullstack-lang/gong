package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Command__000000_Gongsim_Command_Singloton := (&models.Command{}).Stage(stage)

	__Engine__000000_Simulation_Engine := (&models.Engine{}).Stage(stage)

	__Status__000000_Gongsim_Status_Singloton := (&models.Status{}).Stage(stage)

	// Setup of values

	__Command__000000_Gongsim_Command_Singloton.Name = `Gongsim Command Singloton`
	__Command__000000_Gongsim_Command_Singloton.Command = models.COMMAND_PAUSE
	__Command__000000_Gongsim_Command_Singloton.CommandDate = `1741668921015`

	__Engine__000000_Simulation_Engine.Name = `Simulation Engine`
	__Engine__000000_Simulation_Engine.EndTime = `01 Jan 1680`
	__Engine__000000_Simulation_Engine.CurrentTime = `22 Feb 1676`
	__Engine__000000_Simulation_Engine.DisplayFormat = `02 Jan 2006`
	__Engine__000000_Simulation_Engine.SecondsSinceStart = 4536000.000000
	__Engine__000000_Simulation_Engine.Fired = 3921494
	__Engine__000000_Simulation_Engine.State = models.PAUSED
	__Engine__000000_Simulation_Engine.Speed = 1296000.000000

	__Status__000000_Gongsim_Status_Singloton.Name = `Gongsim Status Singloton`
	__Status__000000_Gongsim_Status_Singloton.CurrentCommand = models.COMMAND_PAUSE
	__Status__000000_Gongsim_Status_Singloton.CompletionDate = ``
	__Status__000000_Gongsim_Status_Singloton.CurrentSpeedCommand = models.COMMAND_SPEED_STEADY
	__Status__000000_Gongsim_Status_Singloton.SpeedCommandCompletionDate = ``

	// Setup of pointers
	// setup of Command instances pointers
	__Command__000000_Gongsim_Command_Singloton.Engine = __Engine__000000_Simulation_Engine
	// setup of Engine instances pointers
	// setup of Status instances pointers
}
