// generated code - do not edit
package clock

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type ClockAbstract_WOP struct {
	// insertion point

	Name string

	RadialRepetitions int

	Transparency float64

	RelativeTubeDiameter float64

	RelativeHeight3DTorus float64

	ClockTorusVerticalScale float64

	RelativeHeight float64

	ProjectionAngle float64
}

func (from *ClockAbstract) GongCopyBasicFields(to *ClockAbstract) {
	// insertion point
	*to = *from
}

// end of insertion point
