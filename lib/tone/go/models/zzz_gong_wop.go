// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Freqency_WOP struct {
	// insertion point

	Name string
}

func (from *Freqency) GongCopyBasicFields(to *Freqency) {
	// insertion point
	*to = *from
}

type Note_WOP struct {
	// insertion point

	Name string

	Start float64

	Duration float64

	Velocity float64

	Info string
}

func (from *Note) GongCopyBasicFields(to *Note) {
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

func (from *Player) GongCopyBasicFields(to *Player) {
	// insertion point
	*to = *from
}

// end of insertion point
