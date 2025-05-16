// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Cursor_WOP struct {
	// insertion point

	Name string

	StartX float64

	EndX float64

	Y1 float64

	Y2 float64

	DurationSeconds float64

	Color string

	FillOpacity float64

	Stroke string

	StrokeOpacity float64

	StrokeWidth float64

	StrokeDashArray string

	StrokeDashArrayWhenSelected string

	Transform string

	IsPlaying bool
}

func (from *Cursor) CopyBasicFields(to *Cursor) {
	// insertion point
	to.Name = from.Name
	to.StartX = from.StartX
	to.EndX = from.EndX
	to.Y1 = from.Y1
	to.Y2 = from.Y2
	to.DurationSeconds = from.DurationSeconds
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.IsPlaying = from.IsPlaying
}

