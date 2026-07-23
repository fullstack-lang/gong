package models

// Plant
// In botanical phyllotaxy, N and M denote a contact parastichy pair, which represents
// the number of visible spirals winding in opposite directions around a plant's central
// axis. By standard convention, N and M are consecutive Fibonacci numbers ordered such
// that N < M (e.g., 8 and 13). These variables do not strictly dictate clockwise or
// counter-clockwise directions; rather, they quantify the two opposing sets of spirals.
// Depending on the specific chirality (handedness) of the plant specimen, N may
// represent the number of clockwise spirals while M represents the counter-clockwise
// spirals, or vice versa.
type Plant struct {
	Name string

	N int
	M int

	StackHeight int // height of growth curve stack

	// RhombusInsideAngle is set by the user. It represents the inside angle (in degrees) of the
	// fundamental diamond (rhombus) shape that makes up the grid of leaves on the plant's surface.
	// This angle determines the geometric direction of the two main lattice paths (up-right and up-left).
	RhombusInsideAngle float64

	// RelativeVerticalThickness of the growth curve. when growth curve are stacked, each is separate from the next
	// the vertical thickness is RelativeVerticalThickness x RhombusSideLength
	RelativeVerticalThickness float64

	// thickness allong the radius
	// radial with of the solid torus is RelativeRadialThickness x RhombusSideLength
	RelativeRadialThickness float64

	// RhombusSideLength is set by the user. It represents the physical length of the side of the
	// fundamental diamond (rhombus) shape. It acts as the scale or distance for each step
	// taken along the lattice paths.
	RhombusSideLength float64

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

	// ThreeDModulo impacts the conversion from 2D ribbons to 3D torus.
	// The 3D torus  is construed by 2Pi/ThreeDModulo each time it ranges over the  GrowthCurve2DRibbon
	// If ThreeDModule is 1, the 3D torus wraps 360 degrees while ranging over the GrowthCurve2DRibbon
	ThreeDModulo int

	// RelativeTrajectoryOffsetX x PlantCircumferenceShape.Length
	// RelativeTrajectoryOffsetY x PlantCircumferenceShape.Length
	// is taken into account for drawing the PartiallyGrowthCurve2DTrajectory
	RelativeTrajectoryOffsetX, RelativeTrajectoryOffsetY float64

	LibraryAbstractFields
	AbstractTypeFields

	IsSelected bool

	IsPlantDiagramsNodeExpanded bool
	PlantDiagrams               []*PlantDiagram

	AxesShape *AxesShape

	RhombusStuff *RhombusStuff

	GrowthVectorShape *GrowthVectorShape

	PerpendicularVectorGrid        *PerpendicularVectorGrid
	PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway
	BaseVectorShapeGrid            *BaseVectorShapeGrid
	ArcNormalVectorShapeGrid       *ArcNormalVectorShapeGrid

	// used to construe the xxxHalfwayArcShapeGrid (x)
	StartArcShapeGrid    *StartArcShapeGrid
	TopStartArcShapeGrid *TopStartArcShapeGrid
	EndArcShapeGrid      *EndArcShapeGrid
	TopEndArcShapeGrid   *TopEndArcShapeGrid

	ShiftedBottomTopStartArcShapeGrid *ShiftedBottomTopStartArcShapeGrid
	MidArcVectorShapeGrid             *MidArcVectorShapeGrid
	TopMidArcVectorShapeGrid          *TopMidArcVectorShapeGrid

	StartHalfwayArcShapeGrid    *StartHalfwayArcShapeGrid
	TopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
	EndHalfwayArcShapeGrid      *EndHalfwayArcShapeGrid
	TopEndHalfwayArcShapeGrid   *TopEndHalfwayArcShapeGrid

	StackOfRotatedGrowthCurve2D    *StackOfRotatedGrowthCurve2D
	TopStackOfRotatedGrowthCurve2D *TopStackOfRotatedGrowthCurve2D

	GrowthCurve2D    *GrowthCurve2D
	TopGrowthCurve2D *TopGrowthCurve2D

	StackOfGrowthCurve2D    *StackOfGrowthCurve2D
	TopStackOfGrowthCurve2D *TopStackOfGrowthCurve2D

	StackOfGrowthCurve2DRibbon        *StackOfGrowthCurve2DRibbon
	StackOfRotatedGrowthCurve2DRibbon *StackOfRotatedGrowthCurve2DRibbon
	GrowthCurve2DRibbon               *GrowthCurve2DRibbon
	ShiftedRightGrowthCurve2DRibbon   *ShiftedRightGrowthCurve2DRibbon

	PartiallyGrowthCurve2DRibbon     *PartiallyGrowthCurve2DRibbon
	PartiallyGrowthCurve2DTrajectory *PartiallyGrowthCurve2DTrajectory
	PartiallyGrowthCurve2DTrajectoryP1P2 *PartiallyGrowthCurve2DTrajectoryP1P2
}

