// generated code - do not edit
package stool

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type StoolAbstract_WOP struct {
	// insertion point

	Name string

	RadialRepetitions int

	Transparency float64

	RelativeTubeDiameter float64

	RelativeHeight3DTorus float64

	StoolTorusVerticalScale float64

	RelativeHeight float64

	RelativeSeatThickness float64

	ProjectionAngle float64

	RelativeEyeSeparationCriteria float64

	RelativeEyeCornerControlVectorStrength float64
}

func (from *StoolAbstract) GongCopyBasicFields(to *StoolAbstract) {
	// insertion point
	*to = *from
}

// end of insertion point
