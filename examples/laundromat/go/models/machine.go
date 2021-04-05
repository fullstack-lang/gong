package models

import (
	"fmt"
	"log"
	"time"

	"github.com/fullstack-lang/gong/examples/laundromat/go/events"
	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

// DrumCapacity limits drum fill up
const DrumCapacity = 5.0 // 5 KG

// ProgramDuration is fixed
const ProgramDuration = time.Minute * 80

// Machine is a sim agent
// swagger:model Machine
type Machine struct {
	// Agent
	// swagger:ignore
	models.Agent

	// Name is a mandatory field with metab
	Name string

	// DrumLoad in kg
	DrumLoad float64

	// Remaining Time
	RemainingTime        time.Duration
	RemainingTimeMinutes int

	// Cleanedlaundry indicate wether the laundry in the drum is wet (it has been cleand)
	Cleanedlaundry bool

	// State of the machine
	State MachineStateEnum
}

func (machine *Machine) AppendToSingloton() (res *Machine) {
	models.AppendToSingloton(machine)
	res = machine
	return
}

// MachineStateEnum ..
// swagger:enum MachineStateEnum
type MachineStateEnum string

// state
const (
	MACHINE_DOOR_OPEN           MachineStateEnum = "MACHINE_DOOR_OPEN"
	MACHINE_DOOR_CLOSED_RUNNING MachineStateEnum = "MACHINE_DOOR_CLOSED_RUNNING"
	MACHINE_DOOR_CLOSED_IDLE    MachineStateEnum = "MACHINE_DOOR_CLOSED_IDLE"
)

// FireNextEvent fire next Event
func (machine *Machine) FireNextEvent() {

	event, _ := machine.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *models.UpdateState:
		checkStateEvent := event.(*models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		machine.QueueEvent(checkStateEvent)

		// update state vector
		switch machine.State {
		case MACHINE_DOOR_CLOSED_RUNNING:
			machine.RemainingTime = machine.RemainingTime - checkStateEvent.Period
			if machine.RemainingTime < 0.0 {
				machine.RemainingTime = 0.0
			}
			if machine.RemainingTime == 0.0 {
				machine.State = MACHINE_DOOR_CLOSED_IDLE
				machine.Cleanedlaundry = true
			}
		}
	case *events.OpenDoor:
	case *events.CloseDoor:
	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}

}
