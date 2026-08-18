package models

type VaseAbstract struct {
	Name string

	// RelativeVerticalThickness of the growth curve. when growth curve are stacked, each is separate from the next
	// the vertical thickness is RelativeVerticalThickness x RhombusSideLength
	RelativeVerticalThickness float64

	// thickness allong the radius
	// radial with of the solid torus is RelativeRadialThickness x RhombusSideLength
	RelativeRadialThickness float64

	// For laser cutting the torus forms, they will be stack on top of another, without the rotation
	// RelativeCuttedStackFloorHeight x RhombusSideLength is the distance between each ribbon
	RelativeCuttedStackFloorHeight float64

	// 3D rotated torus stack on top of another, with 	 rotation
	// RelativeRotatedTorusSeparation x RhombusSideLength is an additionnal distance between each ribbon
	RelativeRotatedTorusSeparation float64

	// RotationRatio is the ratio of rotation between 0.0 and 1.0.
	// 0.0 means no rotation
	// 1.0 means the rotation of 2xPi x (GrowthVectorShape.X /  PlantCircumferenceShape.Length)
	RotationRatio float64

	heightAtRotRatio0 float64
	heightAtRotRatio1 float64

	// RadialRepetitions impacts the conversion from 2D ribbons to 3D torus.
	// The 3D torus  is construed by 2Pi/RadialRepetitions each time it ranges over the  GrowthCurve2DRibbon
	// If ThreeDModule is 1, the 3D torus wraps 360 degrees while ranging over the GrowthCurve2DRibbon
	RadialRepetitions int

	// Transparency of the 3D torus material (between 0.0 for opaque and 1.0 for transparent)
	Transparency float64

	// HasAlternatingRingColors when true alternates colors between rings in the stack
	HasAlternatingRingColors bool

	// RelativeTrajectoryOffsetX x PlantCircumferenceShape.Length
	// RelativeTrajectoryOffsetY x PlantCircumferenceShape.Length
	// is taken into account for drawing the PartiallyGrowthCurve2DTrajectory
	RelativeTrajectoryOffsetX, RelativeTrajectoryOffsetY float64

	// NbStepP1P2 is the number of interpolation points for P1, P2 pairs
	NbStepP1P2 int

	// ChosenStep is
	ChosenStep int

	// RelativeHorizontalRingsHeight x RhombusSideLength is the manual height of the bottom/top horizontal rings.
	// If 0.0, the height automatically matches the wave amplitude of the bottom/top ring.
	RelativeHorizontalRingsHeight float64

	OffsetKeyX float64
	OffsetKeyY float64
	HeightKey  float64
	WidthKey   float64

	// 1.0 is full hey hole occupency
	RelativeKeySize float64

	// MovieNbFrames is the number of frames for movie recording.
	// The rotation increment per frame is 1.0 / MovieNbFrames.
	MovieNbFrames int

	PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway

	// used to construe the xxxHalfwayArcShapeGrid (x)
	TopStartArcShapeGrid *TopStartArcShapeGrid
	TopEndArcShapeGrid   *TopEndArcShapeGrid

	ShiftedBottomTopStartArcShapeGrid *ShiftedBottomTopStartArcShapeGrid
	TopMidArcVectorShapeGrid          *TopMidArcVectorShapeGrid

	StartHalfwayArcShapeGrid    *StartHalfwayArcShapeGrid
	TopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
	EndHalfwayArcShapeGrid      *EndHalfwayArcShapeGrid
	TopEndHalfwayArcShapeGrid   *TopEndHalfwayArcShapeGrid

	StackOfRotatedGrowthCurve2D    *StackOfRotatedGrowthCurve2D
	TopStackOfRotatedGrowthCurve2D *TopStackOfRotatedGrowthCurve2D

	TopGrowthCurve2D *TopGrowthCurve2D

	StackOfGrowthCurve2D    *StackOfGrowthCurve2D
	TopStackOfGrowthCurve2D *TopStackOfGrowthCurve2D

	StackOfGrowthCurve2DRibbon        *StackOfGrowthCurve2DRibbon
	StackOfRotatedGrowthCurve2DRibbon *StackOfRotatedGrowthCurve2DRibbon
	GrowthCurve2DRibbon               *GrowthCurve2DRibbon
	ShiftedRightGrowthCurve2DRibbon   *ShiftedRightGrowthCurve2DRibbon
	ShiftedLeftGrowthCurve2DRibbon    *ShiftedLeftGrowthCurve2DRibbon

	PartiallyGrowthCurve2DRibbon            *PartiallyGrowthCurve2DRibbon
	ShiftedLeftPartiallyGrowthCurve2DRibbon *ShiftedLeftPartiallyGrowthCurve2DRibbon
	PartiallyGrowthCurve2DTrajectory        *PartiallyGrowthCurve2DTrajectory
	PartiallyGrowthCurve2DTrajectoryP1P2    *PartiallyGrowthCurve2DTrajectoryP1P2
	PxShape                                 *PxShape
	ChosenP1P2PairShape                     *ChosenP1P2PairShape
	KeyHoleShape                            *KeyHoleShape
}
