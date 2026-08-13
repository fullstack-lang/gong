package models

// PlantAbstract
// In botanical phyllotaxy, N and M denote a contact parastichy pair, which represents
// the number of visible spirals winding in opposite directions around a plant's central
// axis. By standard convention, N and M are consecutive Fibonacci numbers ordered such
// that N < M (e.g., 8 and 13). These variables do not strictly dictate clockwise or
// counter-clockwise directions; rather, they quantify the two opposing sets of spirals.
type ViewType string

const (
	VIEW_PLANT_2D            ViewType = "Plant 2D"
	VIEW_VASE_FORM           ViewType = "Vase Form"
	VIEW_VASE_2D             ViewType = "Vase 2D"
	VIEW_VASE_3D             ViewType = "Vase 3D"
	VIEW_STOOL_3D            ViewType = "Stool 3D"
	VIEW_CLOCK_3D            ViewType = "Clock 3D"
	VIEW_ABOUT_SPIRAL_PLANTS ViewType = "About Spiral Plants"
)

type PlantType string

const (
	Plant PlantType = "Plant"
	Vase  PlantType = "Vase"
	Stool PlantType = "Stool"
	Clock PlantType = "Clock"
)

// Depending on the specific chirality (handedness) of the plant specimen, N may
// represent the number of clockwise spirals while M represents the counter-clockwise
// spirals, or vice versa.
type PlantAbstract struct {
	Name string

	N int
	M int

	StackHeight int // height of growth curve stack

	// RhombusInsideAngle is set by the user. It represents the inside angle (in degrees) of the
	// fundamental diamond (rhombus) shape that makes up the grid of leaves on the plant's surface.
	// This angle determines the geometric direction of the two main lattice paths (up-right and up-left).
	RhombusInsideAngle float64

	// RhombusSideLength is set by the user. It represents the physical length of the side of the
	// fundamental diamond (rhombus) shape. It acts as the scale or distance for each step
	// taken along the lattice paths.
	RhombusSideLength float64

	PlantType PlantType

	VaseAbstract  *VaseAbstract
	StoolAbstract *StoolAbstract
	ClockAbstract *ClockAbstract

	CurrentView ViewType

	LibraryAbstractFields
	AbstractTypeFields

	IsSelected bool

	IsPlantDiagramsNodeExpanded bool
	PlantDiagrams               []*PlantDiagram

	AxesShape *AxesShape

	RhombusStuff *RhombusStuff

	GrowthVectorShape *GrowthVectorShape

	PerpendicularVectorGrid  *PerpendicularVectorGrid
	BaseVectorShapeGrid      *BaseVectorShapeGrid
	ArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid

	StartArcShapeGrid     *StartArcShapeGrid
	MidArcVectorShapeGrid *MidArcVectorShapeGrid
	EndArcShapeGrid       *EndArcShapeGrid

	GrowthCurve2D *GrowthCurve2D

	// StackOfGrowthCurve2DByGrowthVector draws StackHeight copies of GrowthCurve2D,
	// each translated by the growth vector (k * GrowthVectorShape).
	StackOfGrowthCurve2DByGrowthVector *StackOfGrowthCurve2DByGrowthVector
}
