// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

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

func (from *Cursor) GongCopyBasicFields(to *Cursor) {
	// insertion point
	*to = *from
}

// end of insertion point
