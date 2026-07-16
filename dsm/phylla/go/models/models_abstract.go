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
	RadialThickness float64

	// RhombusSideLength is set by the user. It represents the physical length of the side of the
	// fundamental diamond (rhombus) shape. It acts as the scale or distance for each step
	// taken along the lattice paths.
	RhombusSideLength float64

	LibraryAbstractFields
	AbstractTypeFields

	IsSelected bool

	IsPlantDiagramsNodeExpanded bool
	PlantDiagrams               []*PlantDiagram

	AxesShape               *AxesShape
	ReferenceRhombus        *RhombusShape
	PlantCircumferenceShape *PlantCircumferenceShape
	GridPathShape           *GridPathShape

	InitialRhombusGridShape *InitialRhombusGridShape

	ExplanationTextShape    *ExplanationTextShape
	RotatedReferenceRhombus *RhombusShape

	RotatedPlantCircumferenceShape *PlantCircumferenceShape
	RotatedGridPathShape           *GridPathShape

	RotatedRhombusGridShape2 *RotatedRhombusGridShape

	GrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape
	GrowthVectorShape           *GrowthVectorShape

	PerpendicularVectorGrid        *PerpendicularVectorGrid
	PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway
	BaseVectorShapeGrid            *BaseVectorShapeGrid
	ArcNormalVectorShapeGrid       *ArcNormalVectorShapeGrid
	StartArcShapeGrid              *StartArcShapeGrid
	TopStartArcShapeGrid           *TopStartArcShapeGrid
	EndArcShapeGrid                *EndArcShapeGrid
	TopEndArcShapeGrid             *TopEndArcShapeGrid
	GrowthCurveBezierShapeGrid     *GrowthCurveBezierShapeGrid

	StackOfGrowthCurve    *StackOfGrowthCurve
	TopStackOfGrowthCurve *TopStackOfGrowthCurve
	ShiftedLeftStackOfGrowthCurve *ShiftedLeftStackOfGrowthCurve
	ShiftedLeftStackOfNormalVector *ShiftedLeftStackOfNormalVector

	GrowthCurve2D    *GrowthCurve2D
	TopGrowthCurve2D *TopGrowthCurve2D
}
