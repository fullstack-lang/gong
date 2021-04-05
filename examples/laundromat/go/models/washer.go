package models

import (
	"fmt"
	"log"
	"time"

	"github.com/fullstack-lang/gong/examples/laundromat/go/events"
	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

// DrumLoadPerMinute rate of drum load
const DrumLoadPerMinute = 1.2 // 1.2 kg per minutes

// Washer is a sim agent
// swagger:model Washer
type Washer struct {
	// Agent
	models.Agent

	// Name is a mandatory field with metab
	Name string

	// DirtyLaundryWeight is the weigh of the dirty laundry to clean
	DirtyLaundryWeight float64

	// State
	State WasherStateEnum

	// Machine allocated to the washer
	// swagger:ignore
	Machine *Machine

	// CleanedLaundryWeight is the cumulative cleaned laundry
	// the Washer. in kilo
	CleanedLaundryWeight float64
}

func (washer *Washer) AppendToSingloton() (res *Washer) {
	models.AppendToSingloton(washer)
	res = washer
	return
}

// WasherStateEnum ..
// swagger:enum WasherStateEnum
type WasherStateEnum string

// state
const (
	WASHER_IDLE             WasherStateEnum = "WASHER_IDLE"
	WASHER_LOAD_DRUM        WasherStateEnum = "WASHER_LOAD_DRUM"
	WASHER_OPEN_DOOR        WasherStateEnum = "WASHER_OPEN_DOOR"
	WASHER_WAIT_PROGRAM_END WasherStateEnum = "WASHER_WAIT_PROGRAM_END"
	WASHER_CLOSE_DOOR       WasherStateEnum = "WASHER_CLOSE_DOOR"
	WASHER_UNLOAD_DRUM      WasherStateEnum = "WASHER_UNLOAD_DRUM"
	WASHER_START_PROGRAM    WasherStateEnum = "WASHER_START_PROGRAM"
)

// FireNextEvent fire next Event
func (washer *Washer) FireNextEvent() {

	event, _ := washer.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *models.UpdateState:
		checkStateEvent := event.(*models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		washer.TransfertAndQueueEvent(checkStateEvent)

		// compute next state
		washer.State = washer.ComputeState()

		// update internal & external data according to state
		switch washer.State {
		case WASHER_IDLE:
		case WASHER_LOAD_DRUM:
			// how much can be loaded in the drum
			periodInMinutes := float64(checkStateEvent.Period) / float64(time.Minute)
			load := DrumLoadPerMinute * periodInMinutes

			// check that we do not load drum above capacity
			// we do not check if load can be negative
			if (washer.Machine.DrumLoad + load) > DrumCapacity {
				load = DrumCapacity - washer.Machine.DrumLoad
			}

			// laundry cannot go below zero
			if washer.DirtyLaundryWeight-load < 0 {
				load = washer.DirtyLaundryWeight
			}

			washer.Machine.DrumLoad += load
			washer.DirtyLaundryWeight -= load

			// drum is loaded with dirty laundry
			washer.Machine.Cleanedlaundry = false

		case WASHER_OPEN_DOOR:
			washer.Machine.State = MACHINE_DOOR_OPEN

			var openDoorEvent events.OpenDoor
			openDoorEvent.Name = "open door"
			openDoorEvent.SetFireTime(event.GetFireTime())
			washer.Machine.QueueEvent(&openDoorEvent)

		case WASHER_WAIT_PROGRAM_END:
		case WASHER_CLOSE_DOOR:

			washer.Machine.State = MACHINE_DOOR_CLOSED_IDLE
			var closeDoorEvent events.OpenDoor
			closeDoorEvent.Name = "close door"
			closeDoorEvent.SetFireTime(event.GetFireTime())
			washer.Machine.QueueEvent(&closeDoorEvent)

		case WASHER_UNLOAD_DRUM:
			// how much can be unloaded in the drum
			periodInMinutes := float64(checkStateEvent.Period) / float64(time.Minute)
			unload := DrumLoadPerMinute * periodInMinutes

			// check that we do not unload he drum below 0
			if (washer.Machine.DrumLoad - unload) < 0.0 {
				unload = washer.Machine.DrumLoad
			}

			washer.Machine.DrumLoad -= unload
			washer.CleanedLaundryWeight += unload

		case WASHER_START_PROGRAM:
			washer.Machine.RemainingTime = ProgramDuration
			washer.Machine.State = MACHINE_DOOR_CLOSED_RUNNING
		default:
			log.Panic("unknow state")
		}

	case *events.NewDirtyLaundry:
		addLaundryEvent := event.(*events.NewDirtyLaundry)

		// add laudry
		washer.DirtyLaundryWeight += events.LAUNDRY_LOAD_PER_EVENT

		// post next event
		addLaundryEvent.SetFireTime(addLaundryEvent.GetFireTime().Add(addLaundryEvent.Period))
		washer.QueueEvent(addLaundryEvent)

	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

/* decision tree

doorOpen/

	drum empty/

			laundry empty --> close door // to go idle after

			laudry not empty --> load drum

	drum not empty/

		cleanedLaundry (wet) --> unload drum

		dirtyLandry (dry)/

			drum full  --> close door // to start programm

			drum not full/

				laundry empty --> close door // to start program

				laudry not empty --> load drum // until laundry empty or drum full

doorClosed/

	running --> wait end

	idle/

		drum not empty/

			dirty laundry in drum/

				laundry not empty/

					drum not full --> open door // to fill it up

					drum full --> start program

				laundry empty --> start program

			clean laundry in drum --> open door // to empty it

		drum empty/

			laundry not empty --> open door

			laundry empty --> idle
*/

// ComputeState computes the state
// eliminate element of the state vector one after the other
func (washer *Washer) ComputeState() WasherStateEnum {

	if washer.Machine.State == MACHINE_DOOR_OPEN { // 	doorOpen/

		if washer.Machine.DrumLoad == 0 { // 	drum empty/

			if washer.DirtyLaundryWeight == 0.0 { // 	laundry empty --> close door
				return WASHER_CLOSE_DOOR // to go idle later

			}
			// laudry not empty --> load drum
			return WASHER_LOAD_DRUM

		}
		// drum not empty/

		if washer.Machine.Cleanedlaundry { // cleanedLaundry (wet) --> unload drum
			return WASHER_UNLOAD_DRUM

		} else { // dirtyLandry (dry)

			if washer.Machine.DrumLoad == DrumCapacity { // drum full  --> close door // to start programm
				return WASHER_CLOSE_DOOR

			} else { // drum not full/
				if washer.DirtyLaundryWeight == 0.0 { // 	laundry empty --> close door // to start program
					return WASHER_CLOSE_DOOR

				} else { // 	laudry not empty --> load drum // until laundry empty or drum full
					return WASHER_LOAD_DRUM

				}
			}
		}
	} else { // doorClosed

		if washer.Machine.State == MACHINE_DOOR_CLOSED_RUNNING { // 	running --> wait end
			return WASHER_WAIT_PROGRAM_END
		} else { // 	idle/

			if washer.Machine.DrumLoad > 0.0 { // drum not empty

				if !(washer.Machine.Cleanedlaundry) { // dirty laundry in drum/

					if washer.Machine.DrumLoad >= DrumCapacity { // drum full --> start program
						return WASHER_START_PROGRAM
					} else { // drum not full
						if washer.DirtyLaundryWeight > 0 { // laundry not empty --> open door // to fill it up
							return WASHER_OPEN_DOOR

						} else { // laundry empty --> start program
							return WASHER_START_PROGRAM
						}
					}

				} else { // clean laundry in drum --> open door // to empty drum
					return WASHER_OPEN_DOOR
				}

			} else { // drum empty

				if washer.DirtyLaundryWeight > 0.0 { // laundry not empty, open door to load drum
					return WASHER_OPEN_DOOR
				} else { // if laundry == 0 --> IDLE
					return WASHER_IDLE
				}
			}
		}
	}

	// log.Printf("LaundryWeight %f\n", washer.LaundryWeight)
	// log.Printf("Machine DrumLoad %f\n", washer.Machine.DrumLoad)
	// log.Printf("Cleaned Laundry in drum ? %t\n", washer.Machine.Cleanedlaundry)
	// log.Printf("machine state %s\n", washer.Machine.State)
	// log.Panic("cannot compute state for washer")
	// return WASHER_FAILED
}
