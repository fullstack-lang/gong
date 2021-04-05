package events

import (
	"time"

	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

// FillDrum is an event whose role is fill
// the Drum with 1kG
type FillDrum struct {
	models.Event

	// it takes a certain time
	Duration *time.Duration
}

// NewFillDrumEvent inits default Drum Event
func NewFillDrumEvent() FillDrum {
	var FillDrumEvent FillDrum
	*FillDrumEvent.Duration = time.Minute

	return FillDrumEvent
}
