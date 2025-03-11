package models

import (
	"time"
)

// UpdateState is an event whose role is to
// ask the agant to compute its new vector state
// swagger:model UpdateState
type UpdateState struct {
	Event

	// Period is the time between two fires of FireTime
	Period time.Duration
}
