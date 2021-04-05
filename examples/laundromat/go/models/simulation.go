package models

import (
	"log"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/examples/laundromat/go/events"

	gongdoc_models "github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	gongsim_models "github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

// Simulation is the callback support for
// events that happens on the generic engine
type Simulation struct {
	Name string

	Machine *Machine
	Washer  *Washer

	machineState MachineStateEnum
	washerState  WasherStateEnum

	Etats_Machine *gongdoc_models.Umlsc
	Etats_Washer  *gongdoc_models.Umlsc

	LastCommitNb int
}

func (simulation *Simulation) setInitialStateVectorOfAgentsAndSimulation() {

	// seven days of simulation for gongsim
	gongsim_models.EngineSingloton.SetStartTime(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
	gongsim_models.EngineSingloton.SetCurrentTime(gongsim_models.EngineSingloton.GetStartTime())
	gongsim_models.EngineSingloton.State = gongsim_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", gongsim_models.EngineSingloton.GetStartTime())

	gongsim_models.EngineSingloton.SetEndTime(time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", gongsim_models.EngineSingloton.GetEndTime())

	// seven days of simulation
	gongsim_models.EngineSingloton.SetStartTime(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
	gongsim_models.EngineSingloton.SetCurrentTime(gongsim_models.EngineSingloton.GetStartTime())
	gongsim_models.EngineSingloton.State = gongsim_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", gongsim_models.EngineSingloton.GetStartTime())

	gongsim_models.EngineSingloton.SetEndTime(time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", gongsim_models.EngineSingloton.GetEndTime())

	gongsim_models.EngineSingloton.Speed = 1.0

	cleaned := true
	simulation.Machine = (&Machine{
		Agent: gongsim_models.Agent{
			TechName: "machine-1",
		},
		State:          MACHINE_DOOR_CLOSED_IDLE,
		Cleanedlaundry: cleaned,
		Name:           "machine-1",
	}).StageCopy().AppendToSingloton()

	simulation.Machine.Cleanedlaundry = cleaned

	simulation.Washer = (&Washer{
		Name:    "washer-1",
		Machine: simulation.Machine,
		Agent: gongsim_models.Agent{
			TechName: "washer-1",
		},
		State: WASHER_IDLE,
	})
}

// NewSimulation ...
func NewSimulation() (simulation *Simulation) {

	simulation = &Simulation{
		Name: "Laundromat at the street corner",
	}

	simulation.setInitialStateVectorOfAgentsAndSimulation()

	// append Machine & Washer
	simulation.Washer.AppendToSingloton().Stage()
	// simulation.Machine.Stage()
	// simulation.Washer.Stage()

	simulation.CreateInitialEvents()

	simulation.machineState = simulation.Machine.State
	simulation.washerState = simulation.Washer.State

	for umlsc := range gongdoc_models.Stage.Umlscs {
		if strings.Contains(umlsc.Name, "Machine") {
			simulation.Etats_Machine = umlsc
		}
		if strings.Contains(umlsc.Name, "Washer") {
			simulation.Etats_Washer = umlsc
		}
	}

	return
}

// CreateInitialEvents ...
func (simulation *Simulation) CreateInitialEvents() {

	simulation.Washer.Reset()
	simulation.Washer.QueueUpdateEvent(30 * time.Second)

	// washer add laundry periodic event (every 12 hours)
	var addLaundry events.NewDirtyLaundry
	addLaundry.SetFireTime(gongsim_models.EngineSingloton.GetStartTime())
	addLaundry.Period = events.LAUNDRY_LOAD_PERIOD
	addLaundry.Name = "addLaundry"

	simulation.Washer.QueueEvent(&addLaundry)

	simulation.Machine.Reset()
	simulation.Machine.QueueUpdateEvent(30 * time.Second)
}

// HasAnyStateChanged ...
func (simulation *Simulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {

	if simulation.Washer.State != simulation.washerState ||
		simulation.Machine.State != simulation.machineState {
		simulation.washerState = simulation.Washer.State
		simulation.machineState = simulation.Machine.State

		simulation.Etats_Machine.Activestate = string(simulation.machineState)
		simulation.Etats_Washer.Activestate = string(simulation.washerState)

		gongdoc_models.Stage.Commit()
		log.Printf("EngineSpecific: HasAnyStateChanged %s", engine.GetCurrentTime().String())
		return true
	}

	return false
}

// CommitAgents commit all staged agents to the back repo
func (simulation *Simulation) CommitAgents(engine *gongsim_models.Engine) {

	// update DB with new values
	// convert time.Duration in minutes
	simulation.Machine.RemainingTimeMinutes = int(simulation.Machine.RemainingTime.Minutes())

	// commit the simulation agent states
	Stage.Commit()

	// update then commit the state charts diagrams
	simulation.Etats_Machine.Activestate = string(simulation.Machine.State)
	simulation.Etats_Washer.Activestate = string(simulation.Washer.State)
	gongdoc_models.Stage.Commit()
}

// CheckoutAgents checkout all staged agents to the back repo
func (simulation *Simulation) CheckoutAgents(engine *gongsim_models.Engine) {

	// checkout the simulation agent states
	Stage.Checkout()
}

func (simulation *Simulation) GetLastCommitNb() (commitNb uint) {

	if Stage.BackRepo != nil {
		commitNb = Stage.BackRepo.GetLastCommitNb()
	}

	return commitNb
}

// Reset simulation
func (simulation *Simulation) Reset(engine *gongsim_models.Engine) {

	// set initial conditions
	simulation.setInitialStateVectorOfAgentsAndSimulation()

	// add initial events
	simulation.CreateInitialEvents()

	simulation.washerState = simulation.Washer.State
	simulation.machineState = simulation.Machine.State
	simulation.Etats_Machine.Activestate = string(simulation.machineState)
	simulation.Etats_Washer.Activestate = string(simulation.washerState)

	log.Printf("Simulation reset")
}
