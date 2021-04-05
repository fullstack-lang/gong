package models

import (
	"log"
	"time"
)

// AgentInterface is the interface that must be followed by part of the simulation
// swagger:model AgentInterface
type AgentInterface interface {

	// GetNextEvent provides the event with earliest fire time
	GetNextEvent() (event EventInterface, time time.Time)

	// FireNextEvent fire next Event
	FireNextEvent()

	// Events fire next Event
	Events() []EventInterface

	SetEngine(engine *Engine)

	// Reset
	Reset()

	// get technical name
	GetTechName() string

	QueueEvent(event EventInterface)
	TransfertAndQueueEvent(event EventInterface)
}

// Agent is the empty struct to perform
// generic agents chores
type Agent struct {
	TechName string

	// list of events, in increasing fire time
	// swagger:ignore
	events []EventInterface `gorm:"-"`

	// timeLastChecked, time at which the event check state was called
	// swagger:ignore
	timeLastChecked time.Time

	// last event time
	// swagger:ignore
	lastEventTime time.Time

	// usefull to append an agent to the Engine from an agent
	// swagger:ignore
	Engine *Engine
}

func (agent *Agent) SetEngine(engine *Engine) {
	agent.Engine = engine
}

// AppendAgentToEngine append an agent to the engine
func (agent *Agent) AppendAgentToEngine(newAgent AgentInterface) {
	agent.Engine.AppendAgent(newAgent)
}

// RemoveAgentToEngine append an agent to the engine
func (agent *Agent) RemoveAgentToEngine(newAgent AgentInterface) {
	agent.Engine.RemoveAgent(newAgent)
}

// Events append an agent to the engine
func (agent *Agent) Events() []EventInterface {
	return agent.events
}

// Events append an agent to the engine
func (agent *Agent) GetTechName() string {
	return agent.TechName
}

// Reset removes all events from the agent and resets internal checks
func (agent *Agent) Reset() {

	agent.events = nil
	agent.lastEventTime = time.Time{}
}

// GetNextEvent provides the next event from a time point of view
// by convention 0 means infinity
func (agent *Agent) GetNextEvent() (EventInterface, time.Time) {

	if len(agent.events) == 0 {
		return nil, time.Time{}
	}

	return agent.events[0], agent.events[0].GetFireTime()
}

// GetNextEventAndRemoveIt provides the next event from a time point of view
// by convention 0 means infinity
func (agent *Agent) GetNextEventAndRemoveIt() (event EventInterface, t time.Time) {

	event, t = agent.events[0], agent.events[0].GetFireTime()
	if event == nil {
		log.Panic("cannot fire event when no event in queue")
	}

	// remove event
	agent.events = agent.events[1:len(agent.events)]

	//  update last time
	agent.lastEventTime = event.GetFireTime()
	return event, t
}

func (agent *Agent) TransfertAndQueueEvent(event EventInterface) {
	event.SetAgent(agent)
	agent.QueueEvent(event)
}

// QueueEvent is the function by which an agent queues an event from another agent (or of himself)
func (agent *Agent) QueueEvent(event EventInterface) {

	// since an agent can modify the event, one need to make a private copy
	// since we have an interface, passing the argument as value was not possible
	// val := reflect.ValueOf(event)

	// type of the event
	// eventType := reflect.TypeOf(event).Elem()
	// log.Printf("Type of event " + eventType.Name())
	// _ = eventType
	// zero := reflect.Zero(eventType).Interface()

	// eventPtrType := reflect.TypeOf(event)
	// zeroPtr := reflect.Zero(eventPtrType)

	// zeroPtrCast, ok := zeroPtr.Interface().(EventInterface)
	// if !ok {
	// 	log.Printf("Type assertion failed")
	// }

	// log.Printf("%d", zeroPtrCast.GetFireTime().Hour())
	// _ = zero

	// // value of the event
	// privateEventInterface := reflect.New(eventType).Interface()
	// _ = privateEventInterface

	// // new we need to type assert the event to its concret type
	// // https://www.tapirgames.com/blog/golang-interface-implementation
	// // https://medium.com/golangspec/type-assertions-in-go-e609759c42e1
	// if privateEvent, ok := privateEventInterface.(Event); !ok {
	// 	_ = privateEvent
	// 	log.Fatal("unable to type-assert event")
	// }

	if event.GetFireTime().Before(agent.lastEventTime) {
		log.Panic("inserting event in the past")
	}

	// check that this event is not shared by any other agent
	if event.GetAgent() != nil && event.GetAgent() != agent {
		log.Printf("%v - %v", event.GetAgent(), agent)
		log.Panic("inserting a event that already owned by another agent")
	}
	event.SetAgent(agent)

	if len(agent.events) == 0 {
		agent.events = append(agent.events, event)
		return
	}
	// parse all events and insert event when appropriate
	for idx, _event := range agent.events {
		if event.GetFireTime().Before(_event.GetFireTime()) {

			agent.events = append(agent.events, nil)       // Making space for the new element
			copy(agent.events[idx+1:], agent.events[idx:]) // Shifting elements
			agent.events[idx] = event                      // Copying/inserting the value
			return
		}
	}
	// append at the end
	agent.events = append(agent.events, event)
}

// Implements the interface function
func AppendToSingloton(agent AgentInterface) AgentInterface {
	EngineSingloton.AppendAgent(agent)
	return agent
}

// QueueUpdateEvent put an UpdateState event with duration to the agent
// with fire time equals to EngineSingloton current time
//
// return the agent
func (agent *Agent) QueueUpdateEvent(duration time.Duration) *Agent {

	// washer checks its state every 30''
	var event UpdateState
	event.fireTime = EngineSingloton.GetCurrentTime()
	event.Period = duration
	event.Name = "Update event with duration " + duration.String()

	agent.QueueEvent(&event)

	return agent
}
