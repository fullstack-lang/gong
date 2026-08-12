package models

type ClockAbstract struct {
	Name string

	// RadialRepetitions impacts the conversion from 2D ribbons to 3D torus.
	// The 3D torus is construed by 2Pi/RadialRepetitions each time it ranges over the GrowthCurve2DRibbon
	// If ThreeDModule is 1, the 3D torus wraps 360 degrees while ranging over the GrowthCurve2DRibbon
	RadialRepetitions int

	// Transparency of the clock (between 0.0 for opaque and 1.0 for transparent)
	Transparency float64

	// RelativeTubeDiameter of the 3D tube (default is 0.01)
	RelativeTubeDiameter float64

	RelativeHeight3DTorus float64

	// scale applied to 3D Torus sample points
	ClockTorusVerticalScale float64

	// height to the top of the clock (standard is 45 cm)
	RelativeHeight float64

	// angle used to go from the 3D curve on the cylinder to the 2D curve for the clock top
	ProjectionAngle float64
}
