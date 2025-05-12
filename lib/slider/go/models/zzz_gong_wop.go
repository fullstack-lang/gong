// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour
var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Checkbox_WOP struct {
	// insertion point
	Name string
	ValueBool bool
	LabelForTrue string
	LabelForFalse string
}

func (from *Checkbox) CopyBasicFields(to *Checkbox) {
	// insertion point
	to.Name = from.Name
	to.ValueBool = from.ValueBool
	to.LabelForTrue = from.LabelForTrue
	to.LabelForFalse = from.LabelForFalse
}

type Group_WOP struct {
	// insertion point
	Name string
	Percentage float64
}

func (from *Group) CopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.Percentage = from.Percentage
}

type Layout_WOP struct {
	// insertion point
	Name string
}

func (from *Layout) CopyBasicFields(to *Layout) {
	// insertion point
	to.Name = from.Name
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
}

func (from *Slider) CopyBasicFields(to *Slider) {
	// insertion point
	to.Name = from.Name
	to.IsFloat64 = from.IsFloat64
	to.IsInt = from.IsInt
	to.MinInt = from.MinInt
	to.MaxInt = from.MaxInt
	to.StepInt = from.StepInt
	to.ValueInt = from.ValueInt
	to.MinFloat64 = from.MinFloat64
	to.MaxFloat64 = from.MaxFloat64
	to.StepFloat64 = from.StepFloat64
	to.ValueFloat64 = from.ValueFloat64
}

