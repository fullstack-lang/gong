// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Checkbox_WOP struct {
	// insertion point

	Name string

	ValueBool bool

	LabelForTrue string

	LabelForFalse string
}

func (from *Checkbox) GongCopyBasicFields(to *Checkbox) {
	// insertion point
	*to = *from
}

type Group_WOP struct {
	// insertion point

	Name string

	Percentage float64
}

func (from *Group) GongCopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.Percentage = from.Percentage
}

type Layout_WOP struct {
	// insertion point

	Name string

	IsWithCustomGutterSize bool

	GutterSize float64
}

func (from *Layout) GongCopyBasicFields(to *Layout) {
	// insertion point
	to.Name = from.Name
	to.IsWithCustomGutterSize = from.IsWithCustomGutterSize
	to.GutterSize = from.GutterSize
}

type Slider_WOP struct {
	// insertion point

	Name string

	IsFloat64 bool

	IsInt bool

	MinInt int

	MaxInt int

	StepInt int

	ValueInt int

	MinFloat64 float64

	MaxFloat64 float64

	StepFloat64 float64

	ValueFloat64 float64

	IsDisabled bool
}

func (from *Slider) GongCopyBasicFields(to *Slider) {
	// insertion point
	*to = *from
}

// end of insertion point
