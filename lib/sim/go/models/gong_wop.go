// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Command_WOP struct {
	// insertion point
	Name string
	Command CommandType
	CommandDate string
}

func (from *Command) CopyBasicFields(to *Command) {
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

func (from *DummyAgent) CopyBasicFields(to *DummyAgent) {
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

func (from *Engine) CopyBasicFields(to *Engine) {
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

func (from *Event) CopyBasicFields(to *Event) {
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

func (from *Status) CopyBasicFields(to *Status) {
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

func (from *UpdateState) CopyBasicFields(to *UpdateState) {
	// insertion point
	to.Name = from.Name
	to.Duration = from.Duration
	to.Period = from.Period
}

