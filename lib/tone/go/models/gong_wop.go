// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Freqency_WOP struct {
	// insertion point
	Name string
}

func (from *Freqency) CopyBasicFields(to *Freqency) {
	// insertion point
	to.Name = from.Name
}

type Note_WOP struct {
	// insertion point
	Name string
	Start float64
	Duration float64
	Velocity float64
	Info string
}

func (from *Note) CopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.Start = from.Start
	to.Duration = from.Duration
	to.Velocity = from.Velocity
	to.Info = from.Info
}

type Player_WOP struct {
	// insertion point
	Name string
	Status Status
}

func (from *Player) CopyBasicFields(to *Player) {
	// insertion point
	to.Name = from.Name
	to.Status = from.Status
}

