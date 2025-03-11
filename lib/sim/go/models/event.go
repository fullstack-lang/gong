package models

import (
	"time"
)

// Event is the elementary element of a discrete event simulation
// swagger:model Event
type Event struct {
	Name string

	// Fire is the time at which the event is Fired
	fireTime time.Time

	// Duration is the difference between the current time and the fire time of the
	// event. It is handy to compute directly the fire time
	Duration time.Duration

	agent *Agent
}

// GetFireTime ...
func (event *Event) GetFireTime() time.Time {
	return event.fireTime
}

// GetDuration ...
func (event *Event) GetDuration() time.Duration {
	return event.Duration
}

func (event *Event) SetFireTime(t time.Time) {
	event.fireTime = t
}

// EventInterface ...
// swagger:ignore
type EventInterface interface {
	GetFireTime() time.Time
	SetFireTime(time.Time)
	GetDuration() time.Duration
	GetName() string
	SetAgent(agent *Agent)
	GetAgent() *Agent
}

func (event *Event) GetAgent() *Agent      { return event.agent }
func (event *Event) SetAgent(agent *Agent) { event.agent = agent }
