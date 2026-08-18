package models

type StoolAbstract struct {
	Name string

	// RadialRepetitions impacts the conversion from 2D ribbons to 3D torus.
	// The 3D torus  is construed by 2Pi/RadialRepetitions each time it ranges over the  GrowthCurve2DRibbon
	// If ThreeDModule is 1, the 3D torus wraps 360 degrees while ranging over the GrowthCurve2DRibbon
	RadialRepetitions int

	// Transparency of the stool (between 0.0 for opaque and 1.0 for transparent)
	Transparency float64

	// RelativeTubeDiameter of the 3D tube (default is 0.01)
	RelativeTubeDiameter float64

	RelativeHeight3DTorus float64

	// scale applied to 3D Torus sample points
	StoolTorusVerticalScale float64

	// height to the top of the seat (standard is 45 cm)
	RelativeHeight        float64
	RelativeSeatThickness float64 // 5 cm ?

	// angle use to go from the 3D curve on the cylinder to the 2D curves for the seat top, seat bottom and feet bottom
	ProjectionAngle float64

	// Defines what are the points for the eye and those that are not
	RelativeEyeSeparationCriteria float64

	// Strength of the control vector at the eye corner connecting points (default 0.55)
	RelativeEyeCornerControlVectorStrength float64
}
