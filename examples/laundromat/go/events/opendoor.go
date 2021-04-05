package events

import "github.com/fullstack-lang/gong/stacks/gongsim/go/models"

// OpenDoor is an event whose role is open the door
// of the machine
type OpenDoor struct {
	models.Event
}
