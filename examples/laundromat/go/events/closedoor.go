package events

import "github.com/fullstack-lang/gong/stacks/gongsim/go/models"

// CloseDoor is an event whose role is close the door
// of the machine
type CloseDoor struct {
	models.Event
}
