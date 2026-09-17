// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Command_WOP struct {
	// insertion point

	Name string

	Command CommandType

	CommandDate string
}

func (from *Command) GongCopyBasicFields(to *Command) {
	// insertion point
	to.Name = from.Name
	to.Command = from.Command
	to.CommandDate = from.CommandDate
}

type DummyAgent_WOP struct {
	// insertion point

	TechName string

	Name string
}

func (from *DummyAgent) GongCopyBasicFields(to *DummyAgent) {
	// insertion point
	to.TechName = from.TechName
	to.Name = from.Name
}

type Engine_WOP struct {
	// insertion point

	Name string

	EndTime string

	CurrentTime string

	DisplayFormat string

	SecondsSinceStart float64

	Fired int

	ControlMode ControlMode

	State EngineState

	Speed float64
}

func (from *Engine) GongCopyBasicFields(to *Engine) {
	// insertion point
	to.Name = from.Name
	to.EndTime = from.EndTime
	to.CurrentTime = from.CurrentTime
	to.DisplayFormat = from.DisplayFormat
	to.SecondsSinceStart = from.SecondsSinceStart
	to.Fired = from.Fired
	to.ControlMode = from.ControlMode
	to.State = from.State
	to.Speed = from.Speed
}

type Event_WOP struct {
	// insertion point

	Name string

	Duration time.Duration
}

func (from *Event) GongCopyBasicFields(to *Event) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
}

type Status_WOP struct {
	// insertion point

	Name string

	CurrentCommand CommandType

	CompletionDate string

	CurrentSpeedCommand SpeedCommandType

	SpeedCommandCompletionDate string
}

func (from *Status) GongCopyBasicFields(to *Status) {
	// insertion point
	to.Name = from.Name
	to.CurrentCommand = from.CurrentCommand
	to.CompletionDate = from.CompletionDate
	to.CurrentSpeedCommand = from.CurrentSpeedCommand
	to.SpeedCommandCompletionDate = from.SpeedCommandCompletionDate
}

type UpdateState_WOP struct {
	// insertion point

	Name string

	Duration time.Duration

	Period time.Duration
}

func (from *UpdateState) GongCopyBasicFields(to *UpdateState) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
	to.Period = from.Period
}

// end of insertion point
