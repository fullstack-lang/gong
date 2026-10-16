// generated code - do not edit
package barline

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type BarStyleColor_WOP struct {
	// insertion point

	Name string

	Color string

	Value EnumBarStyle
}

func (from *BarStyleColor) GongCopyBasicFields(to *BarStyleColor) {
	// insertion point
	*to = *from
}

type Barline_WOP struct {
	// insertion point

	Name string

	Location EnumRightLeftMiddle

	Segno string

	Coda string

	Divisions string

	Id string
}

func (from *Barline) GongCopyBasicFields(to *Barline) {
	// insertion point
	to.Name = from.Name
	to.Location = from.Location
	to.Segno = from.Segno
	to.Coda = from.Coda
	to.Divisions = from.Divisions
	to.Id = from.Id
}

type Ending_WOP struct {
	// insertion point

	Name string

	Number string

	Type EnumStartStopDiscontinue

	EndLength string

	TextX string

	TextY string

	DefaultX string

	DefaultY string

	RelativeX string

	RelativeY string

	FontFamily string

	FontSize string

	Color string

	System EnumSystemRelation

	Value string
}

func (from *Ending) GongCopyBasicFields(to *Ending) {
	// insertion point
	*to = *from
}

type Repeat_WOP struct {
	// insertion point

	Name string

	Direction EnumBackwardForward

	Times int

	Winged EnumWinged
}

func (from *Repeat) GongCopyBasicFields(to *Repeat) {
	// insertion point
	*to = *from
}

// end of insertion point
