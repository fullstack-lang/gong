package models

import (
	"log"
	"reflect"
	"time"
)

// Engine describes a tiny discrete event simulation engine
// it is responsible for advancing the time
// An engine manages a set of agents
// swagger:model Engine
type Engine struct {
	// Name of the engine "laundramat" for instance
	Name string

	// startTime is the simulation start date
	startTime time.Time

	// endTime is the simulatio end date
	endTime time.Time
	EndTime string

	// currentTime is the simulation current time
	currentTime time.Time
	CurrentTime string

	DisplayFormat string

	// number of the seconds elapsed since the beginning of the simulation
	SecondsSinceStart float64

	// list of engine agents
	agents []AgentInterface `gorm:"-"`

	// Fired events
	Fired int

	// control mode.
	ControlMode ControlMode

	// engine state
	State EngineState

	// LastEvent ...
	LastEvent *EventInterface `gorm:"-"`

	// LastEvent agent
	LastEventAgent *AgentInterface `gorm:"-"`

	// Simulation supportspecific callback
	// on the engine events
	Simulation SimulationInterface `gorm:"-"`

	// management of commits
	//
	// the engine is reponsible for performing commits of agents to the back repo
	// when the nextCommitDate > LastCommitDate
	// when a commit has been performed, the engine updates the LastCommitDate
	// swagger:ignore
	nextCommitDate   time.Time
	nextCheckoutDate time.Time

	// swagger:ignore
	lastCommitDate time.Time

	// swagger:ignore
	lastCheckoutDate time.Time

	//
	// management of realtime simulation (PLAY mode)
	//

	// Speed compared to realtime
	Speed float64

	// when in PLAY mode, both horizon provides input to the main thread to fire events and to compute the sleep time
	// neccessary to cadence the simulation at the desired speed
	// in PLAY mode, the simulation goes from one horizon to the next
	// both horizon are set up by the
	nextRealtimeHorizon      time.Time
	nextSimulatedTimeHorizon time.Time
}

// swagger:enum ControlMode
type ControlMode string

func (engine *Engine) SetStartTime(time time.Time) {
	engine.startTime = time
}
func (engine *Engine) GetStartTime() time.Time {
	return engine.startTime
}

const RFC3339MillisFormat = "2006-01-02T15:04:05.000Z07:00"

func (engine *Engine) SetCurrentTime(time time.Time) {
	engine.currentTime = time
	engine.CurrentTime = engine.currentTime.Format(RFC3339MillisFormat)
	if engine.DisplayFormat != "" {
		engine.CurrentTime = engine.currentTime.Format(engine.DisplayFormat)
	}
}
func (engine *Engine) GetCurrentTime() time.Time {
	return engine.currentTime
}
func (engine *Engine) SetEndTime(time time.Time) {
	engine.endTime = time
	engine.EndTime = engine.endTime.Format(RFC3339MillisFormat)
	if engine.DisplayFormat != "" {
		engine.EndTime = engine.endTime.Format(engine.DisplayFormat)
	}
}
func (engine *Engine) GetEndTime() time.Time {
	return engine.endTime
}

// values for ControlMode
const (
	AUTONOMOUS     ControlMode = "Autonomous" // iota
	CLIENT_CONTROL ControlMode = "ClientControl"
)

// swagger:enum EngineState
type EngineState string

// values for EngineState
const (
	RUNNING EngineState = "RUNNING" // iota
	PAUSED  EngineState = "PAUSED"
	OVER    EngineState = "OVER"
)

// AppendAgent to the engine
func (engine *Engine) AppendAgent(agent AgentInterface) {
	engine.agents = append(engine.agents, agent)
	agent.SetEngine(engine)
}

// RemoveAgent to the engine
func (engine *Engine) RemoveAgent(agent AgentInterface) {
	for index, _agent := range engine.agents {

		// Order is not important
		// If you do not care about ordering, you have the much faster
		// possibility to swap the element to delete with the one at the end of the slice and then return the n-1 first elements:
		// https://stackoverflow.com/a/37335777/5803707
		if _agent == agent {
			engine.agents[index] = engine.agents[len(engine.agents)-1]
			engine.agents = engine.agents[:len(engine.agents)-1]
		}
	}
}

// GetNextEvent ...
func (engine *Engine) GetNextEvent() (agent AgentInterface, nextEventFireTime time.Time, event EventInterface) {

	// engine.ListEvents()
	firstAgent := true
	for _, _agent := range engine.agents {
		_event, agentNextEventFireTime := _agent.GetNextEvent()

		// some agents have no events
		if _event == nil {
			continue
		}

		if firstAgent || agentNextEventFireTime.Before(nextEventFireTime) {
			nextEventFireTime = agentNextEventFireTime
			agent = _agent
			event = _event
		}
		firstAgent = false
	}
	return agent, nextEventFireTime, event
}

// ListEvents
func (engine *Engine) ListEvents() {

	log.Printf("Engine : current time %30s - seconds since start %f",
		engine.currentTime, engine.currentTime.Sub(engine.startTime).Seconds())
	for idx, agent := range engine.agents {
		log.Printf("\tAgent[%d] type %30s - tech name %30s", idx, reflect.TypeOf(agent), agent.GetTechName())

		for idxEvent, event := range agent.Events() {
			log.Printf("\t\tEvent[%d] type %30s - fire time %s - rel %f",
				idxEvent, reflect.TypeOf(event), event.GetFireTime(), event.GetFireTime().Sub(engine.startTime).Seconds())
		}
	}
}

// FireNextEvent fires earliest event
// advances current time
func (engine *Engine) FireNextEvent() (agent AgentInterface, nextTimeEvent time.Time, event EventInterface) {

	agent, nextTimeEvent, event = engine.GetNextEvent()

	if agent == nil {
		return
	}

	agent.FireNextEvent()

	engine.LastEvent = &event
	engine.LastEventAgent = &agent
	engine.SetCurrentTime(nextTimeEvent)
	engine.SecondsSinceStart = (engine.currentTime.Sub(engine.startTime)).Seconds()

	engine.Fired++

	return agent, nextTimeEvent, event
}

// Run will advance time till currentTime > EndTime
func (engine *Engine) Run() {
	log.Printf("time : run\n")

	agent, time, event := engine.GetNextEvent()

	_ = agent
	_ = event
	time_s := time.String()
	_ = time_s

	for (!time.IsZero()) && time.Before(engine.endTime) {
		engine.FireNextEvent()
		agent, time, event = engine.GetNextEvent()

		_ = agent
		_ = event
		time_s := time.String()
		_ = time_s
	}
}

// RunTillAnyStateHasChanged will advance time till currentTime > EndTime or one state Changed in the implementation
func (engine *Engine) RunTillAnyStateHasChanged() {
	log.Printf("time : run\n")

	_, time, _ := engine.GetNextEvent()

	hasAnyStateHasChanged := false

	for (!time.IsZero()) && time.Before(engine.endTime) && !hasAnyStateHasChanged {
		var agent AgentInterface
		agent, _, _ = engine.FireNextEvent()

		if agent == nil {
			return
		}
		_, time, _ = engine.GetNextEvent()
		if engine.Simulation != nil {
			hasAnyStateHasChanged = engine.Simulation.HasAnyStateChanged(engine)
		}
	}
}

func (engine *Engine) Agents() (agents []AgentInterface) {
	return engine.agents
}

func (engine *Engine) GetLastCommitNb() (commitNb uint) {

	if engine.Simulation != nil {
		commitNb = engine.Simulation.GetLastCommitNb()
	}

	return commitNb
}

func (engine *Engine) GetLastCommitNbFromFront() (commitNb uint) {

	if engine.Simulation != nil {
		commitNb = engine.Simulation.GetLastCommitNbFromFront()
	}

	return commitNb
}
