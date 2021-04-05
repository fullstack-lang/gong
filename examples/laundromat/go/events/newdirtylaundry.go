package events

import (
	"time"

	"github.com/fullstack-lang/gong/stacks/gongsim/go/models"
)

const LAUNDRY_LOAD_PER_EVENT = 8           // 8 kg
const LAUNDRY_LOAD_PERIOD = 12 * time.Hour // every 12 hours

// NewDirtyLaundry is an event whose role is to fill up
// laundry basket of washer
type NewDirtyLaundry struct {
	models.Event

	// Period is the time between two fires of FireTime
	Period time.Duration
}
